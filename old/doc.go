// Package core provides shared types, constants, utility functions, and validation logic
// for working with OCPP 1.6 messages in both JSON and SOAP formats.
//
// This package is the foundation for higher-level message handling and includes:
//
//   - Message parsing and validation for CALL, CALLRESULT, and CALLERROR messages
//   - Common types such as CiStringN, enums, and error structures
//   - Plugin support to allow validators to be registered dynamically
//   - JSON and SOAP transport abstraction helpers (e.g., header inspection, payload routing)
//
// Usage:
//
//	core.RegisterValidator("AuthorizeReq", authorize.ValidateAuthorizeReq)
//	core.RegisterValidator("AuthorizeConf", authorize.ValidateAuthorizeConf)
//
//	parsed, err := core.ParseMessage(rawMessage)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	result, err := core.ValidateMessage(parsed)
//
// Validation Hooks:
//
// Pre- and post-validation hooks can be registered globally:
//
//	core.SetPreValidationHook(func(action string, payload any) {
//	    log.Printf("Validating action: %s", action)
//	})
//
//	core.SetPostValidationHook(func(action string, result any) {
//	    log.Printf("Validated payload for %s: %+v", action, result)
//	})
//
// Message Types:
//
// The following message types are defined in the OCPP 1.6 protocol:
//   - CALL (2)
//   - CALLRESULT (3)
//   - CALLERROR (4)
//
// For SOAP support, this package offers basic header parsing and routing,
// but leaves full WS-Security and policy enforcement to higher-level layers.
//
// This package is used internally by the ocpp16_messages package, but can be reused
// in other projects requiring OCPP 1.6 protocol support.
//
// Author: Alexis Sanchez <aasanchez>
package core
