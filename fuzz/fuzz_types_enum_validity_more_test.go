//go:build fuzz

package fuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzChargingProfilePurposeTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		st.ChargePointMaxProfile.String(): {},
		st.TxDefaultProfile.String():      {},
		st.TxProfile.String():             {},
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

		purpose := st.ChargingProfilePurposeType(input)

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
		st.MeasurandCurrentExport.String():                {},
		st.MeasurandCurrentImport.String():                {},
		st.MeasurandCurrentOffered.String():               {},
		st.MeasurandEnergyActiveExportRegister.String():   {},
		st.MeasurandEnergyActiveImportRegister.String():   {},
		st.MeasurandEnergyReactiveExportRegister.String(): {},
		st.MeasurandEnergyReactiveImportRegister.String(): {},
		st.MeasurandEnergyActiveExportInterval.String():   {},
		st.MeasurandEnergyActiveImportInterval.String():   {},
		st.MeasurandEnergyReactiveExportInterval.String(): {},
		st.MeasurandEnergyReactiveImportInterval.String(): {},
		st.MeasurandFrequency.String():                    {},
		st.MeasurandPowerActiveExport.String():            {},
		st.MeasurandPowerActiveImport.String():            {},
		st.MeasurandPowerFactor.String():                  {},
		st.MeasurandPowerOffered.String():                 {},
		st.MeasurandPowerReactiveExport.String():          {},
		st.MeasurandPowerReactiveImport.String():          {},
		st.MeasurandRPM.String():                          {},
		st.MeasurandSoC.String():                          {},
		st.MeasurandTemperature.String():                  {},
		st.MeasurandVoltage.String():                      {},
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

		measurand := st.Measurand(input)

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
		st.PhaseL1.String():   {},
		st.PhaseL2.String():   {},
		st.PhaseL3.String():   {},
		st.PhaseN.String():    {},
		st.PhaseL1N.String():  {},
		st.PhaseL2N.String():  {},
		st.PhaseL3N.String():  {},
		st.PhaseL1L2.String(): {},
		st.PhaseL2L3.String(): {},
		st.PhaseL3L1.String(): {},
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

		phase := st.Phase(input)

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
