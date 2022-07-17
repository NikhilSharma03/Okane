# gen-protobuf generates .pb.go for protobuf files
gen-protobuf:
	protoc -I ./server/api \
		--go_out ./server/pkg --go_opt paths=source_relative \
  		--go-grpc_out ./server/pkg --go-grpc_opt paths=source_relative \
  	   	--grpc-gateway_out ./server/pkg --grpc-gateway_opt paths=source_relative \
  		./server/api/protobuf/*.proto


# clean-protobuf removes generated .pb.go files
clean-protobuf:
	rm -rf server/pkg/protobuf