package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// OCPP 1.6: parse and read compact numbers.
//
// Verifies a valid base-10 string in range
// 0..65535 is accepted and read via Value().
// Useful for small counters and indices.
func Test_sharedtypes_SetInteger(t *testing.T) {
	t.Parallel()

	i, err := st.SetInteger("73")
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	got := i.Value()
	if got != 73 {
		t.Errorf(st.ErrorValueMismatch, 73, got)
	}
}

// Range enforcement: overflow is rejected.
//
// Values above 65535 are invalid for OCPP 1.6
// compact integers. The constructor must fail.
func Test_sharedtypes_SetInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("4294967296")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Unsigned constraint: negatives are invalid.
//
// OCPP 1.6 compact integers are unsigned.
// Negative inputs must return an error.
func Test_sharedtypes_SetInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("-10")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Validation: non-decimal input is rejected.
//
// Only base-10 digits are accepted by SetInteger.
// Alphanumeric strings must fail with an error.
func Test_sharedtypes_SetInteger_Alphanumeric(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("abc")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Validation: decimal notation is not allowed.
//
// The parser accepts integers only. Decimal
// strings like 1.22 must return an error.
func Test_sharedtypes_SetInteger_Decimal(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("1.22")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}
