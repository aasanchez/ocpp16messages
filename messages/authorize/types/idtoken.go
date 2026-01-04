package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdToken represents an OCPP 1.6 identifier token used for authorization.
// It wraps a CiString20Type value that must not be empty.
type IdToken struct {
	value st.CiString20Type
}

// NewIdToken creates a new IdToken from a CiString20Type value.
// Returns an error if the provided value is empty.
func NewIdToken(ciString20Type st.CiString20Type) (IdToken, error) {
	if ciString20Type.Value() == "" {
		return IdToken{}, fmt.Errorf("NewIdToken: %w", st.ErrEmpty("IdToken"))
	}

	return IdToken{value: ciString20Type}, nil
}

// Value returns the underlying CiString20Type value of the IdToken.
func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

// String returns the string representation of the IdToken.
func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
