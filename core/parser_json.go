package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ParsedMessage is a general structure representing an OCPP message.
type ParsedMessage struct {
	TypeID           MessageType     // 2: CALL, 3: CALLRESULT, 4: CALLERROR
	UniqueID         string          // Common to all message types
	Action           string          // CALL only
	Payload          json.RawMessage // CALL and CALLRESULT
	ErrorCode        string          // CALLERROR only
	ErrorDescription string          // CALLERROR only
	ErrorDetails     json.RawMessage // CALLERROR only
}

// ParseJSONMessage parses a raw JSON OCPP message and extracts its fields based on type.
func ParseJSONMessage(data []byte) (*ParsedMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("invalid JSON array: %w", err)
	}

	if len(raw) < 3 {
		return nil, errors.New("invalid OCPP message: must have at least 3 elements")
	}

	var typeID MessageType
	if err := json.Unmarshal(raw[0], &typeID); err != nil {
		return nil, errors.New("invalid message type ID")
	}
	if typeID != CALL && typeID != CALLRESULT && typeID != CALLERROR {
		return nil, fmt.Errorf("unsupported message type ID: %d", typeID)
	}

	var uniqueID string
	if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
		return nil, errors.New("invalid unique ID")
	}

	msg := &ParsedMessage{
		TypeID:   typeID,
		UniqueID: uniqueID,
	}

	switch typeID {
	case CALL:
		if len(raw) != 4 {
			return nil, errors.New("CALL message must have 4 elements")
		}
		var action string
		if err := json.Unmarshal(raw[2], &action); err != nil {
			return nil, errors.New("invalid action")
		}
		msg.Action = action
		msg.Payload = raw[3]

	case CALLRESULT:
		if len(raw) != 3 {
			return nil, errors.New("CALLRESULT message must have 3 elements")
		}
		msg.Payload = raw[2]

	case CALLERROR:
		if len(raw) != 5 {
			return nil, errors.New("CALLERROR message must have 5 elements")
		}
		var errorCode, errorDescription string
		if err := json.Unmarshal(raw[2], &errorCode); err != nil {
			return nil, errors.New("invalid errorCode")
		}
		if err := json.Unmarshal(raw[3], &errorDescription); err != nil {
			return nil, errors.New("invalid errorDescription")
		}
		msg.ErrorCode = errorCode
		msg.ErrorDescription = errorDescription
		msg.ErrorDetails = raw[4]
	}

	return msg, nil
}
