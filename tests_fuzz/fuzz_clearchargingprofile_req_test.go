//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/aasanchez/ocpp16messages/clearchargingprofile"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzClearChargingProfileReq(f *testing.F) {
	f.Add(false, 0, false, 0, false, "", false, 0)
	f.Add(true, 1, true, 0, true, st.TxProfile.String(), true, 1)
	f.Add(true, -1, false, 0, false, "", false, 0)
	f.Add(false, 0, false, 0, true, "bad-purpose", false, 0)

	f.Fuzz(func(
		t *testing.T,
		hasId bool,
		id int,
		hasConnectorId bool,
		connectorId int,
		hasPurpose bool,
		purpose string,
		hasStackLevel bool,
		stackLevel int,
	) {
		if len(purpose) > maxFuzzStringLen {
			t.Skip()
		}

		var idPtr *int
		if hasId {
			idPtr = &id
		}

		var connectorIdPtr *int
		if hasConnectorId {
			connectorIdPtr = &connectorId
		}

		var purposePtr *string
		if hasPurpose {
			purposePtr = &purpose
		}

		var stackLevelPtr *int
		if hasStackLevel {
			stackLevelPtr = &stackLevel
		}

		req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
			Id:                     idPtr,
			ConnectorId:            connectorIdPtr,
			ChargingProfilePurpose: purposePtr,
			StackLevel:             stackLevelPtr,
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

		if hasId {
			if req.Id == nil {
				t.Fatal("Id = nil, want non-nil")
			}
			if id < 0 || id > math.MaxUint16 {
				t.Fatalf("Req succeeded with id=%d", id)
			}
			if got := req.Id.Value(); got != uint16(id) {
				t.Fatalf("Id = %d, want %d", got, id)
			}
		} else if req.Id != nil {
			t.Fatal("Id != nil, want nil")
		}

		if hasConnectorId {
			if req.ConnectorId == nil {
				t.Fatal("ConnectorId = nil, want non-nil")
			}
			if connectorId < 0 || connectorId > math.MaxUint16 {
				t.Fatalf("Req succeeded with connectorId=%d", connectorId)
			}
			if got := req.ConnectorId.Value(); got != uint16(connectorId) {
				t.Fatalf("ConnectorId = %d, want %d", got, connectorId)
			}
		} else if req.ConnectorId != nil {
			t.Fatal("ConnectorId != nil, want nil")
		}

		if hasStackLevel {
			if req.StackLevel == nil {
				t.Fatal("StackLevel = nil, want non-nil")
			}
			if stackLevel < 0 || stackLevel > math.MaxUint16 {
				t.Fatalf("Req succeeded with stackLevel=%d", stackLevel)
			}
			if got := req.StackLevel.Value(); got != uint16(stackLevel) {
				t.Fatalf("StackLevel = %d, want %d", got, stackLevel)
			}
		} else if req.StackLevel != nil {
			t.Fatal("StackLevel != nil, want nil")
		}

		if hasPurpose {
			if req.ChargingProfilePurpose == nil {
				t.Fatal("ChargingProfilePurpose = nil, want non-nil")
			}
			if !req.ChargingProfilePurpose.IsValid() {
				t.Fatalf(
					"ChargingProfilePurpose = %q, want valid",
					req.ChargingProfilePurpose.String(),
				)
			}
			if req.ChargingProfilePurpose.String() != purpose {
				t.Fatalf(
					"ChargingProfilePurpose = %q, want %q",
					req.ChargingProfilePurpose.String(),
					purpose,
				)
			}
		} else if req.ChargingProfilePurpose != nil {
			t.Fatal("ChargingProfilePurpose != nil, want nil")
		}
	})
}
