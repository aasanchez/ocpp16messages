package types

import (
	"errors"
	"testing"
)

// Test helper constants for maximum length validation.
const (
	testMaxLen5  = 5
	testMaxLen10 = 10
	testMaxLen20 = 20
)

// Tests for internal newCiString function (unexported).

func Test_newCiString(t *testing.T) {
	t.Parallel()

	input := "OCPP-Test-123"
	maxLen := testMaxLen20

	cis, err := newCiString(input, maxLen)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorMismatch, got, input)
	}
}

func Test_newCiString_TooLong(t *testing.T) {
	t.Parallel()

	input := "ThisStringIsWayTooLong"
	maxLen := testMaxLen5

	_, err := newCiString(input, maxLen)
	if err == nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if !errors.Is(err, errExceedsMaxLength) {
		t.Errorf("expected errExceedsMaxLength, got: %v", err)
	}
}

func Test_newCiString_NonPrintableASCII(t *testing.T) {
	t.Parallel()

	input := "Valid\x01Invalid"

	_, err := newCiString(input, testMaxLen20)
	if err == nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if !errors.Is(err, errInvalidASCII) {
		t.Errorf("expected errInvalidASCII, got: %v", err)
	}
}

func Test_newCiString_EmptyReturnsError(t *testing.T) {
	t.Parallel()

	input := ""
	maxLen := testMaxLen10

	_, err := newCiString(input, maxLen)
	if err == nil {
		t.Error("newCiString() error = nil, want error for empty string")
	}

	if !errors.Is(err, ErrEmptyValue) {
		t.Errorf("expected ErrEmptyValue, got: %v", err)
	}
}

func Test_newCiString_ASCIIBoundary(t *testing.T) {
	t.Parallel()

	input := " ~"
	maxLen := testMaxLen5

	cis, err := newCiString(input, maxLen)
	if err != nil {
		t.Errorf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorMismatch, got, input)
	}
}
