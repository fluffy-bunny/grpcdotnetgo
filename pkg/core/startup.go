package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/structs"
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo/pkg"
	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	contracts_backgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	contracts_core "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	contracts_grpc "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/grpc"
	contracts_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
	grpcdotnetgo_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin"
	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/config"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
	"github.com/fluffy-bunny/viperEx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/reugn/async"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpclog "google.golang.org/grpc/grpclog"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	grpc_reflection "google.golang.org/grpc/reflection"
)

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(configPath string) error {
	s, err := os.Stat(configPath)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", configPath)
	}
	return nil
}
func LoadConfig(configOptions *contracts_core.ConfigOptions) error {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	var err error
	v.SetConfigType("json")
	// Environment Variables override everything.
	v.AutomaticEnv()

	// 1. Read in as buffer to set a default baseline.
	err = v.ReadConfig(bytes.NewBuffer(configOptions.RootConfig))
	if err != nil {
		log.Err(err).Msg("ConfigDefaultYaml did not read in")
		return err
	}

	environment := os.Getenv("APPLICATION_ENVIRONMENT")

	if len(environment) > 0 && len(configOptions.ConfigPath) != 0 {
		v.AddConfigPath(configOptions.ConfigPath)

		configFile := "appsettings." + environment + ".json"
		configPath := path.Join(configOptions.ConfigPath, configFile)
		err = ValidateConfigPath(configPath)
		if err == nil {
			v.SetConfigFile(configPath)
			err = v.MergeInConfig()
			if err != nil {
				return err
			}
			log.Info().Str("configPath", configPath).Msg("Merging in config")
		} else {
			log.Info().Str("configPath", configPath).Msg("Config file not present")
		}
	}

	// we need to do a viper Unmarshal because that is the only way we get the
	// ENV variables to come in
	err = v.Unmarshal(configOptions.Destination)
	if err != nil {
		return err
	}
	// we do all settings here, becuase a v.AllSettings will NOT bring in the ENV variables
	structs.DefaultTagName = "mapstructure"
	allSettings := structs.Map(configOptions.Destination)

	// normal viper stuff
	myViperEx, err := viperEx.New(allSettings, func(ve *viperEx.ViperEx) error {
		ve.KeyDelimiter = "__"
		return nil
	})
	if err != nil {
		return err
	}
	myViperEx.UpdateFromEnv()
	err = myViperEx.Unmarshal(configOptions.Destination)
	return err
}

// ServerInstance represents an instance of a plugin
type ServerInstance struct {
	StartupManifest contracts_core.StartupManifest
	Server          *grpc.Server
	Future          async.Future[grpcdotnetgoasync.AsyncResponse]

	ServerGRPCGatewayMux *http.Server
	FutureGRPCGatewayMux async.Future[grpcdotnetgoasync.AsyncResponse]

	DotNetGoBuilder *grpcdotnetgo.DotNetGoBuilder
	Endpoints       []interface{}
}

// Runtime type
type Runtime struct {
	ServerInstances []*ServerInstance
	waitChannel     chan os.Signal
}

// NewRuntime returns an instance of a new Runtime
func NewRuntime() *Runtime {
	return &Runtime{
		waitChannel: make(chan os.Signal),
	}
}

// Stop ...
func (s *Runtime) Stop() {
	s.waitChannel <- os.Interrupt
}

