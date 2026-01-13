package types

import (
	st "github.com/aasanchez/ocpp16messages/types"
)

// IdToken represents an OCPP 1.6 identifier token used for authorization.
// It wraps a CiString20Type value that must not be empty.
type IdToken struct {
	value st.CiString20Type
}

// NewIdToken creates a new IdToken from a CiString20Type value.
// CiString20Type already validates that the value is non-empty.
func NewIdToken(ciString20Type st.CiString20Type) IdToken {
	return IdToken{value: ciString20Type}
}

// Value returns the underlying CiString20Type value of the IdToken.
func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

// String returns the string representation of the IdToken.
func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
