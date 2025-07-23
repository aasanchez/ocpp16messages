package sharedtypes

import (
	"errors"
	"testing"
)

func TestInternal_SharedTypes_string_Valid(t *testing.T) {
	t.Parallel()

	input := "OCPP-Test-123"
	maxLen := 20

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if got := cis.val(); got != input {
		t.Errorf("val() = %q; want %q", got, input)
	}
}

func TestInternal_SharedTypes_string_TooLong(t *testing.T) {
	t.Parallel()

	input := "ThisStringIsWayTooLong"
	maxLen := 5

	_, err := setCiString(input, maxLen)
	if err == nil {
		t.Fatal("expected error for exceeding max length, got nil")
	}

	if !errors.Is(err, errExceedsMaxLength) {
		t.Errorf("expected errExceedsMaxLength, got: %v", err)
	}
}

func TestInternal_SharedTypes_string_NonPrintableASCII(t *testing.T) {
	t.Parallel()

	input := "Valid\x01Invalid" // contains ASCII 1 (non-printable)

	_, err := setCiString(input, 20)
	if err == nil {
		t.Fatal("expected error for non-printable ASCII, got nil")
	}

	if !errors.Is(err, errInvalidASCII) {
		t.Errorf("expected errInvalidASCII, got: %v", err)
	}
}

func TestInternal_SharedTypes_string_EmptyIsValid(t *testing.T) {
	t.Parallel()

	input := ""
	maxLen := 10

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error for empty input, got: %v", err)
	}

	if got := cis.val(); got != "" {
		t.Errorf("val() = %q; want empty string", got)
	}
}

func TestInternal_SharedTypes_string_ASCIIBoundary(t *testing.T) {
	t.Parallel()

	// Characters 32 (space) and 126 (~) are valid
	input := " ~"
	maxLen := 5

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if got := cis.val(); got != input {
		t.Errorf("val() = %q; want %q", got, input)
	}
}
