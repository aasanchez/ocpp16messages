package chargePoint

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/models"
)

// AuthorizeReq represents the Authorize.req message sent by the Charge Point
// to validate the identity of a user via an IdTag.
type AuthorizeReq struct {
	IdTag models.CiString20Type `json:"idTag"`
}

// AuthorizeConf represents the Authorize.conf response sent by the Central System.
// It contains the authorization status and optional expiry/parent information.
type AuthorizeConf struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"`
}

// IdTagInfo contains information about the authorization status and validity period.
type IdTagInfo struct {
	Status      enums.AuthorizationStatus `json:"status"`
	ExpiryDate  *time.Time                `json:"expiryDate,omitempty"`
	ParentIdTag *models.CiString20Type    `json:"parentIdTag,omitempty"`
}
