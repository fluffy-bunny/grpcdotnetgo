package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	reflectPackage = protogen.GoImportPath("reflect")
	contextPackage = protogen.GoImportPath("context")
	errorsPackage  = protogen.GoImportPath("errors")
	codesPackage   = protogen.GoImportPath("google.golang.org/grpc/codes")
	statusPackage  = protogen.GoImportPath("google.golang.org/grpc/status")

	grpcPackage                          = protogen.GoImportPath("google.golang.org/grpc")
	grpcStatusPackage                    = protogen.GoImportPath("google.golang.org/grpc/status")
	grpcCodesPackage                     = protogen.GoImportPath("google.golang.org/grpc/codes")
	protoreflectPackage                  = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	diPackage                            = protogen.GoImportPath("github.com/fluffy-bunny/sarulabsdi")
	grpcDIInternalPackage                = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg")
	grpcDIInternalPackageContractsGRPC   = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/grpc")
	grpcDIInternalRuntimePackage         = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/runtime")
	grpcDIProtoError                     = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error")
	diContextPackage                     = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext")
	protocGenGoDiPackage                 = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/pkg")
	protocGenTemporalPackage             = protogen.GoImportPath("go.temporal.io/sdk/temporal")
	protocGenGRPCDNGTemporalUtilsPackage = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/temporal/utils")
)

type genFileContext struct {
	packageName string
	uniqueRunID string
	gen         *protogen.Plugin
	file        *protogen.File
	filename    string
	g           *protogen.GeneratedFile
}

func isServiceIgnored(service *protogen.Service) bool {
	// Look for a comment consisting of "di:ignore"
	const ignore = "di:ignore"
	for _, comment := range service.Comments.LeadingDetached {
		if strings.Contains(string(comment), ignore) {
			return true
		}
	}

	return strings.Contains(string(service.Comments.Leading), ignore)
}
func newGenFileContext(gen *protogen.Plugin, file *protogen.File) *genFileContext {
	ctx := &genFileContext{
		file:        file,
		gen:         gen,
		uniqueRunID: randomString(32),
		packageName: string(file.GoPackageName),
		filename:    file.GeneratedFilenamePrefix + "_di_temporal.pb.go",
	}
	ctx.g = gen.NewGeneratedFile(ctx.filename, file.GoImportPath)
	return ctx
}

// MethodInfo type
type MethodInfo struct {
	NewResponseWithErrorFunc string
	NewResponseFunc          string
	ExecuteFunc              string
}
type methodGenContext struct {
	MethodInfo     *MethodInfo
	ProtogenMethod *protogen.Method
	gen            *protogen.Plugin
	file           *protogen.File
	g              *protogen.GeneratedFile
	service        *protogen.Service
	uniqueRunID    string
}
type serviceGenContext struct {
	packageName     string
	MethodMapGenCtx map[string]*methodGenContext
	gen             *protogen.Plugin
	file            *protogen.File
	g               *protogen.GeneratedFile
	service         *protogen.Service
	uniqueRunID     string
}

func newServiceGenContext(packageName string, uniqueRunId string, gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) *serviceGenContext {
	ctx := &serviceGenContext{
		packageName:     packageName,
		uniqueRunID:     uniqueRunId,
		gen:             gen,
		file:            file,
		g:               g,
		service:         service,
		MethodMapGenCtx: make(map[string]*methodGenContext),
	}
	return ctx
}
func newMethodGenContext(uniqueRunId string, protogenMethod *protogen.Method, gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) *methodGenContext {
	ctx := &methodGenContext{
		uniqueRunID:    uniqueRunId,
		MethodInfo:     &MethodInfo{},
		ProtogenMethod: protogenMethod,
		gen:            gen,
		file:           file,
		g:              g,
		service:        service,
	}
	return ctx
}

