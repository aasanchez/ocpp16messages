// Package core provides the core types and validation infrastructure for OCPP 1.6J messages.
package core

// MessageType represents the type of OCPP message.
type MessageType int

const (
	// Call represents a request message [2, "id", "action", {...}]
	Call MessageType = 2
	// CallResult represents a response message [3, "id", {...}]
	CallResult MessageType = 3
	// CallError represents an error message [4, "id", "errorCode", "errorDescription", {...}]
	CallError MessageType = 4
)

// AuthorizationStatus represents the possible authorization status values.
type AuthorizationStatus string

const (
	// AuthorizationAccepted indicates the authorization was accepted
	AuthorizationAccepted AuthorizationStatus = "Accepted"
	// AuthorizationBlocked indicates the authorization was blocked
	AuthorizationBlocked AuthorizationStatus = "Blocked"
	// AuthorizationExpired indicates the authorization has expired
	AuthorizationExpired AuthorizationStatus = "Expired"
	// AuthorizationInvalid indicates the authorization was invalid
	AuthorizationInvalid AuthorizationStatus = "Invalid"
	// AuthorizationConcurrentTx indicates concurrent transaction
	AuthorizationConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid checks if the AuthorizationStatus is a valid value as per the OCPP specification.
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
