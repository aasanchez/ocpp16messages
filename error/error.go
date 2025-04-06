package core

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

// CallErrorMessage represents an OCPP CALLERROR message [4, uniqueId, errorCode, errorDescription, errorDetails].
type CallErrorMessage struct {
	MessageTypeId    int             // Must be 4
	UniqueID         string          // Matches the UniqueID of the request
	ErrorCode        string          // One of the predefined or custom error codes
	ErrorDescription string          // Human-readable message
	ErrorDetails     json.RawMessage // Additional JSON info (optional)
}

// ValidateCallError validates the structure of a CALLERROR message.
//
// This function expects the message to be already unmarshaled into a []any,
// typically from a JSON array like: [4, "uniqueId", "errorCode", "description", {}]
func ValidateCallError(msg []any) (*CallErrorMessage, error) {
	const noEmptyString = "must be a non-empty string"

	if len(msg) != 5 {
		return nil, types.NewFieldError("CALLERROR", fmt.Sprintf("expected 5 elements, got %d", len(msg)))
	}

	// Validate MessageTypeId
	messageType, ok := msg[0].(int)
	if !ok || types.MessageType(messageType) != types.CALLERROR {
		return nil, types.NewFieldError("MessageTypeId", "must be 4 for CALLERROR")
	}

	uniqueID, ok := msg[1].(string)
	if !ok || uniqueID == "" {
		return nil, types.NewFieldError("UniqueID", noEmptyString)
	}

	errorCode, ok := msg[2].(string)
	if !ok || errorCode == "" {
		return nil, types.NewFieldError("ErrorCode", noEmptyString)
	}

	errorDescription, ok := msg[3].(string)
	if !ok || errorDescription == "" {
		return nil, types.NewFieldError("ErrorDescription", noEmptyString)
	}

	errorDetailsBytes, err := json.Marshal(msg[4])
	if err != nil {
		return nil, types.NewFieldError("ErrorDetails", "must be valid JSON object")
	}

	return &CallErrorMessage{
		MessageTypeId:    int(types.CALLERROR),
		UniqueID:         uniqueID,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorDetails:     errorDetailsBytes,
	}, nil
}
