package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type RequestMessage struct {
	IdTag authorizetypes.IdTokenType
}

func Request(input authorizetypes.RequestMessageInput) (RequestMessage, error) {
	idToken, err := authorizetypes.IdToken(input.IdTag)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("failed to create Authorize.req message, problem with idTag: %s: %w", input.IdTag, err)
	}

	return RequestMessage{IdTag: idToken}, nil
}

func (r RequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("RequestMessage validation failed: %w", err)
	}

	return nil
}
