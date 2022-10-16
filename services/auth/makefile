

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

genproto: ## generates go package from proto files
	@protoc  api/proto/*.proto --go-grpc_out=./internal/delivery/grpc/interface --go_out=./internal/delivery/grpc/interface
	@protoc  ../notifications/api/proto/*.proto --go-grpc_out=./pkg/types/notifications --go_out=./pkg/types/notifications
.PHONY: genproto

run: ## runs app localy
	@go run cmd/app/main.go
.PHONY: run

compose-dev-up: ## testing infrastracture
	@docker-compose -f deployments/docker-compose-dev.yml up & disown
.PHONY: compose-dev-up

compose-down: ## docker compose down
	@docker-compose -f deployments/docker-compose-dev.yml down
.PHONY: compose-down

genmocks: ## generate mocks
	mockgen -source=./internal/useCase/adapters/storage/interface.go -destination=./test/mocks/storage/storage.go -package=storage
	mockgen -source=./internal/useCase/adapters/pubsub/interface.go -destination=./test/mocks/pubsub/pubsub.go -package=pubsub
	mockgen -source=./internal/useCase/adapters/cache/cache.go -destination=./test/mocks/cache/cache.go -package=cache
	mockgen -source=./internal/useCase/interface.go -destination=./test/mocks/useCase/usecase.go -package=useCase
	mockgen -source=./pkg/auth/auth.go -destination=./test/mocks/auth/auth.go -package=auth
	mockgen -source=./pkg/logger/logger.go -destination=./test/mocks/logger/logger.go -package=logger
.PHONY: genmocks

test: ## runs unit and integration tests
	@GOCACHE=off go test ./... -cover
.PHONY: test