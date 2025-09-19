DATE := $(shell date +%Y)

##@ Helpers
help: ## Display this help message, listing all available targets and their descriptions.
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Testing
FUZZTIME ?= 10s
FUZZPKG ?= ./shared/types/fuzz

test: ## Run all tests starting with "Test" and collect coverage.
	@echo "\n--- \033[1;32mExecute Unit Test\033[0m ---"
	@rm -rf reports && mkdir -p reports
	@go clean -cache -testcache -modcache
	@go test -mod=readonly \
		-coverpkg=./... \
		-coverprofile=reports/coverage.out \
		-v ./... >reports/test.txt
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

test-coverage: test ## Generate and open a detailed HTML coverage report in the default browser.
	@go tool cover -html=reports/coverage.out -o reports/coverage.html
	@open -a "Google Chrome" reports/coverage.html

test-example: ## Run documentation-based example tests to verify correctness of usage examples.
	@echo "\n--- \033[1;32mTest Examples\033[0m ---"
	@go test -mod=readonly -v -run '^Example' ./...

test-benchmark: ## Run benchmark tests to measure performance of critical operations.
	@go clean -modcache
	@echo "\n--- \033[1;32mBenchmark\033[0m ---"
	@go test -tags=benchmark -bench=. -benchmem ./...

## Fuzz testing targets
.test-fuzz-header:
	@echo "\n--- \033[1;32mFuzz (time=$(FUZZTIME))\033[0m ---"

test-fuzz-integer: .test-fuzz-header ## Fuzz SetInteger for $(FUZZTIME)
	@echo "fuzz-integer"
	@go test -run '^$$' $(FUZZPKG) -fuzz=FuzzSetInteger -fuzztime=$(FUZZTIME)

test-fuzz-cistring: .test-fuzz-header ## Fuzz CiString types for $(FUZZTIME)
	@echo "fuzz-cistring"
	@go test -run '^$$' $(FUZZPKG) -fuzz=FuzzCiStringTypes -fuzztime=$(FUZZTIME)

test-fuzz: test-fuzz-integer test-fuzz-cistring ## Run all fuzz tests (use FUZZTIME=30s to increase duration)

test-all: test test-example test-fuzz test-benchmark

##@ Code Style and Static Analysis
lint: ## Run static analysis, vetting, and linting using golangci-lint and other tools.
	@golangci-lint cache clean || true
	@golangci-lint --config golangci.yml run ./... || true
	@go vet ./... > reports/govet.json
	@staticcheck ./... > reports/staticcheck

sonar: test lint ## Run test suite and linters, then push code quality reports to SonarQube.
	@sonar-scanner

format: ## Format Go code to maintain consistent styling across the codebase.
	@rg --files -g '*.go' | xargs gci write
	@rg --files -g '*.go' | xargs gofumpt -l -w
	@rg --files -g '*.go' | xargs golines -w
	@gofmt -w .

##@ Documentation
pkgsite: ## Start a local pkgsite server to browse Go documentation interactively.
	@echo "Stopping any running pkgsite processes..."
	@pkill pkgsite || true
	@echo "Tidying up Go modules..."
	@go mod tidy
	@echo "Launching pkgsite in the background..."
	@nohup pkgsite -http=:8080 . > /dev/null 2>&1 &
	@echo "pkgsite server started. You can view the documentation at http://localhost:8080/github.com/aasanchez/ocpp16messages"
	@open -a "Google Chrome" http://localhost:8080/github.com/aasanchez/ocpp16messages
