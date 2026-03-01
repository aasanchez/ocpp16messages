# Changelog

All notable changes to this project will be documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and the project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added (Unreleased)

- Compatibility contract (`COMPATIBILITY.md`) defining SemVer guarantees,
  deprecation process, and what counts as breaking (API, behavior, validation).

### Changed (Unreleased)

- (Nothing yet.)

## [1.0.3] - 2026-03-01

### Added (1.0.3)

- Expanded benchmark suite (opt-in) with high-value constructor benchmarks and
  worst-case coverage for nested payloads and optional-field validation.

### Changed (1.0.3)

- Added comprehensive race test suite (opt-in) for all message constructors and
  core constructors; hardened immutability by eliminating pointer/slice aliasing
  and ensuring pointer/slice getters return defensive copies.
- Renamed opt-in test directories:
  - `./fuzz` -> `./tests_fuzz`
  - `./race` -> `./tests_race`
  - `./benchmark` -> `./tests_benchmark`
- Updated `Makefile`, `README.md`, and VS Code `gopls` configuration to match
  the new opt-in suite directory names and build tags.

## [1.0.2] - 2026-03-01

### Added (1.0.2)

- Comprehensive fuzz test suite under `./fuzz` (guarded by the `fuzz` build
  tag) covering all OCPP 1.6 message constructors (`Req`/`Conf`) and core
  validation types.
- Strict membership fuzzers for every `IsValid()` enum in `types/` and all
  message subpackage `*/types` enums.

### Changed (1.0.2)

- `make test-fuzz` hardened: exact fuzzer selection, configurable per-fuzzer
  time budget, and capped parallelism.
- Weekly GitHub Action schedule runs fuzzers (opt-in suite).

## [1.0.1] - 2026-01-26

### Changed (1.0.1)

- Tests: adjust constants to satisfy revive `magic-number` lint.

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
- CI and go.mod aligned on Go 1.25.7 (documented requirement).

### Removed (1.0.0)

- Legacy opt-in test filenames superseded by the normalized naming scheme.

[1.0.0]: https://github.com/aasanchez/ocpp16messages/releases/tag/v1.0.0
[1.0.1]: https://github.com/aasanchez/ocpp16messages/releases/tag/v1.0.1
[1.0.2]: https://github.com/aasanchez/ocpp16messages/releases/tag/v1.0.2
[1.0.3]: https://github.com/aasanchez/ocpp16messages/releases/tag/v1.0.3
[Unreleased]: https://github.com/aasanchez/ocpp16messages/compare/v1.0.3...HEAD

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
