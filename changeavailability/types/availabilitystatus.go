package types

// AvailabilityStatus represents the result of a ChangeAvailability request
// as defined in OCPP 1.6.
type AvailabilityStatus string

// AvailabilityStatus enumeration values as defined in OCPP 1.6.
const (
	AvailabilityStatusAccepted  AvailabilityStatus = "Accepted"
	AvailabilityStatusRejected  AvailabilityStatus = "Rejected"
	AvailabilityStatusScheduled AvailabilityStatus = "Scheduled"
)

// IsValid checks if the AvailabilityStatus value is valid per OCPP 1.6.
func (a AvailabilityStatus) IsValid() bool {
	switch a {
	case AvailabilityStatusAccepted,
		AvailabilityStatusRejected,
		AvailabilityStatusScheduled:
		return true
	default:
		return false
	}
}

// String returns the string representation of AvailabilityStatus.
func (a AvailabilityStatus) String() string {
	return string(a)
}
