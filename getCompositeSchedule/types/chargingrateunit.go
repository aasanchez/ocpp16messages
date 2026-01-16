package types

// ChargingRateUnit represents the unit of measure for charging rate
// as defined in OCPP 1.6.
type ChargingRateUnit string

// ChargingRateUnit enumeration values as defined in OCPP 1.6.
const (
	ChargingRateUnitWatts   ChargingRateUnit = "W"
	ChargingRateUnitAmperes ChargingRateUnit = "A"
)

// IsValid checks if the ChargingRateUnit value is valid per OCPP 1.6.
func (c ChargingRateUnit) IsValid() bool {
	switch c {
	case ChargingRateUnitWatts,
		ChargingRateUnitAmperes:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingRateUnit.
func (c ChargingRateUnit) String() string {
	return string(c)
}
