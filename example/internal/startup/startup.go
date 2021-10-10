package startup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	backgroundCounterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/cron/counter"
	backgroundWelcomeService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/onetime/welcome"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	singletonService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/singleton"
	transientService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/transient"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oauth2"
	grpcdotnetgo_core_types "github.com/fluffy-bunny/grpcdotnetgo/pkg/core/types"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext/middleware"
	middleware_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/logger"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	middleware_grpc_recovery "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/recovery"
	grpcDIProtoError "github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error"
	mockoidcservice "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/test/mockoidcservice"
	di "github.com/fluffy-bunny/sarulabsdi"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getConfigPath() string {
	var configPath string
	_, err := os.Stat("../etc/config")
	if !os.IsNotExist(err) {
		configPath, _ = filepath.Abs("../etc/config")
		log.Info().Str("path", configPath).Msg("Configuration Root Folder")
	}
	return configPath
}

type Startup struct {
	MockOIDCService interface{}
	ConfigOptions   *grpcdotnetgo_core_types.ConfigOptions
	RootContainer   di.Container
}

func NewStartup() grpcdotnetgo_core_types.IStartup {
	startup := &Startup{}
	startup.ctor()
	return startup
}

func (s *Startup) ctor() {
	s.ConfigOptions = &grpcdotnetgo_core_types.ConfigOptions{
		Destination: &internal.Config{},
		RootConfig:  internal.ConfigDefaultYaml,
		ConfigPath:  getConfigPath(),
	}
}

func (s *Startup) GetConfigOptions() *grpcdotnetgo_core_types.ConfigOptions {
	return s.ConfigOptions
}

func (s *Startup) SetRootContainer(container di.Container) {
	s.RootContainer = container
}

func (s *Startup) GetPort() int {
	config := s.ConfigOptions.Destination.(*internal.Config)
	return config.Example.GRPCPort
}
func (s *Startup) ConfigureServices(builder *di.Builder) {
	// this is how  you get your config before you register your services
	config := s.ConfigOptions.Destination.(*internal.Config)

	var mm = make(map[string]middleware_oidc.EntryPointConfig)

	for k, v := range config.Example.OIDCConfig.EntryPoints {
		mm[k] = v
	}
	for k, v := range mm {
		delete(config.Example.OIDCConfig.EntryPoints, k)
		config.Example.OIDCConfig.EntryPoints[v.FullMethodName] = v
	}
	handlerGreeterService.AddGreeterService(builder)
	handlerGreeterService.AddGreeter2Service(builder)

	singletonService.AddSingletonService(builder)

	transientService.AddTransientService(builder)
	if config.Example.EnableTransient2 {
		transientService.AddTransientService2(builder)
	}

	backgroundCounterService.AddCronCounterJobProvider(builder)
	backgroundWelcomeService.AddOneTimeWelcomeJobProvider(builder)

	mockoidcservice.AddMockOIDCService(builder)

	middleware_oidc.AddOIDCConfigAccessor(builder, config)
	//	backgroundOidcService.AddCronOidcJobProvider(builder)
	//	services_oidc.AddOIDCAuthHandler(builder)

}
func (s *Startup) Configure(unaryServerInterceptorBuilder grpcdotnetgo_core_types.IUnaryServerInterceptorBuilder) {

	// this is how  you get your config before you register your services
	config := s.ConfigOptions.Destination.(*internal.Config)

	grpcFuncAuthConfig := oauth2.NewGrpcFuncAuthConfig(config.Example.OIDCConfig.Authority,
		"bearer", 5)
	for _, v := range config.Example.OIDCConfig.EntryPoints {

		methodClaims := oauth2.MethodClaims{
			OR:  []oauth2.Claim{},
			AND: []oauth2.Claim{},
		}

		for _, vv := range v.ClaimsConfig.AND {
			methodClaims.AND = append(methodClaims.AND, oauth2.Claim{
				Type:  vv.Type,
				Value: vv.Value,
			})

		}

		for _, vv := range v.ClaimsConfig.OR {
			methodClaims.OR = append(methodClaims.OR, oauth2.Claim{
				Type:  vv.Type,
				Value: vv.Value,
			})
		}

		grpcFuncAuthConfig.FullMethodNameToClaims[v.FullMethodName] = methodClaims
	}
	oidcContext, err := oauth2.BuildOpenIdConnectContext(grpcFuncAuthConfig)
	if err != nil {
		panic(err)
	}

	//var recoveryFunc middleware_grpc_recovery.RecoveryHandlerFunc
	recoveryOpts := []middleware_grpc_recovery.Option{
		middleware_grpc_recovery.WithRecoveryHandlerUnary(recoveryUnaryFunc),
	}
	unaryServerInterceptorBuilder.Use(grpc_ctxtags.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureContextLoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureCorrelationIDUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_dicontext.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.LoggingUnaryServerInterceptor())

	//	authHandler := middleware_grpc_auth.GetAuthFuncAccessorFromContainer(serviceProvider.GetContainer())
	//	unaryServerInterceptorBuilder.Use(middleware_grpc_auth.UnaryServerInterceptor(authHandler))

	unaryServerInterceptorBuilder.Use(oauth2.OAuth2UnaryServerInterceptor(oidcContext))
	unaryServerInterceptorBuilder.Use(oauth2.FinalAuthVerificationMiddleware(s.RootContainer))

	unaryServerInterceptorBuilder.Use(middleware_grpc_recovery.UnaryServerInterceptor(recoveryOpts...))

	s.MockOIDCService = mockoidcservice.GetMockOIDCServiceFromContainer(s.RootContainer)
}

// RegisterGRPCEndpoints registeres all our servers with the framework
func (s *Startup) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	var endpoints []interface{}
	endpoints = append(endpoints, pb.RegisterGreeterServerDI(server))
	endpoints = append(endpoints, pb.RegisterGreeter2ServerDI(server))
	return endpoints
}

// GetStartupManifest wrapper
func (s *Startup) GetStartupManifest() grpcdotnetgo_core_types.StartupManifest {
	return grpcdotnetgo_core_types.StartupManifest{
		Name:    "hello",
		Version: "test.1",
	}
}

// OnPreServerStartup wrapper
func (s *Startup) OnPreServerStartup() error {
	return nil
}

// OnPostServerShutdown Wrapper
func (s *Startup) OnPostServerShutdown() {}

func recoveryUnaryFunc(fullMethodName string, p interface{}) (interface{}, error) {
	fmt.Printf("p: %+v\n", p)

	replyFunc := pb.Get_helloworldFullEmptyResponseFromFullMethodName(fullMethodName)
	if replyFunc != nil {
		reply, ok2 := replyFunc().(grpcDIProtoError.IError)
		if ok2 {
			myError := reply.GetError()
			myError.Code = 503
			myError.Message = "Unexpected error2"
			return reply, status.Error(codes.Internal, "Unexpected error2")
		}
	}

	return nil, status.Error(codes.Internal, "Unexpected error1")

}
