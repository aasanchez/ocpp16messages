package test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

// FakeValidator returns a validator function for testing purposes.
func FakeValidator(shouldFail bool) core.MessageValidator {
	return func(payload json.RawMessage) (interface{}, error) {
		if shouldFail {
			return nil, core.NewFieldError("test", errors.New("invalid field"))
		}
		return nil, nil
	}
}

func TestPlugin_RegisterAndValidate(t *testing.T) {
	core.RegisterValidator("CustomAction", FakeValidator(false))

	msg := &core.Message{
		Action:  "CustomAction",
		Payload: json.RawMessage(`{"test": "value"}`),
	}

	err := core.ValidateMessage(msg.Action, msg.Payload)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPlugin_UnknownType(t *testing.T) {
	msg := &core.Message{
		Action:  "UnknownAction",
		Payload: json.RawMessage(`{}`),
	}

	err := core.ValidateMessage(msg.Action, msg.Payload)
	if err == nil {
		t.Error("expected error for unknown validator, got nil")
	}
}

func TestPlugin_ValidationHooks(t *testing.T) {
	var successCalled, failureCalled bool

	core.SetValidationHook(&testHook{
		onSuccess: func(action string, msgType int) {
			successCalled = true
		},
		onFailure: func(action string, msgType int, err error) {
			failureCalled = true
		},
	})

	// Test success case
	core.RegisterValidator("HookAction", FakeValidator(false))
	msg := &core.Message{
		Action:  "HookAction",
		Payload: json.RawMessage(`{}`),
	}
	err := core.ValidateMessage(msg.Action, msg.Payload)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !successCalled {
		t.Error("expected success hook to be called")
	}

	// Test failure case
	core.RegisterValidator("HookAction", FakeValidator(true))
	err = core.ValidateMessage(msg.Action, msg.Payload)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if !failureCalled {
		t.Error("expected failure hook to be called")
	}
}

type testHook struct {
	onSuccess func(action string, msgType int)
	onFailure func(action string, msgType int, err error)
}

func (h *testHook) OnValidationSuccess(action string, msgType int) {
	if h.onSuccess != nil {
		h.onSuccess(action, msgType)
	}
}

func (h *testHook) OnValidationFailure(action string, msgType int, err error) {
	if h.onFailure != nil {
		h.onFailure(action, msgType, err)
	}
}
