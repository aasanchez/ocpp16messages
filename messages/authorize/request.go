package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	IdTag authorizetypes.IdToken
}

func Request(input authorizetypes.RequestPayload) (RequestMessage, error) {
	str, err := sharedtypes.SetCiString20(input.IdTag)
	if err != nil {
		wrapped := fmt.Errorf("request -> invalid idTag -> %w", err)

		return RequestMessage{}, wrapped
	}

	idToken, _ := authorizetypes.SetIdToken(str)

	return RequestMessage{IdTag: idToken}, nil
}
