package enums

type RegistrationStatus string

const (
	RegistrationAccepted RegistrationStatus = "Accepted"
	RegistrationPending  RegistrationStatus = "Pending"
	RegistrationRejected RegistrationStatus = "Rejected"
)

func (s RegistrationStatus) IsValid() bool {
	switch s {
	case RegistrationAccepted, RegistrationPending, RegistrationRejected:
		return true
	default:
		return false
	}
}
