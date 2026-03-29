//go:build fuzz

package fuzz

import (
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzAuthorizationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.AuthorizationStatusAccepted.String():     {},
		types.AuthorizationStatusBlocked.String():      {},
		types.AuthorizationStatusExpired.String():      {},
		types.AuthorizationStatusInvalid.String():      {},
		types.AuthorizationStatusConcurrentTx.String(): {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("accepted")
	f.Add("ConcurrentTX")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := types.AuthorizationStatus(input)

		if got := status.IsValid(); got {
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

func FuzzChargingRateUnitIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.ChargingRateUnitWatts.String():   {},
		types.ChargingRateUnitAmperes.String(): {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("w")
	f.Add("A ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		unit := types.ChargingRateUnit(input)

		if got := unit.IsValid(); got {
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

func FuzzLocationIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.LocationBody.String():   {},
		types.LocationCable.String():  {},
		types.LocationEV.String():     {},
		types.LocationInlet.String():  {},
		types.LocationOutlet.String(): {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("body")
	f.Add("Other")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		loc := types.Location(input)

		if got := loc.IsValid(); got {
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

func FuzzReadingContextIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.ReadingContextInterruptionBegin.String(): {},
		types.ReadingContextInterruptionEnd.String():   {},
		types.ReadingContextOther.String():             {},
		types.ReadingContextSampleClock.String():       {},
		types.ReadingContextSamplePeriodic.String():    {},
		types.ReadingContextTransactionBegin.String():  {},
		types.ReadingContextTransactionEnd.String():    {},
		types.ReadingContextTrigger.String():           {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("Transaction.Begin ")
	f.Add("OtherContext")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		ctx := types.ReadingContext(input)

		if got := ctx.IsValid(); got {
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

func FuzzUnitOfMeasureIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.UnitWh.String():         {},
		types.UnitKWh.String():        {},
		types.UnitVarh.String():       {},
		types.UnitKvarh.String():      {},
		types.UnitW.String():          {},
		types.UnitKW.String():         {},
		types.UnitVA.String():         {},
		types.UnitKVA.String():        {},
		types.UnitVar.String():        {},
		types.UnitKvar.String():       {},
		types.UnitA.String():          {},
		types.UnitV.String():          {},
		types.UnitCelsius.String():    {},
		types.UnitFahrenheit.String(): {},
		types.UnitK.String():          {},
		types.UnitPercent.String():    {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("wh")
	f.Add("kwh")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		unit := types.UnitOfMeasure(input)

		if got := unit.IsValid(); got {
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

func FuzzValueFormatIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		types.ValueFormatRaw.String():        {},
		types.ValueFormatSignedData.String(): {},
	}

	for value := range allowed {
		f.Add(value)
	}

	f.Add("")
	f.Add("raw")
	f.Add("Signeddata")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		format := types.ValueFormat(input)

		if got := format.IsValid(); got {
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
