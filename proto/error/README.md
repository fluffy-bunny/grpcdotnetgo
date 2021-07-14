# ProtoC  

[prerequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)  

## proto requirments.
```

cd c:/work/github/protocolbuffers
git clone https://github.com/protocolbuffers/protobuf.git


cd c:/work/github/gogo
git clone https://github.com/gogo/protobuf.git

cd c:/work/github/fluffy-bunny/grpcdotnetgo/example

```

```powershell

cd example

go get -u github.com/fluffy-bunny/protoc-gen-go-di/cmd/protoc-gen-go-di

protoc --proto_path=. --proto_path=vendor --go_out=. --go_opt=paths=source_relative  proto\error\error.proto
```
