# ocpp16_messages

A high-performance and idiomatic Go package for handling Open Charge Point Protocol (OCPP) 1.6 messages in both JSON and SOAP formats.

---

## ✅ Features

- 🔁 Full parsing and validation of OCPP 1.6J (JSON) and 1.6S (SOAP) message formats.
- 📦 Modular message design (e.g., `authorize`, `bootnotification`, etc.)
- 🧩 Plugin-style validator system for extensibility.
- 🧪 100% test coverage with strong focus on performance.
- 📊 Benchmarking and profiling suite.
- 🧵 Thread-safe validation logic.
- 🧰 Rich set of reusable core types (e.g., `CiString`, `IdToken`, `StatusEnum`).
- 🔌 Examples for JSON and SOAP usage in `example/authorize/json/` and `example/authorize/soap/`.

---

## 🏁 Getting Started

```bash
go get github.com/aasanchez/ocpp16_messages
```

---

## 🧪 Run Tests & Benchmarks

### Run all tests

```bash
make test
```

### Run benchmarks

```bash
go test -bench=. -benchmem ./benchmark
```

---

## 📂 Project Structure

```text
ocpp16_messages/
├── authorize/                  # Implementation of the Authorize.req and Authorize.conf messages
├── core/                       # Shared types, validators, plugins, enums, parsers
├── benchmark/                  # Benchmarks for JSON & SOAP performance analysis
├── example/authorize/json/     # Example for Authorize using JSON
├── example/authorize/soap/     # Example for Authorize using SOAP
├── test/                       # Full coverage test suite
├── ocpp16_messages.go          # Entrypoint to parse, route, and validate OCPP 1.6 messages
├── go.mod / go.sum             # Go module files
└── README.md                   # This file
```

---

## 🔌 Plugin Support

Custom validators can be registered via the plugin system:

```go
core.RegisterValidator("CustomAction", CustomValidator{})
```

Validation hooks can be set for observability:

```go
core.SetPreValidationHook(func(action string, raw json.RawMessage) {
    fmt.Println("Raw received:", action)
})
```

---

## 🔍 Versioning

- Current version: **v0.1.0**
- Next steps: implement BootNotification message and SOAP header validation.

---

## 📄 License

MIT License © Alexis Sanchez
