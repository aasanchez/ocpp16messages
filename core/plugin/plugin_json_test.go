package plugin

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

// mockValidator is a basic implementation of MessageValidator for testing.
type mockValidator struct {
	shouldFail bool
}

func (m *mockValidator) ValidateMessage(payload json.RawMessage) (any, error) {
	if m.shouldFail {
		return nil, errors.New("mock validation error")
	}
	var result map[string]any
	err := json.Unmarshal(payload, &result)
	return result, err
}

func TestRegisterAndGetValidator(t *testing.T) {
	action := "TestAction"
	validator := &mockValidator{}
	RegisterValidator(action, validator)

	retrieved, ok := GetValidator(action)
	if !ok {
		t.Fatalf("expected validator to be registered for action %s", action)
	}
	if retrieved != validator {
		t.Fatalf("expected retrieved validator to match registered validator")
	}
}

func TestValidateRawJSON_Success(t *testing.T) {
	action := "SuccessAction"
	RegisterValidator(action, &mockValidator{})

	payload := json.RawMessage(`{"foo":"bar"}`)
	result, err := ValidateRawJSON(action, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resMap, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected result to be a map, got %T", result)
	}
	if resMap["foo"] != "bar" {
		t.Errorf("expected foo=bar, got foo=%v", resMap["foo"])
	}
}

func TestValidateRawJSON_ValidatorError(t *testing.T) {
	action := "FailAction"
	RegisterValidator(action, &mockValidator{shouldFail: true})

	payload := json.RawMessage(`{"foo":"bar"}`)
	result, err := ValidateRawJSON(action, payload)
	if err == nil {
		t.Fatal("expected error from validator, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestValidateRawJSON_NoValidator(t *testing.T) {
	action := "UnknownAction"
	payload := json.RawMessage(`{"test":123}`)

	result, err := ValidateRawJSON(action, payload)
	if err == nil {
		t.Fatal("expected error due to missing validator")
	}
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestPreAndPostValidationHooks(t *testing.T) {
	var (
		preCalled       bool
		postCalled      bool
		preAction       string
		postAction      string
		capturedPayload json.RawMessage
		capturedResult  any
		capturedError   error
		expectedPayload = json.RawMessage(`{"data":1}`)
		expectedResult  = map[string]any{"data": float64(1)}
		action          = "HookedAction"
	)

	SetPreValidationHook(func(action string, payload json.RawMessage) {
		preCalled = true
		preAction = action
		capturedPayload = payload
	})

	SetPostValidationHook(func(action string, result any, err error) {
		postCalled = true
		postAction = action
		capturedResult = result
		capturedError = err
	})

	RegisterValidator(action, &mockValidator{})

	result, err := ValidateRawJSON(action, expectedPayload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !preCalled || preAction != action || !reflect.DeepEqual(capturedPayload, expectedPayload) {
		t.Errorf("pre-validation hook not called correctly")
	}
	if !postCalled || postAction != action {
		t.Errorf("post-validation hook not called correctly")
	}
	if !reflect.DeepEqual(capturedResult, expectedResult) {
		t.Errorf("expected result %v, got %v", expectedResult, capturedResult)
	}
	if capturedError != nil {
		t.Errorf("expected no error, got %v", capturedError)
	}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("result mismatch: expected %v, got %v", expectedResult, result)
	}
}

func TestPostHookOnValidationFailure(t *testing.T) {
	called := false
	SetPostValidationHook(func(action string, result any, err error) {
		if err != nil && action == "PostFailAction" {
			called = true
		}
	})

	RegisterValidator("PostFailAction", &mockValidator{shouldFail: true})
	_, _ = ValidateRawJSON("PostFailAction", json.RawMessage(`{"x":1}`))

	if !called {
		t.Errorf("expected post-validation hook to be called on failure")
	}
}

func TestPostHookOnMissingValidator(t *testing.T) {
	called := false
	SetPostValidationHook(func(action string, result any, err error) {
		if err != nil && action == "MissingAction" {
			called = true
		}
	})

	_, _ = ValidateRawJSON("MissingAction", json.RawMessage(`{}`))

	if !called {
		t.Errorf("expected post-validation hook to be called on missing validator")
	}
}
