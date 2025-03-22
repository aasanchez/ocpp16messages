# core

The `core` package provides foundational components and shared utilities used across all OCPP 1.6J message types. It is designed to be modular, extensible, and aligned with Go best practices.

---

## Responsibilities

This package includes:

- **Message Parsing** (`parser.go`): Robust utilities to parse raw OCPP JSON arrays (`[2,"uid","Action",{...}]`) into structured messages.
- **Validation Plugin System** (`plugin.go`): A flexible mechanism for registering and executing custom validators for specific OCPP message actions.
- **Error Generation** (`call_error.go`): Utility functions for constructing well-formed OCPP `CALLERROR` responses.
- **Common Types**:
  - `CiString` types (`ci_string.go`): Strongly-typed string definitions with max-length constraints, used throughout the OCPP schema.
  - `Enums` (`enums.go`): Shared enum definitions such as `AuthorizationStatus`, `RegistrationStatus`, etc.
  - `IdToken` (`id_token.go`): Type-safe representation of `idTag` tokens used in multiple OCPP messages.
- **Validation Error Helpers** (`errors.go`): Centralized helpers to produce field-specific validation errors with rich context.

---

## Design Goals

- **Modularity**: Each file focuses on a single area of concern (e.g., string types, enums, plugin registration).
- **Extensibility**: The plugin system allows custom validators and lifecycle hooks for instrumenting or enhancing message validation.
- **Performance**: Core types and utilities are designed to support high-throughput validation scenarios (e.g., proxies or emulators).
- **Reusability**: All data types defined here (e.g., `CiString20Type`, `IdToken`) are reusable across any OCPP 1.6J message implementation.

---

## Usage Example

Registering a plugin validator:

```go
core.RegisterValidator("Authorize", authorize.CustomAuthorizeValidator{})
```

Using a pre-validation hook:

```go
core.SetPreValidationHook(func(action string, payload any) {
    log.Printf("🔍 Validating %s with raw payload: %+v", action, payload)
})
```

Parsing and validating a raw OCPP message:

```go
parsed, err := core.ValidateRawMessage(rawMessageBytes)
if err != nil {
    // handle invalid structure or validation failure
}
```

---

## Related Files

- `plugin.go`: Plugin registration and hooks
- `parser.go`: Raw message parsing
- `call_error.go`: Error response formatting
- `ci_string.go`: Bounded string types
- `enums.go`: Common enums across the protocol
- `id_token.go`: Reusable token wrapper
- `errors.go`: Field-level error helpers

---

## Future Enhancements

- Support for JSON Schema-based validation
- Generic registry for all message types (not only by action)
- Additional hooks: trace IDs, request logging, etc.

---

## Maintainer Notes

This package is core infrastructure for all message definitions. Ensure backward compatibility and thorough test coverage for all changes.
