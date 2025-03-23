package core

import (
	"encoding/json"
	"fmt"
	"sync"
)

// MessageValidator defines an interface for validating raw JSON payloads.
type MessageValidator interface {
	ValidateMessage(json.RawMessage) (any, error)
}

var (
	mu                 sync.RWMutex
	registry           = make(map[string]MessageValidator)
	preValidationHook  func(action string, payload json.RawMessage)
	postValidationHook func(action string, result any, err error)
)

// RegisterValidator registers a MessageValidator for a given action.
func RegisterValidator(action string, validator MessageValidator) {
	mu.Lock()
	defer mu.Unlock()
	if validator == nil {
		delete(registry, action)
		return
	}
	registry[action] = validator
}

// GetValidator retrieves a validator by action name.
func GetValidator(action string) (MessageValidator, bool) {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := registry[action]
	return v, ok
}

// SetPreValidationHook sets a global hook executed before validation.
func SetPreValidationHook(hook func(action string, payload json.RawMessage)) {
	mu.Lock()
	defer mu.Unlock()
	preValidationHook = hook
}

// SetPostValidationHook sets a global hook executed after validation.
func SetPostValidationHook(hook func(action string, result any, err error)) {
	mu.Lock()
	defer mu.Unlock()
	postValidationHook = hook
}

// ValidateRawJSON runs the validation flow for a raw JSON payload by action.
func ValidateRawJSON(action string, payload json.RawMessage) (any, error) {
	mu.RLock()
	hook := preValidationHook
	validator, ok := registry[action]
	mu.RUnlock()

	if hook != nil {
		hook(action, payload)
	}

	if !ok {
		err := fmt.Errorf("no validator registered for action: %s", action)
		if postValidationHook != nil {
			postValidationHook(action, nil, err)
		}
		return nil, err
	}

	result, err := validator.ValidateMessage(payload)

	if post := postValidationHook; post != nil {
		post(action, result, err)
	}

	return result, err
}

// Optional helper for testing: resets registry and hooks.
// func resetPluginSystem() {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	registry = make(map[string]MessageValidator)
// 	preValidationHook = nil
// 	postValidationHook = nil
// }
