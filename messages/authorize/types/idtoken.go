package types

import (
	"fmt"

	sharedTypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdToken represents an OCPP 1.6 identifier token used for authorization.
// It wraps a CiString20 value that must not be empty.
type IdToken struct {
	value sharedTypes.CiString20
}

// SetIdToken creates a new IdToken from a CiString20 value. Returns an error
// if the provided cistring value is empty.
func SetIdToken(cistring sharedTypes.CiString20) (IdToken, error) {
	if cistring.Value() == "" {
		return IdToken{}, fmt.Errorf(
			"SetIdToken: %w",
			sharedTypes.ErrEmpty("IdToken"),
		)
	}

	return IdToken{value: cistring}, nil
}

// Value returns the underlying CiString20 value of the IdToken.
func (idtoken IdToken) Value() sharedTypes.CiString20 {
	return idtoken.value
}

// String returns the string representation of the IdToken.
func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
