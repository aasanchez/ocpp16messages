DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-13s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Basic
.PHONY: test
test: ## is used to run the test suite of the application
	@rm -rf .reports && mkdir -p .reports
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out  -run '^Test' ./... > .reports/test.txt
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=.reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

.PHONY: test-example
test-example: ## is used to run the test suite of the applications
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out  -run '^Example' ./...

.PHONY: test-full
test-full: ## is used to run the test suite of the application
	@rm -rf .reports && mkdir -p .reports
	@go test -mod=readonly -v ./... >.reports/test.txt || true
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out ./... > .reports/test.txt
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=.reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

.PHONY: coverage-html
coverage-html: test ## is used to generate the coverage report of the application
	@go tool cover -html=.reports/coverage.out -o .reports/coverage.html
	@open -a "Google Chrome" .reports/coverage.html

.PHONY: lint
lint:
	@rm -rf .reports/*
	@go test ./... -json > .reports/test-report.out || true
	@go clean -testcache; go test ./... -coverprofile=.reports/coverage.out || true
	@golangci-lint --config golangci.yml run ./... || true
	@go vet ./... >.reports/govet.json
	@staticcheck ./... >.reports/stattickcheck

.PHONY: sonar
sonar: test lint
	@sonar-scanner

.PHONY: format
format: ## is used to format the code of the application
	@gofmt -d .

pkgsite:
	@echo "Stopping any running pkgsite processes..."
	@pkill pkgsite || true
	@echo "Cleaning Go module cache..."
	@go clean -modcache
	@rm -rf $$GOPATH/pkg
	@echo "Tidying up modules..."
	@go mod tidy
	@echo "Starting pkgsite at http://localhost:8080 ..."
	@nohup pkgsite > /dev/null 2>&1 &
	@sleep 2
	@open -a "Google Chrome" http://localhost:8080/github.com/aasanchez/ocpp16messages
