package types_test

import (
	"testing"

	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

const (
	recurrencyDailyStr     = "Daily"
	recurrencyWeeklyStr    = "Weekly"
	methodRecurrencyString = "RecurrencyKindType.String()"
)

func TestRecurrencyKindType_IsValid_Daily(t *testing.T) {
	t.Parallel()

	if !spt.RecurrencyKindDaily.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"RecurrencyKindDaily",
		)
	}
}

func TestRecurrencyKindType_IsValid_Weekly(t *testing.T) {
	t.Parallel()

	if !spt.RecurrencyKindWeekly.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"RecurrencyKindWeekly",
		)
	}
}

func TestRecurrencyKindType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	kind := spt.RecurrencyKindType("")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"RecurrencyKindType(\"\")",
		)
	}
}

func TestRecurrencyKindType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	kind := spt.RecurrencyKindType("Unknown")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"RecurrencyKindType(\"Unknown\")",
		)
	}
}

func TestRecurrencyKindType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	kind := spt.RecurrencyKindType("daily")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"RecurrencyKindType(\"daily\")",
		)
	}
}

func TestRecurrencyKindType_IsValid_Monthly(t *testing.T) {
	t.Parallel()

	kind := spt.RecurrencyKindType("Monthly")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"RecurrencyKindType(\"Monthly\")",
		)
	}
}

func TestRecurrencyKindType_String_Daily(t *testing.T) {
	t.Parallel()

	got := spt.RecurrencyKindDaily.String()
	if got != recurrencyDailyStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodRecurrencyString,
			got,
			recurrencyDailyStr,
		)
	}
}

func TestRecurrencyKindType_String_Weekly(t *testing.T) {
	t.Parallel()

	got := spt.RecurrencyKindWeekly.String()
	if got != recurrencyWeeklyStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodRecurrencyString,
			got,
			recurrencyWeeklyStr,
		)
	}
}
