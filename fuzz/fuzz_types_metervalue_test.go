//go:build fuzz

package fuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewMeterValue(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z", "100")
	f.Add("bad-timestamp", "100")
	f.Add("", "")

	f.Fuzz(func(t *testing.T, timestamp string, value string) {
		_, _ = st.NewMeterValue(st.MeterValueInput{
			Timestamp: timestamp,
			SampledValue: []st.SampledValueInput{
				{Value: value},
			},
		})
	})
}
