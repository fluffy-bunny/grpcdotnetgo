package core

import (
	"bytes"
	"fmt"

	"net"

	"github.com/fatih/structs"
	"github.com/fluffy-bunny/grpcdotnetgo"
	"github.com/fluffy-bunny/viperEx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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

	port := fmt.Sprintf(":%v", startup.GetPort())
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Err(err).
			Str("port", port).
			Msg("failed to listen:")
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryServerInterceptorBuilder.UnaryServerInterceptors...,
		)),
	)
	startup.RegisterGRPCEndpoints(grpcServer)

	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve: ")
	}
}
