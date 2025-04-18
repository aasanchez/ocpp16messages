// Package authorize defines the message structures for the Authorize feature in OCPP 1.6.
//
// This package provides strict, type-safe, and fully specification-compliant
// data models for the Authorize.req and Authorize.conf messages as defined by the
// Open Charge Point Protocol (OCPP) 1.6 Core specification (Section 4.4).
//
// What this package provides:
//
//   - A strongly typed representation of all message fields, including wrapper types
//     for constrained strings (e.g., CiString20Type)
//   - Validation of message structure and required fields
//   - Go-native types that preserve OCPP semantics and constraints
//
// What this package does NOT provide:
//
//   - No authorization decision logic or workflow
//   - No business logic, caching, or state management
//   - No transport or network communication
//
// This package is designed to be used by higher-level systems such as:
//
//   - CSMS implementations
//   - OCPP gateways and proxies
//   - Testing or conformance tooling
//   - Validation and decoding layers for OCPP messages
//
// Example:
//
//	import (
//	    "github.com/aasanchez/ocpp16messages/messages/authorize"
//	    "github.com/aasanchez/ocpp16messages/messages/authorize/types"
//	)
//
//	idTag, _ := types.IdToken("ABC123")
//	request := authorize.RequestMessage{IdTag: idTag}
//
//	if err := request.Validate(); err != nil {
//	    log.Fatalf("invalid request: %v", err)
//	}
//
// Specification Reference:
//   - OCPP 1.6 Core – Section 4.4: Authorize
package authorize
