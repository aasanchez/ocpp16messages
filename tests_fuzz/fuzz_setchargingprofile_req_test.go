//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	scp "github.com/aasanchez/ocpp16messages/setchargingprofile"
	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzSetChargingProfileReq(f *testing.F) {
	f.Add(0, 1, 0, st.ChargePointMaxProfile.String(), "Absolute")
	f.Add(1, 1, 0, st.TxProfile.String(), "Relative")
	f.Add(-1, 1, 0, st.ChargePointMaxProfile.String(), "Absolute")
	f.Add(1, -1, 0, st.ChargePointMaxProfile.String(), "Absolute")
	f.Add(1, 1, 0, "invalid-purpose", "Absolute")
	f.Add(1, 1, 0, st.ChargePointMaxProfile.String(), "invalid-kind")

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		chargingProfileId int,
		stackLevel int,
		purpose string,
		kind string,
	) {
		if len(purpose) > maxFuzzStringLen || len(kind) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := scp.Req(scp.ReqInput{
			ConnectorId: connectorId,
			CsChargingProfiles: spt.ChargingProfileInput{
				ChargingProfileId:      chargingProfileId,
				TransactionId:          nil,
				StackLevel:             stackLevel,
				ChargingProfilePurpose: purpose,
				ChargingProfileKind:    kind,
				RecurrencyKind:         nil,
				ValidFrom:              nil,
				ValidTo:                nil,
				ChargingSchedule: st.ChargingScheduleInput{
					ChargingRateUnit: st.ChargingRateUnitWatts.String(),
					ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
						{
							StartPeriod:  0,
							Limit:        0,
							NumberPhases: nil,
						},
					},
				},
			},
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if got := req.ConnectorId.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorId = %d, want %d", got, connectorId)
		}

		if !req.CsChargingProfiles.ChargingProfilePurpose().IsValid() {
			t.Fatalf(
				"ChargingProfilePurpose() = %q, want valid",
				req.CsChargingProfiles.ChargingProfilePurpose().String(),
			)
		}
	})
}
