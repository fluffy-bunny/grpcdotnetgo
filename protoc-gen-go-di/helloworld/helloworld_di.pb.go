// Code generated by protoc-gen-go-di. DO NOT EDIT.

package helloworld

import (
	context "context"
	pkg "github.com/fluffy-bunny/grpcdotnetgo/pkg"
	grpc1 "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/grpc"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	pkg1 "github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/pkg"
	sarulabsdi "github.com/fluffy-bunny/sarulabsdi"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = pkg.SupportPackageIsVersion7

func setNewField_MYU74zn4kc2UFcLfMVw8lrRTTSk87VSC(dst interface{}, field string) {
	v := reflect.ValueOf(dst).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.New(v.Type().Elem()))
	}
}

type UnimplementedGreeterServerEndpointRegistration struct {
}

func (UnimplementedGreeterServerEndpointRegistration) RegisterGatewayHandler(gwmux *runtime.ServeMux, conn *grpc.ClientConn) {
}

// GreeterEndpointRegistration defines the grpc server endpoint registration
type GreeterEndpointRegistration struct {
}

// GreeterEndpointRegistration defines the grpc server endpoint registration
type GreeterEndpointRegistrationV2 struct {
	UnimplementedGreeterServerEndpointRegistration
}

// TypeGreeterEndpointRegistration reflect type
var TypeGreeterEndpointRegistration = sarulabsdi.GetInterfaceReflectType((*GreeterEndpointRegistration)(nil))

// AddGreeterEndpointRegistration adds a type that implements IServiceEndpointRegistration
func AddGreeterEndpointRegistration(builder *sarulabsdi.Builder, implType reflect.Type) {
	grpc1.AddSingletonIServiceEndpointRegistration(builder, reflect.TypeOf(&GreeterEndpointRegistration{}))
	AddScopedIGreeterService(builder, implType)
}

// AddGreeterEndpointRegistration adds a type that implements IServiceEndpointRegistration
func AddGreeterEndpointRegistrationV2(builder *sarulabsdi.Builder, implType reflect.Type) {
	grpc1.AddSingletonIServiceEndpointRegistration(builder, reflect.TypeOf(&GreeterEndpointRegistrationV2{}))
	AddScopedIGreeterServer(builder, implType)
}

// GetName returns the name of the service
func (s *GreeterEndpointRegistration) GetName() string {
	return "Greeter"
}

// GetName returns the name of the service
func (s *GreeterEndpointRegistrationV2) GetName() string {
	return "Greeter"
}

// GetNewClient returns a new instance of a grpc client
func (s *GreeterEndpointRegistration) GetNewClient(cc grpc.ClientConnInterface) interface{} {
	return NewGreeterClient(cc)
}

// GetNewClient returns a new instance of a grpc client
func (s *GreeterEndpointRegistrationV2) GetNewClient(cc grpc.ClientConnInterface) interface{} {
	return NewGreeterClient(cc)
}

// RegisterEndpoint registers a DI server
func (s *GreeterEndpointRegistration) RegisterEndpoint(server *grpc.Server) interface{} {
	endpoint := RegisterGreeterServerDI(server)
	return endpoint
}

// RegisterEndpoint registers a DI server
func (s *GreeterEndpointRegistrationV2) RegisterEndpoint(server *grpc.Server) interface{} {
	endpoint := RegisterGreeterServerDIV2(server)
	return endpoint
}

// RegisterEndpoint registers a DI server
func (s *GreeterEndpointRegistrationV2) RegisterEndpointV2(server *grpc.Server) interface{} {
	endpoint := RegisterGreeterServerDIV2(server)
	return endpoint
}

// IGreeterServer defines the grpc server
type IGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServerEx defines the grpc server
type UnimplementedGreeterServerEx struct {
	UnimplemtedErrorResponse func() error
}

func (UnimplementedGreeterServerEx) mustEmbedUnimplementedGreeterServer() {}
func (u UnimplementedGreeterServerEx) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	if u.UnimplemtedErrorResponse != nil {
		return nil, u.UnimplemtedErrorResponse()
	}
	return nil, status.Error(codes.Unimplemented, "method SayHello not implemented")
}

// IGreeterService defines the required downstream service interface
type IGreeterService interface {
	SayHello(request *HelloRequest) (*HelloReply, error)
}

// TypeIGreeterServer reflect type
var TypeIGreeterServer = sarulabsdi.GetInterfaceReflectType((*IGreeterService)(nil))

