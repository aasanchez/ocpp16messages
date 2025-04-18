# OCPP 1.6 Messages for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16messages.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16messages)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16messages/branch/main/graph/badge.svg?token=1I9VVL7DWO)](https://codecov.io/gh/aasanchez/ocpp16messages)
[![Go Report Card](https://goreportcard.com/badge/github.com/aasanchez/ocpp16_messages)](https://goreportcard.com/report/github.com/aasanchez/ocpp16_messages)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=bugs)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=aasanchez_ocpp16_messages&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=aasanchez_ocpp16_messages)

## Overview

OCPP (Open Charge Point Protocol) is an open communication standard used between
Electric Vehicle (EV) charging stations and a central management system. It
enables interoperability between different charging station vendors and
management system providers.

This library provides Go data structures and functions for creating, parsing,
and validating OCPP 1.6 messages, simplifying the development of OCPP-compliant
applications in Go.

## Features

Currently implemented OCPP 1.6 messages:

* **Authorize**: Used to request authorization for an idTag.
* **BootNotification**: Sent by the Charge Point to the Central System when it
* boots up.

## Installation

To use this library in your Go project, you can install it using:

```bash
go get github.com/aasanchez/ocpp16messages
```

## Usage

Here's a simple example of how to create an Authorize request message:

```go
package main

import (
  "fmt"
  "github.com/aasanchez/ocpp16messages/messages/authorize"
)

func main() {
  idTag := "B85A-50CBE9678EC6" // Replace with your actual IdTag

  // Create a new Authorize request
  authReq, err := authorize.Request(idTag)
  if err != nil {
    fmt.Printf("Error creating Authorize request: %v\n", err)
    return
  }

  // The authReq object can now be marshalled to JSON and sent to a Central System.
  // For demonstration, we'll just print its string representation.
  fmt.Printf("Authorize Request: %s\n", authReq.String())

  // Example of how you might marshal to JSON (actual sending logic depends on your transport)
  // import "encoding/json"
  // jsonBytes, err := json.Marshal(authReq)
  // if err != nil {
  // fmt.Printf("Error marshalling to JSON: %v\n", err)
  // return
  // }
  // fmt.Printf("Authorize Request JSON: %s\n", string(jsonBytes))
}
```

This example demonstrates how to import the `authorize` package and use its
`Request` constructor to create a new Authorize request. The resulting `authReq`
object then contains the structured data for the OCPP message.
