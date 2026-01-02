package types

import (
	"fmt"

	sharedTypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// IDToken represents an OCPP 1.6 identifier token used for authorization.
// It wraps a CiString20 value that must not be empty.
type IDToken struct {
	value sharedTypes.CiString20
}

// NewIDToken creates a new IDToken from a CiString20 value. Returns an error
// if the provided cistring value is empty.
func NewIDToken(cistring sharedTypes.CiString20) (IDToken, error) {
	if cistring.Value() == "" {
		return IDToken{}, fmt.Errorf(
			"NewIDToken: %w",
			sharedTypes.ErrEmpty("IDToken"),
		)
	}

	return IDToken{value: cistring}, nil
}

// Value returns the underlying CiString20 value of the IDToken.
func (idtoken IDToken) Value() sharedTypes.CiString20 {
	return idtoken.value
}

// String returns the string representation of the IDToken.
func (idtoken IDToken) String() string {
	return idtoken.value.Value()
}
