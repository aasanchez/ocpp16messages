package changeavailabilitytypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// RequestPayload provides fields to create a ChangeAvailability request.
type RequestPayload struct {
	ConnectorId string
	Type        string
}

// Validate checks for missing required fields in the request payload.
func (r RequestPayload) Validate() error {
	if r.ConnectorId == "" {
		return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "missing required field: ConnectorId", sharedtypes.ErrEmptyValueNotAllowed)
	}
	if r.Type == "" {
		return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "missing required field: Type", sharedtypes.ErrEmptyValueNotAllowed)
	}
	return nil
}
