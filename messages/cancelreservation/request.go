package cancelreservation

import (
	"errors"
	"fmt"
	"strconv"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var ErrInvalidRequestIdTag = errors.New("request -> invalid idTag")

type RequestMessage struct {
	ReservationId uint32
}

func Request(input cancelreservationtypes.RequestPayload) (RequestMessage, error) {
	reservationIdTemp, err := strconv.ParseUint(input.ReservationId, 10, 32)
	if err != nil {
		return RequestMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse ReservationId", err)
	}
	var reservationId = uint32(reservationIdTemp)

	return RequestMessage{ReservationId: reservationId}, nil
}
