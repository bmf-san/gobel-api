.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.PHONY: docker-compose-build
docker-compose-build: ## Build containers by docker-compose.
	docker-compose -f docker-compose-local.yml build

.PHONY: docker-compose-up
docker-compose-up: ## Run containers by docker-compose.
	docker-compose -f docker-compose-local.yml up

.PHONY: docker-compose-up-d
docker-compose-up-d: ## Run containers in the background by docker-compose.
	docker-compose -f docker-compose-local.yml up -d

.PHONY: docker-compose-pull
docker-compose-pull: ## Pull images by docker-compose.
	docker-compose -f docker-compose-local.yml pull

.PHONY: build-and-push
build-and-push: ## Build and push image to dockerhub.
	docker build -f app/Dockerfile -t bmfsan/gobel-api ./app/
	docker push bmfsan/gobel-api

.PHONY: tbls
tbls: ## Run tbls for generationg database documents.
	docker run --net gobel_link -it --rm -v $$(pwd)/doc:/doc -w /doc/ k1low/tbls:latest doc --force mysql://root:password@gobel-api-mysql:3306/gobel

.PHONY: mod
mod: ## Run go mod download.
	cd app && go mod download

.PHONY: install-go-cleanarch
install-go-cleanarch: ## Install staticcheck.
ifeq ($(shell command -v go-cleanarch 2> /dev/null),)
	cd app && go install github.com/roblaszczak/go-cleanarch@latest
endif

.PHONY: install-staticcheck
install-staticcheck: ## Install staticcheck.
ifeq ($(shell command -v staticcheck 2> /dev/null),)
	cd app && go install honnef.co/go/tools/cmd/staticcheck@latest
endif

.PHONY: go-cleanarch
go-cleanarch: ## Run go-cleanarch.
	cd app && go-cleanarch -application usecase

.PHONY: staticcheck
staticcheck: ## Run staticcheck.
	cd app && staticcheck ./...

.PHONY: gofmt
gofmt: ## Run gofmt.
	cd app && test -z "$(gofmt -s -l . | tee /dev/stderr)"

.PHONY: vet
vet: ## Run vet.
	cd app && go vet -v ./...

.PHONY: test
test: ## Run unit tests.
	cd app && go test -v -race ./...

.PHONY: test-cover
test-cover: ## Run unit tests with cover options. ex. make test-cover OUT="c.out"
	cd app && go test -v -race -cover -coverprofile=$(OUT) -covermode=atomic ./...

.PHONY: test-api
test-api: ## Run tests for api responses with using db.
	cd app && go test -tags=intefration

.PHONY: build
build: ## Run go build
	cd app && go build