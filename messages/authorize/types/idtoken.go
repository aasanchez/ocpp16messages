package authorizetypes

import (
	"errors"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ErrIdTokenEmpty is a static error indicating that the IdToken value is empty.
var ErrIdTokenEmpty = errors.New("IdToken cannot be empty")

// IdToken represents a validated IdToken value in OCPP 1.6.
type IdToken struct {
	value st.CiString20Type
}

// SetIdToken creates a new IdToken if the input is valid.
func SetIdToken(cistring st.CiString20Type) (IdToken, error) {
	if cistring.Value() == "" {
		return IdToken{}, ErrIdTokenEmpty
	}

	return IdToken{value: cistring}, nil
}

// Value returns the underlying value of the IdToken.
func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

// String returns the string representation of the IdToken.
func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
