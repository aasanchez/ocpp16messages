package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/metervalues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	msrCurrentExportStr              = "Current.Export"
	msrCurrentImportStr              = "Current.Import"
	msrCurrentOfferedStr             = "Current.Offered"
	msrEnergyActiveExportRegisterStr = "Energy.Active.Export.Register"
	msrEnergyActiveImportRegisterStr = "Energy.Active.Import.Register"
	msrVoltageStr                    = "Voltage"
	msrTypeStr                       = "Measurand.String()"
)

func TestMeasurand_IsValid_CurrentExport(t *testing.T) {
	t.Parallel()

	if !mt.CurrentExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "CurrentExport")
	}
}

func TestMeasurand_IsValid_CurrentImport(t *testing.T) {
	t.Parallel()

	if !mt.CurrentImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "CurrentImport")
	}
}

func TestMeasurand_IsValid_CurrentOffered(t *testing.T) {
	t.Parallel()

	if !mt.CurrentOffered.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "CurrentOffered")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	if !mt.EnergyActiveExportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyActiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	if !mt.EnergyActiveImportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyActiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportRegister(t *testing.T) {
	t.Parallel()

	if !mt.EnergyReactiveExportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyReactiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportRegister(t *testing.T) {
	t.Parallel()

	if !mt.EnergyReactiveImportRegister.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyReactiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportInterval(t *testing.T) {
	t.Parallel()

	if !mt.EnergyActiveExportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyActiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportInterval(t *testing.T) {
	t.Parallel()

	if !mt.EnergyActiveImportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyActiveImportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportInterval(t *testing.T) {
	t.Parallel()

	if !mt.EnergyReactiveExportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyReactiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportInterval(t *testing.T) {
	t.Parallel()

	if !mt.EnergyReactiveImportInterval.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "EnergyReactiveImportInterval")
	}
}

func TestMeasurand_IsValid_Frequency(t *testing.T) {
	t.Parallel()

	if !mt.Frequency.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Frequency")
	}
}

func TestMeasurand_IsValid_PowerActiveExport(t *testing.T) {
	t.Parallel()

	if !mt.PowerActiveExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerActiveExport")
	}
}

func TestMeasurand_IsValid_PowerActiveImport(t *testing.T) {
	t.Parallel()

	if !mt.PowerActiveImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerActiveImport")
	}
}

func TestMeasurand_IsValid_PowerFactor(t *testing.T) {
	t.Parallel()

	if !mt.PowerFactor.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerFactor")
	}
}

func TestMeasurand_IsValid_PowerOffered(t *testing.T) {
	t.Parallel()

	if !mt.PowerOffered.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerOffered")
	}
}

func TestMeasurand_IsValid_PowerReactiveExport(t *testing.T) {
	t.Parallel()

	if !mt.PowerReactiveExport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerReactiveExport")
	}
}

func TestMeasurand_IsValid_PowerReactiveImport(t *testing.T) {
	t.Parallel()

	if !mt.PowerReactiveImport.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "PowerReactiveImport")
	}
}

func TestMeasurand_IsValid_RPM(t *testing.T) {
	t.Parallel()

	if !mt.RPM.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RPM")
	}
}

func TestMeasurand_IsValid_SoC(t *testing.T) {
	t.Parallel()

	if !mt.SoC.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "SoC")
	}
}

func TestMeasurand_IsValid_Temperature(t *testing.T) {
	t.Parallel()

	if !mt.Temperature.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Temperature")
	}
}

func TestMeasurand_IsValid_Voltage(t *testing.T) {
	t.Parallel()

	if !mt.Voltage.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Voltage")
	}
}

func TestMeasurand_IsValid_Empty(t *testing.T) {
	t.Parallel()

	msr := mt.Measurand("")
	if msr.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Measurand(\"\")")
	}
}

func TestMeasurand_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	msr := mt.Measurand("Invalid")
	if msr.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Measurand(\"Invalid\")")
	}
}

func TestMeasurand_String_CurrentExport(t *testing.T) {
	t.Parallel()

	got := mt.CurrentExport.String()
	if got != msrCurrentExportStr {
		t.Errorf(st.ErrorMethodMismatch, msrTypeStr, got, msrCurrentExportStr)
	}
}

func TestMeasurand_String_CurrentImport(t *testing.T) {
	t.Parallel()

	got := mt.CurrentImport.String()
	if got != msrCurrentImportStr {
		t.Errorf(st.ErrorMethodMismatch, msrTypeStr, got, msrCurrentImportStr)
	}
}

func TestMeasurand_String_CurrentOffered(t *testing.T) {
	t.Parallel()

	got := mt.CurrentOffered.String()
	if got != msrCurrentOfferedStr {
		t.Errorf(st.ErrorMethodMismatch, msrTypeStr, got, msrCurrentOfferedStr)
	}
}

func TestMeasurand_String_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	got := mt.EnergyActiveExportRegister.String()
	if got != msrEnergyActiveExportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			msrTypeStr,
			got,
			msrEnergyActiveExportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	got := mt.EnergyActiveImportRegister.String()
	if got != msrEnergyActiveImportRegisterStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			msrTypeStr,
			got,
			msrEnergyActiveImportRegisterStr,
		)
	}
}

func TestMeasurand_String_Voltage(t *testing.T) {
	t.Parallel()

	got := mt.Voltage.String()
	if got != msrVoltageStr {
		t.Errorf(st.ErrorMethodMismatch, msrTypeStr, got, msrVoltageStr)
	}
}
