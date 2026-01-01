package types

import (
	"errors"
	"testing"
)

const (
	testMaxLen5  = 5
	testMaxLen10 = 10
	testMaxLen20 = 20
)

// Internal test

func Test_sharedtypes_newCiString_Valid(t *testing.T) {
	t.Parallel()

	input := "OCPP-Test-123"
	maxLen := testMaxLen20

	cis, err := newCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if got := cis.val(); got != input {
		t.Errorf("val() = %q; want %q", got, input)
	}
}

func Test_sharedtypes_newCiString_TooLong(t *testing.T) {
	t.Parallel()

	input := "ThisStringIsWayTooLong"
	maxLen := testMaxLen5

	_, err := newCiString(input, maxLen)
	if err == nil {
		t.Fatal("expected error for exceeding max length, got nil")
	}

	if !errors.Is(err, errExceedsMaxLength) {
		t.Errorf("expected errExceedsMaxLength, got: %v", err)
	}
}

func Test_sharedtypes_newCiString_NonPrintableASCII(t *testing.T) {
	t.Parallel()

	input := "Valid\x01Invalid"

	_, err := newCiString(input, testMaxLen20)
	if err == nil {
		t.Fatal("expected error for non-printable ASCII, got nil")
	}

	if !errors.Is(err, errInvalidASCII) {
		t.Errorf("expected errInvalidASCII, got: %v", err)
	}
}

func Test_sharedtypes_newCiString_EmptyIsValid(t *testing.T) {
	t.Parallel()

	input := ""
	maxLen := testMaxLen10

	cis, err := newCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error for empty input, got: %v", err)
	}

	if got := cis.val(); got != "" {
		t.Errorf("val() = %q; want empty string", got)
	}
}

func Test_sharedtypes_newCiString_ASCIIBoundary(t *testing.T) {
	t.Parallel()

	input := " ~"
	maxLen := testMaxLen5

	cis, err := newCiString(input, maxLen)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if got := cis.val(); got != input {
		t.Errorf("val() = %q; want %q", got, input)
	}
}
