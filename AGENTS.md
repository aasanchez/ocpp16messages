# AGENTS.md

## Prerequist

- Go ≥ 1.24
- Tools used by Makefile: golangci-lint, staticcheck, ripgrep (rg), gci, gofumpt, golines, pkgsite, sonar-scanner
  (optional)

## Build & Modules

- go mod tidy — ensure deps are tidy
- Library package; no binary build target

## Test

- All tests with coverage: make test (artifacts in reports/)
- Open HTML coverage: make test-coverage
- Run a single package: go test -v ./shared/types
- Run a single test: go test -v ./shared/types -run '^TestDateTime$'
- Run example tests: make test-example (or: go test -run '^Example' ./...)
- Benchmarks: make test-benchmark (or: go test -tags=benchmark -bench=. -benchmem ./...)
- Run a single benchmark: go test -run '^$' -bench '^BenchmarkInteger$' ./shared/types/benchmark

## Lint, Vet, Format

- Lint suite: make lint (golangci-lint, go vet → reports/govet.json, staticcheck → reports/staticcheck)
- Format: make format (gci write; gofumpt; golines; gofmt) — run before committing

## Code Style Guidelines

- Imports: managed by gci; group stdlib, module (github.com/aasanchez/ocpp16messages), then others; no unused imports
- Formatting: gofumpt + golines; keep lines concise; run make format to enforce
- Types: prefer precise types in shared/types; time parsing/formatting uses RFC3339/RFC3339Nano (see DateTime)
- Naming: exported identifiers use PascalCase; keep acronyms like ID in caps (revive var-naming allows ID)
- Errors: do not panic in library code; wrap with fmt.Errorf("context: %w", err); provide actionable messages
- Validation: keep constructors/setters returning (T, error) for validation (e.g., SetDateTime)
- Testing: table-driven tests where applicable; example_*.go for documentation tests; keep -race compatible when adding
  tests

## CI & Quality

- Coverage report: reports/coverage.out (go tool cover -func)
- Sonar/quality gate (optional): make sonar after test & lint

## Cursor/Copilot Rules

- No Cursor or Copilot rule files found; nothing to import at this time
