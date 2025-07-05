package authorizetypes

import (
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTokenType struct {
	value sharedtypes.CiString20
}

func IdToken(s sharedtypes.CiString20) (IdTokenType, error) {
	return IdTokenType{value: s}, nil
}

func (id IdTokenType) Value() string {
	return id.value.Value()
}
