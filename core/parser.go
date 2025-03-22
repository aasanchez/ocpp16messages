// Package core provides core parsing logic for OCPP 1.6J messages.
package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ParsedMessage represents a decoded OCPP 1.6J message with its components.
type ParsedMessage struct {
	MessageType MessageType
	UniqueID    string
	Action      string
	Payload     json.RawMessage
}

// ParseMessage parses a raw OCPP JSON message into its constituent parts according to OCPP 1.6J format.
func ParseMessage(data []byte) (*ParsedMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	if len(raw) < 3 {
		return nil, errors.New("message must have at least 3 elements")
	}

	var messageTypeID int
	if err := json.Unmarshal(raw[0], &messageTypeID); err != nil {
		return nil, fmt.Errorf("invalid MessageTypeId: %w", err)
	}

	switch MessageType(messageTypeID) {
	case CALL:
		if len(raw) != 4 {
			return nil, errors.New("CALL message must have 4 elements")
		}
		var uniqueID, action string
		if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
			return nil, fmt.Errorf("invalid UniqueID: %w", err)
		}
		if err := json.Unmarshal(raw[2], &action); err != nil {
			return nil, fmt.Errorf("invalid Action: %w", err)
		}
		return &ParsedMessage{
			MessageType: CALL,
			UniqueID:    uniqueID,
			Action:      action,
			Payload:     raw[3],
		}, nil

	case CALLRESULT:
		if len(raw) != 3 {
			return nil, errors.New("CALLRESULT message must have 3 elements")
		}
		var uniqueID string
		if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
			return nil, fmt.Errorf("invalid UniqueID: %w", err)
		}
		return &ParsedMessage{
			MessageType: CALLRESULT,
			UniqueID:    uniqueID,
			Payload:     raw[2],
		}, nil

	case CALLERROR:
		if len(raw) != 5 {
			return nil, errors.New("CALLERROR message must have 5 elements")
		}
		var uniqueID, errorCode, errorDescription string
		if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
			return nil, fmt.Errorf("invalid UniqueID: %w", err)
		}
		if err := json.Unmarshal(raw[2], &errorCode); err != nil {
			return nil, fmt.Errorf("invalid ErrorCode: %w", err)
		}
		if err := json.Unmarshal(raw[3], &errorDescription); err != nil {
			return nil, fmt.Errorf("invalid ErrorDescription: %w", err)
		}
		return &ParsedMessage{
			MessageType: CALLERROR,
			UniqueID:    uniqueID,
			Action:      errorCode,
			Payload:     raw[4],
		}, nil

	default:
		return nil, fmt.Errorf("unknown MessageTypeId: %d", messageTypeID)
	}
}
