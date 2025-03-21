// Package enums defines OCPP 1.6 enumerated types with validation helpers.
package enums

// RegistrationStatus defines the valid values for the `status` field
// in a BootNotification.conf message from the Central System.
//
// According to the OCPP 1.6 specification, a charge point can be:
//   - "Accepted"  → The Central System accepts the charge point
//   - "Pending"   → The Central System has not made a decision yet
//   - "Rejected"  → The Central System rejects the charge point
type RegistrationStatus string

const (
	// RegistrationAccepted indicates the Central System has accepted the charge point.
	RegistrationAccepted RegistrationStatus = "Accepted"

	// RegistrationPending indicates the Central System has not made a final decision.
	RegistrationPending RegistrationStatus = "Pending"

	// RegistrationRejected indicates the Central System has rejected the charge point.
	RegistrationRejected RegistrationStatus = "Rejected"
)

// IsValid returns true if the RegistrationStatus value
// matches one of the defined constants above.
//
// This is used to validate that the field conforms to
// the allowed values before sending or processing a message.
func (s RegistrationStatus) IsValid() bool {
	switch s {
	case RegistrationAccepted, RegistrationPending, RegistrationRejected:
		return true
	default:
		return false
	}
}
