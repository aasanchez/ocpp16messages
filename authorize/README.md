# authorize Package

The `authorize` package defines the data structures and validation logic for the `Authorize.req` and `Authorize.conf` messages, as specified by the [OCPP 1.6J](https://www.openchargealliance.org/protocols/ocpp-16/) protocol. These messages are used for validating identification tokens before allowing a charging session.

---

## 📥 Authorize.req

The `Authorize.req` message is sent by the Charge Point to the Central System (CSMS) to request authorization for an `idTag`.

### Structure

```json
{
  "idTag": "ABC123456789"
}
```

### Corresponding Go Structure

```go
type Req struct {
    IdTag core.IdToken `json:"idTag"`
}
```

The `IdToken` type wraps the `idTag` as a `CiString20Type` and includes built-in validation.

---

## 📤 Authorize.conf

The `Authorize.conf` message is the response from the CSMS, indicating the result of the authorization request.

### Structure

```json
{
  "idTagInfo": {
    "status": "Accepted",
    "expiryDate": "2025-12-31T23:59:59Z",
    "parentIdTag": "PARENT123"
  }
}
```

### Corresponding Go Structure

```go
type Conf struct {
    IdTagInfo IdTagInfo `json:"idTagInfo"`
}
```

The `IdTagInfo` contains the authorization `status` and optional `expiryDate` and `parentIdTag`.

---

## ✅ Validation

Validation for both request and confirmation messages is handled using idiomatic Go and centralized in:

- `authorize/validator.go`

You can validate a message using:

```go
err := authorize.ValidateReq(req)
err := authorize.ValidateConf(conf)
```

These methods check required fields and ensure all string constraints are respected (e.g., max length for `idTag`).

---

## 🧪 Examples

See the [example/authorization/main.go](../example/authorization/main.go) file for usage examples.

```go
raw := []byte(`[2,"123456","Authorize",{"idTag":"D0431F35"}]`)
parsed, err := ocpp16j.ParseMessage(raw)
// Then validate: ocpp16j.ValidateMessage(parsed)
```

---

## 🔌 Plugin Support

Custom validation logic can be registered dynamically using the plugin API. See [core/plugin.go](../core/plugin.go) for details.

---

## 📚 Related Types

- `core.CiString20Type`
- `core.IdToken`
- `core.AuthorizationStatus`
- `core.FieldError`

---

## 📖 Spec Reference

This package implements the following OCPP 1.6J messages:

- [Authorize.req](https://www.openchargealliance.org/protocols/ocpp-16/ocpp-16-part-2-advanced-features/#152-authorize)
- [Authorize.conf](https://www.openchargealliance.org/protocols/ocpp-16/ocpp-16-part-2-advanced-features/#153-authorizeconf)

---

## 📁 Directory

```
authorize/
├── confirm.go         # Authorize.conf structure
├── request.go         # Authorize.req structure
├── validator.go       # Validation logic
├── README.md          # Package documentation (this file)
```
