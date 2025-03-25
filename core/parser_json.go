package core

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

// ParsedMessage is a general structure representing an OCPP message.
type ParsedMessage struct {
	TypeID           types.MessageType // 2: CALL, 3: CALLRESULT, 4: CALLERROR
	UniqueID         string
	Action           string          // for CALL
	Payload          json.RawMessage // for CALL and CALLRESULT
	ErrorCode        string          // for CALLERROR
	ErrorDescription string          // for CALLERROR
	ErrorDetails     json.RawMessage // for CALLERROR
}

// Define a constant for error formatting
const errUnexpected = "unexpected error: %v"

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
	case types.CALL:
		return parseCallMessage(raw, msg)
	case types.CALLRESULT:
		return parseCallResultMessage(raw, msg)
	case types.CALLERROR:
		return parseCallErrorMessage(raw, msg)
	default:
		return nil, fmt.Errorf("unsupported message type ID: %d", typeID)
	}
}

// parseRawMessage unmarshals the raw JSON into an array.
func parseRawMessage(data []byte) ([]json.RawMessage, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf(errUnexpected, err)
	}
	return raw, nil
}

// extractTypeAndID extracts the message type ID and unique ID.
func extractTypeAndID(raw []json.RawMessage) (types.MessageType, string, error) {
	var typeID types.MessageType
	if err := json.Unmarshal(raw[0], &typeID); err != nil {
		return 0, "", fmt.Errorf(errUnexpected, err)
	}

	var uniqueID string
	if err := json.Unmarshal(raw[1], &uniqueID); err != nil {
		return 0, "", fmt.Errorf(errUnexpected, err)
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
		return nil, fmt.Errorf(errUnexpected, err)
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

	var errorCode string
	if err := json.Unmarshal(raw[2], &errorCode); err != nil {
		return nil, errors.New("invalid errorCode")
	}
	msg.ErrorCode = errorCode

	var errorDescription string
	if err := json.Unmarshal(raw[3], &errorDescription); err != nil {
		return nil, errors.New("invalid errorDescription")
	}
	msg.ErrorDescription = errorDescription

	msg.ErrorDetails = raw[4]
	return msg, nil
}
