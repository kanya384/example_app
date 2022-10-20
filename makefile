#include .env
#export

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

up: ## testing infrastracture
	@docker-compose up & disown
.PHONY: up

test: ## runs unit and integration tests
	@GOCACHE=off go test ./... -cover
.PHONY: test

genproto: ## generates go package from proto files
	@protoc -I ./api ./api/auth.proto --go-grpc_out=./services/auth/internal/delivery/grpc/interface --go_out=./services/auth/internal/delivery/grpc/interface --grpc-gateway_out=./services/auth/internal/delivery/grpc/interface
	

gengateway:
	@protoc -I ./api --openapiv2_out=./gen/openapiv2 --openapiv2_opt allow_merge=true \
   	--go_out ./services/gateway/internal/gateway --go_opt paths=source_relative \
	--go-grpc_out ./services/gateway/internal/gateway --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./services/gateway/internal/gateway --grpc-gateway_opt paths=source_relative \
    ./api/*.proto