// TypeIGreeterService reflect type
var TypeIGreeterService = sarulabsdi.GetInterfaceReflectType((*IGreeterService)(nil))

// ReflectTypeIGreeterServer reflect type
var ReflectTypeIGreeterServer = sarulabsdi.GetInterfaceReflectType((*IGreeterServer)(nil))

// ReflectTypeIGreeterService reflect type
var ReflectTypeIGreeterService = sarulabsdi.GetInterfaceReflectType((*IGreeterService)(nil))

type GetGreeterClient func() (GreeterClient, error)

// AddSingletonIGreeterServerByObj adds a prebuilt obj
func AddSingletonIGreeterServerByObj(builder *sarulabsdi.Builder, obj interface{}) {
	sarulabsdi.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIGreeterServer)
}

// AddSingletonIGreeterServer adds a type that implements IGreeterServer
func AddSingletonIGreeterServer(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIGreeterServer)
}

// AddSingletonIGreeterServerByFunc adds a type by a custom func
func AddSingletonIGreeterServerByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeterServer)
}

// AddSingletonIGreeterServiceByObj adds a prebuilt obj
func AddSingletonIGreeterServiceByObj(builder *sarulabsdi.Builder, obj interface{}) {
	sarulabsdi.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIGreeterService)
}

// AddSingletonIGreeterService adds a type that implements IGreeterService
func AddSingletonIGreeterService(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIGreeterService)
}

// AddSingletonIGreeterServiceByFunc adds a type by a custom func
func AddSingletonIGreeterServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeterService)
}

// AddTransientIGreeterService adds a type that implements IGreeterService
func AddTransientIGreeterService(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIGreeterService)
}

// AddTransientIGreeterServiceByFunc adds a type by a custom func
func AddTransientIGreeterServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeterService)
}

// AddScopedIGreeterServer adds a type that implements IGreeterServer
func AddScopedIGreeterServer(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIGreeterServer)
}

// AddScopedIGreeterService adds a type that implements IGreeterService
func AddScopedIGreeterService(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIGreeterService)
}

// AddScopedIGreeterServiceByFunc adds a type by a custom func
func AddScopedIGreeterServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeterService)
}

// RemoveAllIGreeterService removes all IBillingService from the DI
func RemoveAllIGreeterService(builder *sarulabsdi.Builder) {
	builder.RemoveAllByType(ReflectTypeIGreeterService)
}

// GetGreeterServiceFromContainer fetches the downstream di.Request scoped service
func GetGreeterServiceFromContainer(ctn sarulabsdi.Container) IGreeterService {
	return ctn.GetByType(ReflectTypeIGreeterService).(IGreeterService)
}

// GetGreeterServerFromContainer fetches the downstream di.Request scoped service
func GetGreeterServerFromContainer(ctn sarulabsdi.Container) IGreeterServer {
	return ctn.GetByType(ReflectTypeIGreeterServer).(IGreeterServer)
}

// GetIGreeterServiceFromContainer fetches the downstream di.Request scoped service
func GetIGreeterServiceFromContainer(ctn sarulabsdi.Container) IGreeterService {
	return ctn.GetByType(ReflectTypeIGreeterService).(IGreeterService)
}

// GetIGreeterServerFromContainer fetches the downstream di.Request scoped service
func GetIGreeterServerFromContainer(ctn sarulabsdi.Container) IGreeterServer {
	return ctn.GetByType(ReflectTypeIGreeterServer).(IGreeterServer)
}

// SafeGetIGreeterServiceFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeterServiceFromContainer(ctn sarulabsdi.Container) (IGreeterService, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeterService)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeterService), nil
}

// SafeGetIGreeterServerFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeterServerFromContainer(ctn sarulabsdi.Container) (IGreeterServer, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeterServer)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeterServer), nil
}

// Impl for Greeter server instances
type greeterServer struct {
	UnimplementedGreeterServerEx
}

// Impl for Greeter server instances
type greeterServerV2 struct {
	UnimplementedGreeterServerEx
}

// RegisterGreeterServerDI ...
func RegisterGreeterServerDI(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeterServer{}
	RegisterGreeterServer(s, server)
	return server
}

// RegisterGreeterServerDIV2 ...
func RegisterGreeterServerDIV2(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeterServerV2{}
	RegisterGreeterServer(s, server)
	return server
}

// SayHello...
func (s *greeterServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeterServiceFromContainer(requestContainer)
	return downstreamService.SayHello(request)
}

