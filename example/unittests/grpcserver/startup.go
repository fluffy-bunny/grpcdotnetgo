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
	core_contracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
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

// Startup type
type Startup struct {
	core_contracts.UnimplementedStartup
	MockOIDCService interface{}
	ConfigOptions   *core_contracts.ConfigOptions
	RootContainer   di.Container
}

// NewStartup creates a new IStartup object
func NewStartup() core_contracts.IStartup {
	startup := &Startup{}
	startup.ctor()
	return startup
}

func (s *Startup) ctor() {
	s.ConfigOptions = &core_contracts.ConfigOptions{
		Destination: &contracts_config.Config{},
		RootConfig:  contracts_config.ConfigDefaultJSON,
		ConfigPath:  getConfigPath(),
	}
}

// GetConfigOptions is called by the runtime to determine where to write the configuration information to
func (s *Startup) GetConfigOptions() *core_contracts.ConfigOptions {
	return s.ConfigOptions
}

// SetRootContainer is called by the framework letting us now the root DI container
func (s *Startup) SetRootContainer(container di.Container) {
	s.RootContainer = container
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
	handlerGreeterService.AddGreeterEndpointRegistration(builder)
	handlerGreeterService.AddGreeter2EndpointRegistration(builder)

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
func (s *Startup) Configure(unaryServerInterceptorBuilder core_contracts.IUnaryServerInterceptorBuilder) {
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
				Type:  vv.Claim.Type,
				Value: vv.Claim.Value,
			})
		}

		for _, vv := range v.ClaimsConfig.OR {
			methodClaims.OR = append(methodClaims.OR, claimsprincipalContracts.Claim{
				Type:  vv.Claim.Type,
				Value: vv.Claim.Value,
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

// GetStartupManifest wrapper
func (s *Startup) GetStartupManifest() core_contracts.StartupManifest {

	config := s.ConfigOptions.Destination.(*contracts_config.Config)
	return core_contracts.StartupManifest{
		Name:    "hello",
		Version: "test.1",
		Port:    config.Example.Port,
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
