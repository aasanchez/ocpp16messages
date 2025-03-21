package models

// IdToken represents a unique identifier used to authorize a user.
// It is defined as a CiString20Type in the OCPP 1.6 specification.
type IdToken struct {
	IdTag CiString20Type `json:"idTag"`
}

// Validate checks that the IdTag is not empty and within limits.
func (t IdToken) Validate() error {
	return t.IdTag.Validate()
}
