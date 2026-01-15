# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

OCPP 1.6 Messages for Go - A library providing OCPP 1.6-compliant message types
and validation for EV charging station management. This is a library package
with no binary build target.

## Prerequisites

- Go >= 1.24
- Tools: golangci-lint, staticcheck, ripgrep (rg), gci, gofumpt, golines,
  pkgsite (optional), sonar-scanner (optional)

## Common Commands

### Dependencies

```sh
go mod tidy
```

### Execution Testing

```sh
make test          # Unit and example tests with coverage (reports in reports/)
make test-coverage # Generate and open HTML coverage report
make test-example  # Run example tests (documentation tests)
make test-all      # Run all test types

# Run single test package
go test -v ./types

# Run specific test
go test -v ./types -run '^TestDateTime$'

# Run message package tests
go test -v ./authorize/...
go test -v ./clearChargingProfile/...
```

### Linting and Formatting

```sh
make format # Format all Go code (gci, gofumpt, golines, gofmt)
make lint   # Run golangci-lint, go vet, staticcheck (reports in reports/)
```

### Documentation

```sh
make pkgsite # Start local documentation server at http://localhost:8080
```

## Architecture

### Core Type System (types/)

The library implements OCPP 1.6 strict type validation using constructor
pattern:

- **Integer**: Wraps uint16 with validation via
  `NewInteger(string) (Integer, error)`
- **DateTime**: RFC3339-compliant, auto-normalized to UTC. Parsing uses
  `time.RFC3339`, output uses `time.RFC3339Nano`
  - Constructor: `NewDateTime(string) (DateTime, error)`
- **CiString Types**: Length-validated ASCII printable strings (32-126)
  - CiString20, CiString25, CiString50, CiString255, CiString500
  - All use `NewCiString<N>(string)` constructor that validates length and
    ASCII range

All types use value receivers and immutable fields, designed for thread-safe
concurrent use.

### Message API Design

OCPP messages follow an **Input struct pattern** where users create a simple
struct with raw values and pass it to the constructor. The constructor validates
all fields automatically - there is NO separate `Validate()` method.

#### Naming Convention

Message types use OCPP terminology:

- **Req()**: Constructor for request messages (e.g., `authorize.Req()` for Authorize.req)
- **Conf()**: Constructor for response messages (e.g., `authorize.Conf()` for Authorize.conf)
- **ReqInput/ConfInput**: Input structs for raw values
- **ReqMessage/ConfMessage**: The returned types representing validated messages

#### Usage Pattern

All messages use a `ReqInput`/`ConfInput` struct that accepts raw values
(strings, integers):

```go
// Authorize.req message
req, err := authorize.Req(authorize.ReqInput{
    IdTag: "RFID-ABC123",
})
if err != nil {
    // Handle validation error
}
```

#### Complex Messages (nested structures, arrays)

For messages with complex nested structures (like MeterValues with arrays of
SampledValue), the input struct contains nested input structs:

```go
// MeterValues.req - complex nested structure
req, err := metervalues.Req(metervalues.ReqInput{
    ConnectorId:   1,
    TransactionId: 123,
    MeterValue: []metervalues.MeterValueInput{
        {
            Timestamp: "2025-01-02T15:00:00Z",
            SampledValue: []metervalues.SampledValueInput{
                {Value: "100", Measurand: "Energy.Active.Import.Register"},
            },
        },
    },
})
```

#### Design Principles

- **OCPP naming**: Use `Req()`/`Conf()` to match OCPP terminology
  (Authorize.req, Authorize.conf)
- **Input structs with raw values**: Users create `ReqInput`/`ConfInput`
  structs with simple Go types
- **Single constructor call**: One `Req(input)` call validates everything
- **No separate Validate()**: Validation is built into the constructor
- **Error on construction**: Return `(ReqMessage/ConfMessage, error)` - if
  error is nil, the message is valid
- **Immutable result**: The returned `ReqMessage`/`ConfMessage` contains
  validated, typed fields

#### What NOT to do

```go
// BAD - Don't require separate validation step
input := authorize.ReqInput{IdTag: "ABC123"}
if err := input.Validate(); err != nil { ... }  // NO separate Validate()!
req, err := authorize.Req(input)

// GOOD - Single step, validation built-in
req, err := authorize.Req(authorize.ReqInput{IdTag: "ABC123"})
```

### Test Organization

