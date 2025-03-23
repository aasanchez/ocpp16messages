// Package ocpp16j_messages provides entrypoint parsing and validation for OCPP 1.6J messages.
package ocpp16j_messages

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core"
)

// ParseAndValidateOCPPMessage takes a raw OCPP 1.6J JSON message (CALL, CALLRESULT, or CALLERROR),
// validates the structure and fields, and returns a structured Message or error.
func ParseAndValidateOCPPMessage(data []byte) (*core.Message, error) {
	msg, err := core.ValidateRawMessage(data)
	if err != nil {
		return nil, err
	}

	switch msg.Type {
	case core.Call:
		// Lookup validator for this action
		validator, ok := core.GetRegisteredValidator(msg.Action)
		if !ok {
			return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
		}

		// Run the validator and assign decoded payload
		decoded, err := validator(msg.Payload)
		if err != nil {
			return nil, fmt.Errorf("invalid %s.req payload: %w", msg.Action, err)
		}

		msg.DecodedPayload = decoded
		return msg, nil

	case core.CallResult:
		// Lookup validator for the matching response type
		validator, ok := core.GetRegisteredValidator(msg.Action)
		if !ok {
			return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
		}

		decoded, err := validator(msg.Payload)
		if err != nil {
			return nil, fmt.Errorf("invalid %s.conf payload: %w", msg.Action, err)
		}

		msg.DecodedPayload = decoded
		return msg, nil

	case core.CallError:
		// Validate CALLERROR message structure
		var errMsg core.CallErrorMessage
		if err := json.Unmarshal(msg.Raw, &errMsg); err != nil {
			return nil, fmt.Errorf("invalid CALLERROR structure: %w", err)
		}
		msg.DecodedPayload = errMsg
		return msg, nil

	default:
		return nil, fmt.Errorf("unsupported OCPP message type: %d", msg.Type)
	}
}
