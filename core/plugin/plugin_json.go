// Package core provides infrastructure for validating OCPP messages, including
// support for plugin-style message validators.
package plugin

import (
	"encoding/json"
	"fmt"
	"sync"
)

// MessageValidator defines a function capable of validating a JSON payload
// and returning a typed representation or an error.
type MessageValidator interface {
	ValidateMessage(payload json.RawMessage) (any, error)
}

// validatorRegistry holds registered validators keyed by action name.
var (
	registry           = make(map[string]MessageValidator)
	preValidationHook  func(action string, payload json.RawMessage)
	postValidationHook func(action string, result any, err error)
	mu                 sync.RWMutex
)

// RegisterValidator registers a MessageValidator for a specific OCPP action.
// It is safe for concurrent use.
func RegisterValidator(action string, validator MessageValidator) {
	mu.Lock()
	defer mu.Unlock()
	registry[action] = validator
}

// GetValidator retrieves a validator for the given action name.
func GetValidator(action string) (MessageValidator, bool) {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := registry[action]
	return v, ok
}

// SetPreValidationHook sets a function to be called before validation.
func SetPreValidationHook(hook func(action string, payload json.RawMessage)) {
	preValidationHook = hook
}

// SetPostValidationHook sets a function to be called after validation.
func SetPostValidationHook(hook func(action string, result any, err error)) {
	postValidationHook = hook
}

// ValidateRawJSON validates a raw JSON payload for the given action.
// It invokes pre- and post-hooks if set.
func ValidateRawJSON(action string, payload json.RawMessage) (any, error) {
	if preValidationHook != nil {
		preValidationHook(action, payload)
	}

	validator, ok := GetValidator(action)
	if !ok {
		err := fmt.Errorf("no validator registered for action: %s", action)
		if postValidationHook != nil {
			postValidationHook(action, nil, err)
		}
		return nil, err
	}

	result, err := validator.ValidateMessage(payload)
	if postValidationHook != nil {
		postValidationHook(action, result, err)
	}
	return result, err
}
