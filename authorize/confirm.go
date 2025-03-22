// Package authorize implements the Authorize.conf message as defined in the
// Open Charge Point Protocol (OCPP) 1.6J specification.
//
// This file provides the schema definition for Authorize.conf, which is returned
// by the Central System in response to an Authorize.req.
package authorize

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/core"
)

// AuthorizeConf represents the payload of an Authorize CALLRESULT message
// (MessageTypeId = 3). It is returned by the Central System in response
// to an Authorize.req message.
//
// According to the OCPP 1.6J specification:
//
//	{
//	  "idTagInfo": {
//	    "status": "Accepted",
//	    "expiryDate": "2025-03-22T14:00:00Z",
//	    "parentIdTag": { "idTag": "PARENT001" }
//	  }
//	}
type AuthorizeConf struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"` // Required. Result of the authorization request.
}

// IdTagInfo provides authorization information for the associated idTag.
//
// This structure is also reused in other messages such as StartTransaction.conf.
type IdTagInfo struct {
	Status      core.AuthorizationStatus `json:"status"`                // Required. Authorization status for the idTag.
	ExpiryDate  *time.Time               `json:"expiryDate,omitempty"`  // Optional. Time until this authorization is valid.
	ParentIdTag *core.IdToken            `json:"parentIdTag,omitempty"` // Optional. Linked parent idTag for hierarchical authorization.
}
