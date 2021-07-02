module github.com/fluffy-bunny/grpcdotnetgo

go 1.16

require (
	github.com/fluffy-bunny/sarulabsdi v0.1.9
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/rs/xid v1.3.0
	github.com/rs/zerolog v1.23.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect

)

//replace github.com/fluffy-bunny/sarulabsdi => ../sarulabsdi
