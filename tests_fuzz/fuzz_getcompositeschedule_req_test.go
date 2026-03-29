//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	gc "github.com/aasanchez/ocpp16messages/getcompositeschedule"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzGetCompositeScheduleReq(f *testing.F) {
	f.Add(0, 0, false, "")
	f.Add(1, 60, true, types.ChargingRateUnitWatts.String())
	f.Add(1, 60, true, types.ChargingRateUnitAmperes.String())
	f.Add(-1, 60, false, "")
	f.Add(1, -1, false, "")
	f.Add(1, 60, true, "invalid-unit")

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		duration int,
		hasChargingRateUnit bool,
		chargingRateUnit string,
	) {
		if len(chargingRateUnit) > maxFuzzStringLen {
			t.Skip()
		}

		var chargingRateUnitPtr *string
		if hasChargingRateUnit {
			chargingRateUnitPtr = &chargingRateUnit
		}

		req, err := gc.Req(gc.ReqInput{
			ConnectorId:      connectorId,
			Duration:         duration,
			ChargingRateUnit: chargingRateUnitPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if duration < 0 || duration > math.MaxUint16 {
			t.Fatalf("Req succeeded with duration=%d", duration)
		}

		if got := req.ConnectorId.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorId = %d, want %d", got, connectorId)
		}

		if got := req.Duration.Value(); got != uint16(duration) {
			t.Fatalf("Duration = %d, want %d", got, duration)
		}

		if hasChargingRateUnit {
			if req.ChargingRateUnit == nil {
				t.Fatal("ChargingRateUnit = nil, want non-nil")
			}
			if !req.ChargingRateUnit.IsValid() {
				t.Fatalf("ChargingRateUnit = %q, want valid", req.ChargingRateUnit.String())
			}
		} else if req.ChargingRateUnit != nil {
			t.Fatal("ChargingRateUnit != nil, want nil")
		}
	})
}
