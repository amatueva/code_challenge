MAIN_PKG := github.com/zendesk/recruiting_coding_challenge/cmd/recruiting_coding_challenge
COVERAGE_FILE_NAME := coverage.out
TEST_TARGET ?= ./...
TEST_FLAGS := -race -coverpkg=./...  --coverprofile=$(COVERAGE_FILE_NAME)

default: run

run:
	go run cmd/recruiting_coding_challenge/main.go

build:
	go build -race "$(MAIN_PKG)"

ensure_deps:
	go mod tidy
	go mod vendor

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint_fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...

install_devtools:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_VERSION)
	go install gotest.tools/gotestsum
	go install github.com/go-delve/delve/cmd/dlv
	go install github.com/githubnemo/CompileDaemon

test:
	go test ./...