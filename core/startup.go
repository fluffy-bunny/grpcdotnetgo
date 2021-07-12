package core

import (
	"bytes"
	"fmt"
	"os"

	"net"

	"github.com/fatih/structs"
	"github.com/fluffy-bunny/grpcdotnetgo"
	grpcdotnetgoasync "github.com/fluffy-bunny/grpcdotnetgo/async"
	servicesBackgroundTasks "github.com/fluffy-bunny/grpcdotnetgo/services/backgroundtasks"
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

func loadCoreConfig() (*Config, error) {
	v := viper.New()

	var err error
	v.SetConfigType("yaml")
	// Environment Variables override everything.
	v.AutomaticEnv()

	// 1. Read in as buffer to set a default baseline.
	err = v.ReadConfig(bytes.NewBuffer(coreConfigBaseYaml))
	if err != nil {
		log.Err(err).Msg("ConfigDefaultYaml did not read in")
		return nil, err
	}

	dst := Config{}
	// we need to do a viper Unmarshal because that is the only way we get the
	// ENV variables to come in
	err = v.Unmarshal(&dst)
	if err != nil {
		return nil, err
	}
	// we do all settings here, becuase a v.AllSettings will NOT bring in the ENV variables
	allSettings := structs.Map(dst)

	// normal viper stuff
	myViperEx, err := viperEx.New(allSettings, func(ve *viperEx.ViperEx) error {
		ve.KeyDelimiter = "__"
		return nil
	})
	if err != nil {
		return nil, err
	}
	myViperEx.UpdateFromEnv()
	err = myViperEx.Unmarshal(&dst)
	if err != nil {
		return nil, err
	}
	return &dst, nil

}

func Start(startup IStartup) {
	coreConfig, err := loadCoreConfig()
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
