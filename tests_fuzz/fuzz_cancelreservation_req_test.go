//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/aasanchez/ocpp16messages/cancelReservation"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzCancelReservationReq(f *testing.F) {
	f.Add(1)
	f.Add(0)
	f.Add(-1)
	f.Add(math.MaxUint16 + 1)

	f.Fuzz(func(t *testing.T, reservationId int) {
		req, err := cancelReservation.Req(cancelReservation.ReqInput{
			ReservationId: reservationId,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if reservationId < 0 || reservationId > math.MaxUint16 {
			t.Fatalf("Req succeeded with reservationId=%d", reservationId)
		}

		if got := req.ReservationId.Value(); got != uint16(reservationId) {
			t.Fatalf("ReservationId = %d, want %d", got, reservationId)
		}
	})
}