Tests MUST be organized by API visibility to maintain clear boundaries between
internal implementation and public API:

#### Same-Package Tests (`*_test.go` with `package <name>`)

- **Location**: In the same directory as the code being tested
- **Package**: MUST use same package name (e.g., `package types`)
- **Purpose**: Test private/unexported functions and internal implementation
  details
- **Access**: Can access package internals (unexported functions, types,
  variables)
- **When to create**: ONLY when you have private/unexported functions that need
  testing
- **Example**: `types/cistring_test.go` tests `newCiString()` (private
  constructor)

```go
// types/cistring_test.go
package types  // Same package - can access private functions

func Test_sharedtypes_newCiString_Valid(t *testing.T) {
    // Tests private newCiString() function
    cis, err := newCiString("test", 20)
    // ...
}
```

#### Public API Tests (`tests/*_test.go` with `package <name>_test`)

- **Location**: In a `tests/` subdirectory within the package
- **Package**: MUST use `package <name>_test` (e.g., `package types_test`)
- **Purpose**: Test all exported/public API as external consumers would use it
- **Access**: Black-box testing - imports the package like external code would
- **When to create**: For ALL public constructors, methods, and functions
- **Example**: `types/tests/cistring_test.go` tests
  `NewCiString20Type()` (public constructor)

```go
// types/tests/cistring_test.go
package types_test  // External package - black-box testing

import (
    "testing"

    "github.com/aasanchez/ocpp16messages/types"
)

func Test_sharedtypes_NewCiString20Type(t *testing.T) {
    t.Parallel()
    // Tests public NewCiString20Type() function
    _, err := types.NewCiString20Type("test")
    // ...
}
```

#### Example Tests (`example_*_test.go` with `package <name>_test`)

- **Location**: In the same directory as the code (NOT in `tests/`)
- **Package**: MUST use `package <name>_test` for external perspective
- **Purpose**: Executable documentation for public constructors and complex
  APIs
- **Guidelines**: See "Example Tests - Use Selectively" section below

#### Directory Structure

```text
types/
├── cistring.go                  # Implementation with public/private functions
├── datetime.go                  # DateTime type
├── integer.go                   # Integer type
├── errors.go                    # Shared error constants
├── doc.go                       # Package documentation
├── cistring_test.go             # Tests for PRIVATE (package types)
├── example_cistring_test.go     # Examples for PUBLIC (package types_test)
├── example_datetime_test.go     # Examples for PUBLIC (package types_test)
├── example_integer_test.go      # Examples for PUBLIC (package types_test)
└── tests/
    ├── cistring_test.go         # Tests for PUBLIC API (package types_test)
    ├── datetime_test.go         # Tests for PUBLIC API (package types_test)
    ├── integer_test.go          # Tests for PUBLIC API (package types_test)
    └── errors_test.go           # Tests for error constants

authorize/
├── request.go                   # Authorize.req implementation
├── confirmation.go              # Authorize.conf implementation
├── errors.go                    # Package-level error constants
├── doc.go                       # Package documentation
├── example_request_test.go      # Examples for Req (package authorize_test)
├── example_confirmation_test.go # Examples for Conf (package authorize_test)
├── tests/
│   ├── request_test.go          # Tests for Req (package authorize_test)
│   └── confirmation_test.go     # Tests for Conf (package authorize_test)
└── types/
    ├── idtoken.go               # IdToken type
    ├── idtaginfo.go             # IdTagInfo type
    ├── authorizationstatus.go   # AuthorizationStatus enum
    ├── errors.go                # Type-level error constants
    ├── doc.go                   # Package documentation
    ├── example_idtoken_test.go              # Examples for NewIdToken
    ├── example_idtaginfo_test.go            # Examples for NewIdTagInfo
    ├── example_authorizationstatus_test.go  # Examples for AuthorizationStatus
    └── tests/
        ├── idtoken_test.go               # Tests for IdToken
        ├── idtaginfo_test.go             # Tests for IdTagInfo
        └── authorizationstatus_test.go   # Tests for AuthorizationStatus
```

**Key Guidelines:**

- If a package has ONLY public API (no private functions to test), skip
  same-package tests entirely
- ALL public API MUST have tests in the `tests/` directory using
  `package <name>_test`
- Same-package tests are the exception, not the rule - only create them when
  testing private implementation
