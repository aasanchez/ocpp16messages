package cancelreservationtypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationPayload struct {
	Status string
}

func (r ConfirmationPayload) Validate() error {
	if r.Status == "" {
		return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "missing required field: status", sharedtypes.ErrEmptyValueNotAllowed)
	}

	return nil
}
