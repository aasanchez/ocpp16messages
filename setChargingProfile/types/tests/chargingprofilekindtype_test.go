//nolint:dupl // enum test pattern is intentionally similar across types
package types_test

import (
	"testing"

	spt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

const (
	kindAbsoluteStr  = "Absolute"
	kindRecurringStr = "Recurring"
	kindRelativeStr  = "Relative"
	methodKindString = "ChargingProfileKindType.String()"
)

func TestChargingProfileKindType_IsValid_Absolute(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileKindAbsolute.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileKindAbsolute",
		)
	}
}

func TestChargingProfileKindType_IsValid_Recurring(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileKindRecurring.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileKindRecurring",
		)
	}
}

func TestChargingProfileKindType_IsValid_Relative(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileKindRelative.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileKindRelative",
		)
	}
}

func TestChargingProfileKindType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	kind := spt.ChargingProfileKindType("")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileKindType(\"\")",
		)
	}
}

func TestChargingProfileKindType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	kind := spt.ChargingProfileKindType("Unknown")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileKindType(\"Unknown\")",
		)
	}
}

func TestChargingProfileKindType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	kind := spt.ChargingProfileKindType("absolute")
	if kind.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileKindType(\"absolute\")",
		)
	}
}

func TestChargingProfileKindType_String_Absolute(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileKindAbsolute.String()
	if got != kindAbsoluteStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodKindString,
			got,
			kindAbsoluteStr,
		)
	}
}

func TestChargingProfileKindType_String_Recurring(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileKindRecurring.String()
	if got != kindRecurringStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodKindString,
			got,
			kindRecurringStr,
		)
	}
}

func TestChargingProfileKindType_String_Relative(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileKindRelative.String()
	if got != kindRelativeStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodKindString,
			got,
			kindRelativeStr,
		)
	}
}
