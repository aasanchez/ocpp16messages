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

func NewIdToken(CiString20Type st.CiString20Type) (IdToken, error) {
	if CiString20Type.Value() == "" {
		return IdToken{}, fmt.Errorf("NewIdToken: %w", st.ErrEmpty("IdToken"))
	}

	return IdToken{value: CiString20Type}, nil
}

// Value returns the underlying CiString20Type value of the IdToken.
func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}
