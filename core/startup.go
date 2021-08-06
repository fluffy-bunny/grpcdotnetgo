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
	GRPCPort    int    `mapstructure:"GRPC_PORT"`
}

// ConfigDefaultYaml default yaml
var coreConfigBaseYaml = []byte(`
APPLICATION_ENVIRONMENT: in-environment
GRPC_PORT: 0000
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
func loadConfig(configOptions *ConfigOptions) error {
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
	err = loadConfig(&ConfigOptions{
		Destination:    &dst,
		RootConfigYaml: coreConfigBaseYaml,
	})
	return &dst, err
}

var coreConfig *Config

func Start(startup IStartup) {
	var err error
	coreConfig, err = loadCoreConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(coreConfig.Environment)
	// Create a Builder with the default scopes (App, Request, SubRequest).
	dotNetGoBuilder, err := grpcdotnetgo.NewDotNetGoBuilder()
	if err != nil {
		panic(err)
	}
	startup.Startup()
	configOptions := startup.GetConfigOptions()
	err = loadConfig(configOptions)
	if err != nil {
		panic(err)
	}
	// add the main config into the DI directly
	servicesConfig.AddConfig(dotNetGoBuilder.Builder, configOptions.Destination)

	startup.ConfigureServices(dotNetGoBuilder.Builder)
	dotNetGoBuilder.Build()
	unaryServerInterceptorBuilder := UnaryServerInterceptorBuilder{}
	startup.Configure(grpcdotnetgo.GetContainer(), &unaryServerInterceptorBuilder)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryServerInterceptorBuilder.UnaryServerInterceptors...,
		)),
	)
	startup.RegisterGRPCEndpoints(grpcServer)
	servicesBackgroundTasks.GetBackgroundTasksFromContainer(grpcdotnetgo.GetContainer())

	future := asyncServeGRPC(grpcServer, startup.GetPort())

	sig := utils.WaitSignal()
	log.Info().Str("sig", sig.String()).Msg("Interupt triggered")

	grpcServer.Stop()
	grpcdotnetgo.GetContainer().DeleteWithSubContainers()
	future.Get()

}

func asyncServeGRPC(grpcServer *grpc.Server, port int) async.Future {

	promise := async.NewPromise()

	go func() {
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

		defer promise.Success(&grpcdotnetgoasync.AsyncResponse{
			Message: "End Serve - grpc Server",
			Error:   err,
		})

		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			return
		}

		if err = grpcServer.Serve(lis); err != nil {
			return
		}
		log.Info().Msg("grpc Server has shut down....")

	}()
	return promise.Future()

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
