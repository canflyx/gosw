PROJECT_NAME=switches manaeger
MAIN_FILE=main.go

PKG := "github.com/canflyx/gosw"
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION:=$(shell go version |grep -o 'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"
GIT_TAG:="1.0"

.PHONY: all dep lint vet test test-coverage build clean

all:  build

dep: ## Get the dependencies
	@ go mod tidy

gen: ## gen grpc go Files
	@ protoc -I . *.proto --go_out=. --go-grpc_out=.
	@ protoc-go-inject-tag -input=apps/*/*/*.pb.go

build: dep ## Build the binary file
	@go build -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_TAG=${GIT_TAG}'-X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" -o dist/demo-api.exe $(MAIN_FILE)

linux: dep ## Build the binary file
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" -o dist/demo-api $(MAIN_FILE)
run: # Run Develop server
	@go run $(MAIN_FILE) start -f config.toml

clean: ## Remove previous build
	@rm -f dist/*.exe

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'