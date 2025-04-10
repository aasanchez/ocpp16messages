// Package types contains reusable Open Charge Point Protocol (OCPP) 1.6 data types.
//
// This package is responsible for defining shared core types used across all OCPP 1.6 messages.
// It focuses on strong data modeling, including strict validation of constraints like:
//
// These types are designed to be protocol-agnostic, with no dependency on serialization logic
// (JSON, SOAP, etc.). They are used internally in request and response structures, and can be
// validated independently.
//
// This package should be imported using:
//
//	import "github.com/yourorg/ocpp16messages/types"
package types
