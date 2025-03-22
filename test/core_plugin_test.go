package test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

// FakeValidator is a mock implementation of MessageValidator.
type FakeValidator struct {
	shouldFail bool
	payload    any
}

func (fv *FakeValidator) ValidateMessage(raw json.RawMessage) (any, error) {
	if fv.shouldFail {
		return nil, errors.New("validation failed")
	}
	return fv.payload, nil
}

func TestPlugin_RegisterAndValidate(t *testing.T) {
	core.RegisterValidator("TestAction", &FakeValidator{payload: "valid"})

	payload := json.RawMessage(`"test"`)
	result, err := core.ValidateRawMessage(payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Action != "TestAction" {
		t.Errorf("unexpected action: %s", result.Action)
	}
	if result.Payload != "valid" {
		t.Errorf("unexpected payload: %v", result.Payload)
	}
}

func TestPlugin_UnknownType(t *testing.T) {
	payload := json.RawMessage(`[2, "id", "UnknownAction", {}]`)
	_, err := core.ValidateRawMessage(payload)
	if err == nil || err.Error() != "no validator registered for action: UnknownAction" {
		t.Errorf("expected unknown action error, got: %v", err)
	}
}

func TestPlugin_PreAndPostHooks(t *testing.T) {
	var preCalled, postCalled bool

	core.SetPreValidationHook(func(action string, payload any) {
		preCalled = true
	})
	core.SetPostValidationHook(func(action string, payload any, err error) {
		postCalled = true
	})

	core.RegisterValidator("HookTest", &FakeValidator{payload: "hooked"})

	payload := json.RawMessage(`[2, "id", "HookTest", "data"]`)
	_, err := core.ValidateRawMessage(payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !preCalled {
		t.Error("expected pre-validation hook to be called")
	}
	if !postCalled {
		t.Error("expected post-validation hook to be called")
	}
}
