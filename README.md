# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16messages)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=coverage)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)

---

## Overview

This project provides a robust, modular, and thoroughly tested implementation of OCPP 1.6 message types and payload validation for Go. It is designed for high-quality message serialization, strict type safety, and OCPP-compliant EV charging station management.

### Key Features

- **OCPP 1.6-compliant types**: Strict, validated types for all core OCPP 1.6 messages and fields.
- **Extensive testing**: Includes unit, example, benchmark, fuzz, and race condition tests.
- **High code quality**: Linting, formatting, static analysis, and CI integration.
- **Performance and safety**: Benchmarks for all critical routines, and race/fuzz tests for concurrency and robustness.
- **Idiomatic Go**: Clear error handling, modular structure, and full documentation.

---

## Project Structure

- `shared/types/` — Core types (e.g., `DateTime`, `Integer`, `CiString*`) with validation and OCPP-compliant formatting.
- `shared/types/tests/` — Unit and example tests for all types.
- `shared/types/benchmark/` — Microbenchmarks for performance-critical routines.
- `shared/types/fuzz/` — Fuzz tests for robustness and edge-case discovery.
- `shared/types/race/` — Race condition tests for thread safety under heavy concurrency.
- `.github/workflows/` — CI configuration for build, lint, and test automation.

---

## Getting Started

### Install dependencies

```sh
go mod tidy
```

### Run all tests (unit, example, race, fuzz, benchmark)

```sh
make test           # Unit and example tests
make test-race      # Race detector (thread safety)
make test-fuzz      # Fuzz tests (robustness, edge cases)
make test-benchmark # Benchmarks (performance)
```

### Lint, format, and static analysis

```sh
make lint
make format
```

### Open coverage report

```sh
make test-coverage
```

---

## OCPP Compliance

- All date/time fields are normalized to UTC and formatted per RFC3339, as required by the OCPP 1.6 specification.
- All string types (CiString*) are strictly validated for length and ASCII printability.
- All types are proven safe for concurrent use via race tests.

---

## Contributing

- Ensure full test, fuzz, and benchmark coverage for new logic.
- Use idiomatic Go and document all public types/functions.
- Follow the project's code style and conventions.
- Run `make test-all` before submitting a PR.

---

## License

See [LICENSE](./LICENSE)
