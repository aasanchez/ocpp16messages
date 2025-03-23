package core

import (
	"fmt"
)

// IdToken represents the identifier used by the charge point to authorize a user.
// It wraps a CiString20Type, which restricts the token to a maximum of 20 characters.
type IdToken struct {
	// IdTag is the unique identifier for the user or vehicle.
	IdTag CiString20Type `json:"idTag"`
}

// Validate checks whether the IdToken meets OCPP 1.6J requirements.
func (t IdToken) Validate() error {
	if err := t.IdTag.Validate(); err != nil {
		return NewFieldError("idTag", fmt.Errorf("invalid IdToken: %w", err))
	}
	return nil
}

// String returns the IdTag value as string.
func (t IdToken) String() string {
	return string(t.IdTag)
}

// NewIdToken creates and validates a new IdToken instance from a string.
func NewIdToken(id string) (IdToken, error) {
	tag := CiString20Type(id)
	if err := tag.Validate(); err != nil {
		return IdToken{}, NewFieldError("idTag", err)
	}
	return IdToken{IdTag: tag}, nil
}
