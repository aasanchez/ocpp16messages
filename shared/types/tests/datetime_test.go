package types_test

import (
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func Test_sharedtypes_NewDateTime(t *testing.T) {
	t.Parallel()

	input := "2025-08-30T14:34:56+02:00"

	_, err := st.NewDateTime(input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_sharedtypes_NewDateTime_ReturnsExpectedValue(t *testing.T) {
	t.Parallel()

	input := "2024-08-30T14:34:56+02:00"
	want, _ := time.Parse(time.RFC3339, input)

	dt, _ := st.NewDateTime(input)
	if !dt.Value().Equal(want) {
		t.Errorf("Value() mismatch: got %v, want %v", dt.Value(), want)
	}
}

func Test_sharedtypes_NewDateTime_InvalidReturnsError(t *testing.T) {
	t.Parallel()

	_, err := st.NewDateTime("not-a-time")
	if err == nil {
		t.Error("expected error for invalid datetime, got nil")
	}
}

func Test_sharedtypes_NewDateTime_EmptyError(t *testing.T) {
	t.Parallel()

	dt, _ := st.NewDateTime("")
	if !dt.Value().IsZero() {
		t.Errorf("expected zero value after failed parse, got %v", dt.Value())
	}
}
