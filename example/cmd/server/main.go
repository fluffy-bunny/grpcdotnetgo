// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/core"
	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	"github.com/rs/zerolog/log"

	services_oidc "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/oidc"

	backgroundCounterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/cron/counter"
	backgroundWelcomeService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/background/onetime/welcome"
	handlerGreeterService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/helloworld/handler"
	singletonService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/singleton"
	transientService "github.com/fluffy-bunny/grpcdotnetgo/example/internal/services/transient"
	middleware_grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/middleware/auth"
	middleware_dicontext "github.com/fluffy-bunny/grpcdotnetgo/middleware/dicontext"
	middleware_logger "github.com/fluffy-bunny/grpcdotnetgo/middleware/logger"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc"
	middleware_grpc_recovery "github.com/fluffy-bunny/grpcdotnetgo/middleware/recovery"
	grpcDIProtoError "github.com/fluffy-bunny/grpcdotnetgo/proto/error"
	runtime "github.com/fluffy-bunny/grpcdotnetgo/runtime"
	backgroundOidcService "github.com/fluffy-bunny/grpcdotnetgo/services/oidc"
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	mockoidcservice "github.com/fluffy-bunny/grpcdotnetgo/services/test/mockoidcservice"
	pkg "github.com/fluffy-bunny/protoc-gen-go-di/pkg"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gogo/protobuf/gogoproto"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var version = "development"

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
	port            int
	MockOIDCService interface{}
	ConfigOptions   *grpcdotnetgocore.ConfigOptions
}

func NewStartup() grpcdotnetgocore.IStartup {
	obj := &Startup{}
	obj.ctor()
	return obj
}

func (s *Startup) ctor() {
	s.ConfigOptions = &grpcdotnetgocore.ConfigOptions{
		Destination:    &internal.Config{},
		RootConfigYaml: internal.ConfigDefaultYaml,
		ConfigPath:     getConfigPath(),
	}
}

func (s *Startup) GetConfigOptions() *grpcdotnetgocore.ConfigOptions {
	return s.ConfigOptions
}
func (s *Startup) SetPort(port int) {
	s.port = port
}
func (s *Startup) GetPort() int {
	return s.port
}
func (s *Startup) ConfigureServices(builder *di.Builder) {
	// this is how  you get your config before you register your services
	config := s.ConfigOptions.Destination.(*internal.Config)

	var mm = make(map[string]middleware_oidc.EntryPointConfig)

	for k, v := range config.OIDCConfig.EntryPoints {
		mm[k] = v
	}
	for k, v := range mm {
		delete(config.OIDCConfig.EntryPoints, k)
		config.OIDCConfig.EntryPoints[v.FullMethodName] = v
	}
	handlerGreeterService.AddGreeterService(builder)
	handlerGreeterService.AddGreeter2Service(builder)

	singletonService.AddSingletonService(builder)

	transientService.AddTransientService(builder)
	if config.EnableTransient2 {
		transientService.AddTransientService2(builder)
	}

	backgroundCounterService.AddCronCounterJobProvider(builder)
	backgroundWelcomeService.AddOneTimeWelcomeJobProvider(builder)

	mockoidcservice.AddMockOIDCService(builder)

	middleware_oidc.AddOIDCConfigAccessor(builder, config)
	backgroundOidcService.AddCronOidcJobProvider(builder)
	services_oidc.AddOIDCAuthHandler(builder)

}
func (s *Startup) Configure(
	serviceProvider servicesServiceProvider.IServiceProvider,
	unaryServerInterceptorBuilder *grpcdotnetgocore.UnaryServerInterceptorBuilder) {

	// this is how  you get your config before you register your services
	//config := s.ConfigOptions.Destination.(*internal.Config)

	//var recoveryFunc middleware_grpc_recovery.RecoveryHandlerFunc
	recoveryOpts := []middleware_grpc_recovery.Option{
		middleware_grpc_recovery.WithRecoveryHandlerUnary(recoveryUnaryFunc),
	}
	unaryServerInterceptorBuilder.Use(grpc_ctxtags.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureContextLoggingUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.EnsureCorrelationIDUnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_dicontext.UnaryServerInterceptor())
	unaryServerInterceptorBuilder.Use(middleware_logger.LoggingUnaryServerInterceptor())

	authHandler := middleware_grpc_auth.GetAuthFuncAccessorFromContainer(serviceProvider.GetContainer())
	unaryServerInterceptorBuilder.Use(middleware_grpc_auth.UnaryServerInterceptor(authHandler))
	unaryServerInterceptorBuilder.Use(middleware_grpc_recovery.UnaryServerInterceptor(recoveryOpts...))

	s.MockOIDCService = mockoidcservice.GetMockOIDCService()

}
func (s *Startup) RegisterGRPCEndpoints(server *grpc.Server) {
	pb.RegisterGreeterServerDI(server)
	pb.RegisterGreeter2ServerDI(server)
}

func main() {
	d := gogoproto.E_GoprotoEnumStringer
	if d == nil {
		panic("boo hoo")
	}
	runtime.SetVersion(version)
	fmt.Println("Version:\t", version)

	fmt.Println(internal.PrettyJSON(pkg.NewFullMethodNameToMap(
		func(fullMethodName string) interface{} {
			return make(map[string]interface{})
		},
	)))
	startup := NewStartup()
	runtime.Start(startup)

}

func exampleAuthFunc(ctx context.Context, fullMethodName string) (context.Context, interface{}, error) {

	token, err := middleware_grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil || token == "" {
		replyFunc := pb.Get_helloworldFullEmptyResponseWithErrorFromFullMethodName(fullMethodName)
		if replyFunc != nil {
			reply := replyFunc()
			replyError, ok2 := reply.(grpcDIProtoError.IError)
			if ok2 {
				myError := replyError.GetError()
				myError.Code = 401
				myError.Message = "Unauthorized"
				return ctx, reply, status.Error(codes.Unauthenticated, "Unauthorized")
			}
		}
		return ctx, nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	return ctx, nil, nil
}
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
