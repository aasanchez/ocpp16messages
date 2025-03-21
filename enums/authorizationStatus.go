package enums

type AuthorizationStatus string

const (
	AuthorizationAccepted     AuthorizationStatus = "Accepted"
	AuthorizationBlocked      AuthorizationStatus = "Blocked"
	AuthorizationExpired      AuthorizationStatus = "Expired"
	AuthorizationInvalid      AuthorizationStatus = "Invalid"
	AuthorizationConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case AuthorizationAccepted, AuthorizationBlocked, AuthorizationExpired, AuthorizationInvalid, AuthorizationConcurrentTx:
		return true
	default:
		return false
	}
}
