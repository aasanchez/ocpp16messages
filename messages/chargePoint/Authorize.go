package chargePoint

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/models"
)

type AuthorizeReq struct {
	IdToken models.IdToken `json:"idTag"`
}

type AuthorizeConf struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"`
}

type IdTagInfo struct {
	Status      enums.AuthorizationStatus `json:"status"`
	ExpiryDate  *time.Time                `json:"expiryDate,omitempty"`
	ParentIdTag *models.CiString20Type    `json:"parentIdTag,omitempty"`
}