// SayHello...
func (s *greeterServerV2) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeterServerFromContainer(requestContainer)
	return downstreamService.SayHello(ctx, request)
}

// FullMethodNames for Greeter
const (
	// FMN_Greeter_SayHello
	FMN_Greeter_SayHello = "/helloworld.Greeter/SayHello"
)

type UnimplementedGreeter2ServerEndpointRegistration struct {
}

func (UnimplementedGreeter2ServerEndpointRegistration) RegisterGatewayHandler(gwmux *runtime.ServeMux, conn *grpc.ClientConn) {
}

// Greeter2EndpointRegistration defines the grpc server endpoint registration
type Greeter2EndpointRegistration struct {
}

// Greeter2EndpointRegistration defines the grpc server endpoint registration
type Greeter2EndpointRegistrationV2 struct {
	UnimplementedGreeter2ServerEndpointRegistration
}

// TypeGreeter2EndpointRegistration reflect type
var TypeGreeter2EndpointRegistration = sarulabsdi.GetInterfaceReflectType((*Greeter2EndpointRegistration)(nil))

// AddGreeter2EndpointRegistration adds a type that implements IServiceEndpointRegistration
func AddGreeter2EndpointRegistration(builder *sarulabsdi.Builder, implType reflect.Type) {
	grpc1.AddSingletonIServiceEndpointRegistration(builder, reflect.TypeOf(&Greeter2EndpointRegistration{}))
	AddScopedIGreeter2Service(builder, implType)
}

// AddGreeter2EndpointRegistration adds a type that implements IServiceEndpointRegistration
func AddGreeter2EndpointRegistrationV2(builder *sarulabsdi.Builder, implType reflect.Type) {
	grpc1.AddSingletonIServiceEndpointRegistration(builder, reflect.TypeOf(&Greeter2EndpointRegistrationV2{}))
	AddScopedIGreeter2Server(builder, implType)
}

// GetName returns the name of the service
func (s *Greeter2EndpointRegistration) GetName() string {
	return "Greeter2"
}

// GetName returns the name of the service
func (s *Greeter2EndpointRegistrationV2) GetName() string {
	return "Greeter2"
}

// GetNewClient returns a new instance of a grpc client
func (s *Greeter2EndpointRegistration) GetNewClient(cc grpc.ClientConnInterface) interface{} {
	return NewGreeter2Client(cc)
}

// GetNewClient returns a new instance of a grpc client
func (s *Greeter2EndpointRegistrationV2) GetNewClient(cc grpc.ClientConnInterface) interface{} {
	return NewGreeter2Client(cc)
}

// RegisterEndpoint registers a DI server
func (s *Greeter2EndpointRegistration) RegisterEndpoint(server *grpc.Server) interface{} {
	endpoint := RegisterGreeter2ServerDI(server)
	return endpoint
}

// RegisterEndpoint registers a DI server
func (s *Greeter2EndpointRegistrationV2) RegisterEndpoint(server *grpc.Server) interface{} {
	endpoint := RegisterGreeter2ServerDIV2(server)
	return endpoint
}

// RegisterEndpoint registers a DI server
func (s *Greeter2EndpointRegistrationV2) RegisterEndpointV2(server *grpc.Server) interface{} {
	endpoint := RegisterGreeter2ServerDIV2(server)
	return endpoint
}

// IGreeter2Server defines the grpc server
type IGreeter2Server interface {
	mustEmbedUnimplementedGreeter2Server()
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error)
}

// UnimplementedGreeter2ServerEx defines the grpc server
type UnimplementedGreeter2ServerEx struct {
	UnimplemtedErrorResponse func() error
}

func (UnimplementedGreeter2ServerEx) mustEmbedUnimplementedGreeter2Server() {}
func (u UnimplementedGreeter2ServerEx) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error) {
	if u.UnimplemtedErrorResponse != nil {
		return nil, u.UnimplemtedErrorResponse()
	}
	return nil, status.Error(codes.Unimplemented, "method SayHello not implemented")
}

// IGreeter2Service defines the required downstream service interface
type IGreeter2Service interface {
	SayHello(request *HelloRequest) (*HelloReply2, error)
}

// TypeIGreeter2Server reflect type
var TypeIGreeter2Server = sarulabsdi.GetInterfaceReflectType((*IGreeter2Service)(nil))

