DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help: ## Display this help message, listing all available targets and their descriptions.
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-13s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Testing
test: ## Run all unit tests and generate a code coverage report for critical test paths.
	@rm -rf .reports && mkdir -p .reports
	@go test -mod=readonly -v -coverprofile=.reports/coverage.out -run '^Test' ./... > .reports/test.txt

test-coverage: test ## Generate and open a detailed HTML coverage report in the default browser.
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=.reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

test-coverage-html: test-coverage ## Generate and open a detailed HTML coverage report in the default browser.
	@go tool cover -html=.reports/coverage.out -o .reports/coverage.html
	@open -a "Google Chrome" .reports/coverage.html

.PHONY: benchmark
benchmark: ## Run benchmark tests to measure performance of critical operations.
	@echo "Stopping any running pkgsite processes..."
	@go clean -modcache
	@go test -bench=. -benchmem ./...

##@ Code Style and Static Analysis
.PHONY: lint
lint: ## Run static analysis, vetting, and linting using golangci-lint and other tools.
	@rm -rf .reports/*
	@go test ./... -json > .reports/test-report.out || true
	@go clean -testcache; go test ./... -coverprofile=.reports/coverage.out || true
	@golangci-lint cache clean
	@golangci-lint --config golangci.yml run ./... || true
	@go vet ./... > .reports/govet.json
	@staticcheck ./... > .reports/staticcheck

.PHONY: sonar
sonar: test lint ## Run test suite and linters, then push code quality reports to SonarQube.
	@sonar-scanner

.PHONY: format
format: ## Format Go code to maintain consistent styling across the codebase.
	@gofmt -d .

##@ Documentation
.PHONY: pkgsite
pkgsite: ## Start a local pkgsite server to browse Go documentation interactively.
	@echo "Stopping any running pkgsite processes..."
	@pkill pkgsite || true
	@echo "Cleaning Go module cache..."
	@go clean -modcache
	@rm -rf $$GOPATH/pkg
	@echo "Tidying up Go modules..."
	@go mod tidy
	@echo "Launching pkgsite at http://localhost:8080 ..."
	@nohup pkgsite > /dev/null 2>&1 &
	@sleep 2
	@open -a "Google Chrome" http://localhost:8080/github.com/aasanchez/ocpp16messages
