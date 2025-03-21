# ocpp16_messages

[![codecov](https://codecov.io/gh/aasanchez/ocpp16_messages/branch/main/graph/badge.svg)](https://codecov.io/gh/aasanchez/ocpp16_messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16_messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16_messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)

A Go package for defining and validating OCPP 1.6 message structures.

A protocol-neutral Go package for defining and validating **OCPP 1.6J messages** — request and confirmation types as
defined in the official [Open Charge Alliance specification](https://www.openchargealliance.org/).

This package is ideal for use in both **Charge Point (CP)** and **Central System (CSMS)** implementations, and is
designed to be consumed by external tools, proxies, or full CSMS solutions.

---

## ✨ Features

- 📚 Full set of **OCPP 1.6J message types**
- ✅ **Strict validation** for all required fields and OCPP constraints
- ⚙️ Protocol-neutral (no WebSocket/SOAP binding)
- 🧩 Modular directory structure (chargePoint, centralSystem, models, enums)
- 🔍 Optional values validated when present
- 💡 Cleanly typed `CiString` with max length enforcement
- 🔬 Full test coverage with GitHub Actions & Codecov integration

---

## 📦 Installation

```bash
go get github.com/aasanchez/ocpp16_messages
```
