# ocpp16_messages

[![codecov](https://codecov.io/gh/aasanchez/ocpp16_messages/branch/main/graph/badge.svg)](https://codecov.io/gh/aasanchez/ocpp16_messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16_messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16_messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16_messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16_messages)

## DO NOT USE IN PRODUCTION, STILL UNDER DEVELOPMENT

A Go package for defining and validating OCPP 1.6 message structures.

A protocol-neutral Go package for defining and validating **OCPP 1.6 messages**
— request and confirmation types as defined in the official [Open Charge Alliance specification](https://www.openchargealliance.org/).

This package is ideal for use in both **Charge Point (CP)** and **Central System (CSMS)** implementations, and is
designed to be consumed by external tools, proxies, or full CSMS solutions.

---

## ✨ Features

- 📚 Full set of **OCPP 1.6 message types**
- ✅ **Strict validation** for all required fields and OCPP constraints
- ⚙️ Protocol-neutral (no WebSocket/SOAP/Json binding)
- 🧩 Modular directory structure (chargePoint, centralSystem, models, enums)
- 🔍 Optional values validated when present
- 💡 Cleanly typed `CiString` with max length enforcement
- 🔬 Full test coverage with GitHub Actions & Codecov integration

---

## 📦 Installation

```bash
go get github.com/aasanchez/ocpp16_messages
```

## Usage

```go
package main

import (
  "log"
  "os"
  "time"

  "github.com/aasanchez/ocpp16_messages/messages/chargePoint"
  "github.com/aasanchez/ocpp16_messages/validators"
)

func main() {
  // Configure log output to show timestamps and source file/line
  log.SetFlags(log.LstdFlags | log.Lshortfile)

  // Example BootNotification.req from a Charge Point
  req := chargePoint.BootNotificationReq{
    ChargePointModel:  "ModelX",
    ChargePointVendor: "VendorY",
  }

  // Validate the request and exit if it's invalid
  exitOnError(validators.ValidateBootNotificationReq(req), "BootNotificationReq validation failed")

  log.Println("✅ Valid BootNotification request")

  // Example BootNotification.conf from the Central System
  conf := chargePoint.BootNotificationConf{
    Status:      "Accepted", // Change to an invalid string to test failure
    CurrentTime: time.Now(),
    Interval:    30,
  }

  // Validate the confirmation and exit if it's invalid
  exitOnError(validators.ValidateBootNotificationConf(conf), "BootNotificationConf validation failed")

  log.Println("✅ Valid BootNotification confirmation")
}

// exitOnError logs the given error with context and exits the program with status code 1.
// This is a utility function to simplify error handling in example code.
func exitOnError(err error, context string) {
  if err != nil {
    log.Printf("❌ %s: %v\n", context, err)
    os.Exit(1)
  }
}
```
