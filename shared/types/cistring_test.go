package types

import (
	"strings"
	"testing"
)

const (
	errExpectedError          = "expected error: %v"
	errExpectedNonNilError    = "expected error for %s"
	errExpectedNoError        = "unexpected error: %v"
	errExpectedValueMismatch  = "expected Value() = %q, got %q"
	errExpectedValidationPass = "unexpected validation error: %v"
)

func TestCiString_EmptyFails(t *testing.T) {
	t.Parallel()

	input := ""

	_, err := CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "empty string")
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 21)

	_, err := CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "too long string")
	}
}

func TestCiString_NonPrintableFails(t *testing.T) {
	t.Parallel()

	input := "OCPP\x01"

	_, err := CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "non-printable characters")
	}
}

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "OCPP16-Test"

	_, err := CiString(input, 20)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"

	cs, _ := CiString(input, 20)
	if cs.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString20Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRST"

	_, err := CiString20(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("X", 21)

	_, err := CiString20(input)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "> 20 characters")
	}
}

func TestCiString20Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString20Optional("idTag", nil)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if val != nil {
		t.Errorf("expected nil result for nil input, got: %v", val)
	}
}

func TestCiString20Optional_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("X", 21)
	_, err := CiString20Optional("idTag", &input)

	if err == nil {
		t.Error("expected error for too long input, got nil")
	}
}

func TestCiString20Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "validTag"
	val, err := CiString20Optional("idTag", &input)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, val)
	}
}

func TestCiString25Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	_, err := CiString25(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString25Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString25Optional("parentIdTag", nil)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if val != nil {
		t.Errorf("expected nil result for nil input, got: %v", val)
	}
}

func TestCiString25Optional_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 26)
	_, err := CiString25Optional("parentIdTag", &input)

	if err == nil {
		t.Error("expected error for input > 25 characters")
	}
}

func TestCiString25Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "ParentTagOK"
	val, err := CiString25Optional("parentIdTag", &input)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, val)
	}
}

func TestCiString50Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	_, err := CiString50(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString50Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString50Optional("description", nil)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if val != nil {
		t.Errorf("expected nil result for nil input, got: %v", val)
	}
}

func TestCiString50Optional_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 51)
	_, err := CiString50Optional("description", &input)

	if err == nil {
		t.Error("expected error for input > 50 characters")
	}
}

func TestCiString50Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "ValidDescription50"
	val, err := CiString50Optional("description", &input)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, val)
	}
}

func TestCiString255Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	_, err := CiString255(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString500Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	_, err := CiString500(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString255Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString255Optional("note", nil)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if val != nil {
		t.Errorf("expected nil result for nil input, got: %v", val)
	}
}

func TestCiString255Optional_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 256)
	_, err := CiString255Optional("note", &input)

	if err == nil {
		t.Error("expected error for input > 255 characters")
	}
}

func TestCiString255Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 128)
	val, err := CiString255Optional("note", &input)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, val)
	}
}

func TestCiString500Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString500Optional("message", nil)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
	if val != nil {
		t.Errorf("expected nil result for nil input, got: %v", val)
	}
}

func TestCiString500Optional_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 501)
	_, err := CiString500Optional("message", &input)

	if err == nil {
		t.Error("expected error for input > 500 characters")
	}
}

func TestCiString500Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 300)
	val, err := CiString500Optional("message", &input)

	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, val)
	}
}