// Wait for someone to call stop
func (s *Runtime) Wait() {
	signal.Notify(
		s.waitChannel,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-s.waitChannel
}

// GetServerInstances gets the array or service instances
func (s *Runtime) GetServerInstances() []*ServerInstance {
	return s.ServerInstances
}

// Start to be used in production
func (s *Runtime) Start() {
	s.StartWithListenterAndPlugins(nil, nil)
}

// setup globa logger during runtime
var _setLogger bool
var once sync.Once

// TODO: There is a race condition using the zerologger replacement for grpcLog
// I haven't tracked it down yet so not using it under test.
func ensureLogger() {
	if _setLogger {
		return
	}
	once.Do(func() {
		_setLogger = true
		appEnv := os.Getenv("APPLICATION_ENVIRONMENT")
		appEnv = strings.ToLower(appEnv)
		if appEnv == "production" {
			grpclog.SetLoggerV2(NewGRPCLogger())
		}
	})

}

// StartWithListenterAndPlugins starts up the server
func (s *Runtime) StartWithListenterAndPlugins(lis net.Listener, plugins []contracts_plugin.IGRPCDotNetGoPlugin) {
	if plugins == nil || len(plugins) == 0 {
		plugins = grpcdotnetgo_plugin.GetPlugins() // pull it from the global one
	}
	// start the pprof web server
	pProfServer := NewPProfServer()
	pProfServer.Start()
	defer func() {
		pProfServer.Stop()
	}()
	// start the go profiler
	pprof := NewPProf()
	pprof.Start()
	defer func() {
		pprof.Stop()
	}()

	control := NewControl(s)
	control.Start()
	defer func() {
		control.Stop()
	}()

	logFormat := os.Getenv("LOG_FORMAT")
	if len(logFormat) == 0 {
		logFormat = "json"
	}
	logFileName := os.Getenv("LOG_FILE")
	if len(logFileName) == 0 {
		logFileName = "stderr"
	}
	var logFile *os.File
	// validate log destination
	var target io.Writer
	switch logFileName {
	case "stderr":
		target = os.Stderr
	case "stdout":
		target = os.Stdout
	default:
		// Open the log file
		var err error
		logFileName = fixPath(logFileName)
		if logFile, err = os.Create(logFileName); err != nil {
			log.Fatal().Err(err).Msg("Creating log file")
		}

		// Pass the ioWriter to the logger
		target = logFile
	}
	var err error
	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	prettyLog := false
	prettyLogValue := os.Getenv("PRETTY_LOG")
	if len(prettyLogValue) != 0 {
		b, err := strconv.ParseBool(prettyLogValue)
		if err == nil {
			prettyLog = b
		}
	}
	if prettyLog || logFormat == "pretty" {
		target = zerolog.ConsoleWriter{Out: target}
	}
	log.Logger = log.Output(target)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	switch strings.ToLower(logLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	ensureLogger()
	ctx := context.Background()

	log := zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
	ctx = log.WithContext(ctx)

	for _, plugin := range plugins {
		si := &ServerInstance{}

		// Create a Builder with the default scopes (App, Request, SubRequest).
		si.DotNetGoBuilder, err = grpcdotnetgo.NewDotNetGoBuilder()
		if err != nil {
			panic(err)
		}
		si.DotNetGoBuilder.AddDefaultService()

		startup := plugin.GetStartup()
		startup.SetContext(ctx)
		configOptions := startup.GetConfigOptions()
		err = LoadConfig(configOptions)
		if err != nil {
			panic(err)
		}
		si.StartupManifest = startup.GetStartupManifest()

		// add the main config into the DI directly
		servicesConfig.AddConfig(si.DotNetGoBuilder.Builder, configOptions.Destination)

		startup.ConfigureServices(si.DotNetGoBuilder.Builder)
		si.DotNetGoBuilder.Build()
		rootContainer := si.DotNetGoBuilder.Container
		startup.SetRootContainer(rootContainer)

		unaryServerInterceptorBuilder := NewUnaryServerInterceptorBuilder()
		startup.Configure(unaryServerInterceptorBuilder)

		grpcServer := grpc.NewServer(
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
			}),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				unaryServerInterceptorBuilder.GetUnaryServerInterceptors()...,
			)),
		)
		enableGRPCReflection := utils.BoolEnv("ENABLE_GRPC_SERVER_REFLECTION", false)
		if enableGRPCReflection {
			log.Info().Msg("Enabling GRPC Server Reflection")
			grpc_reflection.Register(grpcServer)
		}
		serverRegistrations, err := contracts_grpc.SafeGetManyIServiceEndpointRegistrationFromContainer(rootContainer)
		if err == nil {
			si.Endpoints = make([]interface{}, 0, len(serverRegistrations))
			for _, serverRegistration := range serverRegistrations {
				endpoint := serverRegistration.RegisterEndpoint(grpcServer)
				si.Endpoints = append(si.Endpoints, endpoint)
			}
			healthServer, _ := contracts_core.SafeGetIHealthServerFromContainer(rootContainer)
			if healthServer != nil {
				health.RegisterHealthServer(grpcServer, healthServer)
				si.Endpoints = append(si.Endpoints, healthServer)
			}
		} else {
			// legacy
			si.Endpoints = startup.RegisterGRPCEndpoints(grpcServer)
		}
		// TODO: Make this a first class abstaction
		// ILifeCycleHook but maybe IStartup can have those
		contracts_backgroundtasks.GetIBackgroundTasksFromContainer(rootContainer)

		err = startup.OnPreServerStartup()
		if err != nil {
			log.Error().Err(err).
				Interface("startupManifest", si.StartupManifest).Msgf("OnPreServerStartup failed")
			panic(err)
		} else {
			if lis == nil {
				if si.StartupManifest.Port == 0 {
					// legacy
					si.StartupManifest.Port = startup.GetPort()
				}
				lis, err = net.Listen("tcp", fmt.Sprintf(":%d", si.StartupManifest.Port))
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to listen")
				}
			}

			future := asyncServeGRPC(grpcServer, lis)
			si.Server = grpcServer
			si.Future = future
			s.ServerInstances = append(s.ServerInstances, si)

			if si.StartupManifest.GRPCGatewayEnabled {
				// Create a client connection to the gRPC server we just started
				// This is where the gRPC-Gateway proxies the requests
				conn, err := grpc.DialContext(
					context.Background(),
					fmt.Sprintf("0.0.0.0:%d", si.StartupManifest.Port),
					grpc.WithBlock(),
					grpc.WithTransportCredentials(insecure.NewCredentials()),
				)
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to dial server")
				}
				gwmux := grpc_gateway_runtime.NewServeMux()
				for _, serverRegistration := range serverRegistrations {
					serverRegistration.RegisterGatewayHandler(gwmux, conn)
				}
				gwServer := &http.Server{
					Addr:    fmt.Sprintf(":%d", si.StartupManifest.RESTPort),
					Handler: gwmux,
				}
				si.ServerGRPCGatewayMux = gwServer
				future := asyncServeGRPCGatewayMux(gwServer)
				si.FutureGRPCGatewayMux = future
			}

		}
	}
	s.Wait()
	log.Info().Msg("Interupt triggered")

	for _, v := range s.ServerInstances {
		// tell all grpc servers to stop
		v.Server.Stop()
		if v.StartupManifest.GRPCGatewayEnabled {
			if v.ServerGRPCGatewayMux != nil {
				v.ServerGRPCGatewayMux.Shutdown(context.Background())
			}
		}
		// tear down the DI Container
		v.DotNetGoBuilder.Container.DeleteWithSubContainers()
	}
	for i := 0; i < len(plugins); i++ {
		plugins[i].GetStartup().OnPostServerShutdown()
	}

	// do a future wait
	for _, v := range s.ServerInstances {
		v.Future.Join()
		if v.StartupManifest.GRPCGatewayEnabled {
			if v.FutureGRPCGatewayMux != nil {
				v.FutureGRPCGatewayMux.Join()
			}
		}
	}
}
func fixPath(fpath string) string {
	if fpath == "" {
		return ""
	}
	if fpath == "stdout" || fpath == "stderr" {
		return fpath
	}

	// Is it already absolute?
	if filepath.IsAbs(fpath) {
		return filepath.Clean(fpath)
	}

	// Make it absolute
	fpath, _ = filepath.Abs(fpath)

	return fpath
}
func asyncServeGRPC(grpcServer *grpc.Server, lis net.Listener) async.Future[grpcdotnetgoasync.AsyncResponse] {

	return grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise[grpcdotnetgoasync.AsyncResponse]) {
		var err error
		log.Info().Msg("gRPC Server Starting up")

		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End Serve - grpc Server",
				Error:   err,
			})
			if err != nil {
				log.Error().Err(err).Msg("gRPC Server exit")
				os.Exit(1)
			}
		}()

		if err = grpcServer.Serve(lis); err != nil {
			return
		}
		log.Info().Msg("grpc Server has shut down....")
	})
}
func asyncServeGRPCGatewayMux(httpServer *http.Server) async.Future[grpcdotnetgoasync.AsyncResponse] {
	return grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise[grpcdotnetgoasync.AsyncResponse]) {
		var err error
		log.Info().Msg("gRPC Server Starting up")

		defer func() {
			promise.Success(&grpcdotnetgoasync.AsyncResponse{
				Message: "End Serve - http Server",
				Error:   err,
			})
			if err != nil {
				log.Error().Err(err).Msg("gRPC Server exit")
				os.Exit(1)
			}
		}()

		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("Failed to listen")
			return
		}
		log.Info().Msg("GRPCGatewayMux Server has shut down....")
	})
}

const bufSize = 1024 * 1024
