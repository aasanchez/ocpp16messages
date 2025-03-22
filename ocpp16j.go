// Package ocpp16_messages provides utilities to validate and decode OCPP 1.6J messages.
// It is designed to be transport-agnostic and focuses on schema-level validation of message structure.
package ocpp16_messages

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

// MessageType represents the OCPP MessageTypeId.
type MessageType int

const (
	Call       MessageType = 2 // Client-to-server request
	CallResult MessageType = 3 // Server-to-client response
	CallError  MessageType = 4 // Server-to-client error
)

// RawMessage represents a parsed OCPP message envelope after decoding.
type RawMessage struct {
	MessageType MessageType // 2, 3, or 4
	UniqueID    string      // Unique identifier for the message
	Action      string      // The OCPP action (e.g. "Authorize") – only used in CALL messages
	Payload     interface{} // The decoded message payload (e.g. AuthorizeReq or AuthorizeConf)
}

// ValidateMessage takes a raw OCPP JSON message (as []byte), parses it,
// validates the structure, and returns a decoded RawMessage or an error.
//
// This function only handles messages with MessageTypeId 2 (CALL) and 3 (CALLRESULT).
// MessageTypeId 4 (CALLERROR) is parsed but not validated.
func ValidateMessage(input []byte) (*RawMessage, error) {
	// First, unmarshal the raw JSON array
	var arr []json.RawMessage
	if err := json.Unmarshal(input, &arr); err != nil {
		return nil, fmt.Errorf("invalid OCPP message format: %w", err)
	}

	if len(arr) < 3 {
		return nil, errors.New("invalid message: must contain at least 3 elements")
	}

	var msgTypeID int
	if err := json.Unmarshal(arr[0], &msgTypeID); err != nil {
		return nil, errors.New("invalid message type ID")
	}

	switch msgTypeID {
	case int(Call):
		return validateCall(arr)
	case int(CallResult):
		return validateCallResult(arr)
	case int(CallError):
		// CALLERROR is not validated but parsed
		var uid, code, desc string
		var details map[string]interface{}
		if len(arr) < 4 {
			return nil, errors.New("invalid CALLERROR message format")
		}
		_ = json.Unmarshal(arr[1], &uid)
		if err := json.Unmarshal(arr[2], &code); err != nil {
			return nil, errors.New("CALLERROR: invalid errorCode field, must be string")
		}
		_ = json.Unmarshal(arr[3], &desc)
		_ = json.Unmarshal(arr[4], &details)
		return &RawMessage{
			MessageType: CallError,
			UniqueID:    uid,
			Action:      code, // Action here holds the error code
			Payload:     desc,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported MessageTypeId: %d", msgTypeID)
	}
}

// validateCall processes a MessageTypeId 2 (CALL) message
func validateCall(arr []json.RawMessage) (*RawMessage, error) {
	if len(arr) != 4 {
		return nil, errors.New("CALL message must contain 4 elements")
	}

	var uid, action string
	if err := json.Unmarshal(arr[1], &uid); err != nil {
		return nil, errors.New("invalid UniqueID in CALL message")
	}
	if err := json.Unmarshal(arr[2], &action); err != nil {
		return nil, errors.New("invalid Action in CALL message")
	}

	switch action {
	case "Authorize":
		var payload authorize.AuthorizeReq
		if err := json.Unmarshal(arr[3], &payload); err != nil {
			return nil, fmt.Errorf("invalid Authorize.req payload: %w", err)
		}
		if err := authorize.ValidateAuthorizeReq(payload); err != nil {
			return nil, fmt.Errorf("Authorize.req validation failed: %w", err)
		}
		return &RawMessage{
			MessageType: Call,
			UniqueID:    uid,
			Action:      action,
			Payload:     payload,
		}, nil
	default:
		// Try plugin registry
		if validator, ok := core.GetRegisteredValidator(action); ok {
			result, err := validator(arr[3])
			if err != nil {
				return nil, fmt.Errorf("%s.req plugin validation failed: %w", action, err)
			}
			return &RawMessage{
				MessageType: Call,
				UniqueID:    uid,
				Action:      action,
				Payload:     result,
			}, nil
		}
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}

// validateCallResult processes a MessageTypeId 3 (CALLRESULT) message
func validateCallResult(arr []json.RawMessage) (*RawMessage, error) {
	if len(arr) != 3 {
		return nil, errors.New("CALLRESULT message must contain 3 elements")
	}

	var uid string
	if err := json.Unmarshal(arr[1], &uid); err != nil {
		return nil, errors.New("invalid UniqueID in CALLRESULT")
	}

	// Currently only supports AuthorizeConf — future messages can be registered via plugin
	var payload authorize.AuthorizeConf
	if err := json.Unmarshal(arr[2], &payload); err != nil {
		return nil, fmt.Errorf("invalid Authorize.conf payload: %w", err)
	}
	if err := authorize.ValidateAuthorizeConf(payload); err != nil {
		return nil, fmt.Errorf("Authorize.conf validation failed: %w", err)
	}
	return &RawMessage{
		MessageType: CallResult,
		UniqueID:    uid,
		Action:      "Authorize",
		Payload:     payload,
	}, nil
}
