package messages

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

type AuthorizeRequestMessage struct {
	IdTag types.IdTag
}

func AuthorizeRequest(idTag string) (AuthorizeRequestMessage, error) {
	tag, err := types.NewIdTag(idTag)
	if err != nil {
		return AuthorizeRequestMessage{}, fmt.Errorf("failed to create AuthorizeRequestMessage: %w", err)
	}

	return AuthorizeRequestMessage{IdTag: tag}, nil
}

func (r AuthorizeRequestMessage) String() string {
	return fmt.Sprintf("AuthorizeRequestMessage{idTag=%s}", r.IdTag.String())
}

func (r AuthorizeRequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("AuthorizeRequestMessage validation failed: %w", err)
	}

	return nil
}
