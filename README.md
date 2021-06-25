# gRPC-dot-net-go

## Contracts  
It all starts with a [proto file](example/internal/grpcContracts/helloworld/helloworld.proto).  In our case we want to have a custom server implementation that understands that the real downstream hanlder is in our DI.  Fortunately we can use a protoc plugin to generate the GO code for that.  

```powershell
cd example
go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative internal\grpcContracts\helloworld\helloworld.proto
```
This will generate all the [go files](example/internal/grpcContracts/helloworld) from our ```helloworld.proto```  

## Make it looke like dotnet core  
In dotnet core when a request comes in the DI has a concept of a scope.  Here we do the same thing by creating that scoped container in a [middleware](middleware/dicontext/dicontext.go)  

```go
    requestContainer, _ := grpcdotnetgo.GetContainer().SubContainer()
		defer requestContainer.Delete()

		ctx = setRequestContainer(ctx, requestContainer)

		contextaccessor := contextaccessor.GetInternalGetContextAccessorFromContainer(requestContainer)
		contextaccessor.SetContext(ctx)
```
The middleware gets a SubContainer from the root container.  When we use this new ```requestContainer``` to fetch objects from it, it will honor scoped services.  The main root container will NOT give you an object that has been registered as scoped for good reason as the root container is global and is there for singletons. 
