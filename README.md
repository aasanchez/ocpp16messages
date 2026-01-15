# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16messages)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16messages/graph/badge.svg?token=1I9VVL7DWO)](https://codecov.io/gh/aasanchez/ocpp16messages)

A type-safe, OCPP 1.6-compliant Go library providing message structures and
validation for Electric Vehicle charging station communication.

## Overview

This library implements OCPP (Open Charge Point Protocol) 1.6 message types
with strict validation, following Go best practices and the official OCPP 1.6
specification. It is designed as a foundation for building OCPP-compliant
charging station management systems and charge point implementations.

**Status:** Active Development (Pre-1.0)

### Key Features

- **Type Safety** - Constructor pattern with validation
  (`New*()` for types, `Req()`/`Conf()` for messages)
- **OCPP 1.6 Compliance** - Strict adherence to protocol specification
- **OCPP Naming** - Uses `Req()`/`Conf()` to match OCPP terminology
  (Authorize.req, Authorize.conf)
- **Immutable Types** - Thread-safe by design with value receivers
- **Comprehensive Testing** - Unit tests and example tests with 100% coverage
- **Zero Panics** - All errors returned, never panicked
- **Well Documented** - Full godoc coverage and examples

## Installation

```bash
go get github.com/aasanchez/ocpp16messages
```

**Requirements:** Go 1.24 or later

## Project Structure

```text
.
├── types/                      # Core OCPP data types
│   ├── cistring.go             # CiString20/25/50/255/500 types
│   ├── datetime.go             # RFC3339 DateTime with UTC normalization
│   ├── integer.go              # Validated uint16 Integer type
│   ├── errors.go               # Shared error constants and sentinels
│   ├── doc.go                  # Package documentation
│   └── tests/                  # Public API tests (black-box)
├── authorize/                  # Authorize message (implemented)
│   ├── request.go              # Authorize.req (Req constructor)
│   ├── confirmation.go         # Authorize.conf (Conf constructor)
│   ├── errors.go               # Package-level error constants
│   ├── doc.go                  # Package documentation
│   ├── tests/                  # Public API tests
│   └── types/                  # Authorize-specific types
│       ├── idtoken.go          # IdToken type
│       ├── idtaginfo.go        # IdTagInfo type
│       ├── authorizationstatus.go
│       └── tests/              # Type tests
├── bootNotification/           # BootNotification message (implemented)
├── cancelReservation/          # CancelReservation message (implemented)
├── changeAvailability/         # ChangeAvailability message (implemented)
├── changeConfiguration/        # ChangeConfiguration message (implemented)
├── clearCache/                 # ClearCache message (implemented)
├── clearChargingProfile/       # ClearChargingProfile message (implemented)
├── dataTransfer/               # DataTransfer message (stub)
├── diagnosticsStatusNotification/  # DiagnosticsStatusNotification (stub)
├── firmwareStatusNotification/ # FirmwareStatusNotification (stub)
├── getCompositeSchedule/       # GetCompositeSchedule (stub)
├── getConfiguration/           # GetConfiguration (stub)
├── getDiagnostics/             # GetDiagnostics (stub)
├── getLocalListVersion/        # GetLocalListVersion (stub)
├── heartbeat/                  # Heartbeat (stub)
├── meterValues/                # MeterValues (stub)
├── remoteStartTransaction/     # RemoteStartTransaction (stub)
├── remoteStopTransaction/      # RemoteStopTransaction (stub)
├── reserveNow/                 # ReserveNow (stub)
├── reset/                      # Reset (stub)
├── sendLocalList/              # SendLocalList (stub)
├── setChargingProfile/         # SetChargingProfile (stub)
├── startTransaction/           # StartTransaction (stub)
├── statusNotification/         # StatusNotification (stub)
├── stopTransaction/            # StopTransaction (stub)
├── triggerMessage/             # TriggerMessage (stub)
├── unlockConnector/            # UnlockConnector (stub)
├── updateFirmware/             # UpdateFirmware (stub)
└── SECURITY.md                 # Security policy and vulnerability reporting
```

## Usage

### Core Types

The library provides validated OCPP 1.6 data types:

```go
import "github.com/aasanchez/ocpp16messages/types"

// CiString types (case-insensitive, ASCII printable, length-validated)
idTag, err := types.NewCiString20Type("RFID-ABC123")
if err != nil {
    // Handle validation error (length > 20 or non-ASCII chars)
}

// DateTime (RFC3339, auto-normalized to UTC)
timestamp, err := types.NewDateTime("2025-01-02T15:04:05Z")
if err != nil {
    // Handle parsing error
}

// Integer (validated uint16)
retryCount, err := types.NewInteger("3")
if err != nil {
    // Handle conversion/range error
}
```

### Message Types

Messages use OCPP terminology with `Req()` for requests and `Conf()` for responses:

```go
import "github.com/aasanchez/ocpp16messages/authorize"

// Create an Authorize.req message using the ReqInput struct
// Validation happens automatically in the constructor
req, err := authorize.Req(authorize.ReqInput{
    IdTag: "RFID-ABC123",
})
if err != nil {
    // Handle validation error (empty, too long, or invalid characters)
}

// Access the validated IdTag
fmt.Println(req.IdTag.String()) // "RFID-ABC123"
```

```go
import "github.com/aasanchez/ocpp16messages/clearChargingProfile"

// ClearChargingProfile.req with optional fields
id := 123
req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
    Id:                     &id,
    ConnectorId:            nil,
    ChargingProfilePurpose: nil,
    StackLevel:             nil,
})
```

The `ReqMessage` type returned by `Req()` contains validated, typed fields
that are immutable and thread-safe.

## Development

