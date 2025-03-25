package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ParsedMessage is a general structure representing an OCPP message.
type ParsedMessage struct {
	TypeID           MessageType // 2: CALL, 3: CALLRESULT, 4: CALLERROR
	UniqueID         string
	Action           string          // for CALL
	Payload          json.RawMessage // for CALL and CALLRESULT
	ErrorCode        string          // for CALLERROR
	ErrorDescription string          // for CALLERROR
	ErrorDetails     json.RawMessage // for CALLERROR
}

// ParseJSONMessage parses a raw JSON OCPP message.
func ParseJSONMessage(data []byte) (*ParsedMessage, error) {
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
	case CALL:
		return parseCallMessage(raw, msg)
	case CALLRESULT:
		return parseCallResultMessage(raw, msg)
	case CALLERROR:
		return parseCallErrorMessage(raw, msg)
	default:
		return nil, fmt.Errorf("unsupported message type ID: %d", typeID)
	}
}

// parseRawMessage unmarshals the raw JSON into an array.
func parseRawMessage(data []byte) ([]json.RawMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON array: %w", err)
	}
	return raw, nil
}

// extractTypeAndID extracts the message type ID and unique ID.
func extractTypeAndID(raw []json.RawMessage) (MessageType, string, error) {
	var typeID MessageType
	if err := json.Unmarshal(raw[0], &typeID); err != nil {
		return 0, "", errors.New("invalid message type ID")
	}

	var uniqueID string
	if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
		return 0, "", errors.New("invalid unique ID")
	}

	return typeID, uniqueID, nil
}

// parseCallMessage parses a CALL message.
func parseCallMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) != 4 {
		return nil, errors.New("CALL message must have 4 elements")
	}
	var action string
	if err := json.Unmarshal(raw[2], &action); err != nil {
		return nil, errors.New("invalid action")
	}
	msg.Action = action
	msg.Payload = raw[3]
	return msg, nil
}

// parseCallResultMessage parses a CALLRESULT message.
func parseCallResultMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) != 3 {
		return nil, errors.New("CALLRESULT message must have 3 elements")
	}
	msg.Payload = raw[2]
	return msg, nil
}

// parseCallErrorMessage parses a CALLERROR message.
func parseCallErrorMessage(raw []json.RawMessage, msg *ParsedMessage) (*ParsedMessage, error) {
	if len(raw) != 5 {
		return nil, errors.New("CALLERROR message must have 5 elements")
	}
	if err := json.Unmarshal(raw[2], &msg.ErrorCode); err != nil {
		return nil, errors.New("invalid errorCode")
	}
	if err := json.Unmarshal(raw[3], &msg.ErrorDescription); err != nil {
		return nil, errors.New("invalid errorDescription")
	}
	msg.ErrorDetails = raw[4]
	return msg, nil
}
