package ocpp16_messages

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	CALL       = 2
	CALLRESULT = 3
	CALLERROR  = 4
)

type Message struct {
	MessageType int
	UniqueID    string
	Action      string
	Payload     any
}

func unmarshalField(raw json.RawMessage, v any) error {
	return json.Unmarshal(raw, v)
}

func processCall(raw []json.RawMessage, uniqueID string) (*Message, error) {
	if len(raw) != 4 {
		return nil, errors.New("CALL message must have 4 elements")
	}

	var action string
	if err := unmarshalField(raw[2], &action); err != nil {
		return nil, fmt.Errorf("invalid action field: %w", err)
	}

	switch action {
	case "Authorize":
		return nil, fmt.Errorf("no implemented action: %s", action)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}

func processCallResult(raw []json.RawMessage, uniqueID string) (*Message, error) {
	if len(raw) != 3 {
		return nil, errors.New("CALLRESULT must have 3 elements")
	}
	return &Message{MessageType: CALLRESULT, UniqueID: uniqueID, Action: "Unknown", Payload: raw[2]}, nil
}

func processCallError(raw []json.RawMessage, uniqueID string) (*Message, error) {
	if len(raw) != 5 {
		return nil, errors.New("CALLERROR must have 5 elements")
	}

	var action, errorCode, errorDescription string
	if err := unmarshalField(raw[2], &action); err != nil {
		return nil, fmt.Errorf("invalid action in CALLERROR: %w", err)
	}
	if err := unmarshalField(raw[3], &errorCode); err != nil {
		return nil, fmt.Errorf("invalid errorCode in CALLERROR: %w", err)
	}
	if err := unmarshalField(raw[4], &errorDescription); err != nil {
		return nil, fmt.Errorf("invalid errorDescription in CALLERROR: %w", err)
	}

	payload := map[string]string{
		"errorCode":        errorCode,
		"errorDescription": errorDescription,
	}
	return &Message{MessageType: CALLERROR, UniqueID: uniqueID, Action: action, Payload: payload}, nil
}

func ParseAndValidate(input []byte) (*Message, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(input, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON array: %w", err)
	}

	if len(raw) < 3 {
		return nil, errors.New("invalid OCPP message: expected at least 3 elements")
	}

	var messageType int
	if err := unmarshalField(raw[0], &messageType); err != nil {
		return nil, fmt.Errorf("invalid MessageTypeId: %w", err)
	}

	var uniqueID string
	if err := unmarshalField(raw[1], &uniqueID); err != nil {
		return nil, fmt.Errorf("invalid UniqueId: %w", err)
	}

	switch messageType {
	case CALL:
		return processCall(raw, uniqueID)
	case CALLRESULT:
		return processCallResult(raw, uniqueID)
	case CALLERROR:
		return processCallError(raw, uniqueID)
	default:
		return nil, fmt.Errorf("unsupported MessageTypeId: %d", messageType)
	}
}
