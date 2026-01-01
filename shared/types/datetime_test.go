package sharedtypes

import (
	"testing"
	"time"
)

func Test_sharedtypes_SetDateTime(t *testing.T) {
	t.Parallel()

	input := "2025-08-30T14:34:56+02:00"

	_, err := SetDateTime(input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_sharedtypes_SetDateTime_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	input := "2024-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)

	dt, _ := SetDateTime(input)
	if !dt.Value().Equal(want) {
		t.Errorf("Value() mismatch: got %v, want %v", dt.Value(), want)
	}
}

func Test_sharedtypes_SetDateTime_StringFormatsRFC3339Nano(t *testing.T) {
	t.Parallel()

	input := "2023-08-30T12:34:56Z"
	want, _ := time.Parse(time.RFC3339, input)
	dt, _ := SetDateTime(input)

	got := dt.String()
	if got != want.Format(time.RFC3339Nano) {
		t.Errorf("String() mismatch: got %q, want %q", got, want.Format(time.RFC3339Nano))
	}
}

func Test_sharedtypes_SetDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := SetDateTime("not-a-time")
	if err == nil {
		t.Error("expected error for invalid datetime, got nil")
	}
}

func Test_sharedtypes_SetDateTime_EmptyError(t *testing.T) {
	t.Parallel()

	dt, _ := SetDateTime("")
	if !dt.Value().IsZero() {
		t.Errorf("expected zero value after failed parse, got %v", dt.Value())
	}
}

func Test_sharedtypes_DateTime_StringZeroValue(t *testing.T) {
	t.Parallel()

	var dt DateTime
	if dt.String() != (time.Time{}).Format(time.RFC3339Nano) {
		t.Errorf("String() zero mismatch: got %q", dt.String())
	}
}
