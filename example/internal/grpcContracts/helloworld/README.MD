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

go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --proto_path=. --proto_path=vendor --proto_path=vendor/github.com/fluffy-bunny  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative internal\grpcContracts\helloworld\helloworld.proto  
```
