package core

import (
	"encoding/json"
	"fmt"
	"sync"
)

var (
	registryLock         sync.RWMutex
	registeredValidators = make(map[string]MessageValidator)

	preValidationHook  func(action string, payload any)
	postValidationHook func(action string, payload any, err error)
)

// MessageValidator defines the interface for a plugin validator.
type MessageValidator interface {
	ValidateMessage(raw json.RawMessage) (any, error)
}

// RegisterValidator registers a plugin validator for a specific action.
func RegisterValidator(action string, validator MessageValidator) {
	registryLock.Lock()
	defer registryLock.Unlock()
	registeredValidators[action] = validator
}

// GetRegisteredValidator returns a registered validator for a given action.
func GetRegisteredValidator(action string) MessageValidator {
	registryLock.RLock()
	defer registryLock.RUnlock()
	return registeredValidators[action]
}

// SetPreValidationHook sets a hook called before message validation.
func SetPreValidationHook(hook func(action string, payload any)) {
	preValidationHook = hook
}

// SetPostValidationHook sets a hook called after message validation.
func SetPostValidationHook(hook func(action string, payload any, err error)) {
	postValidationHook = hook
}

// ValidateRawMessage parses and validates the OCPP message.
func ValidateRawMessage(raw []byte) (*ParsedMessage, error) {
	result, err := ParseMessage(raw)
	if err != nil {
		return nil, err
	}

	validator := GetRegisteredValidator(result.Action)
	if validator == nil {
		return nil, fmt.Errorf("no validator registered for action: %s", result.Action)
	}

	if preValidationHook != nil {
		preValidationHook(result.Action, result.Payload)
	}

	decodedPayload, err := validator.ValidateMessage(result.Payload)

	if postValidationHook != nil {
		postValidationHook(result.Action, result.Payload, err)
	}

	if err != nil {
		return nil, err
	}

	// Attach decoded payload separately
	result.Payload = result.Payload // keep raw for compatibility
	result.Decoded = decodedPayload // new field
	return result, nil
}
