// Package core defines shared OCPP data types and their validation logic.
//
// This file provides the IdToken type, which wraps the idTag field used in
// several messages such as Authorize.req, StartTransaction.req, and others.
package core

// IdToken represents a tag identifier (idTag) that is typically included
// as part of the authorization process for a transaction or user.
//
// This is used in nested fields (not as the top-level message field).
type IdToken struct {
	IdTag CiString20Type `json:"idTag"` // Required. Max 20 characters.
}

// Validate checks that the embedded idTag is valid.
func (t IdToken) Validate() error {
	return t.IdTag.Validate()
}