// TypeIGreeter2Service reflect type
var TypeIGreeter2Service = sarulabsdi.GetInterfaceReflectType((*IGreeter2Service)(nil))

// ReflectTypeIGreeter2Server reflect type
var ReflectTypeIGreeter2Server = sarulabsdi.GetInterfaceReflectType((*IGreeter2Server)(nil))

// ReflectTypeIGreeter2Service reflect type
var ReflectTypeIGreeter2Service = sarulabsdi.GetInterfaceReflectType((*IGreeter2Service)(nil))

type GetGreeter2Client func() (Greeter2Client, error)

// AddSingletonIGreeter2ServerByObj adds a prebuilt obj
func AddSingletonIGreeter2ServerByObj(builder *sarulabsdi.Builder, obj interface{}) {
	sarulabsdi.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIGreeter2Server)
}

// AddSingletonIGreeter2Server adds a type that implements IGreeter2Server
func AddSingletonIGreeter2Server(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIGreeter2Server)
}

// AddSingletonIGreeter2ServerByFunc adds a type by a custom func
func AddSingletonIGreeter2ServerByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeter2Server)
}

// AddSingletonIGreeter2ServiceByObj adds a prebuilt obj
func AddSingletonIGreeter2ServiceByObj(builder *sarulabsdi.Builder, obj interface{}) {
	sarulabsdi.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeIGreeter2Service)
}

// AddSingletonIGreeter2Service adds a type that implements IGreeter2Service
func AddSingletonIGreeter2Service(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeIGreeter2Service)
}

// AddSingletonIGreeter2ServiceByFunc adds a type by a custom func
func AddSingletonIGreeter2ServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeter2Service)
}

// AddTransientIGreeter2Service adds a type that implements IGreeter2Service
func AddTransientIGreeter2Service(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddTransientWithImplementedTypes(builder, implType, ReflectTypeIGreeter2Service)
}

// AddTransientIGreeter2ServiceByFunc adds a type by a custom func
func AddTransientIGreeter2ServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeter2Service)
}

// AddScopedIGreeter2Server adds a type that implements IGreeter2Server
func AddScopedIGreeter2Server(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIGreeter2Server)
}

// AddScopedIGreeter2Service adds a type that implements IGreeter2Service
func AddScopedIGreeter2Service(builder *sarulabsdi.Builder, implType reflect.Type) {
	sarulabsdi.AddScopedWithImplementedTypes(builder, implType, ReflectTypeIGreeter2Service)
}

// AddScopedIGreeter2ServiceByFunc adds a type by a custom func
func AddScopedIGreeter2ServiceByFunc(builder *sarulabsdi.Builder, implType reflect.Type, build func(ctn sarulabsdi.Container) (interface{}, error)) {
	sarulabsdi.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeIGreeter2Service)
}

// RemoveAllIGreeter2Service removes all IBillingService from the DI
func RemoveAllIGreeter2Service(builder *sarulabsdi.Builder) {
	builder.RemoveAllByType(ReflectTypeIGreeter2Service)
}

// GetGreeter2ServiceFromContainer fetches the downstream di.Request scoped service
func GetGreeter2ServiceFromContainer(ctn sarulabsdi.Container) IGreeter2Service {
	return ctn.GetByType(ReflectTypeIGreeter2Service).(IGreeter2Service)
}

// GetGreeter2ServerFromContainer fetches the downstream di.Request scoped service
func GetGreeter2ServerFromContainer(ctn sarulabsdi.Container) IGreeter2Server {
	return ctn.GetByType(ReflectTypeIGreeter2Server).(IGreeter2Server)
}

// GetIGreeter2ServiceFromContainer fetches the downstream di.Request scoped service
func GetIGreeter2ServiceFromContainer(ctn sarulabsdi.Container) IGreeter2Service {
	return ctn.GetByType(ReflectTypeIGreeter2Service).(IGreeter2Service)
}

// GetIGreeter2ServerFromContainer fetches the downstream di.Request scoped service
func GetIGreeter2ServerFromContainer(ctn sarulabsdi.Container) IGreeter2Server {
	return ctn.GetByType(ReflectTypeIGreeter2Server).(IGreeter2Server)
}

// SafeGetIGreeter2ServiceFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeter2ServiceFromContainer(ctn sarulabsdi.Container) (IGreeter2Service, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeter2Service)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeter2Service), nil
}

// SafeGetIGreeter2ServerFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeter2ServerFromContainer(ctn sarulabsdi.Container) (IGreeter2Server, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeter2Server)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeter2Server), nil
}

