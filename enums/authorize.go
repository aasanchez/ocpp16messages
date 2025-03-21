package enums

// AuthorizationStatus defines possible outcomes of an authorization attempt.
type AuthorizationStatus string

const (
	AuthorizationAccepted     AuthorizationStatus = "Accepted"
	AuthorizationBlocked      AuthorizationStatus = "Blocked"
	AuthorizationExpired      AuthorizationStatus = "Expired"
	AuthorizationInvalid      AuthorizationStatus = "Invalid"
	AuthorizationConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid returns true if the status is one of the allowed AuthorizationStatus values.
func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case AuthorizationAccepted, AuthorizationBlocked, AuthorizationExpired,
		AuthorizationInvalid, AuthorizationConcurrentTx:
		return true
	default:
		return false
	}
}
