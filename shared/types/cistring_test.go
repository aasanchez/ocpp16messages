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

	input := "Valid\x01Invalid"

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

// CiString20Type

func TestSetCiString20Type_Valid(t *testing.T) {
	val := "ValidString"
	_, err := SetCiString20Type(val)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSetCiString20Type_MaxLength(t *testing.T) {
	val := "12345678901234567890" // exactly 20 chars
	_, err := SetCiString20Type(val)
	if err != nil {
		t.Fatalf("expected no error for 20-char string, got %v", err)
	}
}

func TestSetCiString20Type_TooLong(t *testing.T) {
	val := "123456789012345678901" // 21 chars
	_, err := SetCiString20Type(val)
	if err == nil {
		t.Fatal("expected error for string longer than 20 chars, got nil")
	}
}

func TestSetCiString20Type_Empty(t *testing.T) {
	val := ""
	_, err := SetCiString20Type(val)
	if err != nil {
		t.Fatalf("expected no error for empty string, got %v", err)
	}
}

// CiString25Type

func TestSetCiString25Type_Valid(t *testing.T) {
	val := "ValidString"
	_, err := SetCiString25Type(val)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSetCiString25Type_MaxLength(t *testing.T) {
	val := "12345678901234567890" // exactly 20 chars
	_, err := SetCiString25Type(val)
	if err != nil {
		t.Fatalf("expected no error for 20-char string, got %v", err)
	}
}

func TestSetCiString25Type_TooLong(t *testing.T) {
	val := "123456789012345678901" // 21 chars
	_, err := SetCiString25Type(val)
	if err == nil {
		t.Fatal("expected error for string longer than 20 chars, got nil")
	}
}

func TestSetCiString25Type_Empty(t *testing.T) {
	val := ""
	_, err := SetCiString25Type(val)
	if err != nil {
		t.Fatalf("expected no error for empty string, got %v", err)
	}
}
