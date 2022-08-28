# Server Streaming example

### Needed dependencies

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

and 

export GO_PATH=~/go
export PATH="$PATH:$(go env GOPATH)/bin"

go get -u google.golang.org/grpc

if any error (not a good practice)

go mod vendor
mod tidy
```

### To generate the needed golang files from the proto files.

```
protoc -I./protofiles/data_streaming        \
  --go_out=./server      \
  --go-grpc_out=require_unimplemented_servers=false:./server \
  --go_out=./client      \
  --go-grpc_out=require_unimplemented_servers=false:./client \
  protofiles/data_streaming/streamingData.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    protofiles/data_streaming/streamingData.proto
  
  ```
