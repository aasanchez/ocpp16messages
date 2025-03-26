// Package authorize contains the data structures and validation logic
// for the Authorize.req and Authorize.conf OCPP 1.6J messages.
package authorize

import (
	"github.com/aasanchez/ocpp16_messages/core/types"
)

// AuthorizeConf represents the Authorize.conf message sent by the central system to the charge point.
type AuthorizeConf struct {
	IdTagInfo types.IdToken `json:"idTagInfo"`
}

// // Validate checks the validity of the AuthorizeConf.
// func (c *AuthorizeConf) Validate() error {
// 	if c.IdTagInfo.Status == "" {
// 		return fmt.Errorf("status is required in IdTagInfo")
// 	}
// 	return nil
// }

// // NewAuthorizeConf creates a new AuthorizeConf with required and optional fields.
// func NewAuthorizeConf(status string, expiryDate, parentIdTag *string) (*AuthorizeConf, error) {
// 	if status == "" {
// 		return nil, fmt.Errorf("status is required")
// 	}
// 	return &AuthorizeConf{
// 		IdTagInfo: IdTagInfo{
// 			ExpiryDate:  expiryDate,
// 			ParentIdTag: parentIdTag,
// 			Status:      status,
// 		},
// 	}, nil
// }
