package types

// ChargingProfileKindType represents the kind of charging profile
// as defined in OCPP 1.6.
type ChargingProfileKindType string

// ChargingProfileKindType enumeration values as defined in OCPP 1.6.
const (
	// ChargingProfileKindAbsolute indicates an absolute schedule with fixed
	// time slots.
	ChargingProfileKindAbsolute ChargingProfileKindType = "Absolute"
	// ChargingProfileKindRecurring indicates a recurring schedule that
	// repeats based on recurrencyKind.
	ChargingProfileKindRecurring ChargingProfileKindType = "Recurring"
	// ChargingProfileKindRelative indicates a schedule relative to the start
	// of the transaction.
	ChargingProfileKindRelative ChargingProfileKindType = "Relative"
)

// IsValid checks if the ChargingProfileKindType value is valid per OCPP 1.6.
func (c ChargingProfileKindType) IsValid() bool {
	switch c {
	case ChargingProfileKindAbsolute,
		ChargingProfileKindRecurring,
		ChargingProfileKindRelative:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingProfileKindType.
func (c ChargingProfileKindType) String() string {
	return string(c)
}
