::run under the workspace directory
protoc --go_out=. --go-grpc_out=. .\message\little_bee.proto
protoc-go-inject-tag -input=.\message\littlebee\little_bee.pb.go
protoc --go_out=. --go-grpc_out=. .\message\query_command.proto