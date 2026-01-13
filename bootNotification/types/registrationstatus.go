package types

// RegistrationStatus represents the result of a BootNotification request.
type RegistrationStatus string

// RegistrationStatus enumeration values as defined in OCPP 1.6.
const (
	RegistrationStatusAccepted RegistrationStatus = "Accepted"
	RegistrationStatusPending  RegistrationStatus = "Pending"
	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

// IsValid checks if the RegistrationStatus value is valid per OCPP 1.6.
func (r RegistrationStatus) IsValid() bool {
	switch r {
	case RegistrationStatusAccepted,
		RegistrationStatusPending,
		RegistrationStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RegistrationStatus.
func (r RegistrationStatus) String() string {
	return string(r)
}
