// Package main implements a server for Greeter service.
package main

import (
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	_ "github.com/fluffy-bunny/grpcdotnetgo/example/internal/plugin"
	runtime "github.com/fluffy-bunny/grpcdotnetgo/pkg/runtime"
	pkg "github.com/fluffy-bunny/protoc-gen-go-di/pkg"
	"github.com/gogo/protobuf/gogoproto"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
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

// Wondering where your grpc app actually is?
// we are a plugin framework
// internal.plugin.plugin.go  look for func init()
