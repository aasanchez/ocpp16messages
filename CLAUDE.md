# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

OCPP 1.6 Messages for Go - A library providing OCPP 1.6-compliant message types and validation for EV charging station management. This is a library package with no binary build target.

## Prerequisites

- Go >= 1.24
- Tools: golangci-lint, staticcheck, ripgrep (rg), gci, gofumpt, golines, pkgsite (optional), sonar-scanner (optional)

## Common Commands

### Dependencies

```sh
go mod tidy
```

### Execution Testing

```sh
make test                   # Unit and example tests with coverage (reports in reports/)
make test-coverage          # Generate and open HTML coverage report
make test-example           # Run example tests (documentation tests)
make test-all               # Run all test types

# Run single test package
go test -v ./shared/types

# Run specific test
go test -v ./shared/types -run '^TestDateTime$'
```

### Linting and Formatting

```sh
make format                 # Format all Go code (gci, gofumpt, golines, gofmt)
make lint                   # Run golangci-lint, go vet, staticcheck (reports in reports/)
```

### Documentation

```sh
make pkgsite                # Start local documentation server at http://localhost:8080
```

## Architecture

### Core Type System (shared/types/)

The library implements OCPP 1.6 strict type validation using constructor pattern:

- **Integer**: Wraps uint16 with validation via `NewInteger(string) (Integer, error)`
- **DateTime**: RFC3339-compliant, auto-normalized to UTC. Parsing uses `time.RFC3339`, output uses `time.RFC3339Nano`
  - Constructor: `NewDateTime(string) (DateTime, error)`
- **CiString Types**: Length-validated ASCII printable strings (32-126)
  - CiString20, CiString25, CiString50, CiString255, CiString500
  - All use `NewCiString<N>(string)` constructor that validates length and ASCII range

All types use value receivers and immutable fields, designed for thread-safe concurrent use.

### Test Organization

- `shared/types/tests/` - Unit tests (table-driven where applicable)
- `shared/types/example_*_test.go` - Documentation/example tests

## Code Style Guidelines

### Imports

- Managed by gci: stdlib first, then github.com/aasanchez/ocpp16messages, then others
- No unused imports

### Formatting

- Enforced by gofumpt + golines
- Always run `make format` before committing

### Type Design

- Use precise types from shared/types for OCPP fields
- Time parsing/formatting strictly uses RFC3339/RFC3339Nano (see DateTime type)
- All constructors/setters return `(T, error)` for validation

### Naming

- Exported identifiers: PascalCase
- Keep acronyms uppercase (e.g., Id not ID - allowed by revive var-naming rule)
- **Constructors**: MUST use `New` prefix following Go idioms
  - Public constructors: `NewType(args) (Type, error)` (e.g., `NewInteger`, `NewDateTime`, `NewCiString20Type`)
  - Private constructors: `newType(args) (type, error)` (lowercase 'n' for unexported)
  - **NEVER** use `Set` prefix for constructors (e.g., `SetInteger` ✗, `NewInteger` ✓)
  - Rationale: "Set" implies mutation; "New" indicates construction (see Effective Go)
- **Getters**: Do NOT use `Get` prefix
  - Use `Value()` not `GetValue()` (see Effective Go getter guidelines)
- **Test naming**: Follow Go conventions
  - Unit tests: `Test_<package>_<Function>` (e.g., `Test_sharedtypes_NewInteger`)
  - Example tests: `Example<Function>` (e.g., `ExampleNewInteger`)
  - Subtests: Use descriptive suffixes (e.g., `Test_sharedtypes_NewInteger_Overflow`)

### Error Handling

- Never panic in library code
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Provide actionable error messages

### Testing

- **CRITICAL: Write atomic, individual test functions** - Each test should verify ONE specific behavior or case
  - ✓ GOOD: `TestIdToken_String_ReturnsValue()`, `TestIdToken_Empty_ReturnsError()`
  - ✗ BAD: Table-driven tests with multiple cases in one function
  - Rationale: Atomic tests are easier to debug, maintain, and understand. Each test failure points directly to the specific case that broke.
  - Exception: Only use table-driven tests when testing the same logic with slight variations (e.g., boundary values), and even then prefer splitting into separate functions when possible.
- Place example tests in `example_*_test.go` for documentation
- Ensure all new tests are race-compatible
- Each test must call `t.Parallel()` for concurrent execution

## CI and Quality

### Coverage

- Coverage reports generated in `reports/coverage.out`
- View with: `go tool cover -func=reports/coverage.out`
- HTML report: `make test-coverage`

### Linter Output

- Execute `make lint` to execute all the lints
- golangci-lint: `reports/golangci-lint.txt`
- go vet: `reports/govet.json`
- staticcheck: `reports/staticcheck`

## OCPP Compliance Notes

- All DateTime fields normalized to UTC per OCPP 1.6 spec
- CiString types enforce ASCII printable characters only (32-126)
- All types proven thread-safe via race tests
- String length validation enforced at construction time