// generateFile generates a _gtm.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	ctx := newGenFileContext(gen, file)
	g := ctx.g

	// Default to skip - will unskip if there is a service to generate
	g.Skip()

	g.P("// Code generated by protoc-gen-go-di. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	ctx.generateFileContent()
	return g
}
func getFileName(file *protogen.File) string {
	v := strings.Split(file.GeneratedFilenamePrefix, "/")
	idx := len(v) - 1
	return v[idx]
}

// generateFileContent generates the DI service definitions, excluding the package statement.
func (s *genFileContext) generateFileContent() {
	gen := s.gen
	file := s.file

	g := s.g

	//proto := file.Proto
	//g.P("/*  file.Proto")
	//g.P(prettyJSON(proto))
	//g.P("*/")
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("const _ = ", grpcDIInternalPackage.Ident("SupportPackageIsVersion7"))
	g.P()

	g.P("func setTemporalNewField_", s.uniqueRunID, "(dst interface{}, field string) {")
	g.P("\tv := ", reflectPackage.Ident("ValueOf"), "(dst).Elem().FieldByName(field)")
	g.P("\tif v.IsValid() {")
	g.P("\t\tv.Set(", reflectPackage.Ident("New"), "(v.Type().Elem()))")
	g.P("\t}")
	g.P("}")

	var serviceGenCtxs []*serviceGenContext
	// Generate each service
	for _, service := range file.Services {
		// Check if this service is ignored for DI purposes
		if isServiceIgnored(service) {
			continue
		}

		// Now we have something to generate
		g.Unskip()

		serviceGenCtx := newServiceGenContext(s.packageName, s.uniqueRunID, gen, file, g, service)
		serviceGenCtx.genService()
		serviceGenCtxs = append(serviceGenCtxs, serviceGenCtx)
	}

	g.P("// New_", getFileName(file), "FullMethodNameSlice create a new map of fullMethodNames to []string")
	g.P("// i.e. /helloworld.Greeter/SayHello ")
	g.P("func New_", getFileName(file), "FullMethodNameSlice() []string {")
	g.P("    slice := []string {")
	for _, sctx := range serviceGenCtxs {
		for k := range sctx.MethodMapGenCtx {
			g.P("        \"", k, "\",")
		}
	}
	g.P("    }")
	g.P("    return slice")
	g.P("}")

	g.P("func init() {")
	g.P("  r := New_", getFileName(file), "FullMethodNameSlice()")
	g.P("  ", protocGenGoDiPackage.Ident("AddFullMethodNameSliceToMap(r)"))
	g.P("}")

	g.P("// ", getFileName(file), "FullMethodNameEmptyResponseMap keys match that of grpc.UnaryServerInfo.FullMethodName")
	g.P("// i.e. /helloworld.Greeter/SayHello ")
	g.P("var ", getFileName(file), "FullMethodNameEmptyResponseMap = map[string]func()interface{} {")
	for _, sctx := range serviceGenCtxs {
		for k, v := range sctx.MethodMapGenCtx {
			g.P("    \"", k, "\": ", v.MethodInfo.NewResponseFunc, ",")
		}
	}
	g.P("}")

	g.P("// Get_", getFileName(file), "FullEmptyResponseFromFullMethodName ...")
	g.P("func Get_", getFileName(file), "FullEmptyResponseFromFullMethodName(fullMethodName string) func() interface{} {")
	g.P("  v,ok := ", getFileName(file), "FullMethodNameEmptyResponseMap[fullMethodName]")
	g.P("  if ok {")
	g.P("    return v")
	g.P("  }")
	g.P("  return nil")
	g.P("}")

	g.P("// ", getFileName(file), "FullMethodNameWithErrorResponseMap keys match that of grpc.UnaryServerInfo.FullMethodName")
	g.P("// i.e. /helloworld.Greeter/SayHello ")
	g.P("var ", getFileName(file), "FullMethodNameWithErrorResponseMap = map[string]func()interface{} {")
	for _, sctx := range serviceGenCtxs {
		for k, v := range sctx.MethodMapGenCtx {
			g.P("    \"", k, "\": ", v.MethodInfo.NewResponseWithErrorFunc, ",")
		}
	}
	g.P("}")

	g.P("// Get_", getFileName(file), "FullEmptyResponseWithErrorFromFullMethodName ...")
	g.P("func Get_", getFileName(file), "FullEmptyResponseWithErrorFromFullMethodName(fullMethodName string) func() interface{} {")
	g.P("  v,ok := ", getFileName(file), "FullMethodNameWithErrorResponseMap[fullMethodName]")
	g.P("  if ok {")
	g.P("    return v")
	g.P("  }")
	g.P("  return nil")
	g.P("}")

	for _, sctx := range serviceGenCtxs {
		g.P("// M_", getFileName(file), "_", sctx.service.GoName, "FullMethodNameExecuteMap keys match that of grpc.UnaryServerInfo.FullMethodName")
		g.P("var M_", getFileName(file), "_", sctx.service.GoName,
			fmt.Sprintf("FullMethodNameExecuteMap = map[string]func(service I%vServer, ctx context.Context, request interface{}) (interface{}, error) {", sctx.service.GoName))
		for k, v := range sctx.MethodMapGenCtx {
			g.P("    \"", k, "\": ", v.MethodInfo.ExecuteFunc, ",")
		}

		g.P("}")
	}
}

