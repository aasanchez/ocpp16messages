package sharedtypes

import (
	"errors"
	"strings"
	"testing"
)

func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	_, err := setCiString(strings.Repeat("A", 21), 20)
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength, got: %v", err)
	}
}

func TestCiString_ValidStringPasses(t *testing.T) {
	t.Parallel()

	_, err := setCiString("OCPP16-Test", 20)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString_EmptyStringPasses(t *testing.T) {
	t.Parallel()

	_, err := setCiString("", 20)
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	input := "OCPP-Valid"
	ciStr, err := setCiString(input, 20)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ciStr.value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.value())
	}
}

func TestCiString_NonPrintableCharacterFails(t *testing.T) {
	t.Parallel()

	// Using ASCII Bell character (code 7)
	_, err := setCiString("Hello\aWorld", 20)
	if !errors.Is(err, ErrNonPrintableASCII) {
		t.Errorf("expected ErrNonPrintableASCII, got: %v", err)
	}
}
