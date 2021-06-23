# ProtoC  

```powershell
go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative internal\grpcContracts\helloworld\helloworld.proto
```
