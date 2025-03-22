package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

// FakeReq represents a mock request payload
type FakeReq struct {
	Name string
}

// FakeValidator is a mock plugin validator
type FakeValidator struct {
	Called bool
	Err    error
}

func (v *FakeValidator) ValidateMessage(msg any) error {
	v.Called = true
	return v.Err
}

func TestPlugin_RegisterAndValidate(t *testing.T) {
	validator := &FakeValidator{}
	core.RegisterValidator("FakeReq", validator)

	msg := FakeReq{Name: "Test"}
	err := core.ValidateMessage("FakeReq", msg)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !validator.Called {
		t.Errorf("expected validator to be called")
	}
}

func TestPlugin_UnknownType(t *testing.T) {
	err := core.ValidateMessage("Unknown", "payload")
	if err == nil || err.Error() != `unknown message type "Unknown"` {
		t.Errorf("expected unknown message type error, got %v", err)
	}
}

func TestPlugin_PreAndPostHooks(t *testing.T) {
	var preCalled, postCalled bool

	core.SetPreValidationHook(func(action string, payload any) {
		if action == "FakeReq" {
			preCalled = true
		}
	})
	core.SetPostValidationHook(func(action string, payload any, err error) {
		if action == "FakeReq" {
			postCalled = true
		}
	})

	_ = core.ValidateMessage("FakeReq", FakeReq{Name: "Demo"})

	if !preCalled {
		t.Errorf("expected pre-validation hook to be called")
	}
	if !postCalled {
		t.Errorf("expected post-validation hook to be called")
	}
}
