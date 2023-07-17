package main

import (
	"flag"
	"fmt"

	_ "github.com/fluffy-bunny/grpcdotnetgo/pkg/proto/error"
	_ "github.com/fluffy-bunny/sarulabsdi"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var version = "development"
var (
	grpcGatewayEnabled *bool
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-di %v\n", version)
		return
	}

	// Additional options
	var flags flag.FlagSet
	grpcGatewayEnabled = flags.Bool("grpc_gateway", false, "enable grpc-gateway")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}
