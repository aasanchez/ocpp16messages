package types

// ReservationStatus represents the result of a ReserveNow request.
type ReservationStatus string

// Accepted is the status value for an accepted reservation.
const Accepted = "Accepted"

// Faulted is the status value when connector or charge point is faulted.
const Faulted = "Faulted"

// Occupied is the status value when connector is occupied or reserved.
const Occupied = "Occupied"

// Rejected is the status value for a rejected reservation.
const Rejected = "Rejected"

// Unavailable is the status value when connector is unavailable.
const Unavailable = "Unavailable"

// ReservationStatus enumeration values as defined in OCPP 1.6.
const (
	ReservationStatusAccepted    ReservationStatus = Accepted
	ReservationStatusFaulted     ReservationStatus = Faulted
	ReservationStatusOccupied    ReservationStatus = Occupied
	ReservationStatusRejected    ReservationStatus = Rejected
	ReservationStatusUnavailable ReservationStatus = Unavailable
)

// IsValid checks if the ReservationStatus value is valid per OCPP 1.6.
func (r ReservationStatus) IsValid() bool {
	switch r {
	case ReservationStatusAccepted,
		ReservationStatusFaulted,
		ReservationStatusOccupied,
		ReservationStatusRejected,
		ReservationStatusUnavailable:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReservationStatus.
func (r ReservationStatus) String() string {
	return string(r)
}
