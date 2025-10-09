package sharedtypes

import (
	"errors"
	"testing"
)

func Test_sharedtypes_setCiString_Valid(t *testing.T) {
	t.Parallel()

	input := "OCPP-Test-123"
	maxLen := 20

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorValueMismatch, got, input)
	}
}

func Test_sharedtypes_setCiString_TooLong(t *testing.T) {
	t.Parallel()

	input := "ThisStringIsWayTooLong"
	maxLen := 5

	_, err := setCiString(input, maxLen)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}

	if !errors.Is(err, errExceedsMaxLength) {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_setCiString_NonPrintableASCII(t *testing.T) {
	t.Parallel()

	input := "Valid\x01Invalid"

	_, err := setCiString(input, 20)
	if err == nil {
		t.Errorf(ErrorExpectedError, err)
	}

	if !errors.Is(err, errInvalidASCII) {
		t.Errorf(ErrorExpectedError, err)
	}
}

func Test_sharedtypes_setCiString_EmptyIsValid(t *testing.T) {
	t.Parallel()

	input := ""
	maxLen := 10

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf(ErrorExpectedError, err)
	}

	if got := cis.val(); got != "" {
		t.Errorf(ErrorValueMismatch, input, got)
	}
}

func Test_sharedtypes_setCiString_ASCIIBoundary(t *testing.T) {
	t.Parallel()

	input := " ~"
	maxLen := 5

	cis, err := setCiString(input, maxLen)
	if err != nil {
		t.Fatalf(ErrorUnexpectedError, err)
	}

	if got := cis.val(); got != input {
		t.Errorf(ErrorValueMismatch, got, input)
	}
}
