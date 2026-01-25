package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestMeasurand_String_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	got := types.MeasurandEnergyActiveImportRegister.String()
	want := "Energy.Active.Import.Register"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestMeasurand_String_Voltage(t *testing.T) {
	t.Parallel()

	got := types.MeasurandVoltage.String()
	want := "Voltage"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestMeasurand_String_CurrentImport(t *testing.T) {
	t.Parallel()

	got := types.MeasurandCurrentImport.String()
	want := "Current.Import"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestMeasurand_String_PowerActiveImport(t *testing.T) {
	t.Parallel()

	got := types.MeasurandPowerActiveImport.String()
	want := "Power.Active.Import"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestMeasurand_String_SoC(t *testing.T) {
	t.Parallel()

	got := types.MeasurandSoC.String()
	want := "SoC"

	if got != want {
		t.Errorf(types.ErrorMismatch, got, want)
	}
}

func TestMeasurand_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	invalidMeasurand := types.Measurand("InvalidMeasurand")

	if invalidMeasurand.IsValid() {
		t.Error("Measurand.IsValid() = true, want false for invalid value")
	}
}
