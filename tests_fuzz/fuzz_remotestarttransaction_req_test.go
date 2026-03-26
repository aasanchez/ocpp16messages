//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzRemoteStartTransactionReq(f *testing.F) {
	f.Add("RFID-ABC123", false, 0)
	f.Add("", false, 0)
	f.Add("RFID-ABC123", true, 1)
	f.Add("RFID-ABC123", true, -1)

	f.Fuzz(func(t *testing.T, idTag string, hasConnectorId bool, connectorId int) {
		if len(idTag) > maxFuzzStringLen {
			t.Skip()
		}

		var connectorIdPtr *int
		if hasConnectorId {
			connectorIdPtr = &connectorId
		}

		req, err := rst.Req(rst.ReqInput{
			IdTag:       idTag,
			ConnectorId: connectorIdPtr,
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

		if idTag == "" {
			t.Fatal("Req succeeded with empty IdTag")
		}

		if req.IdTag.String() != idTag {
			t.Fatalf("IdTag = %q, want %q", req.IdTag.String(), idTag)
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
	})
}
