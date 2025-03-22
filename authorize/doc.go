// Package authorize implements the Authorize.req and Authorize.conf messages
// as defined in the Open Charge Point Protocol (OCPP) 1.6J specification.
//
// This package is part of the ocpp16_messages module and provides:
//
//   - Data structures for Authorize.req (sent from Charge Point to CSMS)
//   - Data structures for Authorize.conf (response from CSMS to Charge Point)
//   - Field-level validation for message payloads
//   - Constructors for safe and idiomatic creation of messages
//
// The Authorize feature is used by a Charge Point to request permission to
// authorize a user based on a presented identifier (idTag). The Central System
// responds with the authorization status, and optionally, information such as
// a parent idTag or expiry time.
//
// Example usage:
//
//	raw := []byte(`[2, "123456", "Authorize", {"idTag":"ABC123"}]`)
//	msg, err := ocpp16_messages.ValidateMessage(raw)
//
// This package is designed to be transport-agnostic and strictly schema-focused.
package authorize
