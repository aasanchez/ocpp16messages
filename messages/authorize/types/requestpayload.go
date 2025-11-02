package authorizetypes

import (
	"errors"
)

// ErrIdTagMismatch indicates the idTag value does not match expected value
var ErrIdTagMismatch = errors.New("idTag value mismatch")

type RequestPayload struct {
	idTag string
}

// Validate checks if the RequestPayload has valid data.
func (input RequestPayload) Validate() error {
	if input.idTag == "" {
		return ErrIdTagMismatch
	}

	return nil
}

// IdTag returns the idTag value.
func (input RequestPayload) IdTag() string {
	return input.idTag
}
