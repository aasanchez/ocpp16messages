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
func CustomValidator(payload json.RawMessage) (interface{}, error) {
	var msg CustomAuthorizeReq
	if err := json.Unmarshal(payload, &msg); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	if msg.Token == "" {
		return nil, core.NewFieldError("token", fmt.Errorf("required"))
	}
	return msg, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Register custom validator for action "CustomAuthorize"
	core.RegisterValidator("CustomAuthorize", CustomValidator)

	// Set up validation hook
	hook := &ValidationHook{}
	core.SetValidationHook(hook)

	raw := []byte(`[2,"20240322112233","CustomAuthorize",{"token":"xyz-123"}]`)

	result, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Fatalf("❌ CustomAuthorize validation failed: %v", err)
	}

	log.Println("✅ CustomAuthorize message validated")
	log.Printf("Action   : %s", result.Action)
	log.Printf("UniqueID : %s", result.ID)
	log.Printf("Payload  : %+v", result.DecodedPayload)
}

// ValidationHook implements core.ValidationHook interface
type ValidationHook struct{}

func (h *ValidationHook) OnValidationSuccess(action string, msgType int) {
	log.Printf("✅ Validation success: Action=%s Type=%d", action, msgType)
}

func (h *ValidationHook) OnValidationFailure(action string, msgType int, err error) {
	log.Printf("❌ Validation failure: Action=%s Type=%d Error=%v", action, msgType, err)
}
