package types

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// Request represents the OCPP 1.6 Authorize request payload structure.
type Request struct {
	IDTag string
}

// Validate checks if the Request is valid according to OCPP 1.6 spec.
// Returns an error if IDTag is empty or exceeds CiString20 length constraints.
func (r Request) Validate() error {
	if r.IDTag == "" {
		return fmt.Errorf("request payload: %w", sharedtypes.ErrInvalidValue)
	}

	_, err := sharedtypes.NewCiString20(r.IDTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

// Value returns the IDTag value from the Request.
func (r Request) Value() string {
	return r.IDTag
}
