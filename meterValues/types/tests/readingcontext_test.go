package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	rcInterruptionBeginStr = "Interruption.Begin"
	rcInterruptionEndStr   = "Interruption.End"
	rcOtherStr             = "Other"
	rcSampleClockStr       = "Sample.Clock"
	rcSamplePeriodicStr    = "Sample.Periodic"
	rcTransactionBeginStr  = "Transaction.Begin"
	rcTransactionEndStr    = "Transaction.End"
	rcTriggerStr           = "Trigger"
	rcTypeString           = "ReadingContext.String()"
)

func TestReadingContext_IsValid_InterruptionBegin(t *testing.T) {
	t.Parallel()

	if !mt.InterruptionBegin.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "InterruptionBegin")
	}
}

func TestReadingContext_IsValid_InterruptionEnd(t *testing.T) {
	t.Parallel()

	if !mt.InterruptionEnd.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "InterruptionEnd")
	}
}

func TestReadingContext_IsValid_Other(t *testing.T) {
	t.Parallel()

	if !mt.Other.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Other")
	}
}

func TestReadingContext_IsValid_SampleClock(t *testing.T) {
	t.Parallel()

	if !mt.SampleClock.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "SampleClock")
	}
}

func TestReadingContext_IsValid_SamplePeriodic(t *testing.T) {
	t.Parallel()

	if !mt.SamplePeriodic.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "SamplePeriodic")
	}
}

func TestReadingContext_IsValid_TransactionBegin(t *testing.T) {
	t.Parallel()

	if !mt.TransactionBegin.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TransactionBegin")
	}
}

func TestReadingContext_IsValid_TransactionEnd(t *testing.T) {
	t.Parallel()

	if !mt.TransactionEnd.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "TransactionEnd")
	}
}

func TestReadingContext_IsValid_Trigger(t *testing.T) {
	t.Parallel()

	if !mt.Trigger.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Trigger")
	}
}

func TestReadingContext_IsValid_Empty(t *testing.T) {
	t.Parallel()

	rc := mt.ReadingContext("")
	if rc.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReadingContext(\"\")")
	}
}

func TestReadingContext_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	rc := mt.ReadingContext("Invalid")
	if rc.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReadingContext(\"Invalid\")")
	}
}

func TestReadingContext_String_InterruptionBegin(t *testing.T) {
	t.Parallel()

	got := mt.InterruptionBegin.String()
	if got != rcInterruptionBeginStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			rcTypeString,
			got,
			rcInterruptionBeginStr,
		)
	}
}

func TestReadingContext_String_InterruptionEnd(t *testing.T) {
	t.Parallel()

	got := mt.InterruptionEnd.String()
	if got != rcInterruptionEndStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			rcTypeString,
			got,
			rcInterruptionEndStr,
		)
	}
}

func TestReadingContext_String_Other(t *testing.T) {
	t.Parallel()

	got := mt.Other.String()
	if got != rcOtherStr {
		t.Errorf(st.ErrorMethodMismatch, rcTypeString, got, rcOtherStr)
	}
}

func TestReadingContext_String_SampleClock(t *testing.T) {
	t.Parallel()

	got := mt.SampleClock.String()
	if got != rcSampleClockStr {
		t.Errorf(st.ErrorMethodMismatch, rcTypeString, got, rcSampleClockStr)
	}
}

func TestReadingContext_String_SamplePeriodic(t *testing.T) {
	t.Parallel()

	got := mt.SamplePeriodic.String()
	if got != rcSamplePeriodicStr {
		t.Errorf(st.ErrorMethodMismatch, rcTypeString, got, rcSamplePeriodicStr)
	}
}

func TestReadingContext_String_TransactionBegin(t *testing.T) {
	t.Parallel()

	got := mt.TransactionBegin.String()
	if got != rcTransactionBeginStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			rcTypeString,
			got,
			rcTransactionBeginStr,
		)
	}
}

func TestReadingContext_String_TransactionEnd(t *testing.T) {
	t.Parallel()

	got := mt.TransactionEnd.String()
	if got != rcTransactionEndStr {
		t.Errorf(st.ErrorMethodMismatch, rcTypeString, got, rcTransactionEndStr)
	}
}

func TestReadingContext_String_Trigger(t *testing.T) {
	t.Parallel()

	got := mt.Trigger.String()
	if got != rcTriggerStr {
		t.Errorf(st.ErrorMethodMismatch, rcTypeString, got, rcTriggerStr)
	}
}
