# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16messages)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16messages/graph/badge.svg?token=1I9VVL7DWO)](https://codecov.io/gh/aasanchez/ocpp16messages)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/aasanchez/ocpp16messages/badge)](https://securityscorecards.dev/viewer/?uri=github.com/aasanchez/ocpp16messages)
[![Supported Go versions](https://img.shields.io/badge/Go-1.25.8%20%7C%201.26.1-00ADD8?logo=go)](#versioning-and-support)
[![Compatibility contract](https://img.shields.io/badge/compatibility-contract-blue)](COMPATIBILITY.md)
[![CLA assistant](https://cla-assistant.io/readme/badge/aasanchez/ocpp16messages)](https://cla-assistant.io/aasanchez/ocpp16messages)

A type-safe Go implementation of the Open Charge Point Protocol (OCPP) 1.6
message set, with validated request/response structures for EV charge points
(EVSE) and Central Systems (CSMS/backends).

## Overview

This library implements OCPP (Open Charge Point Protocol) 1.6 message types
with strict validation, following Go best practices and the official OCPP 1.6
specification. It is designed as a foundation for building OCPP-compliant
charging station management systems and charge point implementations.

**Status:** Stable - v1.0.0 (28/28 OCPP 1.6 JSON messages implemented)

    - API follows SemVer; see [CHANGELOG](CHANGELOG.md) for release notes.
    - Public API is frozen for all v1.x releases; breaking changes will bump
        MAJOR.

**Search terms**: OCPP 1.6, Open Charge Point Protocol, EVSE, CSMS, charge
station backend, Authorize.req, BootNotification, MeterValues, RemoteStart,
Charge Point, Central System.

### Key Features

- **Type Safety** - Constructor pattern with validation
  (`New*()` for types, `Req()`/`Conf()` for messages)
- **OCPP 1.6 Compliance** - Strict adherence to protocol specification
- **OCPP Naming** - Uses `Req()`/`Conf()` to match OCPP terminology
  (Authorize.req, Authorize.conf)
- **Immutable Types** - Thread-safe by design with value receivers
- **Comprehensive Testing** - Unit tests and example tests with high coverage
- **Zero Panics** - All errors returned, never panicked
- **Well Documented** - Full godoc coverage and examples

### Implemented OCPP 1.6 operations

The library covers the full OCPP 1.6 message surface, including:

- Authorize.req / .conf, BootNotification.req / .conf, Heartbeat.req / .conf
- StartTransaction.req / .conf, StopTransaction.req / .conf, StatusNotification
- MeterValues, ClearChargingProfile, SetChargingProfile, TriggerMessage
- RemoteStartTransaction, RemoteStopTransaction, ChangeAvailability,
  ChangeConfiguration, GetConfiguration, DataTransfer
- ReserveNow, CancelReservation, UnlockConnector, Reset, UpdateFirmware,
  GetDiagnostics, DiagnosticsStatusNotification, FirmwareStatusNotification
- SendLocalList, GetLocalListVersion, GetCompositeSchedule

## Installation

    go get github.com/aasanchez/ocpp16messages

**Requirements:** Go 1.25.8 or later (CI and go.mod aligned)

## Project Structure

    .
    ├── types/                      # Core OCPP data types (shared across messages)
    │   ├── cistring.go             # CiString20/25/50/255/500 types
    │   ├── datetime.go             # RFC3339 DateTime with UTC normalization
    │   ├── integer.go              # Validated uint16 Integer type
    │   ├── errors.go               # Shared error constants and sentinels
    │   ├── authorizationstatus.go  # AuthorizationStatus enum
    │   ├── idtoken.go              # IdToken type
    │   ├── idtaginfo.go            # IdTagInfo type
    │   ├── chargingprofilepurposetype.go  # ChargingProfilePurposeType enum
    │   ├── chargingrateunit.go     # ChargingRateUnit enum
    │   ├── chargingschedule.go     # ChargingSchedule type
    │   ├── chargingscheduleperiod.go  # ChargingSchedulePeriod type
    │   ├── metervalue.go           # MeterValue and MeterValueInput types
    │   ├── sampledvalue.go         # SampledValue and SampledValueInput types
    │   ├── measurand.go            # Measurand enum
    │   ├── readingcontext.go       # ReadingContext enum
    │   ├── valueformat.go          # ValueFormat enum
    │   ├── phase.go                # Phase enum
    │   ├── location.go             # Location enum
    │   ├── unitofmeasure.go        # UnitOfMeasure enum
    │   ├── doc.go                  # Package documentation
    │   └── tests/                  # Public API tests (black-box)
    ├── authorize/                       # Authorize message
    ├── bootnotification/                # BootNotification message
    ├── cancelreservation/               # CancelReservation message
    ├── changeavailability/              # ChangeAvailability message
    ├── changeconfiguration/             # ChangeConfiguration message
    ├── clearcache/                      # ClearCache message
    ├── clearchargingprofile/            # ClearChargingProfile message
    ├── datatransfer/                    # DataTransfer message
    ├── diagnosticsstatusnotification/   # DiagnosticsStatusNotification message
    ├── firmwarestatusnotification/      # FirmwareStatusNotification message
    ├── getcompositeschedule/            # GetCompositeSchedule message
    ├── getconfiguration/                # GetConfiguration message
    ├── getdiagnostics/                  # GetDiagnostics message
    ├── getlocallistversion/             # GetLocalListVersion message
    ├── heartbeat/                       # Heartbeat message
    ├── metervalues/                     # MeterValues message
    ├── remotestarttransaction/          # RemoteStartTransaction message
    ├── remotestoptransaction/           # RemoteStopTransaction message
    ├── reservenow/                      # ReserveNow message
    ├── reset/                           # Reset message
    ├── sendlocallist/                   # SendLocalList message
    ├── setchargingprofile/              # SetChargingProfile message
    ├── starttransaction/                # StartTransaction message
    ├── statusnotification/              # StatusNotification message
    ├── stoptransaction/                 # StopTransaction message
    ├── triggermessage/                  # TriggerMessage message
    ├── unlockconnector/                 # UnlockConnector message
    ├── updatefirmware/                  # UpdateFirmware message
    └── SECURITY.md                 # Security policy and vulnerability reporting

## Versioning and support

- Semantic Versioning: API surface follows SemVer starting with v1.0.0.
- Compatibility contract: see [COMPATIBILITY.md](COMPATIBILITY.md) for exact
  SemVer guarantees, deprecation policy, and what counts as breaking.
- Supported Go versions: >= 1.25 (aligned with go.mod and CI).
- Changelog: see [CHANGELOG](CHANGELOG.md) for releases and upgrade notes.

## Usage

### Core Types

The library provides validated OCPP 1.6 data types:

    import "github.com/aasanchez/ocpp16messages/types"

    // CiString types (case-insensitive, ASCII printable, length-validated)
    idTag, err := types.NewCiString20Type("RFID-ABC123")
    if err != nil {
        // Handle validation error (length > 20 or non-ASCII chars)
    }

    // DateTime (RFC3339, must be UTC)
    timestamp, err := types.NewDateTime("2025-01-02T15:04:05Z")
    if err != nil {
        // Handle parsing error
    }

    // Integer (validated uint16)
    retryCount, err := types.NewInteger(3)
    if err != nil {
        // Handle conversion/range error
    }

### Message Types

Messages use OCPP terminology with `Req()` for requests and `Conf()` for responses:

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

    import "github.com/aasanchez/ocpp16messages/clearchargingprofile"

    // ClearChargingProfile.req with optional fields
    id := 123
    req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
        Id:                     &id,
        ConnectorId:            nil,
        ChargingProfilePurpose: nil,
        StackLevel:             nil,
    })

The `ReqMessage` type returned by `Req()` contains validated, typed fields.
Core value types in `types/` are immutable and thread-safe. Message structs
are safe to share between goroutines **as long as they are treated as
read-only** (they have exported fields, so consumers can mutate them).

### Error contract

This library aims to provide stable error identities and flexible error
messages.

- **Stable:** use `errors.Is(err, types.ErrEmptyValue)` and
  `errors.Is(err, types.ErrInvalidValue)` to detect validation failures.
- **Stable:** constructors may return an aggregated error using `errors.Join`;
  callers should use `errors.Is` rather than string matching.
- **Not stable:** exact error strings and formatting (including joined error
  order) may change between releases.

### Common validation errors

All validation happens at construction time. A non-nil error always means the
value/message was not constructed.

Detect the common failure modes via `errors.Is`:

    import (
        "errors"
        "fmt"

        "github.com/aasanchez/ocpp16messages/authorize"
        "github.com/aasanchez/ocpp16messages/types"
    )

    req, err := authorize.Req(authorize.ReqInput{
        IdTag: "",
    })
    if err != nil {
        switch {
        case errors.Is(err, types.ErrEmptyValue):
            fmt.Println("missing required field")
        case errors.Is(err, types.ErrInvalidValue):
            fmt.Println("invalid value (length, charset, range, enum, etc.)")
        default:
            fmt.Println("unexpected error:", err)
        }
    }
    _ = req

Notes:

- Many constructors aggregate multiple field errors with `errors.Join(...)`.
  `errors.Is` still works against joined errors.
- Avoid depending on exact error strings; they may change as long as
  `errors.Is` behavior remains stable.

## Development

### Prerequisites

- Go 1.25+
- golangci-lint
- staticcheck
- gci, gofumpt, golines (formatters)

### Private modules (GOPRIVATE / GONOSUMDB)

If you work in an environment with private forks or private module dependencies,
configure the Go toolchain so it does not use the public checksum database for
those module paths:

    export GOPRIVATE=github.com/<your-org>/*
    export GONOSUMDB=$GOPRIVATE
    export GONOPROXY=$GOPRIVATE

### FAQ

- **Why does `NewDateTime` reject `+02:00` timestamps?**
  - OCPP 1.6 requires UTC. This library enforces UTC-only dateTimes, so provide
    values like `2025-01-02T15:04:05Z`.
- **Nil vs empty slices: what's the difference?**
  - Nil means "field omitted"; empty means "field present but empty". Message
    constructors preserve this distinction.
- **Are messages immutable?**
  - Core types in `types/` are immutable. Message structs have exported fields,
    so they are concurrency-safe only when treated as read-only. See
    `ROADMAP.md` for the v2 plan to make messages structurally immutable.

### Testing philosophy

This repository uses a layered test strategy:

- **Unit tests (default)**: fast, deterministic, run on every CI push/PR via
  `go test ./...` and `make test`.
- **Example tests (default)**: executable docs that demonstrate intended usage
  (`go test -run '^Example' ./...` and `make test-example`).
- **Fuzz tests (opt-in)**: invariants and validation hardening under
  `./tests_fuzz` with build tag `fuzz` (`make test-fuzz`).
- **Race tests (opt-in)**: immutability/aliasing and concurrency guarantees
  under `./tests_race` with build tag `race` (`make test-race`).
- **Benchmarks (opt-in)**: performance regression guards under
  `./tests_benchmark` with build tag `bench` (`make test-bench`).

Opt-in suites are not part of default `go test ./...` by design; they are
heavier and run on a weekly schedule in CI (and on-demand locally).

### Performance discipline

Benchmarks are opt-in and intended to catch regressions in high-value
constructors and worst-case payloads.

Local workflow:

    make test-bench
    go install golang.org/x/perf/cmd/benchstat@latest
    benchstat old.txt new.txt

On releases (tag pushes), CI generates a `benchstat` comparison against the
previous tag and attaches it to the GitHub release assets.

### Adding a new message type

See `ADDING_MESSAGE.md` for a minimal, copy/paste-friendly template that
matches the project's constructor + validation style and test organization.

### Common Commands (including opt-in suites)

    # Install dependencies
    go mod tidy

    # Run tests
    make test                   # Unit tests with coverage
    make test-coverage          # Generate HTML coverage report
    make test-example           # Run example tests (documentation tests)
    make test-all               # Run all test types
    make test-race              # Run race detector with -race (opt-in)
    make test-fuzz              # Run fuzzers in ./tests_fuzz (short budget, opt-in)
    make test-bench             # Run benchmarks in ./tests_benchmark (opt-in)

    # Code quality
    make lint                   # Run all linters (golangci-lint, go vet, staticcheck)
    make format                 # Format code (gci, gofumpt, golines, gofmt)

    # Documentation
    make pkgsite                # Start local documentation server at http://localhost:8080

### Concurrency rules (and required race tests)

This library is designed for safe concurrent use when values are treated as
immutable. To keep this guarantee strong over time, any new public API or
constructor changes must follow these rules.

#### Immutability / aliasing rules

- Never store pointers to caller-owned variables in returned values.
  - Example: if an input field is `*string`, copy `*input` into a new local
    variable and store a pointer to the copy.
- Never expose internal slices directly.
  - If a getter returns a slice, return a copy of the slice header/data.
  - If a constructor stores a slice, allocate a new slice and copy elements.
- Never expose internal pointers directly.
  - If a getter returns a `*T`, return a pointer to a copy of `T`.

These patterns prevent accidental mutation of internal state and eliminate
real race hazards (e.g., aliasing an input pointer, then the caller mutating
that variable concurrently with reads of the returned message/type).

#### Race tests (required for new high-value code)

Race tests live in `./tests_race` behind the `race` build tag (`//go:build race`).
They are not part of default `go test ./...`.

- Add/update race tests for:
  - Every new `Req()` / `Conf()` constructor.
  - Every new exported `New*` constructor.
  - Every new exported getter that returns a pointer or a slice.
- Race tests should be high-signal:
  - Use shared inputs across goroutines to catch unintended mutation.
  - Include "immutability" tests that mutate a returned pointer/slice and
    assert the original value is unchanged.
  - Do not call `t.Fatal*` from goroutines. Use the shared concurrency helper
    in `race/helpers_test.go` (`runConcurrent`).

Run locally via:

    make test-race

#### TODO (future major redesign)

- TODO(v2): Redesign `ReqMessage` / `ConfMessage` structs to be truly immutable
  (unexported fields + getters and/or deep-copy semantics). Today, message
  structs have exported fields, so they are concurrency-safe only when treated
  as read-only.

### Fuzz testing

Fuzzers live in `./tests_fuzz` and are guarded by the `fuzz` build tag, so they
do not run as part of `go test ./...`.

Run fuzzers via the project helper:

    make test-fuzz

Tuning knobs:

- `FUZZTIME` (default `5s`): per-fuzzer time budget.
- `FUZZPROCS` (default `4`): caps fuzz worker parallelism via `GOMAXPROCS`.

To run a single fuzzer directly:

    go test -tags=fuzz -run=^$ -fuzz=^FuzzAuthorizeReq$ -fuzztime=10s ./tests_fuzz

#### What is covered

The fuzz suite is intentionally "high-scrutiny" (high-signal):

- **All message constructors**: every OCPP operation package has fuzzers for
  both `Req()` and `Conf()` (including empty constructors that must always
  succeed).
- **Core validation types**: fuzzers exist for constructors like `NewDateTime`,
  `NewInteger`, all `NewCiString*Type` variants, `NewChargingSchedule` and
  `NewChargingSchedulePeriod`, `NewIdTagInfo`, `NewSampledValue`, and
  `NewMeterValue` (plus their message-scoped equivalents).
- **Enum correctness**: strict membership fuzzers validate that every
  `IsValid()` enum across `types/` and all `*/types` subpackages returns true
  if and only if the input matches one of the OCPP-spec values.
- **Error semantics**: when construction fails, fuzzers require the returned
  error to wrap the shared sentinel validation errors
  (`types.ErrEmptyValue` / `types.ErrInvalidValue`).
- **Success invariants**: when construction succeeds, fuzzers assert strong
  postconditions (range checks, nil vs empty slice preservation, UTC-only
  `DateTime` values, and round-tripping of string/int fields).

#### Editor support (VS Code)

For local development, `.vscode/settings.json` configures `gopls` to include
`-tags=fuzz,race,bench` so the `./tests_fuzz`, `./tests_race`, and
`./tests_benchmark` test packages are indexed in the editor.

### Weekly CI (opt-in suites)

- Weekly workflow runs `make test-all`, `make test-race`, `make test-fuzz`,
  and `make test-bench` to guard the opt-in suites.

### Test Coverage

Reports are generated in the `reports/` directory:

- `reports/coverage.out` - Coverage data
- `reports/golangci-lint.txt` - Lint results

## OCPP 1.6 Compliance

### Type System

#### Core Types (`types/`)

| OCPP Type       | Go Type                 | Validation                              |
|-----------------|-------------------------|-----------------------------------------|
| CiString20Type  | `types.CiString20Type`  | Length <= 20, ASCII printable (32-126)  |
| CiString25Type  | `types.CiString25Type`  | Length <= 25, ASCII printable (32-126)  |
| CiString50Type  | `types.CiString50Type`  | Length <= 50, ASCII printable (32-126)  |
| CiString255Type | `types.CiString255Type` | Length <= 255, ASCII printable (32-126) |
| CiString500Type | `types.CiString500Type` | Length <= 500, ASCII printable (32-126) |
| dateTime        | `types.DateTime`        | RFC3339, UTC only                       |
| integer         | `types.Integer`         | uint16 (0-65535)                        |

#### Authorization Types (`types/`)

| OCPP Type           | Go Type                     | Description                      |
|---------------------|-----------------------------|----------------------------------|
| IdToken             | `types.IdToken`             | RFID tag identifier (CiString20) |
| IdTagInfo           | `types.IdTagInfo`           | Authorization info with status   |
| AuthorizationStatus | `types.AuthorizationStatus` | Accepted, Blocked, Expired, etc  |

#### Charging Profile Types (`types/`)

| OCPP Type                  | Go Type                            | Description                 |
|----------------------------|------------------------------------|-----------------------------|
| ChargingProfilePurposeType | `types.ChargingProfilePurposeType` | TxDefaultProfile, TxProfile |
| ChargingRateUnit           | `types.ChargingRateUnit`           | W or A                      |
| ChargingSchedule           | `types.ChargingSchedule`           | Schedule with periods       |
| ChargingSchedulePeriod     | `types.ChargingSchedulePeriod`     | Start/limit/phases          |

#### Meter Value Types (`types/`)

| OCPP Type      | Go Type                | Description                          |
|----------------|------------------------|--------------------------------------|
| MeterValue     | `types.MeterValue`     | Timestamp + SampledValue array       |
| SampledValue   | `types.SampledValue`   | Value + optional context/format/etc  |
| Measurand      | `types.Measurand`      | Energy, Power, Current, Voltage, etc |
| ReadingContext | `types.ReadingContext` | Sample.Clock, Sample.Periodic, etc   |
| ValueFormat    | `types.ValueFormat`    | Raw or SignedData                    |
| Phase          | `types.Phase`          | L1, L2, L3, N, L1-N, L2-N, L3-N      |
| Location       | `types.Location`       | Body, Cable, EV, Inlet, Outlet       |
| UnitOfMeasure  | `types.UnitOfMeasure`  | Wh, kWh, varh, kvarh, W, kW, VA, etc |

#### Message-Specific Types

| Package                               | Type                           | Description                        |
|---------------------------------------|--------------------------------|------------------------------------|
| `bootnotification/types`              | `RegistrationStatus`           | Accepted, Pending, Rejected        |
| `cancelreservation/types`             | `CancelReservationStatus`      | Accepted, Rejected                 |
| `changeavailability/types`            | `AvailabilityType`             | Inoperative, Operative             |
| `changeavailability/types`            | `AvailabilityStatus`           | Accepted, Rejected, Scheduled      |
| `changeconfiguration/types`           | `ConfigurationStatus`          | Accepted, Rejected, etc            |
| `clearcache/types`                    | `ClearCacheStatus`             | Accepted, Rejected                 |
| `clearchargingprofile/types`          | `ClearChargingProfileStatus`   | Accepted, Unknown                  |
| `datatransfer/types`                  | `DataTransferStatus`           | Accepted, Rejected, etc            |
| `diagnosticsstatusnotification/types` | `DiagnosticsStatus`            | Idle, Uploaded, UploadFailed, etc  |
| `firmwarestatusnotification/types`    | `FirmwareStatus`               | Downloaded, Installing, etc        |
| `getcompositeschedule/types`          | `GetCompositeScheduleStatus`   | Accepted, Rejected                 |
| `getconfiguration/types`              | `KeyValue`                     | Configuration key-value pair       |
| `getlocallistversion/types`           | `ListVersionNumber`            | Local list version number          |
| `remotestarttransaction/types`        | `RemoteStartTransactionStatus` | Accepted, Rejected                 |
| `remotestoptransaction/types`         | `RemoteStopTransactionStatus`  | Accepted, Rejected                 |
| `reservenow/types`                    | `ReservationStatus`            | Accepted, Faulted, Occupied, etc   |
| `reset/types`                         | `ResetType`                    | Hard, Soft                         |
| `reset/types`                         | `ResetStatus`                  | Accepted, Rejected                 |
| `sendlocallist/types`                 | `UpdateType`                   | Differential, Full                 |
| `sendlocallist/types`                 | `UpdateStatus`                 | Accepted, Failed, etc              |
| `sendlocallist/types`                 | `AuthorizationData`            | IdTag + IdTagInfo                  |
| `setchargingprofile/types`            | `ChargingProfile`              | Complete charging profile          |
| `setchargingprofile/types`            | `ChargingProfileKindType`      | Absolute, Recurring, Relative      |
| `setchargingprofile/types`            | `ChargingProfileStatus`        | Accepted, Rejected, etc            |
| `setchargingprofile/types`            | `RecurrencyKindType`           | Daily, Weekly                      |
| `statusnotification/types`            | `ChargePointErrorCode`         | ConnectorLockFailure, etc          |
| `statusnotification/types`            | `ChargePointStatus`            | Available, Charging, Faulted, etc  |
| `stoptransaction/types`               | `StopReason`                   | EmergencyStop, EVDisconnected, etc |
| `triggermessage/types`                | `MessageTrigger`               | BootNotification, Heartbeat, etc   |
| `triggermessage/types`                | `TriggerMessageStatus`         | Accepted, Rejected, NotImplemented |
| `unlockconnector/types`               | `UnlockStatus`                 | Unlocked, UnlockFailed, etc        |

### Message Implementation Status

| Message                       | Request | Confirmation | Package                         |
|-------------------------------|---------|--------------|---------------------------------|
| Authorize                     | Done    | Done         | `authorize`                     |
| BootNotification              | Done    | Done         | `bootnotification`              |
| CancelReservation             | Done    | Done         | `cancelreservation`             |
| ChangeAvailability            | Done    | Done         | `changeavailability`            |
| ChangeConfiguration           | Done    | Done         | `changeconfiguration`           |
| ClearCache                    | Done    | Done         | `clearcache`                    |
| ClearChargingProfile          | Done    | Done         | `clearchargingprofile`          |
| DataTransfer                  | Done    | Done         | `datatransfer`                  |
| DiagnosticsStatusNotification | Done    | Done         | `diagnosticsstatusnotification` |
| FirmwareStatusNotification    | Done    | Done         | `firmwarestatusnotification`    |
| GetCompositeSchedule          | Done    | Done         | `getcompositeschedule`          |
| GetConfiguration              | Done    | Done         | `getconfiguration`              |
| GetDiagnostics                | Done    | Done         | `getdiagnostics`                |
| GetLocalListVersion           | Done    | Done         | `getlocallistversion`           |
| Heartbeat                     | Done    | Done         | `heartbeat`                     |
| MeterValues                   | Done    | Done         | `metervalues`                   |
| RemoteStartTransaction        | Done    | Done         | `remotestarttransaction`        |
| RemoteStopTransaction         | Done    | Done         | `remotestoptransaction`         |
| ReserveNow                    | Done    | Done         | `reservenow`                    |
| Reset                         | Done    | Done         | `reset`                         |
| SendLocalList                 | Done    | Done         | `sendlocallist`                 |
| SetChargingProfile            | Done    | Done         | `setchargingprofile`            |
| StartTransaction              | Done    | Done         | `starttransaction`              |
| StatusNotification            | Done    | Done         | `statusnotification`            |
| StopTransaction               | Done    | Done         | `stoptransaction`               |
| TriggerMessage                | Done    | Done         | `triggermessage`                |
| UnlockConnector               | Done    | Done         | `unlockconnector`               |
| UpdateFirmware                | Done    | Done         | `updatefirmware`                |

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
- [Compatibility Contract](COMPATIBILITY.md)
- [Development Guide](CLAUDE.md)
- [Contributing](CONTRIBUTING.md)
- [Roadmap](ROADMAP.md)
- [Releasing](RELEASING.md)

## Performance Analysis: Custom Types vs Primitives

Condensed result:

- `PrimitiveDirect` is fastest, but it skips safety checks.
- `Custom` is slower in microbenchmarks, and the overhead is bounded and
  predictable versus `PrimitiveValidated`.
- For OCPP workloads, correctness and developer ergonomics generally justify
  the cost of first-class datatypes.

Benchmark suite structure:

- `analysis_benchmak/bench_micro_types_test.go`: constructor-level comparisons
  (`DateTime`, `ParentIdTag` chain, read paths).
- `analysis_benchmak/bench_macro_messages_test.go`: end-to-end message
  constructor path (`StartTransactionReq`).
- `analysis_benchmak/bench_scaling_sendlocallist_test.go`: slice-heavy scaling
  (`SendLocalListReq`) at `1, 25, 100, 250, 500, 1000`.
- `analysis_benchmak/bench_getconfiguration_scaling_test.go`: key-list scaling
  (`GetConfigurationReq`) at `1, 25, 100, 250, 500, 1000`.
- `analysis_benchmak/bench_common_test.go`: shared primitive baselines and
  validation helpers to keep comparisons consistent.

Single summary chart (`Custom / PrimitiveValidated`, lower is better):

![Custom vs PrimitiveValidated ratio](docs/img/custom_vs_validated_ratio.svg)

Full benchmark report with 6 charts and detailed analysis:
[docs/benchmark.md](docs/benchmark.md)

Reproduce:

<!-- markdownlint-disable MD046 -->
```sh
go run ./scripts/benchreport.go
```
<!-- markdownlint-enable MD046 -->
