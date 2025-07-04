package authorizetypes

import (
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdToken struct {
	value sharedtypes.CiString20
}

func SetIdToken(s sharedtypes.CiString20) (IdToken, error) {
	return IdToken{value: s}, nil
}

func (id IdToken) Value() string {
	return id.value.Value()
}
