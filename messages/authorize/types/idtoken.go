package authorizetypes

import (
	"errors"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

var ErrIdTokenEmpty = errors.New("IdToken cannot be empty")

type IdToken struct {
	value st.CiString20Type
}

func SetIdToken(cistring st.CiString20Type) (IdToken, error) {
	if cistring.Value() == "" {
		return IdToken{}, ErrIdTokenEmpty
	}

	return IdToken{value: cistring}, nil
}

func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
