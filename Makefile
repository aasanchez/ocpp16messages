DATE := $(shell date +%Y)

##@ Helpers
.PHONY: help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1;34m${DOCKER_NAMESPACE}\033[0m\tCopyright (c) ${DATE} Alexis Sanchez\n \n\033[1;32mUsage:\033[0m\n  make \033[1;34m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[1;34m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1;33m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Test
.PHONY: test
test: ## is used to run the test suite of the application
	@go test ./... -v    

.PHONY: bench
bench: ## is used to reset the infrastructure o an inditial state
	@go test -bench=. -benchmem ./benchmark
	@go test -bench=BenchmarkValidateRawMessage -cpuprofile cpu.prof ./benchmark
	@go tool pprof -http=:8080 cpu.prof
