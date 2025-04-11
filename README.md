# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aasanchez/ocpp16messages)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16messages/branch/main/graph/badge.svg?token=1I9VVL7DWO)](https://codecov.io/gh/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16_messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16_messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=bugs)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)

---

## Overview

This project provides a robust, modular implementation of OCPP 1.6 message types and payload validation for Go. It supports high-quality message serialization, validation, and extensibility for OCPP-compliant EV charging station management.

### Features

- Complete set of OCPP 1.6 message domains, with a focus on `authorize` and `bootnotification`
- Strict type definitions and validation for all payloads
- Extensive unit and example tests for correctness
- Benchmark suite for key serialization and validation routines
- Idiomatic Go, clear error handling and modular package structure
- Tooling for linting, formatting, and static analysis

### Project Structure

- `messages/authorize`: Authorize message types & validation
- `messages/bootnotification`: BootNotification message types
- `shared/types`: Shared string, datetime and error types
- Extensive `types` subfolders and benchmark folders for fine-grained coverage

### Getting Started

Install dependencies and run all tests:

```sh
go mod tidy
go test ./...
```

Run a single test or file:

```sh
go test ./messages/authorize/types/requestPayload_test.go
go test -run '^TestAuthorizationStatus_validAccepted$' ./...
```

Run benchmarks:

```sh
make benchmark
```

Lint and format:

```sh
make lint
make format
```

### Contributing

- Follow code standards as documented in [`OpenCode.md`](./OpenCode.md)
- Ensure full test and benchmark coverage for new logic
- Use idiomatic Go and document public types/functions

---

## License

See [LICENSE](./LICENSE)
