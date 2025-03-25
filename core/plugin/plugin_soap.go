package plugin

import (
	"fmt"
	"sync"
)

// SOAPValidator defines the interface for validating raw SOAP payloads.
type SOAPValidator interface {
	ValidateSOAP(payload []byte) (any, error)
}

var (
	soapRegistry   = make(map[string]SOAPValidator)
	soapRegistryMu sync.RWMutex
	preSOAPHook    func(action string, payload []byte)
	postSOAPHook   func(action string, result any, err error)
)

// RegisterSOAPValidator registers a SOAP validator for a given OCPP action.
func RegisterSOAPValidator(action string, validator SOAPValidator) {
	soapRegistryMu.Lock()
	defer soapRegistryMu.Unlock()
	soapRegistry[action] = validator
}

// GetSOAPValidator retrieves the registered SOAP validator for an action.
func GetSOAPValidator(action string) (SOAPValidator, bool) {
	soapRegistryMu.RLock()
	defer soapRegistryMu.RUnlock()
	v, ok := soapRegistry[action]
	return v, ok
}

// SetPreSOAPValidationHook allows a hook to be run before SOAP validation.
func SetPreSOAPValidationHook(hook func(action string, payload []byte)) {
	preSOAPHook = hook
}

// SetPostSOAPValidationHook allows a hook to be run after SOAP validation.
func SetPostSOAPValidationHook(hook func(action string, result any, err error)) {
	postSOAPHook = hook
}

// ValidateRawSOAP runs validation on a raw SOAP payload for a given action.
func ValidateRawSOAP(action string, payload []byte) (any, error) {
	if preSOAPHook != nil {
		preSOAPHook(action, payload)
	}

	validator, ok := GetSOAPValidator(action)
	if !ok {
		err := fmt.Errorf("no SOAP validator registered for action: %s", action)
		if postSOAPHook != nil {
			postSOAPHook(action, nil, err)
		}
		return nil, err
	}

	result, err := validator.ValidateSOAP(payload)

	if postSOAPHook != nil {
		postSOAPHook(action, result, err)
	}

	return result, err
}
