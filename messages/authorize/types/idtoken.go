package authorizetypes

import (
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdToken struct {
	value st.CiString20Type
}

func SetIdToken(cistring st.CiString20Type) (IdToken, error) {
	return IdToken{value: cistring}, nil
}

func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
