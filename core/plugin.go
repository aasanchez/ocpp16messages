// Package core provides extension points and plugin infrastructure for the
// ocpp16_messages package.
//
// This file allows external tools to register custom message validators or
// hook into the validation lifecycle for metrics, debugging, or customization.
package core

import (
	"encoding/json"
	"fmt"
	"sync"
)

// MessageValidator defines a function signature for validating custom OCPP messages.
// The function receives a raw JSON payload (extracted from the CALL message),
// and must return a decoded, validated struct or an error.
type MessageValidator func(payload json.RawMessage) (interface{}, error)

// ValidationHook allows external tools to observe the success or failure
// of a validation process, for metrics, tracing, or debugging purposes.
type ValidationHook interface {
	OnValidationSuccess(action string, msgType int)
	OnValidationFailure(action string, msgType int, err error)
}

var (
	registryMu       sync.RWMutex
	customValidators = make(map[string]MessageValidator)
	validationHook   ValidationHook
)

// RegisterValidator adds a new custom message validator for the given action name.
// If a validator already exists for the action, this will overwrite it.
//
// Use this to support vendor-specific extensions or future OCPP messages.
func RegisterValidator(action string, validator MessageValidator) {
	registryMu.Lock()
	defer registryMu.Unlock()
	customValidators[action] = validator
}

// GetRegisteredValidator looks up a registered validator for the given action.
// Returns the validator and true if found, otherwise nil and false.
func GetRegisteredValidator(action string) (MessageValidator, bool) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	v, ok := customValidators[action]
	return v, ok
}

// SetValidationHook sets a global hook for observing validation outcomes.
// Only one hook is supported at a time (last set wins).
func SetValidationHook(hook ValidationHook) {
	registryMu.Lock()
	defer registryMu.Unlock()
	validationHook = hook
}

// NotifyValidationSuccess emits a success event to the registered validation hook.
func NotifyValidationSuccess(action string, msgType int) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	if validationHook != nil {
		validationHook.OnValidationSuccess(action, msgType)
	}
}

// NotifyValidationFailure emits a failure event to the registered validation hook.
func NotifyValidationFailure(action string, msgType int, err error) {
	registryMu.RLock()
	defer registryMu.RUnlock()
	if validationHook != nil {
		validationHook.OnValidationFailure(action, msgType, err)
	}
}

// ValidateMessage validates a message using the registered validator for the given action.
func ValidateMessage(action string, payload interface{}) error {
	validator, ok := GetRegisteredValidator(action)
	if !ok {
		return fmt.Errorf(`unknown message type "%s"`, action)
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	_, err = validator(raw)
	if err != nil {
		NotifyValidationFailure(action, 0, err)
		return err
	}
	NotifyValidationSuccess(action, 0)
	return nil
}
