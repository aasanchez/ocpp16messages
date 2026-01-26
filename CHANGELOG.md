# Changelog

All notable changes to this project will be documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and the project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-26

### Added (1.0.0)

- Complete OCPP 1.6 message surface (28/28) with validated Req/Conf
  constructors.
- Opt-in quality suites: `make test-race`, `make test-fuzz`,
  `make test-bench`.
- Documentation clarifying UTC-only DateTime, optional slice convention,
  and public API expectations.

### Changed (1.0.0)

- Integer constructors accept `int`, validating uint16 range for protocol
  fields.
- Error handling aligned to shared sentinels and `errors.Join` aggregation.
- MeterValues and StopTransaction constructors preserve empty vs nil slices.
- CI and go.mod aligned on Go 1.24.6 (documented requirement).

### Removed (1.0.0)

- Legacy opt-in test filenames superseded by the normalized naming scheme.

[1.0.0]: https://github.com/aasanchez/ocpp16messages/releases/tag/v1.0.0

## [v0.11.0]

### Added (v0.11.0)

- Stringer support across core types to ease logging/debugging.
- Additional validation coverage for charging profiles and enums.

### Changed (v0.11.0)

- Simplified phase/number validation logic in types.
- Documentation adjustments to reflect high (not perfect) coverage.

## [v0.10.0]

### Added (v0.10.0)

- Major feature wave: Start/StopTransaction, StatusNotification,
  SetChargingProfile, RemoteStart/StopTransaction, ReserveNow, CancelReservation,
  Reset, MeterValues, DataTransfer, TriggerMessage, UpdateFirmware,
  GetDiagnostics, DiagnosticsStatusNotification, FirmwareStatusNotification,
  GetCompositeSchedule, GetLocalListVersion, SendLocalList, UnlockConnector.

### Changed (v0.10.0)

- Shared authorization/meter value types moved to common package.
- README updated to reflect full 28/28 coverage.
- Lint cleanups (exhaustruct, line length, constants).

## [v0.9.0]

### Added (v0.9.0)

- Markdown lint configuration; docs brought in line with new structure.

## [v0.8.0]

### Added (v0.8.0)

- ClearCache.req/conf implementation.

### Changed (v0.8.0)

- Major repo layout refactor (moved packages one level up).

## [v0.7.0]

### Added (v0.7.0)

- ChangeConfiguration, ChangeAvailability, CancelReservation,
  BootNotification messages (req/conf).

## [v0.6.0]

### Added (v0.6.0)

- BootNotification message pair (req/conf) and supporting validation.
- Authorize Conf accumulates validation errors.

### Changed (v0.6.0)

- Test fixtures and lint fixes; doc.go placeholders for messages.

## [v0.5.1]

### Added (v0.5.1)

- Authorize.conf message finalized; consistent error handling.

### Changed (v0.5.1)

- Test refactors and lint fixes; message placeholders and structure
  cleanup.

## [v0.5.0]

### Added (v0.5.0)

- Authorize message initial implementation and security policy.
- 100% coverage target with comprehensive tests.

### Changed (v0.5.0)

- Renamed core types for OCPP compliance (CiString*Type, Id* casing) and
  standardized constructors (Set\* -> New\*).
- Extensive doc and test restructuring; CI/workflow updates.

## Earlier tags

For v0.4.0 and earlier, see the git history; these releases covered the
initial type system, early CI setup, and incremental documentation/testing improvements.
