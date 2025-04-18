package authorizetypes

import (
	"errors"
	"fmt"
)

type RequestMessageInput struct {
	IdTag string
}

var ErrMissingIdTag = errors.New("missing required field: idTag")

func (input RequestMessageInput) Validate() error {
	if input.IdTag == "" {
		return fmt.Errorf("%w", ErrMissingIdTag)
	}
	return nil
}
