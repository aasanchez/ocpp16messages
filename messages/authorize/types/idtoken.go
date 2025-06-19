package authorizetypes

import (
	"errors"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// ErrInvalidIdToken is returned when the input does not conform to CiString20Type constraints.
var ErrInvalidIdToken = errors.New("invalid IdToken")

type IdTokenType struct {
	value sharedtypes.CiString20Type
}

// IdToken validates the input string as CiString20Type and wraps it in an IdTokenType.
func IdToken(s sharedtypes.CiString20Type) (IdTokenType, error) {
	return IdTokenType{value: s}, nil
}

func (id IdTokenType) Value() string {
	return id.value.Value()
}
