SHELL := /bin/bash
.PHONY: help

ifeq ($(CI),true)
    GO_TAGS := ci
else
    GO_TAGS := local
endif

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run locally
	export `cat .env | xargs` && \
	go run main.go

create-dist:
	rm -rf dist || true && \
	mkdir dist

build: create-dist ## Build binary
	GOOS=linux GOARCH=amd64 go build -o dist/monitor

build-multiplatform: build ## Create binaries for all supported platforms
	GOOS=linux GOARCH=amd64 go build -o dist/monitor-linux-amd64 && \
    GOOS=linux GOARCH=386 go build -o dist/monitor-linux-i386 && \
    GOOS=linux GOARCH=arm go build -o dist/monitor-linux-arm && \
    GOOS=linux GOARCH=arm64 go build -o dist/monitor-linux-arm64

coverage: ## Run tests and measure coverage
	go test -tags $(GO_TAGS) -covermode=count -coverprofile=/tmp/count.out -v ./...

test-coverage-report: coverage ## Run test and display coverage report in browser
	go tool cover -html=/tmp/count.out

save-coverage-report: coverage ## Save coverage report to coverage.html
	go tool cover -html=/tmp/count.out -o ../coverage.html
