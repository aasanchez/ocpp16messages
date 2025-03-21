package authorizetypes

import (
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdToken struct {
	value st.CiString20Type
}

func SetIdToken(s st.CiString20Type) (IdToken, error) {
	return IdToken{value: s}, nil
}

func (id IdToken) Value() string {
	return id.value.Value()
}
