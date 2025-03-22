package ocpp16_messages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core"
)

// MessageType represents the type of an OCPP 1.6 message.
type MessageType int

const (
	CALL       MessageType = 2
	CALLRESULT MessageType = 3
	CALLERROR  MessageType = 4
)

// RawOCPPMessage represents a parsed OCPP message in generic format.
type RawOCPPMessage struct {
	TypeID   MessageType
	UniqueID string
	Action   string
	Payload  any
}

// ValidateMessage parses and validates an OCPP 1.6 JSON message.
// It ensures the structure and payload conform to the expected format for the message type.
func ValidateMessage(raw []byte) (*RawOCPPMessage, error) {
	var data []json.RawMessage
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("invalid OCPP message format: %w", err)
	}

	if len(data) < 3 {
		return nil, errors.New("invalid OCPP message: not enough fields")
	}

	var typeID int
	if err := json.Unmarshal(data[0], &typeID); err != nil {
		return nil, fmt.Errorf("invalid MessageTypeId: %w", err)
	}

	switch MessageType(typeID) {
	case CALL:
		return handleCall(data)
	case CALLRESULT:
		return handleCallResult(data)
	case CALLERROR:
		return handleCallError(data)
	default:
		return nil, fmt.Errorf("unsupported MessageTypeId: %d", typeID)
	}
}

func handleCall(data []json.RawMessage) (*RawOCPPMessage, error) {
	if len(data) < 4 {
		return nil, errors.New("CALL message must have at least 4 elements")
	}

	var uniqueID, action string
	if err := json.Unmarshal(data[1], &uniqueID); err != nil {
		return nil, fmt.Errorf("invalid UniqueId: %w", err)
	}
	if err := json.Unmarshal(data[2], &action); err != nil {
		return nil, fmt.Errorf("invalid Action: %w", err)
	}

	// Lookup registered validator
	validator, ok := core.GetRegisteredValidator(action)
	if !ok {
		return nil, fmt.Errorf("no validator registered for action: %s", action)
	}

	// Decode payload into registered struct
	payload, err := core.DecodePayload(action, data[3])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// Validate the decoded message
	if err := validator.ValidateMessage(payload); err != nil {
		return nil, fmt.Errorf("payload validation failed: %w", err)
	}

	return &RawOCPPMessage{
		TypeID:   CALL,
		UniqueID: uniqueID,
		Action:   action,
		Payload:  payload,
	}, nil
}

func handleCallResult(data []json.RawMessage) (*RawOCPPMessage, error) {
	if len(data) < 3 {
		return nil, errors.New("CALLRESULT message must have 3 elements")
	}

	var uniqueID string
	if err := json.Unmarshal(data[1], &uniqueID); err != nil {
		return nil, fmt.Errorf("invalid UniqueId: %w", err)
	}

	var payload any
	if err := json.Unmarshal(data[2], &payload); err != nil {
		return nil, fmt.Errorf("invalid CALLRESULT payload: %w", err)
	}

	return &RawOCPPMessage{
		TypeID:   CALLRESULT,
		UniqueID: uniqueID,
		Payload:  payload,
	}, nil
}

func handleCallError(data []json.RawMessage) (*RawOCPPMessage, error) {
	if len(data) < 4 {
		return nil, errors.New("CALLERROR message must have 4 elements")
	}

	var uniqueID, errorCode, errorDescription string
	if err := json.Unmarshal(data[1], &uniqueID); err != nil {
		return nil, fmt.Errorf("invalid UniqueId: %w", err)
	}
	if err := json.Unmarshal(data[2], &errorCode); err != nil {
		return nil, fmt.Errorf("invalid ErrorCode: %w", err)
	}
	if err := json.Unmarshal(data[3], &errorDescription); err != nil {
		return nil, fmt.Errorf("invalid ErrorDescription: %w", err)
	}

	return &RawOCPPMessage{
		TypeID:   CALLERROR,
		UniqueID: uniqueID,
		Action:   errorCode, // Overloaded to carry errorCode
		Payload:  errorDescription,
	}, nil
}
