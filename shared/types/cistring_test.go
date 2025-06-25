package types

import (
	"strings"
	"testing"
)

func TestCiString_EmptyFails(t *testing.T) {
	t.Parallel()

	_, err := CiString("", 20)
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "empty string")
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString(strings.Repeat("A", 21), 20)
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "too long string")
	}
}

func TestCiString_NonPrintableFails(t *testing.T) {
	t.Parallel()
	//lint:ignore ST1018 intentional control character to test validation of non-printable input
	_, err := CiString("OCPP", 20) //nolint:staticcheck
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "non-printable characters")
	}
}

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	_, err := CiString("OCPP16-Test", 20)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"
	cs, _ := CiString(input, 20)

	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString20Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString20("ABCDEFGHIJKLMNOPQRST")
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("X", 21))
	if err == nil {
		t.Errorf(ErrExpectedNonNilError, "> 20 characters")
	}
}

func TestCiString20Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString20Optional("idTag", nil)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val != nil {
		t.Errorf(ErrExpectedNil, val)
	}
}

func TestCiString20Optional_ErrorIncludesFieldName(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("X", 21)
	val, err := CiString20Optional("idTag", &input)
	_ = val

	if err == nil {
		t.Fatal("expected error for too long idTag input, got nil")
	}

	if !strings.Contains(err.Error(), "idTag") {
		t.Errorf("expected error to include field name 'idTag', got: %v", err)
	}
}

func TestCiString20Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "validTag"
	val, err := CiString20Optional("idTag", &input)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, val)
	}
}

func TestCiString25Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString25(strings.Repeat("B", 25))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString25Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString25Optional("parentIdTag", nil)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val != nil {
		t.Errorf(ErrExpectedNil, val)
	}
}

func TestCiString25Optional_ErrorIncludesFieldName(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("Y", 26)
	val, err := CiString25Optional("parentIdTag", &input)
	_ = val

	if err == nil {
		t.Fatal("expected error for too long parentIdTag input, got nil")
	}

	if !strings.Contains(err.Error(), "parentIdTag") {
		t.Errorf("expected error to include field name 'parentIdTag', got: %v", err)
	}
}

func TestCiString25Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "ParentTagOK"
	val, err := CiString25Optional("parentIdTag", &input)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, val)
	}
}

func TestCiString50Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString50(strings.Repeat("C", 50))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString50Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString50Optional("description", nil)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val != nil {
		t.Errorf(ErrExpectedNil, val)
	}
}

func TestCiString50Optional_ErrorIncludesFieldName(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("Z", 51)
	val, err := CiString50Optional("firmwareVersion", &input)
	_ = val

	if err == nil {
		t.Fatal("expected error for too long firmwareVersion input, got nil")
	}

	if !strings.Contains(err.Error(), "firmwareVersion") {
		t.Errorf("expected error to include field name 'firmwareVersion', got: %v", err)
	}
}

func TestCiString50Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "ValidDescription50"
	val, err := CiString50Optional("description", &input)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, val)
	}
}

func TestCiString255Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString255(strings.Repeat("D", 255))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString255Optional_ErrorIncludesFieldName(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("N", 256)
	val, err := CiString255Optional("note", &input)
	_ = val

	if err == nil {
		t.Fatal("expected error for too long note input, got nil")
	}

	if !strings.Contains(err.Error(), "note") {
		t.Errorf("expected error to include field name 'note', got: %v", err)
	}
}

func TestCiString255Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 128)
	val, err := CiString255Optional("note", &input)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val == nil {
		t.Fatal("expected non-nil result for valid input")
	}

	if val.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, val.Value())
	}
}

func TestCiString500Type_CreateValid(t *testing.T) {
	t.Parallel()

	_, err := CiString500(strings.Repeat("E", 500))
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString500Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString500Optional("message", nil)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val != nil {
		t.Errorf(ErrExpectedNil, val)
	}
}

func TestCiString500Optional_ErrorIncludesFieldName(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("M", 501)
	val, err := CiString500Optional("message", &input)
	_ = val

	if err == nil {
		t.Fatal("expected error for too long message input, got nil")
	}

	if !strings.Contains(err.Error(), "message") {
		t.Errorf("expected error to include field name 'message', got: %v", err)
	}
}

func TestCiString500Optional_ValidPasses(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 300)
	val, err := CiString500Optional("message", &input)

	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val == nil || val.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, val)
	}
}

func TestCiString50Optional_InvalidInputTriggersError(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("X", 51)
	_, err := CiString50Optional("field50", &input)

	if err == nil {
		t.Fatal("expected error for invalid input to CiString50Optional, got nil")
	}

	if !strings.Contains(err.Error(), "field50") {
		t.Errorf("expected error to contain field name 'field50', got: %v", err)
	}
}

func TestCiString255Optional_InvalidInputTriggersError(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("Y", 256)
	_, err := CiString255Optional("field255", &input)

	if err == nil {
		t.Fatal("expected error for invalid input to CiString255Optional, got nil")
	}

	if !strings.Contains(err.Error(), "field255") {
		t.Errorf("expected error to contain field name 'field255', got: %v", err)
	}
}

func TestCiString500Optional_InvalidInputTriggersError(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("Z", 501)
	_, err := CiString500Optional("field500", &input)

	if err == nil {
		t.Fatal("expected error for invalid input to CiString500Optional, got nil")
	}

	if !strings.Contains(err.Error(), "field500") {
		t.Errorf("expected error to contain field name 'field500', got: %v", err)
	}
}

func TestCiString20Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString20("TestString20")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString25Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString25("TestString25ValidDataHere")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString50Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString50("This string is valid and under 50 chars")
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString255Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString255(strings.Repeat("A", 100))
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString500Type_ValidatePasses(t *testing.T) {
	t.Parallel()

	cs, _ := CiString500(strings.Repeat("B", 200))
	if err := cs.Validate(); err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString255Optional_NilInputReturnsNil(t *testing.T) {
	t.Parallel()

	val, err := CiString255Optional("note", nil)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}

	if val != nil {
		t.Errorf(ErrExpectedNil, val)
	}
}
