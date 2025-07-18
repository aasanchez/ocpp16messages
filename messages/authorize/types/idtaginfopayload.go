package authorizetypes

import (
	"errors"
	"fmt"
)

var errIdTagInfoPayload = errors.New("invalid IdTagInfoPayload")

type IdTagInfoPayload struct {
	ExpiryDate  string
	ParentIdTag string
	Status      string
}

func (input IdTagInfoPayload) Validate() error {
	if input.Status == "" {
		return fmt.Errorf("%w: %q", errIdTagInfoPayload, input)
	}

	return nil
}
