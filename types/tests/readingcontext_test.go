package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestReadingContext_String_SamplePeriodic(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.ReadingContextSamplePeriodic,
		isValidFn: types.ReadingContextSamplePeriodic.IsValid,
	}, "Sample.Periodic")
}

func TestReadingContext_String_SampleClock(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.ReadingContextSampleClock,
		isValidFn: types.ReadingContextSampleClock.IsValid,
	}, "Sample.Clock")
}

func TestReadingContext_String_TransactionBegin(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.ReadingContextTransactionBegin,
		isValidFn: types.ReadingContextTransactionBegin.IsValid,
	}, "Transaction.Begin")
}

func TestReadingContext_String_TransactionEnd(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.ReadingContextTransactionEnd,
		isValidFn: types.ReadingContextTransactionEnd.IsValid,
	}, "Transaction.End")
}

func TestReadingContext_String_Trigger(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.ReadingContextTrigger,
		isValidFn: types.ReadingContextTrigger.IsValid,
	}, "Trigger")
}

func TestReadingContext_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidContext := types.ReadingContext("InvalidContext")

	assertEnumInvalid(t, enumValidator{
		value:     invalidContext,
		isValidFn: invalidContext.IsValid,
	})
}
