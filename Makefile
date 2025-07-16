DATE := $(shell date +%Y)
FUZZ_TIME ?= 30s

##@ Helpers
.PHONY: help

help: ## Display this help message, listing all available targets and their descriptions.
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-13s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Testing
test: ## Run all unit tests and generate a code coverage report for critical test paths.
	@rm -rf reports && mkdir -p reports
	@go test -mod=readonly -v \
		-coverpkg=`go list -mod=readonly ./... | grep -v '/tests' | paste -sd ',' -` \
		-coverprofile=reports/coverage.out -run '^Test' ./... > reports/test.txt

test-coverage: test ## Generate and open a detailed HTML coverage report in the default browser.
	@echo "\n--- \033[32mCoverage Percentage\033[0m:"
	@go tool cover -func=reports/coverage.out | tail -1 | awk -F" " '{print $$NF}'

test-coverage-html: test-coverage ## Generate and open a detailed HTML coverage report in the default browser.
	@go tool cover -html=reports/coverage.out -o reports/coverage.html
	@open -a "Google Chrome" reports/coverage.html

test-example: ## Run documentation-based example tests to verify correctness of usage examples.
	@go test -mod=readonly -v -coverprofile=reports/coverage.out -run '^Example' ./...

.PHONY: test-fuzz
test-fuzz: ## Run fuzz tests for each Fuzz* function in all packages
	@echo "Running fuzzing (Fuzz) on each package..."
	@for pkg in $$(go list ./...); do \
		for fuzz in $$(go test -timeout $(FUZZ_TIME) -list ^Fuzz $$pkg | grep ^Fuzz || true); do \
			echo "Fuzzing $$fuzz in package $$pkg (for $(FUZZ_TIME))"; \
			go test -mod=readonly -v -timeout $(FUZZ_TIME) -fuzztime=$(FUZZ_TIME) -fuzz=$$fuzz $$pkg; \
		done; \
	done

.PHONY: test-race
test-race: ## Run race detector across all packages.
	@echo "Running race detector on all packages..."
	@go list ./... | grep race | xargs -n1 -I{} go test -mod=readonly -v -timeout 30s -race {}


.PHONY: benchmark
benchmark: ## Run benchmark tests to measure performance of critical operations.
	@echo "Stopping any running pkgsite processes..."
	@go clean -modcache
	@go test -bench=. -benchmem ./...

##@ Code Style and Static Analysis
.PHONY: lint
lint: ## Run static analysis, vetting, and linting using golangci-lint and other tools.
	@rm -rf reports/*
	@golangci-lint cache clean
	@golangci-lint --config golangci.yml run ./... || true
	@go vet ./... > reports/govet.json
	@staticcheck ./... > reports/staticcheck

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
	@echo "Tidying up Go modules..."
	@go mod tidy
	@echo "Launching pkgsite in the background..."
	@nohup pkgsite -http=:8080 . > /dev/null 2>&1 &
	@echo "pkgsite server started. You can view the documentation at http://localhost:8080/github.com/aasanchez/ocpp16messages"
	@open -a "Google Chrome" http://localhost:8080/github.com/aasanchez/ocpp16messages
