package changeavailabilitytypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfirmationPayload represents the payload of a ChangeAvailability confirmation message.
type ConfirmationPayload struct {
	Status string
}

// Validate ensures that required fields are present in the ConfirmationPayload.
func (c ConfirmationPayload) Validate() error {
	if c.Status == "" {
		return fmt.Errorf(st.ErrFmtFieldWrapped, "missing required field: Status", st.ErrEmptyValueNotAllowed)
	}

	return nil
}
