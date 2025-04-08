package core

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core/types"
	errorType "github.com/aasanchez/ocpp16_messages/error/types"
)

type CallErrorMessage struct {
	MessageTypeId    types.MessageType
	UniqueID         string
	ErrorCode        errorType.ErrorCode
	ErrorDescription string
	ErrorDetails     json.RawMessage
}

func ValidateCallError(msg []any) (*CallErrorMessage, error) {
	const noEmptyString = "must be a non-empty string"

	if len(msg) < 4 {
		return nil, types.NewFieldError("CALLERROR", fmt.Sprintf("expected at least 4 elements, got %d", len(msg)))
	}

	messageTypeVal, ok := msg[0].(int)
	if !ok {
		return nil, types.NewFieldError("MessageTypeId", "must be an integer")
	}
	messageType := types.MessageType(messageTypeVal)
	if !messageType.IsValid() {
		return nil, types.NewFieldError("MessageTypeId", fmt.Sprintf("invalid message type: %d", messageType))
	}
	if messageType != types.CALLERROR {
		return nil, types.NewFieldError("MessageTypeId", fmt.Sprintf("expected CALLERROR (%d), got %d", types.CALLERROR, messageType))
	}

	uniqueID, ok := msg[1].(string)
	if !ok || uniqueID == "" {
		return nil, types.NewFieldError("UniqueID", noEmptyString)
	}

	errorCodeStr, ok := msg[2].(string)
	if !ok || errorCodeStr == "" {
		return nil, types.NewFieldError("ErrorCode", noEmptyString)
	}

	errorCode := errorType.ErrorCode(errorCodeStr)
	if !errorCode.IsValid() {
		return nil, types.NewFieldError("ErrorCode", fmt.Sprintf("invalid error code: %s", errorCode))
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
		MessageTypeId:    types.CALLERROR,
		UniqueID:         uniqueID,
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
		ErrorDetails:     errorDetailsBytes,
	}, nil
}
