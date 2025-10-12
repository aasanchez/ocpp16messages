package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdToken struct {
	value st.CiString20Type
}

func SetIdToken(cistring st.CiString20Type) (IdToken, error) {
	if cistring.Value() == "" {
		return IdToken{}, fmt.Errorf("IdToken cant be empty")
	}
	return IdToken{value: cistring}, nil
}

func (idtoken IdToken) Value() st.CiString20Type {
	return idtoken.value
}

func (idtoken IdToken) String() string {
	return idtoken.value.Value()
}
