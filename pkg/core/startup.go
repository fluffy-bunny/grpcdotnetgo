package core

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/structs"
	grpcdotnetgo "github.com/fluffy-bunny/grpcdotnetgo/pkg"
	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/pkg/async"
	contracts_core "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	contracts_grpc "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/grpc"
	contracts_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"

	grpcdotnetgo_plugin "github.com/fluffy-bunny/grpcdotnetgo/pkg/plugin"
	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/config"
	"github.com/fluffy-bunny/viperEx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/reugn/async"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	grpclog "google.golang.org/grpc/grpclog"
	health "google.golang.org/grpc/health/grpc_health_v1"
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
	Future          async.Future[interface{}]
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
	// Replace the grpc logger
	grpclog.SetLoggerV2(NewGRPCLogger())

	for _, plugin := range plugins {
		si := &ServerInstance{}

		// Create a Builder with the default scopes (App, Request, SubRequest).
		si.DotNetGoBuilder, err = grpcdotnetgo.NewDotNetGoBuilder()
		if err != nil {
			panic(err)
		}
		si.DotNetGoBuilder.AddDefaultService()

		startup := plugin.GetStartup()

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
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				unaryServerInterceptorBuilder.GetUnaryServerInterceptors()...,
			)),
		)
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
		//backgroundtasksContracts.GetIBackgroundTasksFromContainer(rootContainer)

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
					panic(err)
				}
			}

			future := asyncServeGRPC(grpcServer, lis)
			si.Server = grpcServer
			si.Future = future
			s.ServerInstances = append(s.ServerInstances, si)
		}
	}
	s.Wait()
	log.Info().Msg("Interupt triggered")

	for _, v := range s.ServerInstances {
		// tell all grpc servers to stop
		v.Server.Stop()
		// tear down the DI Container
		v.DotNetGoBuilder.Container.DeleteWithSubContainers()
	}
	for i := 0; i < len(plugins); i++ {
		plugins[i].GetStartup().OnPostServerShutdown()
	}

	// do a future wait
	for _, v := range s.ServerInstances {
		v.Future.Join()
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
func asyncServeGRPC(grpcServer *grpc.Server, lis net.Listener) async.Future[interface{}] {
	return grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise[interface{}]) {
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

const bufSize = 1024 * 1024
