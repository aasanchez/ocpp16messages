package types

// AvailabilityType represents the requested availability change for a Charge
// Point or connector as defined in OCPP 1.6.
type AvailabilityType string

// AvailabilityType enumeration values as defined in OCPP 1.6.
const (
	AvailabilityTypeInoperative AvailabilityType = "Inoperative"
	AvailabilityTypeOperative   AvailabilityType = "Operative"
)

// IsValid checks if the AvailabilityType value is valid per OCPP 1.6.
func (a AvailabilityType) IsValid() bool {
	switch a {
	case AvailabilityTypeInoperative,
		AvailabilityTypeOperative:
		return true
	default:
		return false
	}
}

// String returns the string representation of AvailabilityType.
func (a AvailabilityType) String() string {
	return string(a)
}
