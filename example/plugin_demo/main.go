package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aasanchez/ocpp16_messages/core"
)

// CustomPayload defines a sample payload structure.
type CustomPayload struct {
	Name string `json:"name"`
}

// CustomValidator validates CustomPayloads.
type CustomValidator struct{}

func (CustomValidator) ValidateMessage(payload any) error {
	p, ok := payload.(CustomPayload)
	if !ok {
		return fmt.Errorf("expected CustomPayload, got %T", payload)
	}
	if p.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func init() {
	core.RegisterPayloadType("CustomAction", CustomPayload{})
	core.RegisterValidator("CustomAction", CustomValidator{})

	core.SetPreValidationHook(func(action string, payload any) {
		log.Printf("🔍 Pre-validation: action=%s, payload=%+v", action, payload)
	})

	core.SetPostValidationHook(func(action string, payload any, err error) {
		if err != nil {
			log.Printf("❌ Post-validation failed: action=%s, err=%v", action, err)
		} else {
			log.Printf("✅ Post-validation passed: action=%s", action)
		}
	})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Simulate an incoming OCPP message
	raw := []byte(`[2, "demo-123", "CustomAction", {"name": "PluginUser"}]`)

	result, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Printf("❌ Validation failed: %v", err)
		os.Exit(1)
	}

	log.Println("✅ Custom plugin message validated successfully")
	fmt.Printf("📦 Parsed Message:\n  Unique ID: %s\n  Action: %s\n  Payload: %+v\n",
		result.UniqueID, result.Action, result.Payload)
}
