// Package core provides core message parsing utilities for handling OCPP 1.6J messages.
// It includes logic to parse JSON-encoded OCPP messages into structured forms for CALL,
// CALLRESULT, and CALLERROR message types.
package core

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

// ParsedMessage represents a parsed OCPP 1.6J message in its abstract form.
//
// It contains fields that vary depending on the message type:
//   - For CALL messages: Action and Payload are used.
//   - For CALLRESULT messages: Payload is used.
//   - For CALLERROR messages: ErrorCode, ErrorDescription, and ErrorDetails are used.
type ParsedMessage struct {
	TypeID           types.MessageType // 2: CALL, 3: CALLRESULT, 4: CALLERROR
	UniqueID         string            // Unique message identifier
	Action           string            // Used only for CALL messages
	Payload          json.RawMessage   // Payload for CALL and CALLRESULT messages
	ErrorCode        string            // Error code for CALLERROR messages
	ErrorDescription string            // Human-readable description for CALLERROR messages
	ErrorDetails     json.RawMessage   // Additional error details for CALLERROR messages
}

// errUnexpected is a reusable format string for wrapping unexpected errors.
const errUnexpected = "unexpected error: %v"

// ParseJsonMessage parses a raw JSON-encoded OCPP 1.6J message into a ParsedMessage structure.
//
// It supports the three standard message types defined by OCPP 1.6J:
//   - CALL (TypeID 2)
//   - CALLRESULT (TypeID 3)
//   - CALLERROR (TypeID 4)
//
// Returns an error if the message does not conform to expected structure or contains invalid data.
func ParseJsonMessage(data []byte) (*ParsedMessage, error) {
	raw, err := parseRawMessage(data)
	if err != nil {
		return nil, err
	}

	if len(raw) < 3 {
		return nil, errors.New("invalid OCPP message: must have at least 3 elements")
	}

	typeID, uniqueID, err := extractTypeAndID(raw)
	if err != nil {
		return nil, err
	}

	msg := &ParsedMessage{
		TypeID:   typeID,
		UniqueID: uniqueID,
	}

	switch typeID {
	case types.CALL:
		return parseCallMessage(raw, msg)
	case types.CALLRESULT:
		return parseCallResultMessage(raw, msg)
	case types.CALLERROR:
		return parseCallErrorMessage(raw, msg)
	default:
		return nil, fmt.Errorf("unsupported message type ID: %d", typeID)
	}
}

// parseRawMessage unmarshals a raw JSON message into a slice of raw JSON parts.
//
// The top-level JSON must be an array representing the structure of the OCPP message.
func parseRawMessage(data []byte) ([]json.RawMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf(errUnexpected, err)
	}
	return raw, nil
}

// extractTypeAndID extracts the message type ID and unique message identifier from
// the first two elements of the parsed JSON array.
func extractTypeAndID(raw []json.RawMessage) (types.MessageType, string, error) {
	var typeID types.MessageType
	if err := json.Unmarshal(raw[0], &typeID); err != nil {
		return 0, "", fmt.Errorf(errUnexpected, err)
	}

	var uniqueID string
	if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
		return 0, "", fmt.Errorf(errUnexpected, err)
	}

	return typeID, uniqueID, nil
}

// parseCallMessage parses a CALL message and extracts its action and payload.
//
// The message must contain exactly 4 elements: [2, uniqueID, action, payload].
func parseCallMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) != 4 {
		return nil, errors.New("CALL message must have 4 elements")
	}
	var action string
	if err := json.Unmarshal(raw[2], &action); err != nil {
		return nil, fmt.Errorf(errUnexpected, err)
	}
	msg.Action = action
	msg.Payload = raw[3]
	return msg, nil
}

// parseCallResultMessage parses a CALLRESULT message and extracts its payload.
//
// The message must contain exactly 3 elements: [3, uniqueID, payload].
func parseCallResultMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) != 3 {
		return nil, errors.New("CALLRESULT message must have 3 elements")
	}
	msg.Payload = raw[2]
	return msg, nil
}

// parseCallErrorMessage parses a CALLERROR message and extracts its error details.
//
// The message must contain exactly 5 elements:
// [4, uniqueID, errorCode, errorDescription, errorDetails].
func parseCallErrorMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) < 4 {
		return nil, errors.New("CALLERROR message must have at least 4 elements")
	}

	var errorCode string
	if err := json.Unmarshal(raw[2], &errorCode); err != nil {
		return nil, errors.New("invalid errorCode")
	}
	msg.ErrorCode = errorCode

	var errorDescription string
	if err := json.Unmarshal(raw[3], &errorDescription); err != nil {
		return nil, errors.New("invalid errorDescription")
	}
	msg.ErrorDescription = errorDescription

	msg.ErrorDetails = raw[4]
	return msg, nil
}
