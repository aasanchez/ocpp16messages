// Package ocpp16_messages provides the main entrypoint for decoding and validating
// OCPP 1.6J messages in both request and confirmation forms.
package ocpp16_messages

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core"
)

// MessageType defines the OCPP message types according to the OCPP 1.6J spec.
type MessageType int

const (
	// CALL represents a request message (type 2).
	CALL MessageType = 2
	// CALLRESULT represents a confirmation message (type 3).
	CALLRESULT MessageType = 3
	// CALLERROR represents an error message (type 4).
	CALLERROR MessageType = 4
)

// ValidateRawMessage validates a raw JSON OCPP message and returns a ParsedMessage
// that contains the message type, action, unique ID, and the validated payload.
//
// It supports messages of type CALL (2), CALLRESULT (3), and CALLERROR (4).
func ValidateRawMessage(raw []byte) (*core.ParsedMessage, error) {
	msg, err := core.ParseMessage(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OCPP message: %w", err)
	}

	switch msg.TypeID {
	case CALL:
		validator := core.GetRegisteredValidator(msg.Action)
		if validator == nil {
			return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
		}

		payload, err := validator.ValidateMessage(msg.Payload)
		if err != nil {
			return nil, fmt.Errorf("validation failed for action %s: %w", msg.Action, err)
		}

		msg.Payload = payload
		return msg, nil

	case CALLRESULT:
		validator := core.GetRegisteredValidator(msg.Action)
		if validator == nil {
			return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
		}

		payload, err := validator.ValidateMessage(msg.Payload)
		if err != nil {
			return nil, fmt.Errorf("validation failed for confirmation of action %s: %w", msg.Action, err)
		}

		msg.Payload = payload
		return msg, nil

	case CALLERROR:
		// Parse CALLERROR-specific structure (ErrorCode, ErrorDescription, Details)
		var callError core.CallErrorMessage
		if err := json.Unmarshal(msg.Payload, &callError); err != nil {
			return nil, fmt.Errorf("failed to parse CALLERROR message: %w", err)
		}
		msg.Payload = callError
		return msg, nil

	default:
		return nil, fmt.Errorf("unsupported message type ID: %d", msg.TypeID)
	}
}
