# OCPP 1.6 Authorize Message Package

This package defines the **message types** and **data structures** for the
`Authorize.req` and `Authorize.conf` PDUs used in the
[OCPP 1.6](https://www.openchargealliance.org/protocols/ocpp-16/) protocol.

> ❗ This package **does not** implement any business logic, message routing, or
> authorization workflows. Its only purpose is to provide a
> **strict, type-safe,specification-compliant message model** for OCPP 1.6
> messages.

---

## What This Package Does

- Defines Go types for all fields in `Authorize.req` and `Authorize.conf`
- Validates required fields, types, constraints (e.g., max string length)
- Wraps simple types (e.g. `string`) in **spec-compliant named types** like:
  - `idTag` → `IdTokenType` → `CiString20Type`
- Ensures OCPP 1.6 semantics are preserved down to the field level
- Designed to be used by higher-level libraries that implement business logic

---

## What This Package Does *Not* Do

- ✖ No authorization decision logic
- ✖ No charge point state handling
- ✖ No communication with Central Systems or devices
- ✖ No implementation of Local Authorization Lists, Caches, or lookup logic

---

## Example: Type-Safe Field Representation

Instead of:

```go
type AuthorizeRequest struct {
  IdTag string // fragile, lacks validation
}
```

You get:

```go
type RequestMessage struct {
  IdTag IdTokenType // validates format, length, encoding
}

type IdTokenType struct {
  value CiString20Type
}

type CiString20Type struct {
  value string // max 20 chars, ASCII printable
}
```

This ensures

- Schema-level conformance (OCPP 1.6J)
- Reusable, validated types
- Better IDE support and safety

---

### Authorize.req Overview

The Charge Point sends an `Authorize.req` to the Central System to check if a
given `idTag` (e.g. RFID card) is permitted to initiate or stop a charging
session.

```go
req := authorize.RequestMessage{
  IdTag: authorizetypes.MustIdToken("ABC123"), // safe, validated
}
```

### Authorize.conf Overview

The Central System replies with an `Authorize.conf` containing the authorization
status and optional metadata (e.g. expiry, parent ID):

```go
res := authorize.ConfirmationMessage{
  IdTagInfo: authorizetypes.IdTagInfoType{
    Status: mustStatus("Accepted"),
  },
}
```

---

## Use Cases

- Protocol gateways
- OCPP CSMS implementations
- Emulators / test runners
- Proxies and middlewares
- Validation tools
