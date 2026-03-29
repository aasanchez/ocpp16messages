//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewMeterValuesTypeSampledValue(f *testing.F) {
	f.Add("", false, "", false, "", false, "", false, "", false, "", false, "")
	f.Add(
		"100",
		true,
		types.SamplePeriodic.String(),
		true,
		types.Raw.String(),
		true,
		types.EnergyActiveImportRegister.String(),
		true,
		types.L1.String(),
		true,
		types.Outlet.String(),
		true,
		types.Wh.String(),
	)
	f.Add("bad\x01", false, "", false, "", false, "", false, "", false, "", false, "")
	f.Add(
		"100",
		true,
		"invalid-context",
		false,
		"",
		false,
		"",
		false,
		"",
		false,
		"",
		false,
		"",
	)

	f.Fuzz(func(
		t *testing.T,
		value string,
		hasContext bool,
		context string,
		hasFormat bool,
		format string,
		hasMeasurand bool,
		measurand string,
		hasPhase bool,
		phase string,
		hasLocation bool,
		location string,
		hasUnit bool,
		unit string,
	) {
		if len(value) > maxFuzzStringLen ||
			len(context) > maxFuzzStringLen ||
			len(format) > maxFuzzStringLen ||
			len(measurand) > maxFuzzStringLen ||
			len(phase) > maxFuzzStringLen ||
			len(location) > maxFuzzStringLen ||
			len(unit) > maxFuzzStringLen {
			t.Skip()
		}

		var contextPtr *string
		if hasContext {
			contextPtr = &context
		}

		var formatPtr *string
		if hasFormat {
			formatPtr = &format
		}

		var measurandPtr *string
		if hasMeasurand {
			measurandPtr = &measurand
		}

		var phasePtr *string
		if hasPhase {
			phasePtr = &phase
		}

		var locationPtr *string
		if hasLocation {
			locationPtr = &location
		}

		var unitPtr *string
		if hasUnit {
			unitPtr = &unit
		}

		sampledValue, err := types.NewSampledValue(types.SampledValueInput{
			Value:     value,
			Context:   contextPtr,
			Format:    formatPtr,
			Measurand: measurandPtr,
			Phase:     phasePtr,
			Location:  locationPtr,
			Unit:      unitPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) && !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if value == "" {
			t.Fatal("NewSampledValue succeeded with empty Value")
		}

		if sampledValue.Value.String() != value {
			t.Fatalf("Value = %q, want %q", sampledValue.Value.String(), value)
		}

		if hasContext {
			if sampledValue.Context == nil {
				t.Fatal("Context = nil, want non-nil")
			}
			if !sampledValue.Context.IsValid() {
				t.Fatalf("Context = %q, want valid", sampledValue.Context.String())
			}
		} else if sampledValue.Context != nil {
			t.Fatal("Context != nil, want nil")
		}

		if hasFormat {
			if sampledValue.Format == nil {
				t.Fatal("Format = nil, want non-nil")
			}
			if !sampledValue.Format.IsValid() {
				t.Fatalf("Format = %q, want valid", sampledValue.Format.String())
			}
		} else if sampledValue.Format != nil {
			t.Fatal("Format != nil, want nil")
		}

		if hasMeasurand {
			if sampledValue.Measurand == nil {
				t.Fatal("Measurand = nil, want non-nil")
			}
			if !sampledValue.Measurand.IsValid() {
				t.Fatalf(
					"Measurand = %q, want valid",
					sampledValue.Measurand.String(),
				)
			}
		} else if sampledValue.Measurand != nil {
			t.Fatal("Measurand != nil, want nil")
		}

		if hasPhase {
			if sampledValue.Phase == nil {
				t.Fatal("Phase = nil, want non-nil")
			}
			if !sampledValue.Phase.IsValid() {
				t.Fatalf("Phase = %q, want valid", sampledValue.Phase.String())
			}
		} else if sampledValue.Phase != nil {
			t.Fatal("Phase != nil, want nil")
		}

		if hasLocation {
			if sampledValue.Location == nil {
				t.Fatal("Location = nil, want non-nil")
			}
			if !sampledValue.Location.IsValid() {
				t.Fatalf(
					"Location = %q, want valid",
					sampledValue.Location.String(),
				)
			}
		} else if sampledValue.Location != nil {
			t.Fatal("Location != nil, want nil")
		}

		if hasUnit {
			if sampledValue.Unit == nil {
				t.Fatal("Unit = nil, want non-nil")
			}
			if !sampledValue.Unit.IsValid() {
				t.Fatalf("Unit = %q, want valid", sampledValue.Unit.String())
			}
		} else if sampledValue.Unit != nil {
			t.Fatal("Unit != nil, want nil")
		}
	})
}
