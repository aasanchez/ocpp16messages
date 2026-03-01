//go:build fuzz

package fuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzAuthorizationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		st.AuthorizationStatusAccepted.String():     {},
		st.AuthorizationStatusBlocked.String():      {},
		st.AuthorizationStatusExpired.String():      {},
		st.AuthorizationStatusInvalid.String():      {},
		st.AuthorizationStatusConcurrentTx.String(): {},
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

		status := st.AuthorizationStatus(input)

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
		st.ChargingRateUnitWatts.String():   {},
		st.ChargingRateUnitAmperes.String(): {},
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

		unit := st.ChargingRateUnit(input)

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
		st.LocationBody.String():   {},
		st.LocationCable.String():  {},
		st.LocationEV.String():     {},
		st.LocationInlet.String():  {},
		st.LocationOutlet.String(): {},
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

		loc := st.Location(input)

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
		st.ReadingContextInterruptionBegin.String(): {},
		st.ReadingContextInterruptionEnd.String():   {},
		st.ReadingContextOther.String():             {},
		st.ReadingContextSampleClock.String():       {},
		st.ReadingContextSamplePeriodic.String():    {},
		st.ReadingContextTransactionBegin.String():  {},
		st.ReadingContextTransactionEnd.String():    {},
		st.ReadingContextTrigger.String():           {},
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

		ctx := st.ReadingContext(input)

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
		st.UnitWh.String():         {},
		st.UnitKWh.String():        {},
		st.UnitVarh.String():       {},
		st.UnitKvarh.String():      {},
		st.UnitW.String():          {},
		st.UnitKW.String():         {},
		st.UnitVA.String():         {},
		st.UnitKVA.String():        {},
		st.UnitVar.String():        {},
		st.UnitKvar.String():       {},
		st.UnitA.String():          {},
		st.UnitV.String():          {},
		st.UnitCelsius.String():    {},
		st.UnitFahrenheit.String(): {},
		st.UnitK.String():          {},
		st.UnitPercent.String():    {},
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

		unit := st.UnitOfMeasure(input)

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
		st.ValueFormatRaw.String():        {},
		st.ValueFormatSignedData.String(): {},
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

		format := st.ValueFormat(input)

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
