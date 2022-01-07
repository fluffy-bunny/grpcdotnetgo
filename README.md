# gRPC-dot-net-go

## Installation

When used with Go modules, use the following import path:

    go get github.com/fluffy-bunny/grpcdotnetgo

## Stand-Alone Samples

[samples](https://github.com/fluffy-bunny/grpcdotnetgo-samples)  

## Contracts  
It all starts with a [proto file](example/internal/grpcContracts/helloworld/helloworld.proto).  In our case we want to have a custom server implementation that understands that the real downstream hanlder is in our DI.  Fortunately we can use a protoc plugin to generate the GO code for that.  

```powershell
cd example
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
go get -u github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --proto_path=. --proto_path=vendor --proto_path=vendor/github.com/fluffy-bunny  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative ./example/internal/grpcContracts/helloworld/helloworld.proto

```
This will generate all the [go files](example/internal/grpcContracts/helloworld) from our ```helloworld.proto```  

### Generated grpc interfaces  

#### grpc interface
```go
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}
```

#### grpcdotnetgo inteface
```go
// IGreeterService defines the required downstream service interface
type IGreeterService interface {
	SayHello(request *HelloRequest) (*HelloReply, error)
}
```

We simplify the interface by removing the ```context.Context``` argument.  If you still needed it, you can inject the ```IContextAccessor```.  This is following the asp.net model where you inject ```IHttpContextAccessor``` to get at the request context when a Request comes in.  

### Scoped Request  

In asp.net when a request comes in, a scoped di container is created and is active for the duration of the request.  We do the samne thing when a grpc request comes in.   
Looking at the ```IGreeterService```, it is up to your application to implement this interface.  It also must be registered as a scoped object.  In asp.net this is hidden from us, but here we have to do it manually.  

```go
import (
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/example/internal"
	pb "github.com/fluffy-bunny/grpcdotnetgo/example/internal/grpcContracts/helloworld"
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	contracts_logger "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/logger"
	contracts_request "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/request"
	grpc_error "github.com/fluffy-bunny/grpcdotnetgo/pkg/grpc/error"
	"google.golang.org/grpc/codes"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	Request         contracts_request.IRequest                 `inject:""`
	ClaimsPrincipal contracts_claimsprincipal.IClaimsPrincipal `inject:""`
	Logger          contracts_logger.ILogger                   `inject:""`
	Config          *internal.Config                           `inject:""`
}
// Ctor if it exists is called when the service is created
func (s *Service) Ctor() {
	s.Logger.Info().Msg("Ctor")
}

// Close if it exists is called when the container is torn down
func (s *Service) Close() {
	s.Logger.Info().Msg("Close")
}


// AddScopedIGreeterService adds service to the DI container
func AddScopedIGreeterService(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddScopedIGreeterService")
	pb.AddScopedIGreeterService(builder, reflect.TypeOf(&Service{}))
}
```

In asp.net core simply by adding an interface into our constructor, the framwork figures out by type what we need and injects it.  
```c#
public class MyGreeterService : IGreeterService
{
        private IHttpContextAccessor _httpContextAccessor;
        private ILogger<MyGreeterService> _logger;

        public MyGreeterService(
            IHttpContextAccessor httpContextAccessor,
            ILogger<MyGreeterService> logger)
        {
            _httpContextAccessor = httpContextAccessor;
            _logger = logger;
        }
}
```

The ```AddGreeterService``` func is our CTOR.  It is responsible for creating the object and as you can see it pulls the services it needs direct from the DI.  

# In Summary  
We map the go grpc endpoints, ```GreeterServer``` to our simpler interface ```IGreeterService```.  When a grpc call comes in, we create a scoped container, pull the matching interface ```IGreeterService``` and call it after we have made sure that relied upon scoped objects are setup correctly ```ILogger,IClaimsPrincipal,IContextAccessor``` etc.   






 
