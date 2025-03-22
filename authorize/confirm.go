// Package authorize contains the data structures and validation logic
// for the Authorize.req and Authorize.conf messages defined in the OCPP 1.6J specification.
package authorize

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/core"
)

// Conf represents the Authorize.conf message sent from the CSMS to the Charge Point.
//
// It includes the IdTagInfo structure, which holds authorization status,
// and optionally, expiration and parent identifiers.
//
// According to OCPP 1.6J, the Authorize.conf format is:
//
//	{
//	  "idTagInfo": {
//	    "status": "Accepted",
//	    "expiryDate": "2024-10-01T10:00:00Z",
//	    "parentIdTag": "ABC123"
//	  }
//	}
type Conf struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"`
}

// IdTagInfo describes the status and related metadata for an authorization token.
//
// This structure is used in multiple OCPP messages.
type IdTagInfo struct {
	// Status indicates whether the token is accepted, expired, blocked, etc.
	Status core.AuthorizationStatus `json:"status"`

	// ExpiryDate defines the date and time when the token authorization expires.
	// Optional.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`

	// ParentIdTag is a reference to another token that authorized this one.
	// Optional.
	ParentIdTag *core.CiString20Type `json:"parentIdTag,omitempty"`
}
