install-grpc-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc 

install-grpc-client:
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
	grpcui -plaintext localhost:5000
