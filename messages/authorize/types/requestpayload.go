package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestPayload struct {
	IdTag string
}

func (r RequestPayload) Validate() error {
	if r.IdTag == "" {
		return fmt.Errorf("request payload: %w", sharedtypes.ErrInvalidIdTag)
	}

	_, err := sharedtypes.CiString20(r.IdTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

func (r RequestPayload) Value() string {
	return r.IdTag
}
