package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdToken struct {
	value sharedtypes.CiString20Type
}

func NewIdToken(s string) (IdToken, error) {
	ci, err := sharedtypes.CiString20(s)
	if err != nil {
		return IdToken{}, fmt.Errorf("invalid IdToken: %w", err)
	}
	return IdToken{value: ci}, nil
}

func (id IdToken) Value() string {
	return id.value.Value()
}
