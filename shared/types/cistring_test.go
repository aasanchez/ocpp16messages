package sharedtypes

import (
	"errors"
	"testing"
)

// Valid CiString construction: ASCII within bound.
// OCPP 1.6: printable ASCII (32..126), length <= max.
// Verifies the value round-trips without mutation.
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

// Length overflow is rejected with a wrapped error.
// OCPP 1.6: bounded strings fail fast on overflow.
// Verifies errExceedsMaxLength is reported.
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

// Non-printable ASCII is rejected.
// OCPP 1.6 requires codes 32..126 only.
// Verifies errInvalidASCII is reported.
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
