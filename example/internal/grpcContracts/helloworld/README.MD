# ProtoC  

[prerequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)  

## proto requirments

```powershell
cd c:/work/github/protocolbuffers
git clone https://github.com/protocolbuffers/protobuf.git

cd c:/work/github/gogo
git clone https://github.com/gogo/protobuf.git

cd c:/work/github/fluffy-bunny/grpcdotnetgo/example
```

```powershell
cd example

go mod vendor 

go get -u github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative,grpc_gateway=true example/internal/grpcContracts/helloworld/helloworld.proto   
```
