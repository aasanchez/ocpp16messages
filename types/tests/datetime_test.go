package types_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testDateTimeUTC    = "2025-01-15T10:30:00Z"
	testDateTimeNonUTC = "2025-08-30T14:34:56+02:00"
)

func TestNewDateTime_AcceptsUTC(t *testing.T) {
	t.Parallel()

	dateTime, err := st.NewDateTime(testDateTimeUTC)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	want, _ := time.Parse(time.RFC3339, testDateTimeUTC)
	if !dateTime.Value().Equal(want) {
		t.Errorf(st.ErrorMismatch, dateTime.Value(), want)
	}

	if dateTime.Value().Location() != time.UTC {
		t.Errorf("location = %v, want UTC", dateTime.Value().Location())
	}
}

func TestNewDateTime_RejectsNonUTC(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime(testDateTimeNonUTC)
	if err == nil {
		t.Fatal("NewDateTime() error = nil, want non-UTC error")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}

	if !strings.Contains(err.Error(), "UTC") {
		t.Errorf(st.ErrorWantContains, err, "UTC")
	}
}

func TestNewDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("not-a-time")
	if err == nil {
		t.Fatal("expected error for invalid datetime, got nil")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestNewDateTime_EmptyError(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("")
	if err == nil {
		t.Fatal("expected error for empty datetime, got nil")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestDateTime_String(t *testing.T) {
	t.Parallel()

	dateTime, _ := st.NewDateTime(testDateTimeUTC)
	got := dateTime.String()

	if got != testDateTimeUTC {
		t.Errorf(st.ErrorMismatch, got, testDateTimeUTC)
	}

	_, err := time.Parse(time.RFC3339Nano, got)
	if err != nil {
		t.Errorf("String() output should be valid RFC3339Nano: %v", err)
	}
}
