// Code generated by protoc-gen-go-di. DO NOT EDIT.

package helloworld

import (
	context "context"
	pkg "github.com/fluffy-bunny/grpcdotnetgo/pkg"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	pkg1 "github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/pkg"
	sarulabsdi "github.com/fluffy-bunny/sarulabsdi"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

/*  file.Proto
{
    "name": "protoc-gen-go-di/helloworld/helloworld.proto",
    "package": "helloworld",
    "message_type": [
        {
            "name": "HelloRequest",
            "field": [
                {
                    "name": "name",
                    "number": 1,
                    "label": 1,
                    "type": 9,
                    "json_name": "name"
                }
            ]
        },
        {
            "name": "HelloReply",
            "field": [
                {
                    "name": "message",
                    "number": 1,
                    "label": 1,
                    "type": 9,
                    "json_name": "message"
                }
            ]
        },
        {
            "name": "HelloReply2",
            "field": [
                {
                    "name": "message",
                    "number": 1,
                    "label": 1,
                    "type": 9,
                    "json_name": "message"
                }
            ]
        },
        {
            "name": "WritePropsRequest"
        },
        {
            "name": "WritePropsResponse"
        }
    ],
    "service": [
        {
            "name": "Greeter",
            "method": [
                {
                    "name": "SayHello",
                    "input_type": ".helloworld.HelloRequest",
                    "output_type": ".helloworld.HelloReply",
                    "options": {}
                }
            ]
        },
        {
            "name": "Greeter2",
            "method": [
                {
                    "name": "SayHello",
                    "input_type": ".helloworld.HelloRequest",
                    "output_type": ".helloworld.HelloReply2",
                    "options": {}
                }
            ]
        },
        {
            "name": "EdgeControlService",
            "method": [
                {
                    "name": "WriteProps",
                    "input_type": ".helloworld.WritePropsRequest",
                    "output_type": ".helloworld.WritePropsResponse",
                    "server_streaming": true
                }
            ]
        }
    ],
    "options": {
        "java_package": "io.grpc.examples.helloworld",
        "java_outer_classname": "HelloWorldProto",
        "java_multiple_files": true,
        "go_package": "google.golang.org/grpc/examples/helloworld/helloworld"
    },
    "source_code_info": {
        "location": [
            {
                "span": [
                    14,
                    0,
                    60,
                    1
                ]
            },
            {
                "path": [
                    12
                ],
                "span": [
                    14,
                    0,
                    18
                ],
                "leading_detached_comments": [
                    " Copyright 2015 gRPC authors.\n\n Licensed under the Apache License, Version 2.0 (the \"License\");\n you may not use this file except in compliance with the License.\n You may obtain a copy of the License at\n\n     http://www.apache.org/licenses/LICENSE-2.0\n\n Unless required by applicable law or agreed to in writing, software\n distributed under the License is distributed on an \"AS IS\" BASIS,\n WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n See the License for the specific language governing permissions and\n limitations under the License.\n"
                ]
            },
            {
                "path": [
                    8
                ],
                "span": [
                    17,
                    0,
                    76
                ]
            },
            {
                "path": [
                    8,
                    11
                ],
                "span": [
                    17,
                    0,
                    76
                ]
            },
            {
                "path": [
                    8
                ],
                "span": [
                    18,
                    0,
                    34
                ]
            },
            {
                "path": [
                    8,
                    10
                ],
                "span": [
                    18,
                    0,
                    34
                ]
            },
            {
                "path": [
                    8
                ],
                "span": [
                    19,
                    0,
                    52
                ]
            },
            {
                "path": [
                    8,
                    1
                ],
                "span": [
                    19,
                    0,
                    52
                ]
            },
            {
                "path": [
                    8
                ],
                "span": [
                    20,
                    0,
                    48
                ]
            },
            {
                "path": [
                    8,
                    8
                ],
                "span": [
                    20,
                    0,
                    48
                ]
            },
            {
                "path": [
                    2
                ],
                "span": [
                    22,
                    0,
                    19
                ]
            },
            {
                "path": [
                    6,
                    0
                ],
                "span": [
                    25,
                    0,
                    28,
                    1
                ],
                "leading_comments": " The greeting service definition.\n"
            },
            {
                "path": [
                    6,
                    0,
                    1
                ],
                "span": [
                    25,
                    8,
                    15
                ]
            },
            {
                "path": [
                    6,
                    0,
                    2,
                    0
                ],
                "span": [
                    27,
                    2,
                    53
                ],
                "leading_comments": " Sends a greeting\n"
            },
            {
                "path": [
                    6,
                    0,
                    2,
                    0,
                    1
                ],
                "span": [
                    27,
                    6,
                    14
                ]
            },
            {
                "path": [
                    6,
                    0,
                    2,
                    0,
                    2
                ],
                "span": [
                    27,
                    16,
                    28
                ]
            },
            {
                "path": [
                    6,
                    0,
                    2,
                    0,
                    3
                ],
                "span": [
                    27,
                    39,
                    49
                ]
            },
            {
                "path": [
                    6,
                    1
                ],
                "span": [
                    29,
                    0,
                    32,
                    1
                ]
            },
            {
                "path": [
                    6,
                    1,
                    1
                ],
                "span": [
                    29,
                    8,
                    16
                ]
            },
            {
                "path": [
                    6,
                    1,
                    2,
                    0
                ],
                "span": [
                    31,
                    2,
                    54
                ],
                "leading_comments": " Sends a greeting\n"
            },
            {
                "path": [
                    6,
                    1,
                    2,
                    0,
                    1
                ],
                "span": [
                    31,
                    6,
                    14
                ]
            },
            {
                "path": [
                    6,
                    1,
                    2,
                    0,
                    2
                ],
                "span": [
                    31,
                    16,
                    28
                ]
            },
            {
                "path": [
                    6,
                    1,
                    2,
                    0,
                    3
                ],
                "span": [
                    31,
                    39,
                    50
                ]
            },
            {
                "path": [
                    4,
                    0
                ],
                "span": [
                    34,
                    0,
                    36,
                    1
                ],
                "leading_comments": " The request message containing the user's name.\n"
            },
            {
                "path": [
                    4,
                    0,
                    1
                ],
                "span": [
                    34,
                    8,
                    20
                ]
            },
            {
                "path": [
                    4,
                    0,
                    2,
                    0
                ],
                "span": [
                    35,
                    2,
                    18
                ]
            },
            {
                "path": [
                    4,
                    0,
                    2,
                    0,
                    5
                ],
                "span": [
                    35,
                    2,
                    8
                ]
            },
            {
                "path": [
                    4,
                    0,
                    2,
                    0,
                    1
                ],
                "span": [
                    35,
                    9,
                    13
                ]
            },
            {
                "path": [
                    4,
                    0,
                    2,
                    0,
                    3
                ],
                "span": [
                    35,
                    16,
                    17
                ]
            },
            {
                "path": [
                    4,
                    1
                ],
                "span": [
                    39,
                    0,
                    41,
                    1
                ],
                "leading_comments": " The response message containing the greetings\n"
            },
            {
                "path": [
                    4,
                    1,
                    1
                ],
                "span": [
                    39,
                    8,
                    18
                ]
            },
            {
                "path": [
                    4,
                    1,
                    2,
                    0
                ],
                "span": [
                    40,
                    2,
                    21
                ]
            },
            {
                "path": [
                    4,
                    1,
                    2,
                    0,
                    5
                ],
                "span": [
                    40,
                    2,
                    8
                ]
            },
            {
                "path": [
                    4,
                    1,
                    2,
                    0,
                    1
                ],
                "span": [
                    40,
                    9,
                    16
                ]
            },
            {
                "path": [
                    4,
                    1,
                    2,
                    0,
                    3
                ],
                "span": [
                    40,
                    19,
                    20
                ]
            },
            {
                "path": [
                    4,
                    2
                ],
                "span": [
                    42,
                    0,
                    45,
                    1
                ]
            },
            {
                "path": [
                    4,
                    2,
                    1
                ],
                "span": [
                    42,
                    8,
                    19
                ]
            },
            {
                "path": [
                    4,
                    2,
                    2,
                    0
                ],
                "span": [
                    43,
                    2,
                    21
                ]
            },
            {
                "path": [
                    4,
                    2,
                    2,
                    0,
                    5
                ],
                "span": [
                    43,
                    2,
                    8
                ]
            },
            {
                "path": [
                    4,
                    2,
                    2,
                    0,
                    1
                ],
                "span": [
                    43,
                    9,
                    16
                ]
            },
            {
                "path": [
                    4,
                    2,
                    2,
                    0,
                    3
                ],
                "span": [
                    43,
                    19,
                    20
                ]
            },
            {
                "path": [
                    6,
                    2
                ],
                "span": [
                    49,
                    0,
                    54,
                    1
                ],
                "leading_comments": " Mapped client connector auth service.\n di:ignore\n"
            },
            {
                "path": [
                    6,
                    2,
                    1
                ],
                "span": [
                    49,
                    8,
                    26
                ]
            },
            {
                "path": [
                    6,
                    2,
                    2,
                    0
                ],
                "span": [
                    53,
                    2,
                    72
                ],
                "leading_comments": " Write a property to one or more known subtended device(s)\n Request: Unary\n Response: Streaming\n"
            },
            {
                "path": [
                    6,
                    2,
                    2,
                    0,
                    1
                ],
                "span": [
                    53,
                    6,
                    16
                ]
            },
            {
                "path": [
                    6,
                    2,
                    2,
                    0,
                    2
                ],
                "span": [
                    53,
                    17,
                    34
                ]
            },
            {
                "path": [
                    6,
                    2,
                    2,
                    0,
                    6
                ],
                "span": [
                    53,
                    45,
                    51
                ]
            },
            {
                "path": [
                    6,
                    2,
                    2,
                    0,
                    3
                ],
                "span": [
                    53,
                    52,
                    70
                ]
            },
            {
                "path": [
                    4,
                    3
                ],
                "span": [
                    56,
                    0,
                    57,
                    1
                ]
            },
            {
                "path": [
                    4,
                    3,
                    1
                ],
                "span": [
                    56,
                    8,
                    25
                ]
            },
            {
                "path": [
                    4,
                    4
                ],
                "span": [
                    59,
                    0,
                    60,
                    1
                ]
            },
            {
                "path": [
                    4,
                    4,
                    1
                ],
                "span": [
                    59,
                    8,
                    26
                ]
            }
        ]
    },
    "syntax": "proto3"
}
*/
// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = pkg.SupportPackageIsVersion7

func setNewField_BpLnfgDsc2WD8F2qNfHK5a84jjJkwzDk(dst interface{}, field string) {
	v := reflect.ValueOf(dst).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.New(v.Type().Elem()))
	}
}

