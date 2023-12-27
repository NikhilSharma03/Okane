## Proto files commands
.PHONY: generate-proto-protoc
generate-proto-protoc:
	@echo "Generating proto files"
	protoc -I ./api \
		--go_out ./pkg --go_opt paths=source_relative \
		--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
		./api/protobuf/*.proto ./api/google/api/*.proto ./api/google/type/*.proto

.PHONY: generate-proto
generate-proto:
	@echo "Generating proto files using buf"
	buf generate api

.PHONY: clean-proto
clean-proto:
	@echo "Cleaning generated proto files"
	rm -rf pkg/{protobuf,google}

## CLI commands
.PHONY: build-cli
build-cli:
	@echo "Building okane cli app"
	go build -o okane_cli ./cli

.PHONY: rm-cli
rm-cli:
	@echo "Removing okane cli app"
	rm -rf okane_cli

## Server commands
.PHONY: lint
lint:
	@echo "Linting code using golangci-lint"
	golangci-lint run ./...

.PHONY: run-server
run-server:
	@echo "Starting Okane server"
	go run cmd/main.go

## Development Server docker-compose commands
DOCKER_COMPOSE_DEV_FILE=docker-compose-dev.yml
DOCKER_COMPOSE_DEV_SERVER_SERVICE=okane_api

.PHONY: compose-dev-build
compose-dev-build:
	@echo "Running Okane API docker compose dev build"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) build

.PHONY: compose-dev-up
compose-dev-up:
	@echo "Running Okane API docker compose dev up in detach mode"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) up -d

.PHONY: compose-dev-logs
compose-dev-logs:
	@echo "Running Okane API docker compose dev logs"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) logs $(DOCKER_COMPOSE_DEV_SERVER_SERVICE) -f

.PHONY: compose-dev-down
compose-dev-down:
	@echo "Running Okane API docker compose dev down"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) down
