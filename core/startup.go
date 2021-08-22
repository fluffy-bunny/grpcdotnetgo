package core

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"net"

	"github.com/fatih/structs"
	"github.com/fluffy-bunny/grpcdotnetgo"
	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/async"
	"github.com/fluffy-bunny/grpcdotnetgo/core/types"
	grpcdotnetgo_plugin "github.com/fluffy-bunny/grpcdotnetgo/plugin"
	servicesBackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/services/backgroundtasks"
	servicesConfig "github.com/fluffy-bunny/grpcdotnetgo/services/config"

	"github.com/fluffy-bunny/grpcdotnetgo/utils"
	"github.com/fluffy-bunny/viperEx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/reugn/async"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Config struct {
	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
}

// ConfigDefaultYaml default yaml
var coreConfigBaseYaml = []byte(`
APPLICATION_ENVIRONMENT: in-environment
`)

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
func loadConfig(configOptions *types.ConfigOptions) error {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	var err error
	v.SetConfigType("yaml")
	// Environment Variables override everything.
	v.AutomaticEnv()

	// 1. Read in as buffer to set a default baseline.
	err = v.ReadConfig(bytes.NewBuffer(configOptions.RootConfigYaml))
	if err != nil {
		log.Err(err).Msg("ConfigDefaultYaml did not read in")
		return err
	}

	environment := os.Getenv("APPLICATION_ENVIRONMENT")

	if len(environment) > 0 && len(configOptions.ConfigPath) != 0 {
		v.AddConfigPath(configOptions.ConfigPath)

		configFile := "appsettings." + coreConfig.Environment + ".yml"
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
	changeAllKeysToLowerCase(allSettings)

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

func loadCoreConfig() (*Config, error) {
	var err error
	dst := Config{}
	err = loadConfig(&types.ConfigOptions{
		Destination:    &dst,
		RootConfigYaml: coreConfigBaseYaml,
	})
	return &dst, err
}

var coreConfig *Config

type ServerInstance struct {
	Server          *grpc.Server
	Future          async.Future
	DotNetGoBuilder *grpcdotnetgo.DotNetGoBuilder
	Endpoints       []interface{}
}

var serverInstances []*ServerInstance

func GetServerInstances() []*ServerInstance {
	return serverInstances
}
func Start() {

	plugins := grpcdotnetgo_plugin.GetPlugins()
	var err error
	coreConfig, err = loadCoreConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(coreConfig.Environment)

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
		err = loadConfig(configOptions)
		if err != nil {
			panic(err)
		}
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
		si.Endpoints = startup.RegisterGRPCEndpoints(grpcServer)
		servicesBackgroundTasks.GetBackgroundTasksFromContainer(rootContainer)

		future := asyncServeGRPC(grpcServer, startup.GetPort())
		si.Server = grpcServer
		si.Future = future

		serverInstances = append(serverInstances, si)
	}
	sig := utils.WaitSignal()
	log.Info().Str("sig", sig.String()).Msg("Interupt triggered")

	for _, v := range serverInstances {
		// tell all grpc servers to stop
		v.Server.Stop()
		// tear down the DI Container
		v.DotNetGoBuilder.Container.DeleteWithSubContainers()
	}

	// do a future wait
	for _, v := range serverInstances {
		v.Future.Get()
	}

}

func asyncServeGRPC(grpcServer *grpc.Server, port int) async.Future {
	return grpcdotnetgoasync.ExecuteWithPromiseAsync(func(promise async.Promise) {
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

		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			return
		}

		if err = grpcServer.Serve(lis); err != nil {
			return
		}
		log.Info().Msg("grpc Server has shut down....")
	})
}

func changeAllKeysToLowerCase(m map[string]interface{}) {
	var lcMap = make(map[string]interface{})
	var currentKeys []string
	for key, value := range m {
		currentKeys = append(currentKeys, key)
		lcMap[strings.ToLower(key)] = value
	}
	// delete original values
	for _, k := range currentKeys {
		delete(m, k)
	}
	// put the lowercase ones in the original map
	for key, value := range lcMap {
		m[key] = value
		vMap, ok := value.(map[string]interface{})
		if ok {
			// if the current value is a map[string]interface{}, keep going
			changeAllKeysToLowerCase(vMap)
		}
	}
}
