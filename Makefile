DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-13s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Basic
.PHONY: test
test: ## is used to run the test suite of the application
	@go clean -testcache; go test ./...

.PHONY: test-verbose
test-verbose: ## is used to run the test suite of the application in verbose mode
	@go test ./... -v

.PHONY: coverage
coverage: test ## is used to generate the coverage report of the application
	@go clean -testcache; go test ./... -coverprofile=.reports/coverage.out; go tool cover -func=.reports/coverage.out

.PHONY: coverage-html
coverage-html: coverage ## is used to generate the coverage report of the application
	@open -a "Google Chrome" .reports/coverage.html

.PHONY: lint
lint:
	@rm -rf .reports/*
	@go test ./... -json > .reports/test-report.out || true
	@go clean -testcache; go test ./... -coverprofile=.reports/coverage.out || true
	@golangci-lint run ./... --config dev/golangci.yml || true
	@go vet ./... >.reports/govet.json
	@staticcheck ./... >.reports/stattickcheck

.PHONY: sonar
sonar: lint
	@sonar-scanner


.PHONY: format
format: ## is used to format the code of the application
	@gofmt -d .

doc: ## is used to format the code of the application
	@echo "Starting godoc on http://localhost:6060 ..."
	@nohup godoc -http=:6060 > /dev/null 2>&1 &
	@sleep 1
	@open -a "Google Chrome" http://localhost:6060/pkg/github.com/aasanchez/ocpp16messages/
