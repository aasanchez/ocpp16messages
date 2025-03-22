// Package core defines shared data structures and validation logic used across OCPP 1.6 message types.
package core

// IdToken represents an identification token used in the Authorize.req message.
// It wraps a CiString20Type and enforces OCPP 1.6 constraints.
type IdToken struct {
	// IdTag is a case-insensitive string identifier, max 20 characters.
	IdTag CiString20Type `json:"idTag"`
}

// Validate checks whether the IdToken is valid according to the OCPP 1.6 specification.
// It ensures the IdTag field is not empty and does not exceed the maximum allowed length.
func (t IdToken) Validate() error {
	if err := t.IdTag.Validate(); err != nil {
		return NewFieldError("idTag", err)
	}
	return nil
}
