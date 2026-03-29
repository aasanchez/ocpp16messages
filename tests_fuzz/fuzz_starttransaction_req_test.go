//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	stt "github.com/aasanchez/ocpp16messages/starttransaction"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzStartTransactionReq(f *testing.F) {
	f.Add(0, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(-1, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "RFID-ABC123", -1, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "RFID-ABC123", 0, "bad-timestamp", false, 0)
	f.Add(0, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", true, -1)

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		idTag string,
		meterStart int,
		timestamp string,
		hasReservationId bool,
		reservationId int,
	) {
		if len(idTag) > maxFuzzStringLen || len(timestamp) > maxFuzzStringLen {
			t.Skip()
		}

		var reservationIdPtr *int
		if hasReservationId {
			reservationIdPtr = &reservationId
		}

		req, err := stt.Req(stt.ReqInput{
			ConnectorId:   connectorId,
			IdTag:         idTag,
			MeterStart:    meterStart,
			Timestamp:     timestamp,
			ReservationId: reservationIdPtr,
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

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if got := req.ConnectorId.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorId = %d, want %d", got, connectorId)
		}

		if idTag == "" {
			t.Fatal("Req succeeded with empty IdTag")
		}

		if req.IdTag.String() != idTag {
			t.Fatalf("IdTag = %q, want %q", req.IdTag.String(), idTag)
		}

		if meterStart < 0 || meterStart > math.MaxUint16 {
			t.Fatalf("Req succeeded with meterStart=%d", meterStart)
		}

		if got := req.MeterStart.Value(); got != uint16(meterStart) {
			t.Fatalf("MeterStart = %d, want %d", got, meterStart)
		}

		if req.Timestamp.Value().Location() != time.UTC {
			t.Fatalf("Timestamp location = %v, want UTC", req.Timestamp.Value().Location())
		}

		if hasReservationId {
			if req.ReservationId == nil {
				t.Fatal("ReservationId = nil, want non-nil")
			}

			if reservationId < 0 || reservationId > math.MaxUint16 {
				t.Fatalf("Req succeeded with reservationId=%d", reservationId)
			}

			if got := req.ReservationId.Value(); got != uint16(reservationId) {
				t.Fatalf("ReservationId = %d, want %d", got, reservationId)
			}
		} else if req.ReservationId != nil {
			t.Fatal("ReservationId != nil, want nil")
		}
	})
}
