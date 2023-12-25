.PHONY: generate-proto-protoc
generate-proto-protoc:
	protoc -I ./api \
		--go_out ./pkg --go_opt paths=source_relative \
		--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
		./api/protobuf/*.proto ./api/google/api/*.proto ./api/google/type/*.proto

.PHONY: generate-proto
generate-proto:
	buf generate api

.PHONY: clean-proto
clean-proto:
	rm -rf pkg/protobuf pkg/google

.PHONY: run-server
run-server:
	go run cmd/main.go
