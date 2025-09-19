# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=bugs)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=coverage)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)

---

## Overview

This project provides a robust, modular implementation of OCPP 1.6 message types and payload validation for Go. It supports high-quality message serialization, validation, and extensibility for OCPP-compliant EV charging station management.

### Features

- A growing set of OCPP 1.6 message types.
- Strict type definitions and validation for all payloads.
- Extensive unit and example tests for correctness.
- Benchmark suite for key serialization and validation routines.
- Idiomatic Go, clear error handling and modular package structure.
- Tooling for linting, formatting, and static analysis.

### Project Structure

The project is organized by message type, with shared utilities in a separate package.

- `messages/authorize`: Handles authorization requests and confirmations.
- `messages/bootnotification`: Manages boot notifications from charging stations.
- `messages/cancelreservation`: Implements the cancellation of reservations.
- `shared/types`: Contains common data types (e.g., for strings and datetimes) used across all messages.

### Getting Started

Install dependencies and run all tests:

```sh
go mod tidy
make test
```

Run benchmarks:

```sh
make benchmark
```

Lint and format the code:

```sh
make lint
make format
```

### Contributing

- Ensure full test and benchmark coverage for new logic.
- Use idiomatic Go and document public types/functions.
- Follow the project's code style and conventions.

---

## License

See [LICENSE](./LICENSE)
