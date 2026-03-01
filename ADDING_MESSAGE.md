# Adding a new message type

This repository implements OCPP 1.6 messages using a strict "constructor +
validation" style:

- Public entry points are `Req(input)` and `Conf(input)`.
- Inputs use `ReqInput` / `ConfInput` with raw Go values.
- Constructors validate everything and return a typed message or an error.
- Constructors should accumulate field errors with `errors.Join`.
- Returned values must not alias caller-owned pointers/slices.

If you add a new OCPP operation package, follow this template.

## Minimal package template

Directory layout:

```text
<operation>/
├── doc.go
├── request.go
├── confirmation.go
├── example_request_test.go
├── example_confirmation_test.go
└── tests/
    ├── request_test.go
    └── confirmation_test.go
```

If the operation has message-specific enums or types, add:

```text
<operation>/types/
├── doc.go
├── <type>.go
├── example_<type>_test.go
└── tests/
    └── <type>_test.go
```

## Code template (Req / Conf)

`request.go`:

```go
package <operation>

import (
  "errors"
  "fmt"

  st "github.com/aasanchez/ocpp16messages/types"
)

type ReqInput struct {
  // Raw inputs (string/int/*int/[]T/etc)
}

type ReqMessage struct {
  // Validated, typed fields (may be exported; treat as read-only)
}

func Req(input ReqInput) (ReqMessage, error) {
  var errs []error

  // Validate each field, accumulating errors:
  //
  // value, err := st.NewCiString20Type(input.IdTag)
  // if err != nil {
  //   errs = append(errs, fmt.Errorf("IdTag: %w", err))
  // }

  if len(errs) > 0 {
    return ReqMessage{}, errors.Join(errs...)
  }

  return ReqMessage{
    // Set every field explicitly (exhaustruct)
  }, nil
}
```

`confirmation.go` mirrors the same pattern with `ConfInput`, `ConfMessage`,
and `Conf(input)`.

## Immutability / aliasing checklist

In constructors and getters:

- Copy any pointer input before storing it.
- Copy any slice input before storing it.
- If returning pointers or slices, return defensive copies.

Add or update race tests in `./tests_race` when you change any of these rules.

## Test template

Public API tests live in `tests/` and must use `package <operation>_test`.

Keep tests atomic:

- One behavior per test function.
- `t.Parallel()` in every test.

Example:

```go
package <operation>_test

import (
  "testing"

  "github.com/aasanchez/ocpp16messages/<operation>"
)

func Test_<operation>_Req_Valid(t *testing.T) {
  t.Parallel()

  _, err := <operation>.Req(<operation>.ReqInput{
    // valid minimal input
  })
  if err != nil {
    t.Fatalf("expected nil error, got %v", err)
  }
}
```

## Fuzz / race / bench (high-value expectations)

When adding a new message, also add:

- Fuzzers for `Req` and `Conf` in `./tests_fuzz` (`//go:build fuzz`).
- Race/immutability tests in `./tests_race` (`//go:build race`).
- Benchmarks for the constructor on worst-case payloads in `./tests_benchmark`
  (`//go:build bench`), when the message is complex or frequently used.
