// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	_ "github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	middleware_grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/auth"
	grpcDIProtoError "github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error"
	runtime "github.com/fluffy-bunny/grpcdotnetgo/pkg/runtime"
	pkg "github.com/fluffy-bunny/protoc-gen-go-di/pkg"
	"github.com/gogo/protobuf/gogoproto"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var version = "development"

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

	runtime.Start()

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
