package plugin

import (
	"bytes"
	"errors"
	"testing"
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

	RegisterSOAPValidator(action, validator)

	got, ok := GetSOAPValidator(action)
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
	RegisterSOAPValidator("Authorize", validator)

	result, err := ValidateRawSOAP("Authorize", []byte(`<Authorize><idTag>abc</idTag></Authorize>`))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil {
		t.Error("expected result to be non-nil")
	}
}

func TestValidateRawSOAP_NoValidator(t *testing.T) {
	_, err := ValidateRawSOAP("UnknownAction", []byte(`<Unknown/>`))
	if err == nil {
		t.Error("expected error for unregistered validator")
	}
}

func TestSOAPHooks(t *testing.T) {
	var preCalled, postCalled bool

	SetPreSOAPValidationHook(func(action string, payload []byte) {
		if action == "HookAction" {
			preCalled = true
		}
	})

	SetPostSOAPValidationHook(func(action string, result any, err error) {
		if action == "HookAction" {
			postCalled = true
		}
	})

	validator := &mockSOAPValidator{}
	RegisterSOAPValidator("HookAction", validator)

	_, _ = ValidateRawSOAP("HookAction", []byte(`<Test/>`))

	if !preCalled {
		t.Error("expected pre-validation hook to be called")
	}
	if !postCalled {
		t.Error("expected post-validation hook to be called")
	}
}

func TestPostSOAPHookOnMissingValidator(t *testing.T) {
	called := false

	SetPostSOAPValidationHook(func(action string, result any, err error) {
		if action == "MissingAction" && err != nil {
			called = true
		}
	})

	_, _ = ValidateRawSOAP("MissingAction", []byte(`<nope/>`))

	if !called {
		t.Error("expected post SOAP hook to be called on missing validator")
	}
}
