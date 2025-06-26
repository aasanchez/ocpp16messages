# OpenCode Guidelines for ocpp16messages

## Build, Lint, and Test

- **Build/Install:** Standard `go build ./...`
- **Run all tests:** `make test` or `go test ./...`
- **Run a single test file:** `go test ./path/to/file_test.go`
- **Run a specific test:** `go test -run '^Test<Name>$' ./...`
- **Example tests:** `make test-example`
- **Lint:** `make lint` or `golangci-lint run ./...`
- **Format:** `make format` or `gofmt -w .`
- **Benchmarks:** `make benchmark` or `go test -bench=. ./...`
- **Coverage HTML:** `make coverage-html`

## Code Style

- **Imports:** Group stdlib, third-party, then local, no unused imports.
- **Formatting:** Always use `gofmt`.
- **Types and Naming:**
  - Use PascalCase for exported and camelCase for unexported.
  - Suffix types with `Type` for strong typing (`AuthorizationStatusType`), interfaces without `I` prefix.
  - Constants use CamelCase, consider logical groupings (e.g., `Accepted`).
- **Error Handling:** Errors must be wrapped with context, use error variables for common error types (`ErrInvalidIdTag`).
- **Testing:** Use `t.Parallel()` where possible. Prefer table-driven tests for variations. Coverage is required for new code.
- **Packages:** Keep package structure flat and semantically grouped; use `_test` suffix for external test packages.
- **Linting:** All code—manual or agent-generated—must pass `golangci-lint` using the current repository configuration (`golangci.yml`). Compliance with `staticcheck` is also required.
- **Docs:** Document all exported functions/types. Use Go-style doc comments.
- **Markdown:** All markdown documentation, including this file, should conform to the default rules of [markdownlint](https://github.com/DavidAnson/markdownlint).
- **No Cursor or Copilot-specific rules present at this time.**
