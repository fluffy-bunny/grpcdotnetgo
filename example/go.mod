module github.com/fluffy-bunny/grpcdotnetgo/example

go 1.16

require (
	github.com/fluffy-bunny/grpcdotnetgo v0.1.26
	github.com/fluffy-bunny/protoc-gen-go-di v0.0.21
	github.com/fluffy-bunny/sarulabsdi v0.1.12
	github.com/fluffy-bunny/viperEx v0.0.12
	github.com/gogo/googleapis v1.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/gogo/status v1.1.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jnewmano/grpc-json-proxy v0.0.3
	github.com/rs/zerolog v1.23.0
	github.com/spf13/viper v1.8.1
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/fluffy-bunny/grpcdotnetgo => ../

//replace github.com/fluffy-bunny/sarulabsdi => ../../sarulabsdi
