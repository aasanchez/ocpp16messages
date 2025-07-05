package types

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

func TestCiString_ValidPasses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		maxLen int
	}{
		{
			name:  "valid string",
			value: "OCPP16-Test",
			maxLen: 20,
		},
		{
			name:  "empty string",
			value: "",
			maxLen: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := setCiString(tt.value, tt.maxLen)
			if err != nil {
				t.Errorf(ErrExpectedNoError, err)
			}
		})
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
