package authorizetypes

import (
	"errors"
	"fmt"
)

var ErrInvalidIdTagInfo = errors.New("invalid idTagInfo")

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
