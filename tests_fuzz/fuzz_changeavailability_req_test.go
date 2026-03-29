//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/aasanchez/ocpp16messages/changeavailability"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzChangeAvailabilityReq(f *testing.F) {
	f.Add(0, types.AvailabilityTypeOperative.String())
	f.Add(1, types.AvailabilityTypeInoperative.String())
	f.Add(-1, types.AvailabilityTypeOperative.String())
	f.Add(math.MaxUint16+1, types.AvailabilityTypeOperative.String())
	f.Add(0, "bad-type")

	f.Fuzz(func(t *testing.T, connectorId int, availabilityType string) {
		if len(availabilityType) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := changeavailability.Req(changeavailability.ReqInput{
			ConnectorId: connectorId,
			Type:        availabilityType,
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
		if !req.Type.IsValid() {
			t.Fatalf("Type = %q, want valid", req.Type.String())
		}
		if req.Type.String() != availabilityType {
			t.Fatalf("Type = %q, want %q", req.Type.String(), availabilityType)
		}
	})
}
