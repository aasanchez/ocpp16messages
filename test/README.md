# test/

This directory contains the unit tests for the `ocpp16_messages` package. The tests ensure correct functionality, high coverage, and compliance with the OCPP 1.6J specification.

## Contents

- `ocpp16j_authorize_full_test.go`: Full end-to-end validation for `Authorize.req` and `Authorize.conf` messages using raw OCPP 1.6J array-style JSON.
- `core_errors_test.go`: Tests for structured error types, including field-specific errors and standard call errors.
- `core_plugin_test.go`: Tests for plugin registration, custom validators, and pre/post validation hooks.
- `core_types_test.go`: Unit tests for core types such as `CiString` and `IdToken`.

## Notes

- All tests are written using idiomatic Go testing patterns.
- Coverage goal is **100%** for all critical paths.
- Integration of testable benchmarks and profiling is handled under the `benchmark/` directory.

## How to Run

```bash
go test ./test -v
```

To run with coverage:

```bash
go test -coverprofile=coverage.out ./test
go tool cover -html=coverage.out
```

Contribution
If you are adding a new message type or extending validation logic, corresponding tests should be added here.
