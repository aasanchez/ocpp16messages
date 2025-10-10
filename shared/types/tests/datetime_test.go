package sharedtypes_test

import (
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// OCPP 1.6: accept RFC3339 timestamps.
// Verifies parsing succeeds for a valid input.
// Internally the value is normalized to UTC.
func Test_sharedtypes_SetDateTime(t *testing.T) {
	t.Parallel()

	input := "2025-08-30T14:34:56+02:00"

	_, err := st.SetDateTime(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

// Equality uses absolute time.
// OCPP 1.6 inputs may include offsets; Value()
// is in UTC but Equal compares instants.
func Test_sharedtypes_SetDateTime_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	input := "2024-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)

	dt, _ := st.SetDateTime(input)
	if !dt.Value().Equal(want) {
		t.Errorf(st.ErrorValueMismatch, dt.Value(), want)
	}
}

// Canonical rendering uses RFC3339Nano in UTC.
// Useful for logs, signatures, and testing.
func Test_sharedtypes_SetDateTime_StringFormatsRFC3339Nano(t *testing.T) {
	t.Parallel()

	input := "2023-08-30T12:34:56Z"
	want, _ := time.Parse(time.RFC3339, input)
	dt, _ := st.SetDateTime(input)

	got := dt.String()
	if got != want.Format(time.RFC3339Nano) {
		t.Errorf(st.ErrorStringMismatch, got, want.Format(time.RFC3339Nano))
	}
}

// Invalid inputs return a wrapped parse error.
// Allows callers to inspect the root cause.
func Test_sharedtypes_SetDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := st.SetDateTime("not-a-time")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

// Zero-value semantics on failure.
// When parsing fails, the returned struct is
// the zero value; Value() is time.Time{}.
func Test_sharedtypes_SetDateTime_EmptyError(t *testing.T) {
	t.Parallel()

	dt, _ := st.SetDateTime("")
	if !dt.Value().IsZero() {
		expected := "0001-01-01 00:00:00 +0000 UTC"
		t.Errorf(st.ErrorValueMismatch, expected, dt.Value())
	}
}

// Zero DateTime string form.
// String() returns RFC3339Nano for time.Time{}.
func Test_sharedtypes_DateTime_StringZeroValue(t *testing.T) {
	t.Parallel()

	var dt st.DateTime
	if dt.String() != (time.Time{}).Format(time.RFC3339Nano) {
		t.Errorf(st.ErrorStringMismatch, "0001-01-01T00:00:00Z", dt.String())
	}
}
