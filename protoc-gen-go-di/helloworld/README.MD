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

go build .\protoc-gen-go-di\cmd\protoc-gen-go-di\  

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative --go-di_out=. --go-di_opt=paths=source_relative,grpc_gateway=true protoc-gen-go-di/helloworld/helloworld.proto

```

```bash
git clone https://github.com/googleapis/googleapis.git
# copy to google/api is so that the linter can find the files in visual studio. 
# leave the googleapis folder in place for the protoc to find the files
md ./google/api
cp ./googleapis/google/api/annotations.proto ./google/api/annotations.proto
cp ./googleapis/google/api/field_behavior.proto ./google/api/field_behavior.proto
cp ./googleapis/google/api/http.proto ./google/api/http.proto
cp ./googleapis/google/api/httpbody.proto ./google/api/httpbody.proto
```
