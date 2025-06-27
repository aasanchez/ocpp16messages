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
	ReservationId uint64
}

func Request(input cancelreservationtypes.RequestPayload) (RequestMessage, error) {
	reservationId, err := strconv.ParseUint(input.ReservationId, 10, 64)
	if err != nil {
		return RequestMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse ReservationId", err)
	}
	return RequestMessage{ReservationId: reservationId}, nil
}
