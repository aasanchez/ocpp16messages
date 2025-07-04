package cancelreservationtypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestPayload struct {
	ReservationId string
}

func (r RequestPayload) Validate() error {
	if r.ReservationId == "" {
		return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "missing required field: ReservationId", sharedtypes.ErrEmptyValueNotAllowed)
	}

	return nil
}
