# gen-protobuf generates .pb.go for protobuf files
gen-protobuf:
	protoc -I ./api \
		--go_out ./pkg --go_opt paths=source_relative \
  		--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
  	   	--grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
  		./api/protobuf/*.proto ./api/google/api/*.proto ./api/google/type/*.proto


# clean-protobuf removes generated .pb.go files
clean-protobuf:
	rm -rf pkg/protobuf pkg/google

# run-server starts the grpc server
run-server:
	go run cmd/main.go

# build-cli builds the cli app
build-cli:
	cd cli; go build -o okane; mv okane ./../okane-cli
