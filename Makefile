# gen-protobuf generates .pb.go for protobuf files
gen-protobuf:
	protoc --go_out=./server/pkg --go-grpc_out=./server/pkg server/api/google/type/*.proto server/api/protobuf/*.proto

# clean-protobuf removes generated .pb.go files
clean-protobuf:
	rm -rf server/pkg/okanepb server/pkg/google.golang.org