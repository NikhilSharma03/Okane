# gen-protobuf generates .pb.go for protobuf files
gen-protobuf:
	protoc --go_out=./server/pkg --go-grpc_out=./server/pkg server/api/protobuf/*