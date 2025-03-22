package core

import (
	"encoding/json"
	"fmt"
)

// ParsedMessage represents a parsed raw OCPP message array.
type ParsedMessage struct {
	MessageTypeID int // 2 (CALL), 3 (CALLRESULT), 4 (CALLERROR)
	UniqueID      string
	Action        string
	Payload       json.RawMessage // Original JSON payload
	Decoded       any             // Optional decoded object (populated by plugin or validator)
}

// ParseMessage parses a raw OCPP 1.6J JSON message array into ParsedMessage.
func ParseMessage(data []byte) (*ParsedMessage, error) {
	var raw []any
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("malformed OCPP message: %w", err)
	}

	if len(raw) < 3 {
		return nil, fmt.Errorf("incomplete OCPP message")
	}

	messageType, ok := raw[0].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid messageTypeId format")
	}

	uniqueID, ok := raw[1].(string)
	if !ok {
		return nil, fmt.Errorf("invalid uniqueID format")
	}

	switch int(messageType) {
	case 2: // CALL
		if len(raw) != 4 {
			return nil, fmt.Errorf("CALL message must have 4 elements")
		}
		action, ok := raw[2].(string)
		if !ok {
			return nil, fmt.Errorf("invalid action format")
		}
		payloadBytes, err := json.Marshal(raw[3])
		if err != nil {
			return nil, fmt.Errorf("failed to re-marshal payload: %w", err)
		}
		return &ParsedMessage{
			MessageTypeID: int(messageType),
			UniqueID:      uniqueID,
			Action:        action,
			Payload:       payloadBytes,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported messageTypeId: %v", messageType)
	}
}
