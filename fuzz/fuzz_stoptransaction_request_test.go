//go:build fuzz

package fuzz

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	stp "github.com/aasanchez/ocpp16messages/stopTransaction"
)

func FuzzStopTransactionReq(f *testing.F) {
	f.Add(1, 100, "2025-01-15T10:30:00Z")
	f.Add(-1, -1, "bad-timestamp")

	f.Fuzz(func(t *testing.T, txID int, meterStop int, ts string) {
		_, _ = stp.Req(stp.ReqInput{
			TransactionId: txID,
			MeterStop:     meterStop,
			Timestamp:     ts,
			TransactionData: []mt.MeterValueInput{
				{
					Timestamp: ts,
					SampledValue: []mt.SampledValueInput{
						{Value: "100"},
					},
				},
			},
		})
	})
}