### Prerequisites

- Go 1.24+
- golangci-lint
- staticcheck
- gci, gofumpt, golines (formatters)

### Common Commands

```bash
# Install dependencies
go mod tidy

# Run tests
make test                   # Unit tests with coverage
make test-coverage          # Generate HTML coverage report
make test-example           # Run example tests (documentation tests)
make test-all               # Run all test types

# Code quality
make lint                   # Run all linters (golangci-lint, go vet, staticcheck)
make format                 # Format code (gci, gofumpt, golines, gofmt)

# Documentation
make pkgsite                # Start local documentation server at http://localhost:8080
```

### Test Coverage

Reports are generated in the `reports/` directory:

- `reports/coverage.out` - Coverage data
- `reports/golangci-lint.txt` - Lint results

## OCPP 1.6 Compliance

### Type System

| OCPP Type       | Go Type                 | Validation                              |
|-----------------|-------------------------|-----------------------------------------|
| CiString20Type  | `types.CiString20Type`  | Length <= 20, ASCII printable (32-126)  |
| CiString25Type  | `types.CiString25Type`  | Length <= 25, ASCII printable (32-126)  |
| CiString50Type  | `types.CiString50Type`  | Length <= 50, ASCII printable (32-126)  |
| CiString255Type | `types.CiString255Type` | Length <= 255, ASCII printable (32-126) |
| CiString500Type | `types.CiString500Type` | Length <= 500, ASCII printable (32-126) |
| dateTime        | `types.DateTime`        | RFC3339, normalized to UTC              |
| integer         | `types.Integer`         | uint16 (0-65535)                        |

### Message Implementation Status

| Message                       | Request | Confirmation | Package                         |
|-------------------------------|---------|--------------|---------------------------------|
| Authorize                     | Done    | Done         | `authorize`                     |
| BootNotification              | Done    | Done         | `bootNotification`              |
| CancelReservation             | Done    | Done         | `cancelReservation`             |
| ChangeAvailability            | Done    | Done         | `changeAvailability`            |
| ChangeConfiguration           | Done    | Done         | `changeConfiguration`           |
| ClearCache                    | Done    | Done         | `clearCache`                    |
| ClearChargingProfile          | Done    | Done         | `clearChargingProfile`          |
| DataTransfer                  | Planned | Planned      | `dataTransfer`                  |
| DiagnosticsStatusNotification | Planned | Planned      | `diagnosticsStatusNotification` |
| FirmwareStatusNotification    | Planned | Planned      | `firmwareStatusNotification`    |
| GetCompositeSchedule          | Planned | Planned      | `getCompositeSchedule`          |
| GetConfiguration              | Planned | Planned      | `getConfiguration`              |
| GetDiagnostics                | Planned | Planned      | `getDiagnostics`                |
| GetLocalListVersion           | Planned | Planned      | `getLocalListVersion`           |
| Heartbeat                     | Planned | Planned      | `heartbeat`                     |
| MeterValues                   | Planned | Planned      | `meterValues`                   |
| RemoteStartTransaction        | Planned | Planned      | `remoteStartTransaction`        |
| RemoteStopTransaction         | Planned | Planned      | `remoteStopTransaction`         |
| ReserveNow                    | Planned | Planned      | `reserveNow`                    |
| Reset                         | Planned | Planned      | `reset`                         |
| SendLocalList                 | Planned | Planned      | `sendLocalList`                 |
| SetChargingProfile            | Planned | Planned      | `setChargingProfile`            |
| StartTransaction              | Planned | Planned      | `startTransaction`              |
| StatusNotification            | Planned | Planned      | `statusNotification`            |
| StopTransaction               | Planned | Planned      | `stopTransaction`               |
| TriggerMessage                | Planned | Planned      | `triggerMessage`                |
| UnlockConnector               | Planned | Planned      | `unlockConnector`               |
| UpdateFirmware                | Planned | Planned      | `updateFirmware`                |

### Design Principles

1. **OCPP Naming** - Messages use `Req()`/`Conf()` to match OCPP terminology
2. **Constructor Validation** - All types require constructors that validate input
3. **Input Struct Pattern** - Raw values passed via `ReqInput`/`ConfInput`
   structs, validated automatically
4. **Immutability** - Types use private fields and value receivers
5. **Error Accumulation** - Constructors report all validation errors at once
   using `errors.Join()`
6. **Error Wrapping** - Context preserved via `fmt.Errorf` with `%w`
7. **No Panics** - Library never panics; all errors returned
8. **Thread Safety** - Designed for safe concurrent use
9. **Go Conventions** - Follows [Effective Go](https://go.dev/doc/effective_go)
   guidelines

## Security

Security is critical for EV charging infrastructure. This library:

- Validates all input at construction time
- Prevents injection attacks via strict type constraints
- Provides clear error messages without exposing internals
- Uses immutable types to prevent tampering
- Is designed for safe concurrent use

**Reporting vulnerabilities:** See [SECURITY.md](SECURITY.md) for our security
policy and responsible disclosure process.

## Contributing

We welcome contributions! Please:

1. Follow Go best practices and [Effective Go](https://go.dev/doc/effective_go)
2. Add tests for all new functionality
3. Ensure `make test-all` passes
4. Run `make lint` and `make format` before committing
5. Document all exported types and functions
6. Follow the existing code style

See [CLAUDE.md](CLAUDE.md) for detailed development guidelines.

## License

See [LICENSE](./LICENSE)

## Resources

- [OCPP 1.6 Specification](https://www.openchargealliance.org/protocols/ocpp-16/)
- [Go Package Documentation](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
- [Security Policy](SECURITY.md)
- [Development Guide](CLAUDE.md)
