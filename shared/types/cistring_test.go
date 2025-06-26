package types

import (
	"errors"
	"strings"
	"testing"
)

func TestCiString_EmptyFails(t *testing.T) {
	t.Parallel()

	_, err := CiString("", 20)
	if !errors.Is(err, ErrEmptyValueNotAllowed) {
		t.Errorf("expected ErrEmptyValueNotAllowed, got: %v", err)
	}
}

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString(strings.Repeat("A", 21), 20)
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength, got: %v", err)
	}
}

func TestCiString_NonPrintableFails(t *testing.T) {
	t.Parallel()

	//lint:ignore ST1018 intentional control character
	_, err := CiString("OCPP\x01", 20) //nolint:staticcheck
	if !errors.Is(err, ErrNonPrintableASCII) {
		t.Errorf("expected ErrNonPrintableASCII, got: %v", err)
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
	cs, err := CiString(input, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

// --- Type-specific creation + validation tests ---

func TestCiString20Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRST"
	cs, err := CiString20(input)
	if err != nil {
		t.Fatalf("CiString20 creation failed: %v", err)
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("CiString20 validation failed: %v", err)
	}
	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	_, err := CiString20(strings.Repeat("X", 21))
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength for >20 chars, got: %v", err)
	}
}

func TestCiString25Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)
	cs, err := CiString25(input)
	if err != nil {
		t.Fatalf("CiString25 creation failed: %v", err)
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("CiString25 validation failed: %v", err)
	}
	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString50Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)
	cs, err := CiString50(input)
	if err != nil {
		t.Fatalf("CiString50 creation failed: %v", err)
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("CiString50 validation failed: %v", err)
	}
	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString255Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)
	cs, err := CiString255(input)
	if err != nil {
		t.Fatalf("CiString255 creation failed: %v", err)
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("CiString255 validation failed: %v", err)
	}
	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}

func TestCiString500Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)
	cs, err := CiString500(input)
	if err != nil {
		t.Fatalf("CiString500 creation failed: %v", err)
	}
	if err := cs.Validate(); err != nil {
		t.Errorf("CiString500 validation failed: %v", err)
	}
	if cs.Value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, cs.Value())
	}
}
