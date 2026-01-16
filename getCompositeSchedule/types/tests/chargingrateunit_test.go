package types_test

import (
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	unitWattsStr   = "W"
	unitAmperesStr = "A"
	unitTypeString = "ChargingRateUnit.String()"
)

func TestChargingRateUnit_IsValid_Watts(t *testing.T) {
	t.Parallel()

	if !gt.ChargingRateUnitWatts.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargingRateUnitWatts")
	}
}

func TestChargingRateUnit_IsValid_Amperes(t *testing.T) {
	t.Parallel()

	if !gt.ChargingRateUnitAmperes.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ChargingRateUnitAmperes")
	}
}

func TestChargingRateUnit_IsValid_Empty(t *testing.T) {
	t.Parallel()

	unit := gt.ChargingRateUnit("")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"\")")
	}
}

func TestChargingRateUnit_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	unit := gt.ChargingRateUnit("X")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"X\")")
	}
}

func TestChargingRateUnit_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	unit := gt.ChargingRateUnit("w")
	if unit.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ChargingRateUnit(\"w\")")
	}
}

func TestChargingRateUnit_String_Watts(t *testing.T) {
	t.Parallel()

	got := gt.ChargingRateUnitWatts.String()
	if got != unitWattsStr {
		t.Errorf(st.ErrorMethodMismatch, unitTypeString, got, unitWattsStr)
	}
}

func TestChargingRateUnit_String_Amperes(t *testing.T) {
	t.Parallel()

	got := gt.ChargingRateUnitAmperes.String()
	if got != unitAmperesStr {
		t.Errorf(st.ErrorMethodMismatch, unitTypeString, got, unitAmperesStr)
	}
}
