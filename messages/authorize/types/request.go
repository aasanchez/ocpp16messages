package types

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// Request represents the OCPP 1.6 Authorize request payload structure.
type Request struct {
	IdTag string
}

// Validate checks if the Request is valid per OCPP 1.6 spec.
// Returns an error if IdTag is empty or exceeds CiString20Type length.
func (r Request) Validate() error {
	if r.IdTag == "" {
		return fmt.Errorf("request payload: %w", sharedtypes.ErrInvalidValue)
	}

	_, err := sharedtypes.NewCiString20Type(r.IdTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

// Value returns the IdTag value from the Request.
func (r Request) Value() string {
	return r.IdTag
}
