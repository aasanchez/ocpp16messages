package authorize

import (
	"errors"
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var ErrInvalidRequestIdTag = errors.New("request -> invalid IdTag")

type RequestMessage struct {
	IdTag authorizetypes.IdTokenType
}

func Request(input authorizetypes.RequestPayload) (RequestMessage, error) {
	str, err := sharedtypes.SetCiString20Type(input.IdTag)
	if err != nil {
		wrapped := fmt.Errorf("request -> invalid idTag -> %w", err)

		return RequestMessage{}, wrapped
	}

	idToken, _ := authorizetypes.IdToken(str)

	return RequestMessage{IdTag: idToken}, nil
}
