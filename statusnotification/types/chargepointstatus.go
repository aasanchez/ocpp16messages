package types

// ChargePointStatus represents the status of a connector or the entire Charge
// Point as defined in OCPP 1.6.
type ChargePointStatus string

// ChargePointStatus enumeration values as defined in OCPP 1.6.
const (
	ChargePointStatusAvailable     ChargePointStatus = "Available"
	ChargePointStatusPreparing     ChargePointStatus = "Preparing"
	ChargePointStatusCharging      ChargePointStatus = "Charging"
	ChargePointStatusSuspendedEV   ChargePointStatus = "SuspendedEV"
	ChargePointStatusSuspendedEVSE ChargePointStatus = "SuspendedEVSE"
	ChargePointStatusFinishing     ChargePointStatus = "Finishing"
	ChargePointStatusReserved      ChargePointStatus = "Reserved"
	ChargePointStatusUnavailable   ChargePointStatus = "Unavailable"
	ChargePointStatusFaulted       ChargePointStatus = "Faulted"
)

// IsValid checks if the ChargePointStatus value is valid per OCPP 1.6.
func (c ChargePointStatus) IsValid() bool {
	switch c {
	case ChargePointStatusAvailable,
		ChargePointStatusPreparing,
		ChargePointStatusCharging,
		ChargePointStatusSuspendedEV,
		ChargePointStatusSuspendedEVSE,
		ChargePointStatusFinishing,
		ChargePointStatusReserved,
		ChargePointStatusUnavailable,
		ChargePointStatusFaulted:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargePointStatus.
func (c ChargePointStatus) String() string {
	return string(c)
}
