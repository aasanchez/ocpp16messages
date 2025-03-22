// Command plugin_demo demonstrates how to register a custom validator
// and a validation hook in the ocpp16_messages package to extend its functionality.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/core"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register a validator for a custom OCPP action
	core.RegisterValidator("CustomAction", validateCustomAction)

	// Attach a validation hook to observe lifecycle events
	core.SetValidationHook(LoggerHook{})

	// Send a valid CustomAction message
	raw := []byte(`[2, "uid123", "CustomAction", {"customField":"ok"}]`)
	result, err := ocpp16_messages.ValidateMessage(raw)
	exitOnError(err, "CustomAction validation failed")

	// Cast the payload back to your custom type
	if payload, ok := result.Payload.(CustomPayload); ok {
		fmt.Println("✅ CustomAction payload:", payload.CustomField)
	} else {
		log.Fatalf("❌ Payload is not CustomPayload")
	}
}

// CustomPayload is an example structure for a new OCPP action
type CustomPayload struct {
	CustomField string `json:"customField"`
}

// validateCustomAction is a custom validator for the "CustomAction" message.
func validateCustomAction(payload json.RawMessage) (interface{}, error) {
	var cp CustomPayload
	if err := json.Unmarshal(payload, &cp); err != nil {
		return nil, err
	}
	if cp.CustomField == "" {
		return nil, fmt.Errorf("customField is required")
	}
	return cp, nil
}

// LoggerHook implements the ValidationHook interface to log validation events.
type LoggerHook struct{}

// OnValidationSuccess logs when a message passes validation.
func (LoggerHook) OnValidationSuccess(action string, msgType int) {
	log.Printf("✅ Hook: %s validated successfully (msgType: %d)", action, msgType)
}

// OnValidationFailure logs when a message fails validation.
func (LoggerHook) OnValidationFailure(action string, msgType int, err error) {
	log.Printf("❌ Hook: %s failed validation (msgType: %d): %v", action, msgType, err)
}

// exitOnError logs and terminates the program on fatal validation error.
func exitOnError(err error, context string) {
	if err != nil {
		log.Fatalf("❌ %s: %v", context, err)
	}
}
