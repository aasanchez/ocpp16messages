//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewMeterValue(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z", "100")
	f.Add("bad-timestamp", "100")
	f.Add("", "")
	f.Add("2025-01-15T10:30:00Z", "")

	f.Fuzz(func(t *testing.T, timestamp string, value string) {
		if len(timestamp) > maxFuzzStringLen || len(value) > maxFuzzStringLen {
			t.Skip()
		}

		meterValue, err := types.NewMeterValue(types.MeterValueInput{
			Timestamp: timestamp,
			SampledValue: []types.SampledValueInput{
				{Value: value},
			},
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

		roundTrip := meterValue.Timestamp.String()
		if _, parseErr := time.Parse(time.RFC3339Nano, roundTrip); parseErr != nil {
			t.Fatalf("Timestamp.String() not RFC3339Nano: %v", parseErr)
		}

		if len(meterValue.SampledValue) == 0 {
			t.Fatal("SampledValue is empty, want at least one")
		}

		if meterValue.SampledValue[0].Value.String() != value {
			t.Fatalf(
				"SampledValue[0].Value = %q, want %q",
				meterValue.SampledValue[0].Value.String(),
				value,
			)
		}
	})
}