// IGreeterServer defines the grpc server
type IGreeterServer interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
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

// GetIGreeterServiceFromContainer fetches the downstream di.Request scoped service
func GetIGreeterServiceFromContainer(ctn sarulabsdi.Container) IGreeterService {
	return ctn.GetByType(ReflectTypeIGreeterService).(IGreeterService)
}

// SafeGetIGreeterServiceFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeterServiceFromContainer(ctn sarulabsdi.Container) (IGreeterService, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeterService)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeterService), nil
}

// Impl for Greeter server instances
type greeterServer struct {
	UnimplementedGreeterServer
}

// RegisterGreeterServerDI ...
func RegisterGreeterServerDI(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeterServer{}
	RegisterGreeterServer(s, server)
	return server
}

// SayHello...
func (s *greeterServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeterServiceFromContainer(requestContainer)
	return downstreamService.SayHello(request)
}

// FullMethodNames for Greeter
const (
	// FMN_Greeter_SayHello
	FMN_Greeter_SayHello = "/helloworld.Greeter/SayHello"
)

// IGreeter2Server defines the grpc server
type IGreeter2Server interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error)
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

// GetIGreeter2ServiceFromContainer fetches the downstream di.Request scoped service
func GetIGreeter2ServiceFromContainer(ctn sarulabsdi.Container) IGreeter2Service {
	return ctn.GetByType(ReflectTypeIGreeter2Service).(IGreeter2Service)
}

// SafeGetIGreeter2ServiceFromContainer fetches the downstream di.Request scoped service
func SafeGetIGreeter2ServiceFromContainer(ctn sarulabsdi.Container) (IGreeter2Service, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGreeter2Service)
	if err != nil {
		return nil, err
	}
	return obj.(IGreeter2Service), nil
}

// Impl for Greeter2 server instances
type greeter2Server struct {
	UnimplementedGreeter2Server
}

// RegisterGreeter2ServerDI ...
func RegisterGreeter2ServerDI(s grpc.ServiceRegistrar) interface{} {
	// Register the server
	var server = &greeter2Server{}
	RegisterGreeter2Server(s, server)
	return server
}

// SayHello...
func (s *greeter2Server) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply2, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := GetGreeter2ServiceFromContainer(requestContainer)
	return downstreamService.SayHello(request)
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
		setNewField_BpLnfgDsc2WD8F2qNfHK5a84jjJkwzDk(ret, "Error")
		return ret
	},
	"/helloworld.Greeter2/SayHello": func() interface{} {
		ret := &HelloReply2{}
		setNewField_BpLnfgDsc2WD8F2qNfHK5a84jjJkwzDk(ret, "Error")
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
