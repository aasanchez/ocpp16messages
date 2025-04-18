package authorizetypes

import (
	"errors"
	"fmt"
)

var ErrMissingStatus = errors.New("missing required field: idTagInfo.Status")

type ConfirmationMessageInput struct {
	IdTagInfo IdTagInfoInput
}

type IdTagInfoInput struct {
	Status      string
	ExpiryDate  *string
	ParentIdTag *string
}

func (input ConfirmationMessageInput) Validate() error {
	if input.IdTagInfo.Status == "" {
		return fmt.Errorf("%w", ErrMissingStatus)
	}
	return nil
}
