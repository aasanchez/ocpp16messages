package authorize

import (
	"errors"
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var ErrInvalidRequestIdTag = errors.New("request: invalid idTag")

type RequestMessage struct {
	IdTag authorizetypes.IdTokenType
}

func Request(input authorizetypes.RequestPayload) (RequestMessage, error) {
	ci, err := sharedtypes.CiString20(input.IdTag)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("%w: %v", ErrInvalidRequestIdTag, err)
	}

	idToken, _ := authorizetypes.IdToken(ci)

	return RequestMessage{IdTag: idToken}, nil
}
