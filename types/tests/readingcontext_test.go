package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestReadingContext_String_SamplePeriodic(t *testing.T) {
	t.Parallel()

	got := types.ReadingContextSamplePeriodic.String()
	want := "Sample.Periodic"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestReadingContext_String_SampleClock(t *testing.T) {
	t.Parallel()

	got := types.ReadingContextSampleClock.String()
	want := "Sample.Clock"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestReadingContext_String_TransactionBegin(t *testing.T) {
	t.Parallel()

	got := types.ReadingContextTransactionBegin.String()
	want := "Transaction.Begin"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestReadingContext_String_TransactionEnd(t *testing.T) {
	t.Parallel()

	got := types.ReadingContextTransactionEnd.String()
	want := "Transaction.End"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestReadingContext_String_Trigger(t *testing.T) {
	t.Parallel()

	got := types.ReadingContextTrigger.String()
	want := "Trigger"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestReadingContext_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidContext := types.ReadingContext("InvalidContext")

	if invalidContext.IsValid() {
		t.Error("ReadingContext.IsValid() = true, want false for invalid value")
	}
}
