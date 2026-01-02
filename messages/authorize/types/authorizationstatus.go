package types

// AuthorizationStatus represents the status in a response to an Authorize
// request.
type AuthorizationStatus string

// AuthorizationStatus enumeration values as defined in OCPP 1.6.
const (
	AuthorizationStatusAccepted     AuthorizationStatus = "Accepted"
	AuthorizationStatusBlocked      AuthorizationStatus = "Blocked"
	AuthorizationStatusExpired      AuthorizationStatus = "Expired"
	AuthorizationStatusInvalid      AuthorizationStatus = "Invalid"
	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid checks if the AuthorizationStatus value is valid per OCPP 1.6.
func (a AuthorizationStatus) IsValid() bool {
	switch a {
	case AuthorizationStatusAccepted,
		AuthorizationStatusBlocked,
		AuthorizationStatusExpired,
		AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx:
		return true
	default:
		return false
	}
}

// String returns the string representation of AuthorizationStatus.
func (a AuthorizationStatus) String() string {
	return string(a)
}
