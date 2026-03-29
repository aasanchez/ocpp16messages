//go:build fuzz

package fuzz

import (
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzChargingProfilePurposeTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.ChargePointMaxProfile.String(): {},
		types.TxDefaultProfile.String():      {},
		types.TxProfile.String():             {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("txprofile")
	f.Add("ChargePointMaxProfile ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		purpose := types.ChargingProfilePurposeType(input)

		if got := purpose.IsValid(); got {
			if _, ok := allowed[input]; !ok {
				t.Fatalf("IsValid() = true for %q, want false", input)
			}
		} else {
			if _, ok := allowed[input]; ok {
				t.Fatalf("IsValid() = false for %q, want true", input)
			}
		}
	})
}

func FuzzMeasurandIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.MeasurandCurrentExport.String():                {},
		types.MeasurandCurrentImport.String():                {},
		types.MeasurandCurrentOffered.String():               {},
		types.MeasurandEnergyActiveExportRegister.String():   {},
		types.MeasurandEnergyActiveImportRegister.String():   {},
		types.MeasurandEnergyReactiveExportRegister.String(): {},
		types.MeasurandEnergyReactiveImportRegister.String(): {},
		types.MeasurandEnergyActiveExportInterval.String():   {},
		types.MeasurandEnergyActiveImportInterval.String():   {},
		types.MeasurandEnergyReactiveExportInterval.String(): {},
		types.MeasurandEnergyReactiveImportInterval.String(): {},
		types.MeasurandFrequency.String():                    {},
		types.MeasurandPowerActiveExport.String():            {},
		types.MeasurandPowerActiveImport.String():            {},
		types.MeasurandPowerFactor.String():                  {},
		types.MeasurandPowerOffered.String():                 {},
		types.MeasurandPowerReactiveExport.String():          {},
		types.MeasurandPowerReactiveImport.String():          {},
		types.MeasurandRPM.String():                          {},
		types.MeasurandSoC.String():                          {},
		types.MeasurandTemperature.String():                  {},
		types.MeasurandVoltage.String():                      {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("Energy.Active.Import.Register ")
	f.Add("energy.active.import.register")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		measurand := types.Measurand(input)

		if got := measurand.IsValid(); got {
			if _, ok := allowed[input]; !ok {
				t.Fatalf("IsValid() = true for %q, want false", input)
			}
		} else {
			if _, ok := allowed[input]; ok {
				t.Fatalf("IsValid() = false for %q, want true", input)
			}
		}
	})
}

func FuzzPhaseIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.PhaseL1.String():   {},
		types.PhaseL2.String():   {},
		types.PhaseL3.String():   {},
		types.PhaseN.String():    {},
		types.PhaseL1N.String():  {},
		types.PhaseL2N.String():  {},
		types.PhaseL3N.String():  {},
		types.PhaseL1L2.String(): {},
		types.PhaseL2L3.String(): {},
		types.PhaseL3L1.String(): {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("L1N")
	f.Add("L1-N ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		phase := types.Phase(input)

		if got := phase.IsValid(); got {
			if _, ok := allowed[input]; !ok {
				t.Fatalf("IsValid() = true for %q, want false", input)
			}
		} else {
			if _, ok := allowed[input]; ok {
				t.Fatalf("IsValid() = false for %q, want true", input)
			}
		}
	})
}
