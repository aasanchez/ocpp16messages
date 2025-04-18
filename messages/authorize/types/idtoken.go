package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTokenType struct {
	value sharedtypes.CiString20Type
}

func IdToken(s string) (IdTokenType, error) {
	ci, err := sharedtypes.CiString20(s)
	if err != nil {
		return IdTokenType{}, fmt.Errorf("IdToken: invalid IdToken, %w", err)
	}

	return IdTokenType{value: ci}, nil
}

func (id IdTokenType) Validate() error {
	if err := id.value.Validate(); err != nil {
		return fmt.Errorf("IdToken validation failed: %w", err)
	}

	return nil
}
