# Core Package

The `core` package provides shared types, constants, utilities, and parsers that are fundamental for handling Open Charge Point Protocol (OCPP) 1.6 messages in both JSON and SOAP formats. This package is designed to be agnostic to specific OCPP message types and instead focuses on providing foundational building blocks that other sub-packages (like `authorize`) rely on.

---

## 📦 Responsibilities

- Define OCPP message types and their constants (`CALL`, `CALLRESULT`, `CALLERROR`).
- Provide reusable data types such as `CiString20`, `IdToken`, and their validators.
- Parse and validate raw OCPP messages from JSON (`parser_json.go`) and, optionally, SOAP in the future.
- Define and handle errors consistently using custom error types such as `FieldError`.
- Handle CALLERROR message structure and validation.

---

## 📁 Files Overview

| File             | Description                                                                         |
|------------------|-------------------------------------------------------------------------------------|
| `call_error.go`  | Validates and constructs CALLERROR messages.                                        |
| `ci_string.go`   | Implements CiString types (e.g., `CiString20`) with length validation.              |
| `enums.go`       | Contains shared enums like `MessageType` used throughout OCPP message handling.     |
| `errors.go`      | Custom error handling, including `FieldError` for validation errors.                |
| `id_token.go`    | Contains the `IdToken` type with support for strict validation.                     |
| `parser_json.go` | Parses JSON-formatted OCPP messages and extracts structured message data.           |
| `validator.go`   | Provides the message validator registry and plugin hooks for extensible validation. |
| `plugin.go`      | Allows third-party registration of custom validators and validation hooks.          |
| `doc.go`         | Package-level documentation for compatibility with pkg.go.dev.                      |
| `test/`          | Contains unit tests for each of the above components.                               |

---

## 🧪 Testing

All core components are tested in isolation using Go's standard `testing` package. The tests are organized in `core/test` to separate concerns and improve clarity.

To run tests:

```bash
go test ./core/test -v
```

Code coverage is targeted to be 100% for this package.

---

## 🔧 Example Usage

### Parsing a JSON OCPP CALL message:

```go
import "github.com/aasanchez/ocpp16_messages/core"

raw := []byte(`[2, "abc123", "Authorize", {"idTag": "ABC123"}]`)
msg, err := core.ParseJSONMessage(raw)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Parsed message action:", msg.Action)
```

### Validating a CALLERROR message:

```go
msg := []any{
    4,
    "id123",
    "ProtocolError",
    "Unsupported operation",
    map[string]any{"field": "value"},
}

validated, err := core.ValidateCallError(msg)
if err != nil {
    log.Fatal("CALLERROR validation failed:", err)
}
```

---

## 🧩 Design Notes

- This package only uses Go’s standard library. No third-party dependencies are introduced for portability, minimalism, and auditability.
- SOAP support is planned and will be included in separate files (e.g., `parser_soap.go`) to preserve separation of concerns.

---

## 📝 License

MIT © Alexis Sanchez, 2025
