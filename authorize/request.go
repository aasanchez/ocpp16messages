// Package authorize implements the Authorize.req message as defined in the
// Open Charge Point Protocol (OCPP) 1.6J specification.
//
// This file provides the schema definition and validation logic for Authorize.req,
// which is used by the Charge Point to request authorization for an idTag.
package authorize

// AuthorizeReq represents the payload of an Authorize CALL message (MessageTypeId = 2).
// According to the OCPP 1.6J spec, this message contains a single field:
//
//	{
//	  "idTag": "ABC123456"
//	}
//
// The idTag must be a non-empty, case-insensitive string with a maximum length of 20 characters.
type AuthorizeReq struct {
	IdTag string `json:"idTag"` // Required. IdTag to be authorized by the Central System.
}