- The `tests/` directory is the primary location for understanding how the
  package works from a user perspective

**Rationale:**

- Clear separation between internal implementation tests and public API tests
- External developers can focus on `tests/` directory to understand public API
  behavior
- Internal developers can find implementation tests alongside the code
- `package <name>_test` enforces black-box testing and prevents accidental
  coupling to internals
- Prevents accidental breaking of public API by making external-facing tests
  explicit

## Code Style Guidelines

### Imports

- Managed by gci: stdlib first, then github.com/aasanchez/ocpp16messages,
  then others
- No unused imports
- **Prefer full package names over aliases** in test files for readability:
  - GOOD: `authorize.Req(authorize.ReqInput{...})`
  - AVOID: `ma.Req(ma.Input{...})`
- Use short aliases only when:
  - Package name conflicts with another import
  - Import line exceeds 80 characters (revive line-length-limit)
- Common aliases when needed:
  - `st` for `github.com/aasanchez/ocpp16messages/types`
  - `at` for `github.com/aasanchez/ocpp16messages/authorize/types`
  - `ct` for `github.com/aasanchez/ocpp16messages/clearChargingProfile/types`
- **Line length constraint**: Import lines must be <= 80 characters. For long
  package paths like `clearChargingProfile/types`, use short 2-character aliases
  (e.g., `ct`) to stay within the limit.

### Formatting

- Enforced by gofumpt + golines
- Always run `make format` before committing

### Markdown Linting

- Run `markdownlint README.md CLAUDE.md SECURITY.md` after modifying any
  Markdown file
- Configuration in `.markdownlint.json`:
  - Line length: 80 characters max
  - Tables and code blocks: exempt from line length
- Fix all errors before committing Markdown changes

### Type Design

- Use precise types from types/ for OCPP fields
- Time parsing/formatting strictly uses RFC3339/RFC3339Nano (see DateTime type)
- All constructors/setters return `(T, error)` for validation

### Naming

- Exported identifiers: PascalCase
- Keep acronyms uppercase (e.g., Id not ID - allowed by revive var-naming rule)
- **Constructors**: MUST use `New` prefix following Go idioms
  - Public constructors: `NewType(args) (Type, error)` (e.g., `NewInteger`,
    `NewDateTime`, `NewCiString20Type`)
  - Private constructors: `newType(args) (type, error)` (lowercase 'n'
    for unexported)
  - **NEVER** use `Set` prefix for constructors
    (e.g., `SetInteger` BAD, `NewInteger` GOOD)
  - Rationale: "Set" implies mutation; "New" indicates
    construction(see Effective Go)
- **Getters**: Do NOT use `Get` prefix
  - Use `Value()` not `GetValue()` (see Effective Go getter guidelines)
- **Test naming**: Follow Go conventions
  - Unit tests: `Test_<package>_<Function>`
    (e.g., `Test_sharedtypes_NewInteger`)
  - Example tests: `Example<Function>` (e.g., `ExampleNewInteger`)
  - Subtests: Use descriptive suffixes
    (e.g., `Test_sharedtypes_NewInteger_Overflow`)

### Error Handling

- Never panic in library code
- Wrap errors with context: `fmt.Errorf("context: %w", err)`
- Provide actionable error messages
- **Error Constants and Reusability**:
  - ALWAYS check if an error constant or format string already exists before
    creating a new one
  - Use sentinel errors (`ErrEmptyValue`, `ErrInvalidValue`) from
    `types/errors.go` with `fmt.Errorf` and `ErrorFieldFormat`:

    ```go
    // Example: wrapping a sentinel error with field context
    return fmt.Errorf("NewIdToken: "+st.ErrorFieldFormat, "IdToken", st.ErrEmptyValue)
    ```

  - Prefer reusing error constants from:
    1. `types/errors.go` - For generic validation errors
       (ErrorMismatch, ErrorExpectedError, ErrorFieldFormat, ErrEmptyValue,
       ErrInvalidValue)
    2. Package-level `errors.go` - For package-specific errors
       (e.g., `authorize/types/errors.go`)
    3. Parent package `errors.go` - For message-level errors
       (e.g., `authorize/errors.go`)
  - When creating new error constants:
    - Place in `types/errors.go` if applicable across multiple packages
    - Place in package `errors.go` if specific to that package
    - Document the parameters and usage clearly
  - Avoid string literal duplication in tests - use constants from
    errors.go files
  - **Do NOT create helper functions** that just wrap `fmt.Errorf` - use
    `fmt.Errorf` directly with sentinel errors and format constants
