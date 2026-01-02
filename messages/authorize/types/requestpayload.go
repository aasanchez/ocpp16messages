package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// RequestPayload represents the OCPP 1.6 Authorize request payload structure.
type RequestPayload struct {
	IDTag string
}

// Validate checks if the RequestPayload is valid according to OCPP 1.6 spec.
// Returns an error if IDTag is empty or exceeds CiString20 length constraints.
func (r RequestPayload) Validate() error {
	if r.IDTag == "" {
		return fmt.Errorf("request payload: %w", ErrInvalidIDTag)
	}

	_, err := st.NewCiString20(r.IDTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

// Value returns the IDTag value from the RequestPayload.
func (r RequestPayload) Value() string {
	return r.IDTag
}
