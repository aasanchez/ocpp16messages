# Migration Plan: Replace Duplicated Types with `ocpp16types`

## Problem

`ocpp16messages` duplicates every OCPP 1.6 type locally instead of importing from `ocpp16types`. Types are scattered across two layers:

- **`types/`** — shared types (CiString, DateTime, Integer, IdToken, IdTagInfo, AuthorizationStatus, SampledValue, MeterValue, ChargingSchedule, ChargingSchedulePeriod, and 9 enums)
- **`<message>/types/`** — 23 message-specific type subdirectories containing enums and composite types that already exist in `ocpp16types`

Both implementations have been verified as API-identical (same constructors, same field names, same method signatures). The migration is a drop-in replacement.

---

## Phase 1: Foundation

### 1.1 Add `ocpp16types` dependency

```bash
go get github.com/aasanchez/ocpp16types@latest
```

This adds the module to `go.mod` and `go.sum`.

### 1.2 Update `golangci.yml` depguard

Add `github.com/aasanchez/ocpp16types` to the depguard allow list:

```yaml
depguard:
  rules:
    main:
      allow:
        - $gostd
        - github.com/aasanchez/ocpp16messages
        - github.com/aasanchez/ocpp16types    # <-- add this
```

### 1.3 Define the import alias convention

All message packages will import `ocpp16types` with the alias `st` (matching the existing convention):

```go
import (
    st "github.com/aasanchez/ocpp16types"
)
```

---

## Phase 2: Remove `types/` (shared types directory)

The entire `types/` directory at the package root is a duplicate of `ocpp16types`. It contains 19 files:

| File                            | Type(s)                                             | Replacement in `ocpp16types`  |
|---------------------------------|-----------------------------------------------------|-------------------------------|
| `authorizationstatus.go`        | AuthorizationStatus                                 | authorizationstatus.go        |
| `chargingprofilepurposetype.go` | ChargingProfilePurposeType                          | chargingprofilepurposetype.go |
| `chargingrateunit.go`           | ChargingRateUnit                                    | chargingrateunit.go           |
| `chargingschedule.go`           | ChargingSchedule, ChargingScheduleInput             | chargingschedule.go           |
| `chargingscheduleperiod.go`     | ChargingSchedulePeriod, ChargingSchedulePeriodInput | chargingscheduleperiod.go     |
| `cistring.go`                   | CiString20/25/50/255/500Type                        | cistring.go                   |
| `datetime.go`                   | DateTime                                            | datetime.go                   |
| `errors.go`                     | ErrEmptyValue, ErrInvalidValue                      | errors.go                     |
| `idtaginfo.go`                  | IdTagInfo                                           | idtaginfo.go                  |
| `idtoken.go`                    | IdToken                                             | idtoken.go                    |
| `integer.go`                    | Integer                                             | integer.go                    |
| `location.go`                   | Location                                            | location.go                   |
| `measurand.go`                  | Measurand                                           | measurand.go                  |
| `metervalue.go`                 | MeterValue, MeterValueInput                         | metervalue.go                 |
| `phase.go`                      | Phase                                               | phase.go                      |
| `readingcontext.go`             | ReadingContext                                      | readingcontext.go             |
| `sampledvalue.go`               | SampledValue, SampledValueInput                     | sampledvalue.go               |
| `unitofmeasure.go`              | UnitOfMeasure                                       | unitofmeasure.go              |
| `valueformat.go`                | ValueFormat                                         | valueformat.go                |

**Action:** Delete all `.go` files in `types/` (including `doc.go` and test files). Remove the `types/` directory entirely.

---

## Phase 3: Remove message-specific `types/` subdirectories

23 message packages have their own `types/` subdirectory. Every type in these directories exists in `ocpp16types`.

