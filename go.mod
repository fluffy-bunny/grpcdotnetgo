module github.com/fluffy-bunny/grpcdotnetgo

go 1.16

require (
	github.com/bamzi/jobrunner v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/fluffy-bunny/sarulabsdi v0.1.12
	github.com/fluffy-bunny/viperEx v0.0.12
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/reugn/async v0.0.0-20200819063434-15e5b3951cd7
	github.com/robfig/cron/v3 v3.0.0
	github.com/rs/xid v1.3.0
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect

)

//replace github.com/fluffy-bunny/sarulabsdi => ../sarulabsdi
