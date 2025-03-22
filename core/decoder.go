package core

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// registeredTypes maps OCPP action names to struct types for decoding.
var registeredTypes = make(map[string]reflect.Type)

// RegisterPayloadType registers a payload type for a given OCPP action.
// Typically called by init() in the message-specific package (e.g. authorize).
func RegisterPayloadType(action string, example any) {
	registeredTypes[action] = reflect.TypeOf(example)
}

// DecodePayload uses the registered type for the given action to decode the payload.
func DecodePayload(action string, raw json.RawMessage) (any, error) {
	typ, ok := registeredTypes[action]
	if !ok {
		return nil, fmt.Errorf("no payload type registered for action: %s", action)
	}

	ptr := reflect.New(typ).Interface()
	if err := json.Unmarshal(raw, ptr); err != nil {
		return nil, fmt.Errorf("failed to decode %s payload: %w", action, err)
	}

	return reflect.ValueOf(ptr).Elem().Interface(), nil
}
