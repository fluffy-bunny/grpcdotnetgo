module github.com/fluffy-bunny/grpcdotnetgo/example

go 1.16

require (
	github.com/fluffy-bunny/grpcdotnetgo v0.1.4
	github.com/fluffy-bunny/sarulabsdi v0.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jnewmano/grpc-json-proxy v0.0.3
	github.com/rs/zerolog v1.23.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/fluffy-bunny/grpcdotnetgo => ../

replace github.com/fluffy-bunny/sarulabsdi => ../../sarulabsdi
