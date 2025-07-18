package authorizetypes

import (
	"errors"
	"fmt"
)

var errMissingIdTag = errors.New("error: Missing IdTag ")

type RequestPayload struct {
	idTag string
}

func (input RequestPayload) Validate() error {
	if input.idTag == "" {
		return fmt.Errorf("%w: (%q)", errMissingIdTag, input)
	}

	return nil
}

func (input RequestPayload) IdTag() string {
	return input.idTag
}
