package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationPayload struct {
	IdTagInfo IdTagInfoPayload
}

func (c ConfirmationPayload) Validate() error {
	if c.IdTagInfo.Status == "" {
		return fmt.Errorf("confirmation payload: %w", sharedtypes.ErrInvalidIdTagInfo)
	}

	return nil
}

func (c ConfirmationPayload) Value() ConfirmationPayload {
	return c
}
