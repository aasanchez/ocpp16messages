package sharedtypes_test

import (
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// --- Valid input: basic properties ---
func TestSetDateTime_ReturnsNoError(t *testing.T) {
	t.Parallel()

	input := "2025-08-30T14:34:56+02:00"

	_, err := st.SetDateTime(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSetDateTime_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	input := "2024-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)

	dt, _ := st.SetDateTime(input)
	if !dt.Value().Equal(want) {
		t.Fatalf("Value() mismatch: got %v, want %v", dt.Value(), want)
	}
}

func TestSetDateTime_StringFormatsRFC3339Nano(t *testing.T) {
	t.Parallel()

	input := "2023-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)
	dt, _ := st.SetDateTime(input)

	got := dt.String()
	if got != want.Format(time.RFC3339Nano) {
		t.Fatalf("String() mismatch: got %q, want %q", got, want.Format(time.RFC3339Nano))
	}
}

// --- Invalid input: error and zero value ---
func TestSetDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := st.SetDateTime("not-a-time")
	if err == nil {
		t.Fatal("expected error for invalid datetime, got nil")
	}
}

func TestSetDateTime_InvalidYieldsZeroValue(t *testing.T) {
	t.Parallel()

	dt, _ := st.SetDateTime("not-a-time")
	if !dt.Value().IsZero() {
		t.Fatalf("expected zero value after failed parse, got %v", dt.Value())
	}
}

// --- Zero value formatting ---
func TestDateTime_StringZeroValue(t *testing.T) {
	t.Parallel()

	var dt st.DateTime
	if dt.String() != (time.Time{}).Format(time.RFC3339Nano) {
		t.Fatalf("String() zero mismatch: got %q", dt.String())
	}
}
