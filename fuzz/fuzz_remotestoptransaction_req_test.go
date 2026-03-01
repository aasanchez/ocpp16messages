//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/aasanchez/ocpp16messages/remoteStopTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzRemoteStopTransactionReq(f *testing.F) {
	f.Add(1)
	f.Add(0)
	f.Add(-1)
	f.Add(math.MaxUint16 + 1)

	f.Fuzz(func(t *testing.T, transactionId int) {
		req, err := remoteStopTransaction.Req(remoteStopTransaction.ReqInput{
			TransactionId: transactionId,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if transactionId < 0 || transactionId > math.MaxUint16 {
			t.Fatalf("Req succeeded with transactionId=%d", transactionId)
		}

		if got := req.TransactionId.Value(); got != uint16(transactionId) {
			t.Fatalf("TransactionId = %d, want %d", got, transactionId)
		}
	})
}