- **Error Accumulation in Constructors**:
  - Message constructors (`Req()`, `Conf()`) SHOULD accumulate all validation
    errors and return them together using `errors.Join()`
  - This maximizes verbosity and efficiency by showing all issues at once
  - Pattern:

    ```go
    func Conf(input ConfInput) (ConfMessage, error) {
        var errs []error

        // Validate each field, accumulating errors
        if err := validateField1(); err != nil {
            errs = append(errs, fmt.Errorf("field1: %w", err))
        }
        if err := validateField2(); err != nil {
            errs = append(errs, fmt.Errorf("field2: %w", err))
        }

        // Return all accumulated errors at once
        if len(errs) > 0 {
            return ConfMessage{}, errors.Join(errs...)
        }

        return ConfMessage{...}, nil
    }
    ```

  - Benefits:
    - Users see all validation issues in one call, not one at a time
    - Reduces debugging cycles for malformed input
    - Each error is prefixed with the field name for clarity

### Testing

- **CRITICAL: Write atomic, individual test functions** - Each test should
  verify ONE specific behavior or case
  - GOOD: `TestIdToken_String_ReturnsValue()`,
    `TestIdToken_Empty_ReturnsError()`
  - BAD: Table-driven tests with multiple cases in one function
  - Rationale: Atomic tests are easier to debug, maintain, and understand.
    Each test failure points directly to the specific case that broke.
  - Exception: Only use table-driven tests when testing the same logic with
    slight variations (e.g., boundary values), and even then prefer splitting
    into separate functions when possible.
- Ensure all new tests are race-compatible
- Each test must call `t.Parallel()` for concurrent execution
- **Cognitive complexity**: Keep test functions simple. If revive reports
  cognitive complexity > 7, split the test into smaller focused tests.
- **Magic numbers**: Use named constants for test values instead of raw numbers
  (e.g., `valueZero`, `valueOne`, `valueNegative`, `valueExceedsMax`).
- **Variable names**: Avoid short variable names like `p`, `id`. Use descriptive
  names like `purposeType`, `profileId` to pass varnamelen linter.

#### Example Tests - Use Selectively

Example tests serve as executable documentation and appear in `go doc` and
pkgsite. **Use examples strategically**, not for every function.

**CREATE examples for:**

- **Public constructors** - Primary API entry points (e.g., `Req`, `Conf`,
  `NewIdToken`, `NewIdTagInfo`)
  - Place in `example_*_test.go` in the same package
  - Show both success cases and common error scenarios
- **Complex public APIs** - Where usage patterns aren't immediately obvious
- **Package-level functions** - Functions that users will call directly from
  outside the package

**SKIP examples for:**

- **Error Management files** - files related to consolidation of error messages,
  and not need it.
- **Simple getter methods** - Methods like `Value()`, `String()` are
  self-explanatory
- **Internal/private types** - Types and methods not exported or used outside
  the package
- **Implementation details** - Methods like `Validate()` that are called
  internally, not by end users
- **Obvious behavior** - If unit test names clearly describe the behavior

**Rationale:**

- Examples for **public APIs**, unit tests for **implementation**
- Good developers can understand simple methods from clear unit test names
- Over-documenting creates maintenance burden without adding value
- Focus examples on the actual developer workflow, not internal mechanics

**Example organization:**

```text
authorize/
├── example_request_test.go      # Examples for Req (public API)
├── example_confirmation_test.go # Examples for Conf (public API)
└── types/
    ├── example_idtoken_test.go  # Examples for NewIdToken (public API)
    └── tests/
        └── idtoken_test.go      # Unit tests for IdToken (public API)
```

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

### Linter Rules

Key linter rules enforced by golangci-lint:

- **line-length-limit**: 80 characters max (revive)
- **cognitive-complexity**: max 7 (revive)
- **varnamelen**: variable names must be descriptive
- **wsl_v5**: whitespace linter for consistent formatting
- **exhaustruct**: all struct fields must be explicitly initialized

## OCPP Compliance Notes

- All DateTime fields normalized to UTC per OCPP 1.6 spec
- CiString types enforce ASCII printable characters only (32-126)
- All types proven thread-safe via race tests
- String length validation enforced at construction time
