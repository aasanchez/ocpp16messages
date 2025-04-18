package types

import (
	"errors"
	"testing"
	"time"
)

func TestDateTime_parsesValidRFC3339(t *testing.T) {
	t.Parallel()

	_, err := DateTime("2026-04-12T10:03:04-04:00")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDateTime_failsOnInvalidFormat(t *testing.T) {
	t.Parallel()

	_, err := DateTime("not-a-date")
	if err == nil {
		t.Error("expected error for invalid RFC3339 string, got nil")
	}
}

func TestDateTime_validateSucceedsForNonZeroTime(t *testing.T) {
	t.Parallel()

	dt, _ := DateTime("2027-04-12T10:03:04-04:00")
	if err := dt.Validate(); err != nil {
		t.Errorf("expected validation to succeed, got: %v", err)
	}
}

func TestDateTime_validateFailsForZeroTime(t *testing.T) {
	t.Parallel()

	dt := DateTimeType{value: time.Time{}}
	err := dt.Validate()

	if err == nil {
		t.Error("expected validation to fail for zero time, got nil")
	} else if !errors.Is(err, ErrZeroDateTime) {
		t.Errorf("expected ErrZeroDateTime, got: %v", err)
	}
}

func TestDateTime_valueReturnsCorrectTime(t *testing.T) {
	t.Parallel()

	raw := "2027-04-12T09:03:04-04:00"
	dt, _ := DateTime(raw)

	if dt.Value().IsZero() {
		t.Error("expected non-zero time from Value(), got zero")
	}
}
