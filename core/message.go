// Package core provides the core types and validation infrastructure for OCPP 1.6J messages.
package core

import (
	"encoding/json"
	"fmt"
)

// Message represents a parsed OCPP message.
type Message struct {
	Type           MessageType
	ID             string
	Action         string
	Payload        json.RawMessage
	DecodedPayload interface{}
	Raw            []byte
}

// ValidateRawMessage parses and validates a raw OCPP message.
func ValidateRawMessage(raw []byte) (*Message, error) {
	var msgType MessageType
	var id string
	var action string
	var payload json.RawMessage

	// Parse the message array
	if err := json.Unmarshal(raw, &[]interface{}{&msgType, &id, &action, &payload}); err != nil {
		return nil, fmt.Errorf("invalid message format: %w", err)
	}

	// Validate message type
	if msgType != Call && msgType != CallResult && msgType != CallError {
		return nil, fmt.Errorf("invalid message type: %d", msgType)
	}

	// Validate ID
	if id == "" {
		return nil, fmt.Errorf("message ID is required")
	}

	// Validate action for Call messages
	if msgType == Call && action == "" {
		return nil, fmt.Errorf("action is required for Call messages")
	}

	// Validate payload
	if len(payload) == 0 {
		return nil, fmt.Errorf("payload is required")
	}

	return &Message{
		Type:    msgType,
		ID:      id,
		Action:  action,
		Payload: payload,
		Raw:     raw,
	}, nil
}
