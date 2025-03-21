package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTagInfoPayload struct {
	ExpiryDate  string
	ParentIdTag string
	Status      string
}

func (input IdTagInfoPayload) Validate() error {
	if input.Status == "" {
		return fmt.Errorf(
			st.ErrFmtFieldWrapped,
			"missing required field: Status",
			st.ErrEmptyValueNotAllowed,
		)
	}

	return nil
}
