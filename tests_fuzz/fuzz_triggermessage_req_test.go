//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	tm "github.com/aasanchez/ocpp16messages/triggerMessage"
	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzTriggerMessageReq(f *testing.F) {
	f.Add(tt.MessageTriggerBootNotification.String(), false, 0)
	f.Add(tt.MessageTriggerHeartbeat.String(), true, 1)
	f.Add("invalid-trigger", false, 0)
	f.Add(tt.MessageTriggerMeterValues.String(), true, -1)

	f.Fuzz(func(
		t *testing.T,
		requestedMessage string,
		hasConnectorId bool,
		connectorId int,
	) {
		if len(requestedMessage) > maxFuzzStringLen {
			t.Skip()
		}

		var connectorIdPtr *int
		if hasConnectorId {
			connectorIdPtr = &connectorId
		}

		req, err := tm.Req(tm.ReqInput{
			RequestedMessage: requestedMessage,
			ConnectorId:      connectorIdPtr,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !req.RequestedMessage.IsValid() {
			t.Fatalf("RequestedMessage = %q, want valid", req.RequestedMessage.String())
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
