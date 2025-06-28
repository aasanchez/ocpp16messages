package changeavailabilitytypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// RequestPayload provides fields to create a ChangeAvailability request.
type RequestPayload struct {
	ConnectorId string
	Type        string
}

// Validate checks for missing required fields in the request payload.
func (r RequestPayload) Validate() error {
	if r.ConnectorId == "" {
		return fmt.Errorf(st.ErrFmtFieldWrapped, "missing required field: ConnectorId", st.ErrEmptyValueNotAllowed)
	}
	if r.Type == "" {
		return fmt.Errorf(st.ErrFmtFieldWrapped, "missing required field: Type", st.ErrEmptyValueNotAllowed)
	}
	return nil
}
