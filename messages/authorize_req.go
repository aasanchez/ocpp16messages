package authorize

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// AuthorizeRequestMessage represents the OCPP 1.6J Authorize.req message payload.
// This message is sent from the Charge Point to the Central System
// to request authorization for a given idTag.
type AuthorizeRequestMessage struct {
	IdTag types.IdTagType
}

// AuthorizeRequest creates a new Authorize.req payload using a raw idTag string.
// It validates the idTag and returns a fully-constructed message or an error.
func AuthorizeRequest(idTag string) (AuthorizeRequestMessage, error) {
	tag, err := types.IdTag(idTag)
	if err != nil {
		return AuthorizeRequestMessage{}, fmt.Errorf("failed to create AuthorizeRequestMessage: %w", err)
	}

	return AuthorizeRequestMessage{IdTag: tag}, nil
}

// String returns a debug-friendly representation of the request.
func (r AuthorizeRequestMessage) String() string {
	return fmt.Sprintf("AuthorizeRequestMessage{idTag=%s}", r.IdTag.String())
}

// Validate re-validates the message according to OCPP rules.
func (r AuthorizeRequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("AuthorizeRequestMessage validation failed: %w", err)
	}

	return nil
}
