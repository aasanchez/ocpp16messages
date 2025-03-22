// Package core contains common enumerations used across multiple OCPP 1.6 message types.
// These enums are aligned with the OCPP 1.6J specification.
package core

// AuthorizationStatus represents the possible values for IdTagInfo.status in Authorize.conf.
type AuthorizationStatus string

// MessageType represents the OCPP 1.6 message types.
type MessageType int

const (
	CALL       MessageType = 2
	CALLRESULT MessageType = 3
	CALLERROR  MessageType = 4
)

const (
	// AuthorizationAccepted indicates that the idTag is valid for charging.
	AuthorizationAccepted AuthorizationStatus = "Accepted"

	// AuthorizationBlocked indicates that the idTag is valid but has been blocked.
	AuthorizationBlocked AuthorizationStatus = "Blocked"

	// AuthorizationExpired indicates that the idTag has expired and is no longer valid.
	AuthorizationExpired AuthorizationStatus = "Expired"

	// AuthorizationInvalid indicates that the idTag is not recognized.
	AuthorizationInvalid AuthorizationStatus = "Invalid"

	// AuthorizationConcurrentTx indicates that the idTag is valid, but charging is not allowed due to an active transaction.
	AuthorizationConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid returns true if the AuthorizationStatus value is a recognized constant.
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
