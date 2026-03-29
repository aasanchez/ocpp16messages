package cancelreservation

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a CancelReservation.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	ReservationId int // Required: ID of the reservation to cancel
}

// ReqMessage represents an OCPP 1.6 CancelReservation.req message.
type ReqMessage struct {
	ReservationId types.Integer
}

// Req creates a CancelReservation.req message from the given input.
// It validates all fields and returns an error if:
//   - ReservationId is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	reservationId, err := types.NewInteger(input.ReservationId)
	if err != nil {
		return ReqMessage{}, fmt.Errorf("reservationId: %w", err)
	}

	return ReqMessage{ReservationId: reservationId}, nil
}
