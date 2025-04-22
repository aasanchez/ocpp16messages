package authorizetypes

type AuthorizationStatus string

const (
	Accepted     AuthorizationStatus = "Accepted"
	Blocked      AuthorizationStatus = "Blocked"
	Expired      AuthorizationStatus = "Expired"
	Invalid      AuthorizationStatus = "Invalid"
	ConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case Accepted, Blocked, Expired, Invalid, ConcurrentTx:
		return true
	default:
		return false
	}
}
