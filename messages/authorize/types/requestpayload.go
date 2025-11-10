package authorizetypes

import (
	"errors"
)

var ErrIdTagMismatch = errors.New("idTag value mismatch")

type RequestPayload struct {
	idTag string
}

func (input RequestPayload) Validate() error {
	if input.idTag == "" {
		return ErrIdTagMismatch
	}

	return nil
}

func (input RequestPayload) IdTag() string {
	return input.idTag
}
