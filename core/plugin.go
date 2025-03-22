// Package core provides extensibility mechanisms for OCPP message validation,
// including plugin registration and lifecycle hooks.
package core

import (
	"encoding/json"
	"fmt"
	"sync"
)

// MessageValidator defines the interface that any custom validator must implement.
// It receives a raw JSON payload and returns a decoded message and error.
type MessageValidator interface {
	ValidateMessage(json.RawMessage) (any, error)
}

var (
	mu                 sync.RWMutex
	validators         = make(map[string]MessageValidator)
	preValidationHook  func(action string, payload any)
	postValidationHook func(action string, payload any, err error)
)

// RegisterValidator registers a custom validator for a specific OCPP action name.
func RegisterValidator(action string, validator MessageValidator) {
	mu.Lock()
	defer mu.Unlock()
	validators[action] = validator
}

// GetRegisteredValidator retrieves the validator for a given OCPP action, if registered.
func GetRegisteredValidator(action string) MessageValidator {
	mu.RLock()
	defer mu.RUnlock()
	return validators[action]
}

// SetPreValidationHook installs a callback hook that runs before any validation.
// Useful for instrumentation or debugging.
func SetPreValidationHook(hook func(action string, payload any)) {
	mu.Lock()
	defer mu.Unlock()
	preValidationHook = hook
}

// SetPostValidationHook installs a callback hook that runs after validation.
// It provides access to the action, payload, and validation result.
func SetPostValidationHook(hook func(action string, payload any, err error)) {
	mu.Lock()
	defer mu.Unlock()
	postValidationHook = hook
}

// ValidateRawMessage parses the raw message, finds the correct validator, and validates it.
func ValidateRawMessage(raw []byte) (*ParsedMessage, error) {
	msg, err := ParseMessage(raw)
	if err != nil {
		return nil, err
	}

	if msg.MessageType != CALL && msg.MessageType != CALLRESULT {
		return nil, nil // Not a message with a validator
	}

	validator := GetRegisteredValidator(msg.Action)
	if validator == nil {
		return nil, fmt.Errorf("no validator registered for action: %s", msg.Action)
	}

	if preValidationHook != nil {
		preValidationHook(msg.Action, msg.Payload)
	}

	decodedPayload, err := validator.ValidateMessage(msg.Payload)
	if postValidationHook != nil {
		postValidationHook(msg.Action, decodedPayload, err)
	}

	if err != nil {
		return nil, err
	}

	msg.Payload = nil // Remove raw payload
	return msg, nil
}
