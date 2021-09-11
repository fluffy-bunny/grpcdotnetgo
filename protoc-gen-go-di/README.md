# gRPC DotNetGo ProtoCGen

## Dependencies  

[sarulabsdi](https://github.com/fluffy-bunny/sarulabsdi)  
Modified DI library that accounts for registering many of a given type.  

[grpcdotnetgo](https://github.com/fluffy-bunny/grpcdotnetgo)  
Core library that relies on sarulabsdi as it DI.  grpcdotnetgo brings in shared services that are similar to asp.net core;  

```go  
IServiceProvier
ILogger
ClaimsPrincipal
IContextAccessor (similar to IHttpRequestAccessor)
etc.
```

This protoc plugin builds out the grpc implementation of a given grpc server.  The concept is that an application registers an implementation of a published interface based on the proto server.  

In short the design mirrors how asp.net core works.   A request comes in, a container is created that only offers up scoped registered objects, and removes the need to pass context.Context into every method.  You can still get the context by pulling the scoped IContextAccessor from the container.  

## Individual outputs  

```powershell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di
protoc --go_out=. --go_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative helloworld/helloworld.proto 
```

## One hit wonder

```powershell
go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative helloworld/helloworld.proto 
```

## Example project  

[grpcdotnetgo example](https://github.com/fluffy-bunny/grpcdotnetgo/tree/main/example)  
