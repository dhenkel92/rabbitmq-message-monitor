.PHONY: build
build: ## Builds a production read binary without symbols and debug flags
	go build -ldflags "-s -w" $(ARGS)

.PHONY: style
style: ## Executes the pre-commit configuration
	pre-commit run -a

.PHONY: run
run: ## Runs the application
	go run main.go -c "amqp://guest:guest@localhost:5672/" exchange --binding "amqp.topic=#" $(ARGS)

.PHONY: start-test
start-test:
	docker-compose up -d
	sleep 15
	docker-compose -f docker-compose-test.yaml up

.PHONY: clean
clean:
	docker-compose down
	docker-compose -f docker-compose-test.yaml down

.PHONY: test
test: ## Executes unit tests
	go test -v ./...

.PHONY: help
help: ## Get all commands with short explanation
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
