DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-13s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Basic
.PHONY: test
test: ## is used to run the test suite of the application
	@rm -rf .reports && mkdir -p .reports
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out  -run '^Test' ./... > .reports/coverage.txt
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=.reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

##@ Basic
.PHONY: test-full
test-full: ## is used to run the test suite of the application
	@rm -rf .reports && mkdir -p .reports
	@go test -mod=readonly -v ./... >.reports/test.txt || true
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out ./... > .reports/coverage.txt
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
	@golangci-lint run ./... || true
	@go vet ./... >.reports/govet.json
	@staticcheck ./... >.reports/stattickcheck

.PHONY: sonar
sonar: test lint
	@sonar-scanner

.PHONY: format
format: ## is used to format the code of the application
	@gofmt -d .

doc: ## is used to format the code of the application
	@echo "Starting godoc on http://localhost:6060 ..."
	@nohup godoc -http=:6060 > /dev/null 2>&1 &
	@sleep 1
	@open -a "Google Chrome" http://localhost:6060/pkg/github.com/aasanchez/ocpp16messages/
