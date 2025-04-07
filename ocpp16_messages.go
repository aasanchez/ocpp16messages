// Package ocpp16_messages provides parsing, validation, and introspection utilities
// for OCPP 1.6J and 1.6S messages over JSON (and optionally SOAP).
//
// This package is designed as a high-performance, extensible foundation for building
// OCPP 1.6-based Charging Station Management Systems (CSMS) and Charge Point applications.
//
// Core Features:
//
//   - Full support for OCPP 1.6J message structures over JSON arrays
//   - Strict message type checking for CALL, CALLRESULT, and CALLERROR
//   - Plugin-based extensibility for registering custom message validators
//   - Graceful error handling and informative feedback during message decoding
//
// This package currently supports message parsing and limited validation. Validation for
// specific message types (e.g., Authorize) can be implemented as plugins or extensions.
//
// Usage examples and test cases can be found under the `example/authorization/json` directory.
//
// Recommended Workflow:
//
//  1. Use ParseAndValidate to decode an incoming JSON OCPP message.
//  2. Examine the returned ParsedMessage, including its MessageType, UniqueID, and Payload.
//  3. If message validation is required, register a plugin or handle the action externally.
package ocpp16_messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MessageTypeID constants represent the three core message types defined by the OCPP 1.6 specification.
//
// The OCPP protocol uses numbered arrays for transmission. The first element in the array is always
// the MessageTypeId, which determines how the message should be parsed.
//
//	2 = CALL       - Request from client (e.g., a Charge Point) to server (CSMS)
//	3 = CALLRESULT - Successful response from server to a CALL
//	4 = CALLERROR  - Error response from server to a CALL
const (
	CALL       = 2
	CALLRESULT = 3
	CALLERROR  = 4
)

// ParsedMessage represents a fully decoded and partially interpreted OCPP message.
//
// This structure is used as a common representation of CALL, CALLRESULT, and CALLERROR messages,
// after being decoded from their JSON array structure.
//
// Fields:
//   - MessageType: one of CALL (2), CALLRESULT (3), CALLERROR (4)
//   - UniqueID: the correlation ID used to match requests and responses
//   - Action: for CALL and CALLERROR types, this indicates the operation (e.g., "Authorize")
//   - Payload: the message body or payload, unmarshaled as-is for further processing
type ParsedMessage struct {
	MessageType int
	UniqueID    string
	Action      string
	Payload     any
}

// ParseAndValidate parses an incoming OCPP 1.6J JSON message array,
// determines the message type, and validates its structure.
//
// The function supports the following MessageTypeId values:
//   - CALL (2): expects 4 elements [2, UniqueID, Action, Payload]
//   - CALLRESULT (3): expects 3 elements [3, UniqueID, Payload]
//   - CALLERROR (4): expects 5 elements [4, UniqueID, Action, ErrorCode, ErrorDescription]
//
// If the message type is CALL and the action is known but not yet implemented,
// the function will return a specific "not implemented" error.
//
// This function does not currently dispatch message payloads to detailed schema
// validation routines — that is expected to be handled via plugins or extensions.
//
// Returns:
//   - A ParsedMessage struct containing metadata and raw payload
//   - An error if the message is invalid, incomplete, or contains unsupported types
//
// Example usage:
//
//	input := []byte(`[2, "12345", "Authorize", {"idTag": "ABC123"}]`)
//	parsed, err := ParseAndValidate(input)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Received %s message with ID %s\n", parsed.Action, parsed.UniqueID)
func ParseAndValidate(input []byte) (*ParsedMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(input, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON array: %w", err)
	}

	if len(raw) < 3 {
		return nil, errors.New("invalid OCPP message: expected at least 3 elements")
	}

	// Extract MessageTypeId
	var messageType int
	if err := json.Unmarshal(raw[0], &messageType); err != nil {
		return nil, fmt.Errorf("invalid MessageTypeId: %w", err)
	}

	// Extract UniqueId
	var uniqueID string
	if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
		return nil, fmt.Errorf("invalid UniqueId: %w", err)
	}

	switch messageType {
	case CALL:
		if len(raw) != 4 {
			return nil, errors.New("CALL message must have 4 elements")
		}
		var action string
		if err := json.Unmarshal(raw[2], &action); err != nil {
			return nil, fmt.Errorf("invalid action field: %w", err)
		}

		switch action {
		case "Authorize":
			return nil, fmt.Errorf("no implemented action: %s", action)
		default:
			return nil, fmt.Errorf("unsupported action: %s", action)
		}

	case CALLRESULT:
		if len(raw) != 3 {
			return nil, errors.New("CALLRESULT must have 3 elements")
		}
		return &ParsedMessage{MessageType: CALLRESULT, UniqueID: uniqueID, Action: "Unknown", Payload: raw[2]}, nil

	case CALLERROR:
		if len(raw) != 5 {
			return nil, errors.New("CALLERROR must have 5 elements")
		}
		var action, errorCode, errorDescription string
		if err := json.Unmarshal(raw[2], &action); err != nil {
			return nil, fmt.Errorf("invalid action in CALLERROR: %w", err)
		}
		if err := json.Unmarshal(raw[3], &errorCode); err != nil {
			return nil, fmt.Errorf("invalid errorCode in CALLERROR: %w", err)
		}
		if err := json.Unmarshal(raw[4], &errorDescription); err != nil {
			return nil, fmt.Errorf("invalid errorDescription in CALLERROR: %w", err)
		}
		payload := map[string]string{
			"errorCode":        errorCode,
			"errorDescription": errorDescription,
		}
		return &ParsedMessage{MessageType: CALLERROR, UniqueID: uniqueID, Action: action, Payload: payload}, nil

	default:
		return nil, fmt.Errorf("unsupported MessageTypeId: %d", messageType)
	}
}
