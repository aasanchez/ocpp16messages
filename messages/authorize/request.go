package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type RequestMessage struct {
	IdTag authorizetypes.IdTokenType
}

func Request(idTag string) (RequestMessage, error) {
	IdToken, err := authorizetypes.IdToken(idTag)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("failed to create Authorize.req message, problem with idTag: %s: %w", idTag, err)
	}

	return RequestMessage{IdTag: IdToken}, nil
}

func (r RequestMessage) String() string {
	return fmt.Sprintf("{idTag:%s}", r.IdTag.String())
}

func (r RequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("RequestMessage validation failed: %w", err)
	}

	return nil
}
