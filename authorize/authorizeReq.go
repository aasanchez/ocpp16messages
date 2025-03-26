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

// AuthorizeConf represents the Authorize.conf message sent by the central system to the charge point.
type AuthorizeConf struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"`
}

// IdTagInfo contains status and optional fields returned in Authorize.conf.
type IdTagInfo struct {
	Status      string  `json:"status"`
	ExpiryDate  *string `json:"expiryDate,omitempty"`
	ParentIdTag *string `json:"parentIdTag,omitempty"`
}

// Validate checks the validity of the AuthorizeConf.
func (c *AuthorizeConf) Validate() error {
	if c.IdTagInfo.Status == "" {
		return fmt.Errorf("status is required in IdTagInfo")
	}
	return nil
}

// NewAuthorizeConf creates a new AuthorizeConf with required and optional fields.
func NewAuthorizeConf(status string, expiryDate, parentIdTag *string) (*AuthorizeConf, error) {
	if status == "" {
		return nil, fmt.Errorf("status is required")
	}
	return &AuthorizeConf{
		IdTagInfo: IdTagInfo{
			Status:      status,
			ExpiryDate:  expiryDate,
			ParentIdTag: parentIdTag,
		},
	}, nil
}
