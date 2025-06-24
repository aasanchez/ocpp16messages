package authorizetypes

import (
	"fmt"
)

type ConfirmationPayload struct {
	IdTagInfo IdTagInfoPayload
}

func (c ConfirmationPayload) Validate() error {
	if c.IdTagInfo.Status == "" {
		return fmt.Errorf("confirmation payload: %w", ErrInvalidIdTagInfo)
	}

	return nil
}

func (c ConfirmationPayload) Value() ConfirmationPayload {
	return c
}
