# Agent Guidelines for ocpp16messages

## Build, Lint, & Test

- **Build:** `go build ./...`
- **Run all tests:** `make test` or `go test ./...`
- **Run a single test file:** `go test ./path/to/file_test.go`
- **Run a single test:** `go test -run '^Test<Name>$' ./path/to/file_test.go`
- **Example tests:** `make test-example`
- **Lint:** `make lint` or `golangci-lint run ./...`
- **Format:** `make format` or `gofmt -w .`
- **Benchmarks:** `make benchmark` or `go test -bench=. ./...`
- **Coverage HTML:** `make coverage-html`

## Code Style

- **Imports:** Stdlib, then third-party, then local; no unused imports.
- **Formatting:** Always run `gofmt`.
- **Types/Naming:** PascalCase for exported, camelCase for unexported. Suffix "Type" for strong types (e.g., `AuthorizationStatusType`). Do not prefix interfaces with "I". Group constants logically using CamelCase.
- **Error Handling:** Wrap errors with context; use error vars for common cases (e.g., `ErrInvalidIdTag`).
- **Testing:** All tests must be atomic, minimal, and simple: one test, one case. Prefer `t.Parallel()`; avoid table-driven tests unless necessary. Achieve code 100% coverage for new code.
- **Packages:** Organize flat, semantically; use `_test` package suffix for external tests.
- **Linting:** Pass `golangci-lint`/`staticcheck` per `golangci.yml`.
- **Docs:** All exported APIs/types must have Go-style doc comments.
- **Markdown:** Follow [markdownlint](https://github.com/DavidAnson/markdownlint) for all docs.
- **Other:** No Cursor or Copilot rules currently.
