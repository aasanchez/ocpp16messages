package core_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

type mockSOAPValidator struct {
	called bool
}

func (v *mockSOAPValidator) ValidateSOAP(payload []byte) (any, error) {
	v.called = true
	if bytes.Contains(payload, []byte("fail")) {
		return nil, errors.New("validation failed")
	}
	return map[string]string{"parsed": "ok"}, nil
}

func TestRegisterAndGetSOAPValidator(t *testing.T) {
	action := "Authorize"
	validator := &mockSOAPValidator{}

	core.RegisterSOAPValidator(action, validator)

	got, ok := core.GetSOAPValidator(action)
	if !ok {
		t.Fatalf("expected validator to be found")
	}

	_, err := got.ValidateSOAP([]byte(`<Authorize><idTag>abc</idTag></Authorize>`))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !validator.called {
		t.Error("expected ValidateSOAP to be called")
	}
}

func TestValidateRawSOAP_Success(t *testing.T) {
	validator := &mockSOAPValidator{}
	core.RegisterSOAPValidator("Authorize", validator)

	result, err := core.ValidateRawSOAP("Authorize", []byte(`<Authorize><idTag>abc</idTag></Authorize>`))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Error("expected result to be non-nil")
	}
}

func TestValidateRawSOAP_NoValidator(t *testing.T) {
	_, err := core.ValidateRawSOAP("UnknownAction", []byte(`<Unknown/>`))
	if err == nil {
		t.Error("expected error for unregistered validator")
	}
}

func TestSOAPHooks(t *testing.T) {
	var preCalled, postCalled bool

	core.SetPreSOAPValidationHook(func(action string, payload []byte) {
		if action == "HookAction" {
			preCalled = true
		}
	})

	core.SetPostSOAPValidationHook(func(action string, result any, err error) {
		if action == "HookAction" {
			postCalled = true
		}
	})

	validator := &mockSOAPValidator{}
	core.RegisterSOAPValidator("HookAction", validator)

	_, _ = core.ValidateRawSOAP("HookAction", []byte(`<Test/>`))

	if !preCalled {
		t.Error("expected pre-validation hook to be called")
	}
	if !postCalled {
		t.Error("expected post-validation hook to be called")
	}
}
