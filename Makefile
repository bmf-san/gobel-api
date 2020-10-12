.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

docker-compose-build: ## Build containers by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose-ci.yml build
else
	docker-compose -f docker-compose.yml build
endif

docker-compose-up: ## Run containers by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose-ci.yml up
else
	docker-compose -f docker-compose.yml up
endif

docker-compose-up-d: ## Run containers in the background by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose-ci.yml up -d
else
	docker-compose -f docker-compose.yml up -d
endif

docker-compose-pull: ## Pull images by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose-ci.yml pull
else
	docker-compose -f docker-compose.yml pull
endif

lint: ## Run golint.
	docker exec -it gobel-api golint ./...

test: ## Run tests.
	docker exec -it gobel-api go test -v ./...

tbls: ## Update database documents.
	tbls doc -f

build: ## Run go build
	cd app && GOOS=linux GOARCH=amd64 go build -o app