| Message Package                        | Types to Remove                                                                                  | `ocpp16types` Equivalent                             |
|----------------------------------------|--------------------------------------------------------------------------------------------------|------------------------------------------------------|
| `bootnotification/types/`              | RegistrationStatus                                                                               | registrationstatus.go                                |
| `cancelreservation/types/`             | CancelReservationStatus                                                                          | cancelreservationstatus.go                           |
| `changeavailability/types/`            | AvailabilityStatus, AvailabilityType                                                             | availabilitystatus.go, availabilitytype.go           |
| `changeconfiguration/types/`           | ConfigurationStatus                                                                              | configurationstatus.go                               |
| `clearcache/types/`                    | ClearCacheStatus                                                                                 | clearcachestatus.go                                  |
| `clearchargingprofile/types/`          | ClearChargingProfileStatus                                                                       | clearchargingprofilestatus.go                        |
| `datatransfer/types/`                  | DataTransferStatus                                                                               | datatransferstatus.go                                |
| `diagnosticsstatusnotification/types/` | DiagnosticsStatus                                                                                | diagnosticsstatus.go                                 |
| `firmwarestatusnotification/types/`    | FirmwareStatus                                                                                   | firmwarestatus.go                                    |
| `getcompositeschedule/types/`          | GetCompositeScheduleStatus                                                                       | getcompositeschedulestatus.go                        |
| `getconfiguration/types/`              | KeyValue, KeyValueInput                                                                          | keyvalue.go                                          |
| `getlocallistversion/types/`           | ListVersionNumber                                                                                | listversionnumber.go                                 |
| `metervalues/types/`                   | Location, Measurand, MeterValue, Phase, ReadingContext, SampledValue, UnitOfMeasure, ValueFormat | Already in shared types                              |
| `remotestarttransaction/types/`        | RemoteStartTransactionStatus                                                                     | remotestarttransactionstatus.go                      |
| `remotestoptransaction/types/`         | RemoteStopTransactionStatus                                                                      | remotestoptransactionstatus.go                       |
| `reservenow/types/`                    | ReservationStatus                                                                                | reservationstatus.go                                 |
| `reset/types/`                         | ResetStatus, ResetType                                                                           | resetstatus.go, resettype.go                         |
| `sendlocallist/types/`                 | AuthorizationData, UpdateStatus, UpdateType                                                      | authorizationdata.go, updatestatus.go, updatetype.go |
| `setchargingprofile/types/`            | ChargingProfile, ChargingProfileKindType, ChargingProfileStatus, RecurrencyKindType              | chargingprofile.go, etc.                             |
| `statusnotification/types/`            | ChargePointErrorCode, ChargePointStatus                                                          | chargepointerrorcode.go, chargepointstatus.go        |
| `stoptransaction/types/`               | StopReason (named Reason in spec)                                                                | stopreason.go                                        |
| `triggermessage/types/`                | MessageTrigger, TriggerMessageStatus                                                             | messagetrigger.go, triggermessagestatus.go           |
| `unlockconnector/types/`               | UnlockStatus                                                                                     | unlockstatus.go                                      |

**Action:** Delete all 23 `<message>/types/` subdirectories entirely (source, tests, doc.go, examples).

---

## Phase 4: Update imports in all 28 message packages

Each message package has `request.go`, `confirmation.go`, example tests, and test files that import from the local types. All imports must be rewritten.

### 4.1 Packages that only import shared `types/` (5 packages)

These packages have no message-specific types subdirectory:

- `authorize`
- `heartbeat`
- `starttransaction`
- `stoptransaction` *(note: stoptransaction does have a types/ dir but it's for StopReason only)*
- `updatefirmware`

**Change:**
```go
// Before
import st "github.com/aasanchez/ocpp16messages/types"

// After
import st "github.com/aasanchez/ocpp16types"
```

### 4.2 Packages that import shared `types/` + message-specific `types/` (23 packages)

These packages currently import from two locations. Both collapse into one:

```go
// Before
import (
    mbt "github.com/aasanchez/ocpp16messages/bootnotification/types"
    st  "github.com/aasanchez/ocpp16messages/types"
)

// After
import (
    st "github.com/aasanchez/ocpp16types"
)
```

**All message-specific type references change their prefix:**
```go
// Before
mbt.RegistrationStatusAccepted

// After
st.RegistrationStatusAccepted
```

### 4.3 Update example tests

Every `example_request_test.go` and `example_confirmation_test.go` in the 28 message packages imports types. Update all import paths.

### 4.4 Update test files

Test files in `<message>/tests/` subdirectories also import types. Update all import paths.

---

## Phase 5: Migration of special cases

### 5.1 `metervalues/types/` — heavy duplication

This directory re-exports the full SampledValue/MeterValue type graph (8 files + errors.go). All are identical to `ocpp16types`. Delete entirely and update `metervalues/request.go` and `metervalues/confirmation.go` to import from `ocpp16types`.

### 5.2 `setchargingprofile/types/chargingprofile.go` — composite type

This file contains the full ChargingProfile composite type with validation. It imports from `types/` for ChargingSchedule, DateTime, Integer, etc. The identical implementation exists in `ocpp16types`. Delete and update imports.

### 5.3 `sendlocallist/types/authorizationdata.go` — composite type

Contains AuthorizationData with IdToken + IdTagInfo construction. Identical implementation in `ocpp16types`. Delete and update imports.

### 5.4 `getconfiguration/types/keyvalue.go` — composite type

Contains KeyValue with CiString50Type + CiString500Type. Identical implementation in `ocpp16types`. Delete and update imports.

### 5.5 `getlocallistversion/types/listversionnumber.go` — composite type

Contains ListVersionNumber with int32 validation. Identical implementation in `ocpp16types`. Delete and update imports.

---

## Phase 6: Verify and clean up

### 6.1 Run tests

```bash
make test-all
```

### 6.2 Run linter

```bash
make lint
```

### 6.3 Tidy modules

```bash
go mod tidy
```

### 6.4 Verify no local type references remain

```bash
grep -r "ocpp16messages/types" --include="*.go" .
grep -r "ocpp16messages/.*/types" --include="*.go" .
```

Both should return zero results.

### 6.5 Delete empty directories

Confirm all 23 message-specific `types/` directories and the root `types/` directory are removed.

---

## Execution Order

Work message-by-message to keep the codebase compilable at each step:

1. **Phase 1** — Add dependency, update linter config
2. **Phase 2** — Migrate one simple message first (e.g., `authorize`) as a proof-of-concept, verify tests pass
3. **Phase 3** — Migrate remaining 27 message packages in batches:
   - Simple packages (no message-specific types): `heartbeat`, `updatefirmware`, `starttransaction`
   - Enum-only message types: `bootnotification`, `cancelreservation`, `changeavailability`, etc.
   - Complex message types: `setchargingprofile`, `sendlocallist`, `getconfiguration`, `getlocallistversion`, `metervalues`
4. **Phase 4** — Delete root `types/` directory
5. **Phase 5** — Run full test suite, lint, tidy
6. **Phase 6** — Commit and push

---

## Files Changed Summary

| Category                        | Count                | Action                     |
|---------------------------------|----------------------|----------------------------|
| `go.mod`                        | 1                    | Add ocpp16types dependency |
| `golangci.yml`                  | 1                    | Update depguard allow list |
| Root `types/` directory         | ~19 files + tests    | Delete entirely            |
| Message-specific `types/` dirs  | ~23 dirs, ~80+ files | Delete entirely            |
| Message `request.go` files      | 28                   | Update imports             |
| Message `confirmation.go` files | 28                   | Update imports             |
| Message example test files      | ~56                  | Update imports             |
| Message test files              | ~56+                 | Update imports             |
| **Total files deleted**         | **~100+**            |                            |
| **Total files modified**        | **~170+**            |                            |

---

## Risks and Mitigations

| Risk                       | Mitigation                                                                                  |
|----------------------------|---------------------------------------------------------------------------------------------|
| API drift between packages | Verified: constructors, fields, methods are identical                                       |
| Constant name differences  | Verified: all enum constant names match exactly                                             |
| Error message differences  | Verified: same sentinel errors and format strings                                           |
| Import alias collisions    | Standardize on `st` alias everywhere                                                        |
| Test output changes        | Run full test suite after each batch                                                        |
| Downstream consumers       | `ocpp16messages` is the downstream consumer of `ocpp16types` — no further downstream impact |