// Impl for Greeter2 server instances
type greeter2Server struct {
	UnimplementedGreeter2ServerEx
}

// Impl for Greeter2 server instances
type greeter2ServerV2 struct {
	UnimplementedGreeter2ServerEx
}

// RegisterGreeter2ServerDI ...
func RegisterGreeter2ServerDI(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeter2Server{}
	RegisterGreeter2Server(s, server)
	return server
}

// RegisterGreeter2ServerDIV2 ...
func RegisterGreeter2ServerDIV2(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeter2ServerV2{}
	RegisterGreeter2Server(s, server)
	return server
}

// SayHello...
func (s *greeter2Server) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeter2ServiceFromContainer(requestContainer)
	return downstreamService.SayHello(request)
}

// SayHello...
func (s *greeter2ServerV2) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeter2ServerFromContainer(requestContainer)
	return downstreamService.SayHello(ctx, request)
}

// FullMethodNames for Greeter2
const (
	// FMN_Greeter2_SayHello
	FMN_Greeter2_SayHello = "/helloworld.Greeter2/SayHello"
)

// New_helloworldFullMethodNameSlice create a new map of fullMethodNames to []string
// i.e. /helloworld.Greeter/SayHello
func New_helloworldFullMethodNameSlice() []string {
	slice := []string{
		"/helloworld.Greeter/SayHello",
		"/helloworld.Greeter2/SayHello",
	}
	return slice
}
func init() {
	r := New_helloworldFullMethodNameSlice()
	pkg1.AddFullMethodNameSliceToMap(r)
}

// helloworldFullMethodNameEmptyResponseMap keys match that of grpc.UnaryServerInfo.FullMethodName
// i.e. /helloworld.Greeter/SayHello
var helloworldFullMethodNameEmptyResponseMap = map[string]func() interface{}{
	"/helloworld.Greeter/SayHello": func() interface{} {
		ret := &HelloReply{}
		return ret
	},
	"/helloworld.Greeter2/SayHello": func() interface{} {
		ret := &HelloReply2{}
		return ret
	},
}

// Get_helloworldFullEmptyResponseFromFullMethodName ...
func Get_helloworldFullEmptyResponseFromFullMethodName(fullMethodName string) func() interface{} {
	v, ok := helloworldFullMethodNameEmptyResponseMap[fullMethodName]
	if ok {
		return v
	}
	return nil
}

// helloworldFullMethodNameWithErrorResponseMap keys match that of grpc.UnaryServerInfo.FullMethodName
// i.e. /helloworld.Greeter/SayHello
var helloworldFullMethodNameWithErrorResponseMap = map[string]func() interface{}{
	"/helloworld.Greeter/SayHello": func() interface{} {
		ret := &HelloReply{}
		setNewField_MYU74zn4kc2UFcLfMVw8lrRTTSk87VSC(ret, "Error")
		return ret
	},
	"/helloworld.Greeter2/SayHello": func() interface{} {
		ret := &HelloReply2{}
		setNewField_MYU74zn4kc2UFcLfMVw8lrRTTSk87VSC(ret, "Error")
		return ret
	},
}

// Get_helloworldFullEmptyResponseWithErrorFromFullMethodName ...
func Get_helloworldFullEmptyResponseWithErrorFromFullMethodName(fullMethodName string) func() interface{} {
	v, ok := helloworldFullMethodNameWithErrorResponseMap[fullMethodName]
	if ok {
		return v
	}
	return nil
}

// M_helloworld_GreeterFullMethodNameExecuteMap keys match that of grpc.UnaryServerInfo.FullMethodName
var M_helloworld_GreeterFullMethodNameExecuteMap = map[string]func(service IGreeterServer, ctx context.Context, request interface{}) (interface{}, error){
	"/helloworld.Greeter/SayHello": func(service IGreeterServer, ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*HelloRequest)
		return service.SayHello(ctx, req)
	},
}

// M_helloworld_Greeter2FullMethodNameExecuteMap keys match that of grpc.UnaryServerInfo.FullMethodName
var M_helloworld_Greeter2FullMethodNameExecuteMap = map[string]func(service IGreeter2Server, ctx context.Context, request interface{}) (interface{}, error){
	"/helloworld.Greeter2/SayHello": func(service IGreeter2Server, ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*HelloRequest)
		return service.SayHello(ctx, req)
	},
}
