// Package authorize defines the structure and validation for the Authorize.req message.
package authorize

import "github.com/aasanchez/ocpp16_messages/core"

// Req represents the Authorize.req message as defined in OCPP 1.6J.
//
// The Charge Point sends this message to the Central System to request
// authorization for a specific idTag before a transaction begins.
type Req struct {
	// IdTag is the identifier used for authorization. Required. Max 20 characters.
	IdTag core.IdToken `json:"idTag"`
}
