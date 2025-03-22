// Package core contains shared constants, enumerations, and data types
// used across multiple OCPP 1.6J messages.
//
// This file defines enumerations such as AuthorizationStatus for use in
// messages like Authorize.conf, StartTransaction.conf, etc.
package core

// AuthorizationStatus represents the result of an idTag authorization attempt
// as returned in an Authorize.conf or StartTransaction.conf message.
type AuthorizationStatus string

const (
	// AuthorizationAccepted means the idTag is accepted and ready for use.
	AuthorizationAccepted AuthorizationStatus = "Accepted"

	// AuthorizationBlocked means the idTag is known but has been disabled by the operator.
	AuthorizationBlocked AuthorizationStatus = "Blocked"

	// AuthorizationExpired means the idTag is known but no longer valid (e.g. expired subscription).
	AuthorizationExpired AuthorizationStatus = "Expired"

	// AuthorizationInvalid means the idTag is unknown to the CSMS.
	AuthorizationInvalid AuthorizationStatus = "Invalid"

	// AuthorizationConcurrentTx means the idTag is valid, but another ongoing transaction is already using it.
	AuthorizationConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid checks whether the status is a valid AuthorizationStatus.
func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case AuthorizationAccepted,
		AuthorizationBlocked,
		AuthorizationExpired,
		AuthorizationInvalid,
		AuthorizationConcurrentTx:
		return true
	default:
		return false
	}
}
