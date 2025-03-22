# OCPP 1.6 Message Validator & Parser Library

This Go package provides a robust, extensible, and idiomatic toolkit for decoding, validating, and extending Open Charge Point Protocol (OCPP) 1.6J messages. It is designed to be protocol-agnostic and can be used with both JSON and SOAP representations of OCPP messages.

---

## ✨ Features

- ✅ Full decoding and validation support for the `Authorize.req` and `Authorize.conf` messages.
- 🧩 Plugin system for custom message types, validators, and validation hooks.
- 📈 Benchmark and profiling suite for analyzing performance.
- 🧪 High test coverage across all components (target: 100%).
- 📚 Fully documented with GoDoc-compatible comments.
- 🛠 Forward-compatible layout for future OCPP 1.6J messages.

---

## 📦 Installation

```bash
go get github.com/aasanchez/ocpp16_messages
```

---

## 🚀 Quick Start

```go
package main

import (
	"log"
	"github.com/aasanchez/ocpp16_messages/core"
)

func main() {
	raw := []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)

	result, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Fatalf("❌ Invalid message: %v", err)
	}

	log.Printf("✅ Valid message. Action: %s | ID: %s", result.Action, result.UniqueID)
}
```

---

## 📁 Project Structure

```
ocpp16_messages/
├── authorize/          # Authorize.req and Authorize.conf structs and validation
├── core/               # Shared core types, plugin registry, enums, utils
├── benchmark/          # Benchmark & profiling tests
├── example/            # Example usage demos (authorization & plugin usage)
├── test/               # Full test suite (unit, integration, plugins)
├── ocpp16j.go          # Central entrypoint to validate OCPP messages
├── go.mod, go.sum      # Module dependencies
└── README.md           # You are here!
```

---

## 🧪 Testing

```bash
go test ./... -v
```

Generate coverage:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

---

## 🧩 Plugin Support

You can register custom validators or hooks via the `core` plugin registry:

```go
core.RegisterValidator("MyCustomAction", myValidator)
core.SetPreValidationHook(func(action string, payload any) {
	log.Printf("🔎 About to validate action %s", action)
})
```

See [`example/plugin_demo`](example/plugin_demo/main.go) for more.

---

## 📊 Benchmarking & Profiling

Run detailed performance benchmarks:

```bash
go test -bench=. -benchmem ./benchmark
```

Generate profiling data:

```bash
go test -bench=ValidateMessage -cpuprofile cpu.prof ./benchmark
go tool pprof -http=:8080 cpu.prof
```

---

## 📄 License

[MIT License](LICENSE)

---

## 🙌 Contributing

Contributions are welcome! Open issues, suggest ideas, or send pull requests. For major changes, please open a discussion first to propose what you’d like to change.

---

## 📚 Documentation

GoDoc: [pkg.go.dev/github.com/aasanchez/ocpp16_messages](https://pkg.go.dev/github.com/aasanchez/ocpp16_messages)

Each internal package includes its own `README.md` with detailed purpose and scope.
