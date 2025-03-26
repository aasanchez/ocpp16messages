// Package authorize contains the data structures and validation logic
// for the Authorize.req and Authorize.conf OCPP 1.6J messages.
package authorize

import (
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core/types"
)

// AuthorizeReq represents the Authorize.req message sent by the charge point to the central system.
type AuthorizeReq struct {
	IdTag types.IdToken `json:"idTag"`
}

// NewAuthorizeReq creates a new instance of AuthorizeReq.
func NewAuthorizeReq(idTag string) (*AuthorizeReq, error) {
	it, err := types.NewIdToken(idTag)
	if err != nil {
		return nil, fmt.Errorf("invalid idTag: %w", err)
	}
	return &AuthorizeReq{IdTag: *it}, nil
}
