package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	reflectPackage        = protogen.GoImportPath("reflect")
	contextPackage        = protogen.GoImportPath("context")
	errorsPackage         = protogen.GoImportPath("errors")
	grpcPackage           = protogen.GoImportPath("google.golang.org/grpc")
	protoreflectPackage   = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	diPackage             = protogen.GoImportPath("github.com/fluffy-bunny/sarulabsdi")
	grpcDIInternalPackage = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg")
	grpcDIProtoError      = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error")
	diContextPackage      = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext")
	protocGenGoDiPackage  = protogen.GoImportPath("github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/pkg")
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
		filename:    file.GeneratedFilenamePrefix + "_di.pb.go",
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
	proto := file.Proto
	g := s.g
	g.P("/*  file.Proto")
	g.P(prettyJSON(proto))
	g.P("*/")
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("const _ = ", grpcDIInternalPackage.Ident("SupportPackageIsVersion7"))
	g.P()

	g.P("func setNewField_", s.uniqueRunID, "(dst interface{}, field string) {")
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

	interfaceServerName := fmt.Sprintf("I%vServer", service.GoName)
	g.P("// ", interfaceServerName, " defines the grpc server")
	g.P("type ", interfaceServerName, " interface {")
	for _, method := range service.Methods {
		methodGenCtx := newMethodGenContext(s.uniqueRunID, method, gen, file, g, service)
		g.P(methodGenCtx.serverSignature())
	}
	g.P("}")
	g.P()

	interfaceDownstreamServiceName := fmt.Sprintf("I%vService", service.GoName)
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

	typeServerInterfaceName := fmt.Sprintf("Type%s", interfaceServerName)
	typeDownstreamServiceInterfaceName := fmt.Sprintf("Type%s", interfaceDownstreamServiceName)
	// user reflection once to record the type
	g.P("// ", typeServerInterfaceName, " reflect type")
	g.P("var ", typeServerInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	g.P("// ", typeDownstreamServiceInterfaceName, " reflect type")
	g.P("var ", typeDownstreamServiceInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	// making type look like sarulabsdi genny types
	typeServerInterfaceName = fmt.Sprintf("ReflectType%v", interfaceServerName)
	g.P("// ", typeServerInterfaceName, " reflect type")
	g.P("var ", typeServerInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceServerName, ")(nil))")

	typeDownstreamServiceInterfaceName = fmt.Sprintf("ReflectType%v", interfaceDownstreamServiceName)
	g.P("// ", typeDownstreamServiceInterfaceName, " reflect type")
	g.P("var ", typeDownstreamServiceInterfaceName, " = ", diPackage.Ident("GetInterfaceReflectType"), "((*", interfaceDownstreamServiceName, ")(nil))")

	// DI Helpers

	// making type look like sarulabsdi genny types

	g.P("// AddSingleton", interfaceServerName, "ByObj adds a prebuilt obj")
	g.P("func AddSingleton", interfaceServerName, "ByObj(builder *", diPackage.Ident("Builder"), ", obj interface{})", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypesByObj"), "(builder,obj,", typeServerInterfaceName, ",)")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceServerName, " adds a type that implements ", interfaceServerName)
	g.P("func AddSingleton", interfaceServerName, "(builder *", diPackage.Ident("Builder"), ",implType ", reflectPackage.Ident("Type"), ")", " {")
	g.P(diPackage.Ident("AddSingletonWithImplementedTypes"), "(builder,implType,", typeServerInterfaceName, ")")
	g.P("}")
	g.P()

	g.P("// AddSingleton", interfaceServerName, "ByFunc adds a type by a custom func")
	g.P("func AddSingleton", interfaceServerName, "ByFunc(builder *", diPackage.Ident("Builder"), ", implType ", reflectPackage.Ident("Type"), ", build func(ctn ", diPackage.Ident("Container"), ") (interface{}, error)) {")
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
	g.P("Unimplemented", service.GoName, "Server")
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
				setNewField_%v(ret, "Error")
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
