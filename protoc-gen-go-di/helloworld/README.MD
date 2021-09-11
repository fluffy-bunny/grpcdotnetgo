# ProtoC  

## Individual outputs  

go mod vendor

```powershell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
go get -u github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/cmd/protoc-gen-go-di
protoc --go_out=. --go_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative helloworld/helloworld.proto 
```

## One hit wonder

go mod vendor

```powershell
go get -u github.com/fluffy-bunny/grpcdotnetgo/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --proto_path=. --proto_path=vendor --proto_path=vendor/github.com/fluffy-bunny  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative .\protoc-gen-go-di\helloworld\helloworld.proto  

```
