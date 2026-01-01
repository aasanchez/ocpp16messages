package sharedtypes

import (
	"errors"
	"strings"
	"testing"
)

// Internal test

func Test_sharedtypes_setCiString_Valid(t *testing.T) {
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

func Test_sharedtypes_setCiString_TooLong(t *testing.T) {
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

func Test_sharedtypes_setCiString_NonPrintableASCII(t *testing.T) {
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

func Test_sharedtypes_setCiString_EmptyIsValid(t *testing.T) {
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

func Test_sharedtypes_setCiString_ASCIIBoundary(t *testing.T) {
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

// External Test

// CiString20Type
func Test_sharedtypes_SetCiString20Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 20)

	_, err := SetCiString20Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 20-character string: %v", err)
	}
}

func Test_sharedtypes_SetCiString20Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString20Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString20Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("A", 21)

	_, err := SetCiString20Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 20 characters; got none")
	}
}

func Test_sharedtypes_SetCiString20Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, _ := SetCiString20Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString25Type
func Test_sharedtypes_SetCiString25Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 25)

	_, err := SetCiString25Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 25-character string: %v", err)
	}
}

func Test_sharedtypes_SetCiString25Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString25Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString25Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("B", 26)

	_, err := SetCiString25Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 25 characters; got none")
	}
}

func Test_sharedtypes_SetCiString25Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, _ := SetCiString25Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString50Type
func Test_sharedtypes_SetCiString50Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 50)

	_, err := SetCiString50Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 50-character string: %v", err)
	}
}

func Test_sharedtypes_SetCiString50Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString50Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString50Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("C", 51)

	_, err := SetCiString50Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 50 characters; got none")
	}
}

func Test_sharedtypes_SetCiString50Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, _ := SetCiString50Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString255Type
func Test_sharedtypes_SetCiString255Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 255)

	_, err := SetCiString255Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 255-character string: %v", err)
	}
}

func Test_sharedtypes_SetCiString255Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString255Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString255Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("D", 256)

	_, err := SetCiString255Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 255 characters; got none")
	}
}

func Test_sharedtypes_SetCiString255Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, _ := SetCiString255Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}

// CiString500Type
func Test_sharedtypes_SetCiString500Type(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 500)

	_, err := SetCiString500Type(val)
	if err != nil {
		t.Errorf("unexpected error for a 500-character string: %v", err)
	}
}

func Test_sharedtypes_SetCiString500Type_Empty(t *testing.T) {
	t.Parallel()

	val := ""

	_, err := SetCiString500Type(val)
	if err != nil {
		t.Errorf("empty string should be accepted; got error: %v", err)
	}
}

func Test_sharedtypes_SetCiString500Type_TooLong(t *testing.T) {
	t.Parallel()

	val := strings.Repeat("E", 501)

	_, err := SetCiString500Type(val)
	if err == nil {
		t.Error("expected an error for a string longer than 500 characters; got none")
	}
}

func Test_sharedtypes_SetCiString500Type_TestValue(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, _ := SetCiString500Type(input)
	if ciStr.Value() != input {
		t.Errorf("value mismatch: want %q, got %q", input, ciStr.Value())
	}
}
