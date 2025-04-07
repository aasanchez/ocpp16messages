// Package ocpp16_messages offers a modular and extensible framework for decoding,
// validating, and working with Open Charge Point Protocol (OCPP) 1.6 messages.
//
// This package is intended for use in Central Systems, OCPP proxies, test harnesses,
// simulators, and any tooling that needs to parse and validate OCPP 1.6J messages.
//
// Features:
//
//   - Decodes CALL, CALLRESULT, and CALLERROR messages from JSON
//   - Strict type-safe modeling of message fields (e.g., CiStringN types)
//   - Modular message definitions (Authorize, BootNotification, etc.)
//   - Plugin system for custom validator registration
//   - Clear separation between core types, message logic, and transport details
//
// Directory Structure:
//
//	core/            → shared types, validation infrastructure
//	messages/        → message-specific types and validation logic
//	example/         → runnable examples for encoding/decoding/validation
//
// Getting Started:
//
// Use ParseAndValidate to parse a raw OCPP 1.6J message:
//
//	msg, err := ocpp16_messages.ParseAndValidate(input)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Then, use registered plugins or custom logic to interpret the message:
//
//	if msg.Action == "Authorize" {
//	    var req authorize.AuthorizeRequest
//	    json.Unmarshal(msg.Payload, &req)
//	    // validate or process req...
//	}
//
// Version:
//
// This documentation refers to version 0.1.0 of ocpp16_messages.
//
// Author: Alexis Sanchez <aasanchez>
package ocpp16_messages
