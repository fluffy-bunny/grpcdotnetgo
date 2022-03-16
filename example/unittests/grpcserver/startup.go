package grpcserver

import (
	"fmt"
	"os"
	"path/filepath"

	contracts_config "github.com/fluffy-bunny/grpcdotnetgo/example/internal/contracts/config"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	backgroundCounterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/cron/counter"
	backgroundWelcomeService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/onetime/welcome"
	healthService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/health"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	services_scoped "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/scoped"
	services_singleton "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/singleton"
	services_transient "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/transient"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/auth/oauth2"
	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext/middleware"
	middleware_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/logger"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	middleware_grpc_recovery "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/recovery"
	grpcDIProtoError "github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error"
	mockoidcservice "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/test/mockoidcservice"
	di "github.com/fluffy-bunny/sarulabsdi"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // justified
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
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

// Startup type
type Startup struct {
	MockOIDCService interface{}
	ConfigOptions   *coreContracts.ConfigOptions
	RootContainer   di.Container
}

// NewStartup creates a new IStartup object
func NewStartup() coreContracts.IStartup {
	startup := &Startup{}
	startup.ctor()
	return startup
}

func (s *Startup) ctor() {
	s.ConfigOptions = &coreContracts.ConfigOptions{
		Destination: &contracts_config.Config{},
		RootConfig:  contracts_config.ConfigDefaultJSON,
		ConfigPath:  getConfigPath(),
	}
}

// GetConfigOptions is called by the runtime to determine where to write the configuration information to
func (s *Startup) GetConfigOptions() *coreContracts.ConfigOptions {
	return s.ConfigOptions
}

// SetRootContainer is called by the framework letting us now the root DI container
func (s *Startup) SetRootContainer(container di.Container) {
	s.RootContainer = container
}

// GetPort get the port number
func (s *Startup) GetPort() int {
	config := s.ConfigOptions.Destination.(*contracts_config.Config)
	return config.Example.Port
}

// ConfigureServices is where we register our services with the DI
func (s *Startup) ConfigureServices(builder *di.Builder) {
	// this is how  you get your config before you register your services
	config := s.ConfigOptions.Destination.(*contracts_config.Config)

	var mm = make(map[string]*middleware_oidc.EntryPointConfig)

	for k, v := range config.Example.OIDCConfig.EntryPoints {
		mm[k] = v
	}
	for k, v := range mm {
		delete(config.Example.OIDCConfig.EntryPoints, k)
		config.Example.OIDCConfig.EntryPoints[v.FullMethodName] = v
	}
	handlerGreeterService.AddScopedIGreeterService(builder)
	handlerGreeterService.AddScopedIGreeter2Service(builder)

	services_singleton.AddSingletonISingleton(builder)
	services_scoped.AddScopedIScoped(builder)
	services_transient.AddTransientITransient(builder)
	if config.Example.EnableTransient2 {
		services_transient.AddTransientITransient2(builder)
	}

	backgroundCounterService.AddCronCounterJobProvider(builder)
	backgroundWelcomeService.AddOneTimeWelcomeJobProvider(builder)

	mockoidcservice.AddMockOIDCService(builder)

	middleware_oidc.AddOIDCConfigAccessor(builder, config)
	//	backgroundOidcService.AddCronOidcJobProvider(builder)
	//	services_oidc.AddOIDCAuthHandler(builder)

	healthService.AddSingletonHealthService(builder)
}

// Configure setups up our middleware
func (s *Startup) Configure(unaryServerInterceptorBuilder coreContracts.IUnaryServerInterceptorBuilder) {
	// this is how  you get your config before you register your services
	config := s.ConfigOptions.Destination.(*contracts_config.Config)

	grpcFuncAuthConfig := oauth2.NewGrpcFuncAuthConfig(config.Example.OIDCConfig.Authority,
		"bearer", 5)
	for _, v := range config.Example.OIDCConfig.EntryPoints {
		methodClaims := oauth2.MethodClaims{
			OR:  []claimsprincipalContracts.Claim{},
			AND: []claimsprincipalContracts.Claim{},
		}

		for _, vv := range v.ClaimsConfig.AND {
			methodClaims.AND = append(methodClaims.AND, claimsprincipalContracts.Claim{
				Type:  vv.Type,
				Value: vv.Value,
			})
		}

		for _, vv := range v.ClaimsConfig.OR {
			methodClaims.OR = append(methodClaims.OR, claimsprincipalContracts.Claim{
				Type:  vv.Type,
				Value: vv.Value,
			})
		}

		grpcFuncAuthConfig.FullMethodNameToClaims[v.FullMethodName] = methodClaims
	}

	//var recoveryFunc middleware_grpc_recovery.RecoveryHandlerFunc
	recoveryOpts := []middleware_grpc_recovery.Option{
		middleware_grpc_recovery.WithRecoveryHandlerUnary(recoveryUnaryFunc),
	}
	unaryServerInterceptorBuilder.Use(grpc_ctxtags.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureContextLoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureCorrelationIDUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_dicontext.UnaryServerInterceptor(s.RootContainer))
	unaryServerInterceptorBuilder.Use(middleware_logger.LoggingUnaryServerInterceptor())

	unaryServerInterceptorBuilder.Use(middleware_grpc_recovery.UnaryServerInterceptor(recoveryOpts...))
}

// RegisterGRPCEndpoints registeres all our servers with the framework
func (s *Startup) RegisterGRPCEndpoints(server *grpc.Server) []interface{} {
	var endpoints []interface{}
	endpoints = append(endpoints, pb.RegisterGreeterServerDI(server))
	endpoints = append(endpoints, pb.RegisterGreeter2ServerDI(server))
	healthServer, _ := coreContracts.SafeGetIHealthServerFromContainer(s.RootContainer)
	if healthServer != nil {
		health.RegisterHealthServer(server, healthServer)
		endpoints = append(endpoints, healthServer)
	}
	return endpoints
}

// GetStartupManifest wrapper
func (s *Startup) GetStartupManifest() coreContracts.StartupManifest {
	return coreContracts.StartupManifest{
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
