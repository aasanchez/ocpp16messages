package plugin

import (
	"encoding/json"
	"fmt"
	"sync"
)

type MessageValidator interface {
	ValidateMessage(payload json.RawMessage) (any, error)
}

var (
	registry           = make(map[string]MessageValidator)
	preValidationHook  func(action string, payload json.RawMessage)
	postValidationHook func(action string, result any, err error)
	mu                 sync.RWMutex
)

func RegisterValidator(action string, validator MessageValidator) {
	mu.Lock()
	defer mu.Unlock()
	registry[action] = validator
}

func GetValidator(action string) (MessageValidator, bool) {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := registry[action]
	return v, ok
}

func SetPreValidationHook(hook func(action string, payload json.RawMessage)) {
	preValidationHook = hook
}

func SetPostValidationHook(hook func(action string, result any, err error)) {
	postValidationHook = hook
}

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
