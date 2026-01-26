package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const invalidMeasurand = types.Measurand("InvalidMeasurand")

func TestMeasurand_String_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.MeasurandEnergyActiveImportRegister,
		isValidFn: types.MeasurandEnergyActiveImportRegister.IsValid,
	}, "Energy.Active.Import.Register")
}

func TestMeasurand_String_Voltage(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.MeasurandVoltage,
		isValidFn: types.MeasurandVoltage.IsValid,
	}, "Voltage")
}

func TestMeasurand_String_CurrentImport(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.MeasurandCurrentImport,
		isValidFn: types.MeasurandCurrentImport.IsValid,
	}, "Current.Import")
}

func TestMeasurand_String_PowerActiveImport(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.MeasurandPowerActiveImport,
		isValidFn: types.MeasurandPowerActiveImport.IsValid,
	}, "Power.Active.Import")
}

func TestMeasurand_String_SoC(t *testing.T) {
	t.Parallel()

	assertEnumValid(t, enumValidator{
		value:     types.MeasurandSoC,
		isValidFn: types.MeasurandSoC.IsValid,
	}, "SoC")
}

func TestMeasurand_IsValid_InvalidValue(t *testing.T) {
	t.Parallel()

	assertEnumInvalid(t, enumValidator{
		value:     invalidMeasurand,
		isValidFn: invalidMeasurand.IsValid,
	})
}
