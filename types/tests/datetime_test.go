package types_test

import (
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testDateTimeUTC = "2025-01-15T10:30:00Z"
)

func TestNewDateTime(t *testing.T) {
	t.Parallel()

	input := "2025-08-30T14:34:56+02:00"

	_, err := st.NewDateTime(input)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestNewDateTime_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	input := "2024-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)

	dt, _ := st.NewDateTime(input)
	if !dt.Value().Equal(want) {
		t.Errorf(st.ErrorMismatch, dt.Value(), want)
	}
}

func TestNewDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("not-a-time")
	if err == nil {
		t.Error("expected error for invalid datetime, got nil")
	}
}

func TestNewDateTime_EmptyError(t *testing.T) {
	t.Parallel()

	dt, _ := st.NewDateTime("")
	if !dt.Value().IsZero() {
		t.Errorf("expected zero value after failed parse, got %v", dt.Value())
	}
}

func TestDateTime_String(t *testing.T) {
	t.Parallel()

	dt, _ := st.NewDateTime(testDateTimeUTC)
	got := dt.String()

	if got != testDateTimeUTC {
		t.Errorf(st.ErrorMismatch, got, testDateTimeUTC)
	}
}

func TestDateTime_String_WithTimezone(t *testing.T) {
	t.Parallel()

	// Input with timezone offset is converted to UTC
	input := "2025-01-15T12:30:00+02:00"

	dt, _ := st.NewDateTime(input)
	got := dt.String()

	// Should be normalized to UTC (12:30 + 02:00 = 10:30 UTC)
	if got != testDateTimeUTC {
		t.Errorf(st.ErrorMismatch, got, testDateTimeUTC)
	}
}

func TestDateTime_String_WithNanoseconds(t *testing.T) {
	t.Parallel()

	// RFC3339 parsing doesn't preserve nanoseconds, verify output format
	dt, _ := st.NewDateTime(testDateTimeUTC)
	got := dt.String()

	// Verify it can be parsed back as RFC3339
	_, err := time.Parse(time.RFC3339Nano, got)
	if err != nil {
		t.Errorf("String() output should be valid RFC3339Nano: %v", err)
	}
}
