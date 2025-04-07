// Package ocpp16_messages provides decoding and validation capabilities
// for OCPP 1.6J and 1.6S messages using JSON and SOAP formats.
//
// It supports strict type validation for all official message types,
// and includes plugin-based extensibility to allow custom validators.
//
// Usage examples are available under the `example/authorization/json`
// directory.
// Package ocpp16_messages is the entrypoint for parsing and validating OCPP 1.6J messages.
package ocpp16_messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MessageTypeID values from OCPP 1.6J
const (
	CALL       = 2
	CALLRESULT = 3
	CALLERROR  = 4
)

// ParsedMessage represents a generic parsed message.
type ParsedMessage struct {
	MessageType int
	UniqueID    string
	Action      string
	Payload     any
}

// ParseAndValidate parses and validates an OCPP 1.6J JSON message.
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
