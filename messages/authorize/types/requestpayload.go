package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// RequestPayload is the payload for the Authorize.req message.
// It contains the IdTag that needs to be authorized.
// See OCPP 1.6J Part 2, Section 4.1.2.1, Table 12 for details.
type RequestPayload struct {
	IdTag string // IdTag is the identifier that needs to be authorized. (Required)
}

// Validate checks if the RequestPayload is valid according to OCPP 1.6J specification.
// It ensures that the `IdTag` is a valid CiString20Type.
//
// Returns:
//
//	error: `nil` if the payload is valid, otherwise an error detailing the validation failure.
func (r RequestPayload) Validate() error {
	_, err := sharedtypes.SetCiString20Type(r.IdTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

// Value returns the string representation of the IdTag.
//
// Returns:
//
//	string: The string content of the IdTag.
func (r RequestPayload) Value() string {
	return r.IdTag
}
