# blogWebSite
CRUD operation for blog posts

s1 : go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
s2 : go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
s3	: protoc --go_out=./protos ./protos/blog.proto
s4 : protoc --go-grpc_out=./protos ./protos/blog.proto
s4 : go run server.go
s5 : go run sampleGRPCClient.go