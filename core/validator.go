// Package core provides foundational validation logic for OCPP 1.6J messages.
// This file implements the main entrypoint for message validation using
// registered plugin validators and lifecycle hooks.
package core

import (
	"encoding/json"
	"fmt"
)

// ValidateRawMessage parses a raw JSON OCPP message and applies the registered validator.
//
// This is the primary entrypoint for decoding and validating an OCPP CALL message.
//
// Returns the parsed message structure or an error if validation fails.
func ValidateRawMessage(raw []byte) (*ParsedMessage, error) {
	msg, err := ParseMessage(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OCPP message: %w", err)
	}

	// Lookup validator for the action
	validator := GetRegisteredValidator(msg.Action)
	if validator == nil {
		return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
	}

	// Trigger pre-validation hook if available
	if preValidationHook != nil {
		preValidationHook(msg.Action, msg.Payload)
	}

	// Validate payload
	parsedPayload, err := validator.ValidateMessage(msg.Payload)
	if err != nil {
		return nil, fmt.Errorf("invalid %s payload: %w", msg.Action, err)
	}

	// Trigger post-validation hook if available
	if postValidationHook != nil {
		postValidationHook(msg.Action, parsedPayload)
	}

	msg.Payload = parsedPayload
	return msg, nil
}

// ValidateMessage decodes and validates a payload against the validator for the specified action.
//
// This is a lower-level utility and is primarily used by advanced consumers or plugin developers.
func ValidateMessage(action string, payload any) error {
	raw, ok := payload.(json.RawMessage)
	if !ok {
		return fmt.Errorf("expected json.RawMessage but got %T", payload)
	}

	validator := GetRegisteredValidator(action)
	if validator == nil {
		return fmt.Errorf("no validator registered for action: %s", action)
	}

	_, err := validator.ValidateMessage(raw)
	return err
}
