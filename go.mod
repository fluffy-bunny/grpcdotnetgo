module github.com/fluffy-bunny/grpcdotnetgo

go 1.16

require (
	github.com/fatih/structs v1.1.0
	github.com/fluffy-bunny/sarulabsdi v0.1.9
	github.com/fluffy-bunny/viperEx v0.0.12
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/rs/xid v1.3.0
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1 // indirect

)

//replace github.com/fluffy-bunny/sarulabsdi => ../sarulabsdi
