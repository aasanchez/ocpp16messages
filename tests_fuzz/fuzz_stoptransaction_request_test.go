//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	stp "github.com/aasanchez/ocpp16messages/stoptransaction"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzStopTransactionReq(f *testing.F) {
	f.Add(1, 100, "2025-01-15T10:30:00Z", false, "", false, "", uint8(0))
	f.Add(
		1,
		100,
		"2025-01-15T10:30:00Z",
		true,
		"RFID-ABC123",
		true,
		types.ReasonLocal.String(),
		uint8(2),
	)
	f.Add(-1, -1, "bad-timestamp", false, "", false, "", uint8(2))
	f.Add(1, 100, "2025-01-15T10:30:00Z", true, "", false, "", uint8(0))
	f.Add(1, 100, "2025-01-15T10:30:00Z", false, "", true, "bad-reason", uint8(0))

	f.Fuzz(func(
		t *testing.T,
		transactionId int,
		meterStop int,
		timestamp string,
		hasIdTag bool,
		idTag string,
		hasReason bool,
		reason string,
		transactionDataMode uint8,
	) {
		if len(timestamp) > maxFuzzStringLen ||
			len(idTag) > maxFuzzStringLen ||
			len(reason) > maxFuzzStringLen {
			t.Skip()
		}

		var idTagPtr *string
		if hasIdTag {
			idTagPtr = &idTag
		}

		var reasonPtr *string
		if hasReason {
			reasonPtr = &reason
		}

		var transactionData []types.MeterValueInput

		switch transactionDataMode % 3 {
		case 0:
			transactionData = nil
		case 1:
			transactionData = []types.MeterValueInput{}
		default:
			transactionData = []types.MeterValueInput{
				{
					Timestamp: timestamp,
					SampledValue: []types.SampledValueInput{
						{Value: "100"},
					},
				},
			}
		}

		req, err := stp.Req(stp.ReqInput{
			TransactionId:   transactionId,
			IdTag:           idTagPtr,
			MeterStop:       meterStop,
			Timestamp:       timestamp,
			Reason:          reasonPtr,
			TransactionData: transactionData,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if transactionId < 0 || transactionId > math.MaxUint16 {
			t.Fatalf("Req succeeded with transactionId=%d", transactionId)
		}

		if got := req.TransactionId.Value(); got != uint16(transactionId) {
			t.Fatalf("TransactionId = %d, want %d", got, transactionId)
		}

		if meterStop < 0 || meterStop > math.MaxUint16 {
			t.Fatalf("Req succeeded with meterStop=%d", meterStop)
		}

		if got := req.MeterStop.Value(); got != uint16(meterStop) {
			t.Fatalf("MeterStop = %d, want %d", got, meterStop)
		}

		if req.Timestamp.Value().Location() != time.UTC {
			t.Fatalf(
				"Timestamp location = %v, want UTC",
				req.Timestamp.Value().Location(),
			)
		}

		if hasIdTag {
			if req.IdTag == nil {
				t.Fatal("IdTag = nil, want non-nil")
			}
			if idTag == "" {
				t.Fatal("Req succeeded with empty IdTag")
			}
			if req.IdTag.String() != idTag {
				t.Fatalf("IdTag = %q, want %q", req.IdTag.String(), idTag)
			}
		} else if req.IdTag != nil {
			t.Fatal("IdTag != nil, want nil")
		}

		if hasReason {
			if req.Reason == nil {
				t.Fatal("Reason = nil, want non-nil")
			}
			if !req.Reason.IsValid() {
				t.Fatalf("Reason = %q, want valid", req.Reason.String())
			}
		} else if req.Reason != nil {
			t.Fatal("Reason != nil, want nil")
		}

		switch transactionDataMode % 3 {
		case 0:
			if req.TransactionData != nil {
				t.Fatal("TransactionData != nil, want nil")
			}
		case 1:
			if req.TransactionData == nil {
				t.Fatal("TransactionData = nil, want empty slice")
			}
			if len(req.TransactionData) != 0 {
				t.Fatalf("len(TransactionData) = %d, want 0", len(req.TransactionData))
			}
		default:
			if len(req.TransactionData) == 0 {
				t.Fatal("TransactionData is empty, want at least one")
			}
		}
	})
}
