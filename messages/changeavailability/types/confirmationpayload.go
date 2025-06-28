package changeavailabilitytypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfirmationPayload represents the payload of a ChangeAvailability confirmation message.
type ConfirmationPayload struct {
	Status string
}

// Validate ensures that required fields are present in the ConfirmationPayload.
func (c ConfirmationPayload) Validate() error {
	if c.Status == "" {
		return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "missing required field: Status", sharedtypes.ErrEmptyValueNotAllowed)
	}
	return nil
}
