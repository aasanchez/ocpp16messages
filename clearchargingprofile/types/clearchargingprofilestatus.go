package types

// ClearChargingProfileStatus represents the result of a ClearChargingProfile
// request as defined in OCPP 1.6.
type ClearChargingProfileStatus string

// ClearChargingProfileStatus enumeration values as defined in OCPP 1.6.
const (
	ClearChargingProfileStatusAccepted ClearChargingProfileStatus = "Accepted"
	ClearChargingProfileStatusUnknown  ClearChargingProfileStatus = "Unknown"
)

// IsValid checks if the ClearChargingProfileStatus value is valid per OCPP 1.6.
func (c ClearChargingProfileStatus) IsValid() bool {
	switch c {
	case ClearChargingProfileStatusAccepted,
		ClearChargingProfileStatusUnknown:
		return true
	default:
		return false
	}
}

// String returns the string representation of ClearChargingProfileStatus.
func (c ClearChargingProfileStatus) String() string {
	return string(c)
}
