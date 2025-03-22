// Package core provides foundational validation logic for OCPP 1.6J messages.
// This file implements the main entrypoint for message validation using
// registered plugin validators and lifecycle hooks.
package core

import (
	"encoding/json"
	"fmt"
)

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
