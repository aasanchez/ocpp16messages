package enums

// RegistrationStatus represents the status in BootNotification.conf
type RegistrationStatus string

const (
	RegistrationAccepted RegistrationStatus = "Accepted"
	RegistrationPending  RegistrationStatus = "Pending"
	RegistrationRejected RegistrationStatus = "Rejected"
)

// IsValid checks whether the value is a valid registration status
func (s RegistrationStatus) IsValid() bool {
	switch s {
	case RegistrationAccepted, RegistrationPending, RegistrationRejected:
		return true
	default:
		return false
	}
}
