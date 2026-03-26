//go:build race

package race

import (
	"fmt"
	"testing"

	mv "github.com/aasanchez/ocpp16messages/metervalues"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	mvWorkers    = 16
	mvIterations = 200
)

func TestRace_MeterValuesSingleReq(t *testing.T) {
	t.Parallel()

	sampledValues := []st.SampledValueInput{{Value: "100"}}
	metervalues := []st.MeterValueInput{
		{
			Timestamp:    "2025-01-02T15:00:00Z",
			SampledValue: sampledValues,
		},
	}

	runConcurrent(t, mvWorkers, mvIterations, func(_, _ int) error {
		_, err := mv.Req(mv.ReqInput{
			ConnectorId:   1,
			MeterValue:    metervalues,
			TransactionId: nil,
		})
		if err != nil {
			return fmt.Errorf("MeterValues.Req: %w", err)
		}

		return nil
	})
}
