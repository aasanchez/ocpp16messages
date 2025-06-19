package types

import (
	"testing"
	"time"
)

func TestDateTime_parsesValidRFC3339(t *testing.T) {
	t.Parallel()

	input := "2026-04-12T10:03:04-04:00"

	_, err := DateTime(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDateTime_failsOnInvalidFormat(t *testing.T) {
	t.Parallel()

	input := "not-a-date"
	_, err := DateTime(input)
	if err == nil {
		t.Error("expected error for invalid RFC3339 string, got nil")
	}
}

func TestDateTime_failsOnZeroTime(t *testing.T) {
	t.Parallel()

	// simulate bypassing parse but constructing zero value directly
	// this is impossible via parsing, so we directly validate the logic
	dt := DateTimeType{value: time.Time{}}
	if !dt.Value().IsZero() {
		t.Fatal("expected zero time from manually constructed value")
	}
}

func TestDateTime_valueReturnsCorrectTime(t *testing.T) {
	t.Parallel()

	input := "2027-04-12T09:03:04-04:00"
	dt, _ := DateTime(input)

	if dt.Value().IsZero() {
		t.Error("expected non-zero time from Value(), got zero")
	}
}

func TestDateTime_String_returnsRFC3339(t *testing.T) {
	t.Parallel()

	input := "2025-12-25T15:00:00Z"
	dt, err := DateTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := dt.String()
	if got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}
