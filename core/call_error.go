// Package core provides the core types and validation infrastructure for OCPP 1.6J messages.
package core

import (
	"encoding/json"
	"fmt"
)

// CallErrorMessage represents an OCPP CALLERROR message.
type CallErrorMessage struct {
	MessageTypeID    MessageType `json:"-"`
	MessageID        string      `json:"-"`
	ErrorCode        string      `json:"errorCode"`
	ErrorDescription string      `json:"errorDescription"`
	ErrorDetails     interface{} `json:"errorDetails,omitempty"`
}

// ValidateCallError performs structural validation of the CALLERROR message.
func ValidateCallError(raw []byte) (*CallErrorMessage, error) {
	var msgType MessageType
	var id string
	var errCode string
	var errDesc string
	var errDetails interface{}

	// Parse the message array
	if err := json.Unmarshal(raw, &[]interface{}{&msgType, &id, &errCode, &errDesc, &errDetails}); err != nil {
		return nil, fmt.Errorf("invalid CALLERROR format: %w", err)
	}

	// Validate message type
	if msgType != CallError {
		return nil, fmt.Errorf("invalid message type for CALLERROR: %d", msgType)
	}

	// Validate required fields
	if id == "" {
		return nil, fmt.Errorf("message ID is required")
	}
	if errCode == "" {
		return nil, fmt.Errorf("error code is required")
	}
	if errDesc == "" {
		return nil, fmt.Errorf("error description is required")
	}

	return &CallErrorMessage{
		MessageTypeID:    msgType,
		MessageID:        id,
		ErrorCode:        errCode,
		ErrorDescription: errDesc,
		ErrorDetails:     errDetails,
	}, nil
}

// NewCallErrorMessage constructs a new CALLERROR message with standard formatting.
// This is typically used when validation fails or unexpected input is received.
func NewCallErrorMessage(code string, description string, details map[string]any) *CallErrorMessage {
	return &CallErrorMessage{
		MessageTypeID:    CallError,
		MessageID:        "", // Should be filled in by caller if responding to a specific request
		ErrorCode:        code,
		ErrorDescription: description,
		ErrorDetails:     details,
	}
}

// Marshal serializes the CallErrorMessage into a valid OCPP 1.6 CALLERROR message array.
func (e *CallErrorMessage) Marshal() ([]byte, error) {
	message := []any{
		int(CallError),
		e.MessageID,
		e.ErrorCode,
		e.ErrorDescription,
		e.ErrorDetails,
	}
	return json.Marshal(message)
}

// Error implements the error interface for CallErrorMessage.
func (e *CallErrorMessage) Error() string {
	return fmt.Sprintf("OCPP CALLERROR [%s]: %s", e.ErrorCode, e.ErrorDescription)
}
