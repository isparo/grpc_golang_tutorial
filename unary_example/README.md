
##Unary

protoc -I./protofiles/greetpb        \
  --go_out=./server      \
  --go-grpc_out=require_unimplemented_servers=false:./server \
  --go_out=./client      \
  --go-grpc_out=require_unimplemented_servers=false:./client \
  protofiles/greetpb/greet.proto