PROGRAM_NAME = pgcenter

COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags |cut -d- -f1)

LDFLAGS = -ldflags "-X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}"

.PHONY: help clean dep build install uninstall 

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

dep: ## Download the dependencies.
	go mod download

build: dep ## Build pgcenter executable.
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${PROGRAM_NAME} ./cmd

clean: ## Clean build directory.
	rm -f ./bin/${PROGRAM_NAME}
	rmdir ./bin

lint: dep ## Lint the source files
	golangci-lint run --timeout 5m -E golint
	gosec -quiet ./...lint: dep ## Lint the source files
	
test: dep ## Run tests
	go test -race -p 1 -timeout 300s -coverprofile=.test_coverage.txt ./... && \
    	go tool cover -func=.test_coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm .test_coverage.txt

docker-build: ## Build docker image
	docker build -t lesovsky/pgcenter:${TAG} .
	docker image prune --force --filter label=stage=intermediate

docker-push: ## Push docker image to registry
	docker push lesovsky/pgcenter:${TAG}