package types

import (
	"errors"
	"strings"
	"testing"
)

func TestCiString_EmptyFails(t *testing.T) {
	t.Parallel()

	_, err := SetCiString("", 20)
	if !errors.Is(err, ErrEmptyValueNotAllowed) {
		t.Errorf("expected ErrEmptyValueNotAllowed, got: %v", err)
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	_, err := SetCiString(strings.Repeat("A", 21), 20)
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength, got: %v", err)
	}
}

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	_, err := SetCiString("OCPP16-Test", 20)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"
	ciStr, err := SetCiString(input, 20)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ciStr.value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.value())
	}
}

func TestCiString20Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRST"
	ciStr, err := SetCiString20(input)

	if err != nil {
		t.Fatalf("CiString20 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString20 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	_, err := SetCiString20(strings.Repeat("X", 21))
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength for >20 chars, got: %v", err)
	}
}

func TestCiString25Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)
	ciStr, err := SetCiString25(input)

	if err != nil {
		t.Fatalf("CiString25 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString25 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString50Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)
	ciStr, err := SetCiString50(input)

	if err != nil {
		t.Fatalf("CiString50 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString50 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString255Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)
	ciStr, err := SetCiString255(input)

	if err != nil {
		t.Fatalf("CiString255 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString255 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString500Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)
	ciStr, err := SetCiString500(input)

	if err != nil {
		t.Fatalf("CiString500 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString500 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString_NonPrintableCharacterFails(t *testing.T) {
	t.Parallel()

	// Using ASCII Bell character (code 7)
	_, err := SetCiString("Hello\aWorld", 20)
	if !errors.Is(err, ErrNonPrintableASCII) {
		t.Errorf("expected ErrNonPrintableASCII, got: %v", err)
	}
}
