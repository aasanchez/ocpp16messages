// Command plugin_demo demonstrates how to register a custom OCPP message validator
// using the plugin system in the ocpp16_messages/core package.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16_messages/core"
)

type CustomAuthorizeReq struct {
	Token string `json:"token"`
}

// CustomValidator implements the core.MessageValidator interface for a custom action.
type CustomValidator struct{}

// ValidateMessage validates a custom message format for the action "CustomAuthorize".
func (v CustomValidator) ValidateMessage(raw json.RawMessage) (any, error) {
	var msg CustomAuthorizeReq
	if err := json.Unmarshal(raw, &msg); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	if msg.Token == "" {
		return nil, core.NewFieldError("token", "cannot be empty")
	}
	return msg, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register custom validator for action "CustomAuthorize"
	core.RegisterValidator("CustomAuthorize", CustomValidator{})

	// Optionally, set hooks to log pre/post validation activity
	core.SetPreValidationHook(func(action string, payload any) {
		log.Printf("🔍 Pre-validation hook: Action=%s Payload=%v", action, payload)
	})
	core.SetPostValidationHook(func(action string, result any, err error) {
		log.Printf("✅ Post-validation hook: Action=%s Success=%v Error=%v", action, err == nil, err)
	})

	raw := []byte(`[2,"20240322112233","CustomAuthorize",{"token":"xyz-123"}]`)

	result, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Fatalf("❌ CustomAuthorize validation failed: %v", err)
	}

	log.Println("✅ CustomAuthorize message validated")
	log.Printf("Action   : %s", result.Action)
	log.Printf("UniqueID : %s", result.UniqueID)
	log.Printf("Payload  : %+v", result.Payload)
}
