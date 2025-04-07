// Package core provides the foundational infrastructure for working with
// Open Charge Point Protocol (OCPP) 1.6 messages using the JSON transport format.
//
// It is designed as the internal backbone for parsing, validation, and extensibility,
// and is reusable for building protocol-aware systems such as CSMS (Central Systems)
// or OCPP gateways.
//
// Overview:
//
// The core package centralizes the following capabilities:
//
//   - Unified message parsing for OCPP 1.6J messages (CALL, CALLRESULT, CALLERROR)
//   - Extensible message validation via a plugin registry
//   - Shared primitive types and strongly typed string definitions (CiString variants)
//   - Enum definitions for standardized fields (e.g., AuthorizationStatus, ConnectorStatus)
//   - Structured error handling and validation hooks
//
// Usage Example:
//
//	RegisterValidator("AuthorizeReq", authorize.ValidateAuthorizeReq)
//	RegisterValidator("AuthorizeConf", authorize.ValidateAuthorizeConf)
//
//	parsed, err := ParseMessage(rawMessage)
//	if err != nil {
//	    log.Fatalf("Parsing failed: %v", err)
//	}
//
//	result, err := ValidateMessage(parsed)
//	if err != nil {
//	    log.Fatalf("Validation failed: %v", err)
//	}
//
// Hooks:
//
// Developers can inject custom logic before and after validation using hooks:
//
//	SetPreValidationHook(func(action string, payload any) {
//	    log.Printf("[PRE] Validating action: %s", action)
//	})
//
//	SetPostValidationHook(func(action string, result any) {
//	    log.Printf("[POST] Validation complete for %s", action)
//	})
//
// Message Types:
//
// The OCPP 1.6J protocol defines three primary message types, which are recognized here:
//
//   - CALL       = 2  → request from charge point to central system
//   - CALLRESULT = 3  → successful response to a CALL
//   - CALLERROR  = 4  → error response to a CALL
//
// Relationship to Other Packages:
//
// This package is used internally by the ocpp16_messages package and intended
// for use cases where developers need to:
//
//   - Inspect and route OCPP traffic
//   - Register and execute custom validators for different message actions
//   - Handle JSON-based OCPP 1.6 message flow with a unified API
//
// Author:
//
//	Alexis Sanchez <aasanchez>
package core
