package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestPayload struct {
	idTag string
}

func (input RequestPayload) Validate() error {
	if input.idTag == "" {
		return fmt.Errorf(st.ErrFmtFieldWrapped,
			"missing required field: idTag",
			st.ErrEmptyValueNotAllowed,
		)
	}

	return nil
}

func (input RequestPayload) IdTag() string {
	return input.idTag
}