func (s *serviceGenContext) genService() {
	gen := s.gen
	file := s.file
	proto := file.Proto
	g := s.g
	service := s.service

	// IServiceEndpointRegistration
	serviceEndpointRegistrationName := fmt.Sprintf("%vEndpointRegistration", service.GoName)
	interfaceServerActiviteisName := fmt.Sprintf("I%vActivities", service.GoName)
	mustEmbedUnimplementedName := fmt.Sprintf("mustEmbedUnimplemented%vServer", service.GoName)
	interfaceDownstreamServiceName := fmt.Sprintf("I%vService", service.GoName)

	// Define the activites interface
	//----------------------------------------------------------------------------------------------
	g.P("// ", interfaceServerActiviteisName, " defines the activites interface")
	g.P("type ", interfaceServerActiviteisName, " interface {")
	g.P("	", service.GoName, "Server")
	g.P("}")
	g.P()
	serviceStructName := fmt.Sprintf("service%vActivites", service.GoName)
	unimplementedExServerName := fmt.Sprintf("Unimplemented%vServerEx", service.GoName)

	// Define the ServiceEndpointRegistration implementation
	//----------------------------------------------------------------------------------------------
	g.P("// ", serviceStructName, " defines the activites struct")
	g.P("type ", serviceStructName, " struct {")
	g.P(unimplementedExServerName)
	g.P("GetClient ", "Get", service.GoName, "Client", " `inject:\"\"`")

	g.P("}")
	g.P()
	g.P("func (s *", serviceStructName, ") Ctor() {")
	g.P("s.UnimplemtedErrorResponse = func() error {")
	g.P("	return ", protocGenTemporalPackage.Ident("NewNonRetryableApplicationError"), "(\"method not implemented\",nil,nil)")
	g.P("	}")
	g.P("}")
	g.P()

	for _, method := range service.Methods {
		methodGenCtx := newMethodGenContext(s.uniqueRunID, method, gen, file, g, service)
		g.P("func (s *", serviceStructName, ") ", methodGenCtx.serverSignature(), "{")
		g.P("asyncActivity := func() (interface{}, error) {")
		g.P("	grpcClient, err := s.GetClient()")
		g.P("	if err != nil {")
		g.P("		return nil, err")
		g.P("	}")
		g.P("	result, err := grpcClient.", methodGenCtx.getMethodName(), "(context.Background(),request)")
		g.P("	s, ok := ", statusPackage.Ident("FromError"), "(err)")
		g.P("	if ok {")
		g.P("		if s.Code() == ", codesPackage.Ident("NotFound"), "{")
		g.P("			return nil,", protocGenTemporalPackage.Ident("NewNonRetryableApplicationError"), "(\"not found\",nil,nil)")
		g.P("		}")
		g.P("	}")
		g.P("	return result, err")
		g.P("}")
		g.P("res, err := ", protocGenGRPCDNGTemporalUtilsPackage.Ident("DoAsyncGrpcCallWithActivityHeartbeat"), "(context.Background(), asyncActivity)")
		g.P("if err != nil {")
		g.P("	return nil, err")
		g.P("}")
		getMethodResults := methodGenCtx.getMethodResults()
		g.P("return *", getMethodResults[0], ",nil")

		g.P("}")
		g.P()
	}

	// Add the DI Singleton registration
	//----------------------------------------------------------------------------------------------
	typeServiceStructName := fmt.Sprintf("Type%s", serviceStructName)
	g.P("// ", typeServiceStructName, " reflect type")
	g.P("var ", typeServiceStructName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", serviceStructName, ")(nil))")
	g.P()
	g.P("// Add", serviceStructName, " adds a type that implements IServiceEndpointRegistration")
	g.P("func AddSingleton", interfaceServerActiviteisName, "(builder *", diPackage.Ident("Builder"), ")", " {")
	g.P("   ", grpcDIInternalPackageContractsGRPC.Ident("AddSingletonIServiceEndpointRegistration"), "(builder,reflect.TypeOf(&", serviceStructName, "{}))")
	g.P("	AddScoped", interfaceDownstreamServiceName, "(builder,implType)")
	g.P("}")
	g.P()

	g.P("// ", interfaceServerActiviteisName, " defines the grpc server")
	g.P("type ", interfaceServerActiviteisName, " interface {")
	g.P("  	", mustEmbedUnimplementedName, "()")
	for _, method := range service.Methods {
		methodGenCtx := newMethodGenContext(s.uniqueRunID, method, gen, file, g, service)
		g.P(methodGenCtx.serverSignature())
	}
	g.P("}")
	g.P()

	g.P("// ", unimplementedExServerName, " defines the grpc server")
	g.P("type ", unimplementedExServerName, " struct {")
	g.P("  UnimplemtedErrorResponse func() error")
	g.P("}")
	g.P("func (", unimplementedExServerName, ") ",
		"mustEmbedUnimplemented", service.GoName, "Server(){}")

	for _, method := range service.Methods {
		methodGenCtx := newMethodGenContext(s.uniqueRunID, method, gen, file, g, service)
		g.P("func (u ", unimplementedExServerName, ") ", methodGenCtx.serverSignature(), "{")
		g.P("  if u.UnimplemtedErrorResponse != nil {")
		g.P("    return nil, u.UnimplemtedErrorResponse()")
		g.P("  }")
		g.P("  return nil, ", grpcStatusPackage.Ident("Error"), "(", grpcCodesPackage.Ident("Unimplemented"), ",\"method ", method.GoName, " not implemented\")")
		g.P("}")
	}

	g.P("// ", interfaceDownstreamServiceName, " defines the required downstream service interface")
	g.P("type ", interfaceDownstreamServiceName, " interface {")
	for _, method := range service.Methods {
		serverType := method.Parent.GoName
		key := "/" + *proto.Package + "." + serverType + "/" + method.GoName
		methodGenCtx := newMethodGenContext(s.uniqueRunID, method, gen, file, g, service)
		methodGenCtx.genDownstreamMethodSignature()
		s.MethodMapGenCtx[key] = methodGenCtx
	}
	g.P("}")
	g.P()

	typeServerInterfaceName := fmt.Sprintf("Type%s", interfaceServerActiviteisName)
	typeDownstreamServiceInterfaceName := fmt.Sprintf("Type%s", interfaceDownstreamServiceName)
	// user reflection once to record the type
	g.P("// ", typeServerInterfaceName, " reflect type")
	g.P("var ", typeServerInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	g.P("// ", typeDownstreamServiceInterfaceName, " reflect type")
	g.P("var ", typeDownstreamServiceInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	// making type look like sarulabsdi genny types
	typeServerInterfaceName = fmt.Sprintf("ReflectType%v", interfaceServerActiviteisName)
	g.P("// ", typeServerInterfaceName, " reflect type")
	g.P("var ", typeServerInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceServerActiviteisName, ")(nil))")

	typeDownstreamServiceInterfaceName = fmt.Sprintf("ReflectType%v", interfaceDownstreamServiceName)
	g.P("// ", typeDownstreamServiceInterfaceName, " reflect type")
	g.P("var ", typeDownstreamServiceInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	g.P("type Get", service.GoName, "Client func() ", service.GoName, "Client")

	// Client Creation
	g.P("func GetNew", service.GoName, "Client(cc ", grpcPackage.Ident("ClientConnInterface"), ") ", service.GoName, "Client {")
	g.P("return New", service.GoName, "Client(cc)")
	g.P("}")
	g.P()
	// DI Helpers

	// making type look like sarulabsdi genny types

	g.P("// AddSingleton", interfaceServerActiviteisName, "ByObj adds a prebuilt obj")
	g.P("func AddSingleton", interfaceServerActiviteisName, "ByObj(builder *", diPackage.Ident("Builder"), ", obj interface{})", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypesByObj"), "(builder,obj,", typeServerInterfaceName, ",)")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceServerActiviteisName, " adds a type that implements ", interfaceServerActiviteisName)
	g.P("func AddSingleton", interfaceServerActiviteisName, "(builder *", diPackage.Ident("Builder"), ",implType ", reflectPackage.Ident("Type"), ")", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypes"), "(builder,implType,", typeServerInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceServerActiviteisName, "ByFunc adds a type by a custom func")
	g.P("func AddSingleton", interfaceServerActiviteisName, "ByFunc(builder *", diPackage.Ident("Builder"), ", implType ", reflectPackage.Ident("Type"), ", build func(ctn ", diPackage.Ident("Container"), ") (interface{}, error)) {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypesByFunc"), "(builder, implType, build,", typeServerInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceDownstreamServiceName, "ByObj adds a prebuilt obj")
	g.P("func AddSingleton", interfaceDownstreamServiceName, "ByObj(builder *", diPackage.Ident("Builder"), ", obj interface{})", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypesByObj"), "(builder,obj,", typeDownstreamServiceInterfaceName, ",)")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceDownstreamServiceName, " adds a type that implements ", interfaceDownstreamServiceName)
	g.P("func AddSingleton", interfaceDownstreamServiceName, "(builder *", diPackage.Ident("Builder"), ",implType ", reflectPackage.Ident("Type"), ")", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypes"), "(builder,implType,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceDownstreamServiceName, "ByFunc adds a type by a custom func")
	g.P("func AddSingleton", interfaceDownstreamServiceName, "ByFunc(builder *", diPackage.Ident("Builder"), ", implType ", reflectPackage.Ident("Type"), ", build func(ctn ", diPackage.Ident("Container"), ") (interface{}, error)) {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypesByFunc"), "(builder, implType, build,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddTransient", interfaceDownstreamServiceName, " adds a type that implements ", interfaceDownstreamServiceName)
	g.P("func AddTransient", interfaceDownstreamServiceName, "(builder *", diPackage.Ident("Builder"), ",implType ", reflectPackage.Ident("Type"), ")", " {")
	g.P(diPackage.Ident("AddTransientWithImplementedTypes"), "(builder,implType,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddTransient", interfaceDownstreamServiceName, "ByFunc adds a type by a custom func")
	g.P("func AddTransient", interfaceDownstreamServiceName, "ByFunc(builder *", diPackage.Ident("Builder"), ", implType ", reflectPackage.Ident("Type"), ", build func(ctn ", diPackage.Ident("Container"), ") (interface{}, error)) {")
	g.P(diPackage.Ident("AddTransientWithImplementedTypesByFunc"), "(builder, implType, build,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddScoped", interfaceDownstreamServiceName, " adds a type that implements ", interfaceDownstreamServiceName)
	g.P("func AddScoped", interfaceDownstreamServiceName, "(builder *", diPackage.Ident("Builder"), ",implType ", reflectPackage.Ident("Type"), ")", " {")
	g.P(diPackage.Ident("AddScopedWithImplementedTypes"), "(builder,implType,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddScoped", interfaceDownstreamServiceName, "ByFunc adds a type by a custom func")
	g.P("func AddScoped", interfaceDownstreamServiceName, "ByFunc(builder *", diPackage.Ident("Builder"), ", implType ", reflectPackage.Ident("Type"), ", build func(ctn ", diPackage.Ident("Container"), ") (interface{}, error)) {")
	g.P(diPackage.Ident("AddScopedWithImplementedTypesByFunc"), "(builder, implType, build,", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// RemoveAll", interfaceDownstreamServiceName, " removes all IBillingService from the DI")
	g.P("func RemoveAll", interfaceDownstreamServiceName, "(builder *", diPackage.Ident("Builder"), ")  {")
	g.P("builder.RemoveAllByType(", typeDownstreamServiceInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// Get", service.GoName, "ServiceFromContainer fetches the downstream di.Request scoped service")
	g.P("func Get", service.GoName, "ServiceFromContainer(ctn ", diPackage.Ident("Container"), ") ", interfaceDownstreamServiceName, " {")
	g.P("return ctn.GetByType(", typeDownstreamServiceInterfaceName, ").(", interfaceDownstreamServiceName, ")")
	g.P("}")
	g.P()

	// making type look like sarulabsdi genny types
	g.P("// Get", interfaceDownstreamServiceName, "FromContainer fetches the downstream di.Request scoped service")
	g.P("func Get", interfaceDownstreamServiceName, "FromContainer(ctn ", diPackage.Ident("Container"), ") ", interfaceDownstreamServiceName, " {")
	g.P("return ctn.GetByType(", typeDownstreamServiceInterfaceName, ").(", interfaceDownstreamServiceName, ")")
	g.P("}")
	g.P()

	// making type look like sarulabsdi genny types
	g.P("// SafeGet", interfaceDownstreamServiceName, "FromContainer fetches the downstream di.Request scoped service")
	g.P("func SafeGet", interfaceDownstreamServiceName, "FromContainer(ctn ", diPackage.Ident("Container"), ") (", interfaceDownstreamServiceName, ",error) {")
	g.P("obj, err := ctn.SafeGetByType(", typeDownstreamServiceInterfaceName, ")")
	g.P("if err != nil {")
	g.P("    return nil, err")
	g.P("}")
	g.P("return obj.(", interfaceDownstreamServiceName, "),nil")
	g.P("}")
	g.P()

	// Instance Impl
	g.P("// Impl for ", service.GoName, " server instances")
	g.P("type ", strings.ToLower(service.GoName), "Server struct {")
	g.P(unimplementedExServerName)
	g.P("}")

	// Server Registration
	g.P("// Register", service.GoName, "ServerDI ...")
	g.P("func Register", service.GoName, "ServerDI(s ", grpcPackage.Ident("ServiceRegistrar"), ") interface{} {")
	g.P("// Register the server")
	g.P("var server = &", strings.ToLower(service.GoName), "Server{ }")
	g.P("Register", service.GoName, "Server(s, server)")
	g.P("return server")
	g.P("}")
	g.P()

	// Client method implementations.
	for _, method := range service.Methods {
		serverType := method.Parent.GoName
		key := "/" + *proto.Package + "." + serverType + "/" + method.GoName
		methodGenCtx := s.MethodMapGenCtx[key]
		methodGenCtx.genServerMethodShim()
	}
	g.P("// FullMethodNames for ", service.GoName)
	g.P("const (")
	for _, method := range service.Methods {
		serverType := method.Parent.GoName
		key := "/" + *proto.Package + "." + serverType + "/" + method.GoName
		g.P("// FMN_", serverType, "_", method.GoName)
		g.P("FMN_", serverType, "_", method.GoName, " = \"", key, "\"")
	}
	g.P(")")
}

func (s *methodGenContext) genDownstreamMethodSignature() {
	g := s.g
	if s.ProtogenMethod.Desc.IsStreamingClient() || s.ProtogenMethod.Desc.IsStreamingServer() {
		// Explicitly no current support for streaming methods
		panic("Does not currently support streaming methods")
	}
	// Unary method
	g.P(s.downstreamServiceSignature())
	g.P()
}

func (s *methodGenContext) genServerMethodShim() {
	service := s.service
	g := s.g

	method := s.ProtogenMethod

	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		// Explicitly no current support for streaming methods
		panic("Does not currently support streaming methods")
	}

	// Unary method
	serverType := method.Parent.GoName

	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		s.MethodInfo.NewResponseFunc = fmt.Sprintf(
			`func() interface{} {
				ret := &%v{}
				return ret 
			}`, g.QualifiedGoIdent(method.Output.GoIdent))

		s.MethodInfo.NewResponseWithErrorFunc = fmt.Sprintf(
			`func() interface{} {
				ret := &%v{}
				setNewFieldTemporal_%v(ret, "Error")
				return ret 
			}`, g.QualifiedGoIdent(method.Output.GoIdent), s.uniqueRunID)

		s.MethodInfo.ExecuteFunc = fmt.Sprintf(
			`func(service I%vServer, ctx context.Context, request interface{}) (interface{}, error) {
				req := request.(*%v)
				return service.%v(ctx, req)
			}`, service.GoName, g.QualifiedGoIdent(method.Input.GoIdent),
			method.GoName,
		)
	}
	/*
	   var dd = map[string]func(service *greeter2Server, ctx context.Context, request interface{}) (interface{}, error){
	   	"/helloworld.Greeter/SayHello": func(service *greeter2Server, ctx context.Context, request interface{}) (interface{}, error) {
	   		req := request.(*HelloRequest)
	   		return service.SayHello(ctx, req)
	   	},
	   }
	*/

	g.P("// ", s.ProtogenMethod.GoName, "...")
	g.P("func (s *", strings.ToLower(serverType), "Server) ", s.serverSignature(), "{")
	g.P("requestContainer := ", diContextPackage.Ident("GetRequestContainer(ctx)"))
	g.P("downstreamService := Get", service.GoName, "ServiceFromContainer(requestContainer)")
	g.P("return downstreamService.", method.GoName, "(request)")
	g.P("}")
	g.P()
}

func (s *methodGenContext) downstreamServiceSignature() string {
	g := s.g
	method := s.ProtogenMethod
	var reqArgs []string
	ret := "error"
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		ret = "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "request *"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, method.Parent.GoName+"_"+method.GoName+"Server")
	}
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}
func (s *methodGenContext) getMethodName() string {
	return s.ProtogenMethod.GoName
}
func (s *methodGenContext) getMethodResults() []string {
	results := []string{}
	g := s.g
	method := s.ProtogenMethod
	results = append(results, g.QualifiedGoIdent(method.Output.GoIdent))
	results = append(results, "error")
	return results
}

func (s *methodGenContext) serverSignature() string {
	g := s.g
	method := s.ProtogenMethod
	var reqArgs []string
	ret := "error"
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, "ctx "+g.QualifiedGoIdent(contextPackage.Ident("Context")))
		ret = "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "request *"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, method.Parent.GoName+"_"+method.GoName+"Server")
	}
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}
func prettyJSON(obj interface{}) string {
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
