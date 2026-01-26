//go:build race

package race

import (
	"sync"
	"testing"

	mv "github.com/aasanchez/ocpp16messages/meterValues"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	mvWorkers    = 16
	mvIterations = 200
)

func TestRace_MeterValuesSingleReq(t *testing.T) {
	t.Parallel()

	sampled := st.SampledValueInput{Value: "100"}

	var wg sync.WaitGroup
	wg.Add(mvWorkers)

	for i := 0; i < mvWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < mvIterations; j++ {
				_, err := mv.Req(mv.ReqInput{
					ConnectorId: 1,
					MeterValue: []st.MeterValueInput{
						{
							Timestamp: "2025-01-02T15:00:00Z",
							SampledValue: []st.SampledValueInput{
								sampled,
							},
						},
					},
					TransactionId: nil,
				})
				if err != nil {
					t.Fatalf("Req unexpected error: %v", err)
				}
			}
		}()
	}

	wg.Wait()
}
