# ProtoC  

```powershell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative     
helloworld/helloworld.proto 

go get -u github.com/fluffy-bunny/grpcdotnetgo/cmd/protoc-gen-go-di

  
protoc --go_out=. --go_opt=paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative   helloworld/helloworld.proto 

```
