//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewMeterValuesTypeMeterValue(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z", true, "100")
	f.Add("bad-timestamp", true, "100")
	f.Add("", true, "")
	f.Add("2025-01-15T10:30:00Z", false, "")

	f.Fuzz(func(t *testing.T, timestamp string, hasSampledValue bool, value string) {
		if len(timestamp) > maxFuzzStringLen || len(value) > maxFuzzStringLen {
			t.Skip()
		}

		var sampledValues []types.SampledValueInput
		if hasSampledValue {
			sampledValues = []types.SampledValueInput{
				{Value: value},
			}
		}

		meterValue, err := types.NewMeterValue(types.MeterValueInput{
			Timestamp:    timestamp,
			SampledValue: sampledValues,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if meterValue.Timestamp.Value().Location() != time.UTC {
			t.Fatalf(
				"Timestamp location = %v, want UTC",
				meterValue.Timestamp.Value().Location(),
			)
		}

		if len(meterValue.SampledValue) == 0 {
			t.Fatal("SampledValue is empty, want at least one")
		}

		if hasSampledValue {
			if meterValue.SampledValue[0].Value.String() != value {
				t.Fatalf(
					"SampledValue[0].Value = %q, want %q",
					meterValue.SampledValue[0].Value.String(),
					value,
				)
			}
		}
	})
}
