# `core` Package

The `core` package provides shared data types, utilities, and parsing logic used across all OCPP 1.6 message types. It is designed to support **both JSON and SOAP** message serialization formats, with a clear separation of responsibilities.

---

## ✨ Responsibilities

- **Message Type Constants**  
  Defines the OCPP message types (`CALL`, `CALLRESULT`, `CALLERROR`) and their numeric representations.

- **Common Type Definitions**  
  Includes reusable types such as `CiString20`, `IdTag`, and `AuthorizationStatus`.

- **Parsing Logic**  
  - `parser_json.go`: Handles parsing of OCPP messages in JSON format.
  - `parser_soap.go`: Handles minimal parsing of OCPP SOAP envelopes for routing and validation.

- **Error Modeling**  
  Implements standardized error reporting (`FieldError`) and validation helpers.

- **CALLERROR Message Support**  
  Full definition and validation logic for OCPP CALLERROR messages (JSON only).

---

## 📂 File Overview

| File             | Purpose                                                       |
|------------------|---------------------------------------------------------------|
| `call_error.go`  | JSON structure and validation for CALLERROR messages          |
| `ci_string.go`   | CiString-based types (`CiString20`, etc.)                     |
| `enums.go`       | Enum definitions (`MessageType`, `AuthorizationStatus`, etc.) |
| `errors.go`      | Field-level validation errors and utilities                   |
| `id_token.go`    | Shared `IdTag` structure and constraints                      |
| `parser_json.go` | Parsing of JSON OCPP messages into structured form            |
| `parser_soap.go` | SOAP envelope parsing for routing (no full XML binding)       |
| `doc.go`         | Package documentation for GoDoc and pkg.go.dev                |

---

## 🧪 Testing

All test files live under:

```text
core/test/
```

Tests use **only the Go standard library** and aim for **100% coverage**.

---

## 🧼 Design Principles

- ✅ Clear boundaries between JSON and SOAP logic
- ✅ No external dependencies
- ✅ Idiomatic Go (GoDoc comments, standard error patterns, type safety)
- ✅ Reusable and pluggable for different OCPP message implementations

---

## 🔄 Extensibility

This package is designed to be **imported** and reused by specific OCPP message packages (e.g., `authorize`) and is agnostic to the application layer (e.g., charger, CSMS, proxy).

For future extensions:

- Add new enum values to `enums.go`
- Introduce new `CiString`-based constraints in `ci_string.go`
- Add parsing logic if new message types or protocols are supported
