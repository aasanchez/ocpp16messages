package types_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
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

	_, err := types.CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "empty string")
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 21)

	_, err := types.CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "too long string")
	}
}

func TestCiString_NonPrintableFails(t *testing.T) {
	t.Parallel()

	input := "OCPP\x01"

	_, err := types.CiString(input, 20)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "non-printable characters")
	}
}

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	input := "OCPP16-Test"

	_, err := types.CiString(input, 20)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"

	cs, _ := types.CiString(input, 20)
	if cs.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString20Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRST"

	_, err := types.CiString20(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("X", 21)

	_, err := types.CiString20(input)
	if err == nil {
		t.Errorf(errExpectedNonNilError, "> 20 characters")
	}
}

func TestCiString25Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	_, err := types.CiString25(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString50Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	_, err := types.CiString50(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString255Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	_, err := types.CiString255(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString500Type_CreateValid(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	_, err := types.CiString500(input)
	if err != nil {
		t.Errorf(errExpectedNoError, err)
	}
}

func TestCiString20Type_Value(t *testing.T) {
	t.Parallel()

	raw := strings.Repeat("A", 20)

	cs, _ := types.CiString20(raw)
	if cs.Value() != raw {
		t.Errorf(errExpectedValueMismatch, raw, cs.Value())
	}
}

func TestCiString20Type_Validate(t *testing.T) {
	t.Parallel()

	input := "OCPP20"

	cs, _ := types.CiString20(input)
	if err := cs.Validate(); err != nil {
		t.Errorf(errExpectedValidationPass, err)
	}
}

func TestCiString25Type_Value(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	cs, _ := types.CiString25(input)
	if cs.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString25Type_Validate(t *testing.T) {
	t.Parallel()

	input := "TestString"

	cs, _ := types.CiString25(input)
	if err := cs.Validate(); err != nil {
		t.Errorf(errExpectedValidationPass, err)
	}
}

func TestCiString50Type_Value(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	cs, _ := types.CiString50(input)
	if cs.Value() != input {
		t.Errorf(errExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString50Type_Validate(t *testing.T) {
	t.Parallel()

	cs, _ := types.CiString50("ShortValidString")
	if err := cs.Validate(); err != nil {
		t.Errorf(errExpectedValidationPass, err)
	}
}

func TestCiString255Type_Value(t *testing.T) {
	t.Parallel()

	raw := strings.Repeat("D", 255)
	cs, _ := types.CiString255(raw)

	if cs.Value() != raw {
		t.Errorf(errExpectedValueMismatch, raw, cs.Value())
	}
}

func TestCiString255Type_Validate(t *testing.T) {
	t.Parallel()

	cs, _ := types.CiString255("Basic test string")
	if err := cs.Validate(); err != nil {
		t.Errorf(errExpectedValidationPass, err)
	}
}

func TestCiString500Type_Value(t *testing.T) {
	t.Parallel()

	raw := strings.Repeat("E", 500)
	cs, _ := types.CiString500(raw)

	if cs.Value() != raw {
		t.Errorf(errExpectedValueMismatch, raw, cs.Value())
	}
}

func TestCiString500Type_Validate(t *testing.T) {
	t.Parallel()

	cs, _ := types.CiString500("This is a long string but valid for 500 length")
	if err := cs.Validate(); err != nil {
		t.Errorf(errExpectedValidationPass, err)
	}
}
