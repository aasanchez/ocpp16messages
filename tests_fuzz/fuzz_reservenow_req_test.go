//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	rn "github.com/aasanchez/ocpp16messages/reserveNow"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzReserveNowReq(f *testing.F) {
	f.Add(1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(-1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, -1, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, 0, "", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, 0, "RFID-ABC123", "bad-time", false, "")
	f.Add(1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", true, "")

	f.Fuzz(func(
		t *testing.T,
		reservationId int,
		connectorId int,
		idTag string,
		expiryDate string,
		hasParentIdTag bool,
		parentIdTag string,
	) {
		if len(idTag) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIdTag) > maxFuzzStringLen {
			t.Skip()
		}

		var parentIdTagPtr *string
		if hasParentIdTag {
			parentIdTagPtr = &parentIdTag
		}

		req, err := rn.Req(rn.ReqInput{
			ReservationId: reservationId,
			ConnectorId:   connectorId,
			IdTag:         idTag,
			ExpiryDate:    expiryDate,
			ParentIdTag:   parentIdTagPtr,
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

		if reservationId < 0 || reservationId > math.MaxUint16 {
			t.Fatalf("Req succeeded with reservationId=%d", reservationId)
		}

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if idTag == "" {
			t.Fatal("Req succeeded with empty IdTag")
		}

		if req.ReservationId.Value() != uint16(reservationId) {
			t.Fatalf(
				"ReservationId = %d, want %d",
				req.ReservationId.Value(),
				reservationId,
			)
		}

		if req.ConnectorId.Value() != uint16(connectorId) {
			t.Fatalf("ConnectorId = %d, want %d", req.ConnectorId.Value(), connectorId)
		}

		if req.IdTag.String() != idTag {
			t.Fatalf("IdTag = %q, want %q", req.IdTag.String(), idTag)
		}

		if req.ExpiryDate.Value().Location() != time.UTC {
			t.Fatalf(
				"ExpiryDate location = %v, want UTC",
				req.ExpiryDate.Value().Location(),
			)
		}

		if hasParentIdTag {
			if req.ParentIdTag == nil {
				t.Fatal("ParentIdTag = nil, want non-nil")
			}
			if parentIdTag == "" {
				t.Fatal("Req succeeded with empty ParentIdTag")
			}
			if req.ParentIdTag.String() != parentIdTag {
				t.Fatalf(
					"ParentIdTag = %q, want %q",
					req.ParentIdTag.String(),
					parentIdTag,
				)
			}
		} else if req.ParentIdTag != nil {
			t.Fatal("ParentIdTag != nil, want nil")
		}
	})
}
