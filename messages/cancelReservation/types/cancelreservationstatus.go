package types

// CancelReservationStatus represents the result of a CancelReservation request.
type CancelReservationStatus string

// CancelReservationStatus enumeration values as defined in OCPP 1.6.
const (
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"
	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

// IsValid checks if the CancelReservationStatus value is valid per OCPP 1.6.
func (c CancelReservationStatus) IsValid() bool {
	switch c {
	case CancelReservationStatusAccepted,
		CancelReservationStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of CancelReservationStatus.
func (c CancelReservationStatus) String() string {
	return string(c)
}
