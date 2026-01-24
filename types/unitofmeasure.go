package types

// UnitOfMeasure represents the unit of a measurement as defined in OCPP 1.6.
type UnitOfMeasure string

// Alias for shorter constant names within this package.
type uom = UnitOfMeasure

// UnitOfMeasure enumeration values as defined in OCPP 1.6.
const (
	// UnitWh is watt-hours (energy).
	UnitWh uom = "Wh"
	// UnitKWh is kilowatt-hours (energy).
	UnitKWh uom = "kWh"
	// UnitVarh is var-hours (reactive energy).
	UnitVarh uom = "varh"
	// UnitKvarh is kilovar-hours (reactive energy).
	UnitKvarh uom = "kvarh"
	// UnitW is watts (power).
	UnitW uom = "W"
	// UnitKW is kilowatts (power).
	UnitKW uom = "kW"
	// UnitVA is volt-amperes (apparent power).
	UnitVA uom = "VA"
	// UnitKVA is kilovolt-amperes (apparent power).
	UnitKVA uom = "kVA"
	// UnitVar is vars (reactive power).
	UnitVar uom = "var"
	// UnitKvar is kilovars (reactive power).
	UnitKvar uom = "kvar"
	// UnitA is amperes (current).
	UnitA uom = "A"
	// UnitV is volts (voltage).
	UnitV uom = "V"
	// UnitCelsius is degrees Celsius (temperature).
	UnitCelsius uom = "Celsius"
	// UnitFahrenheit is degrees Fahrenheit (temperature).
	UnitFahrenheit uom = "Fahrenheit"
	// UnitK is Kelvin (temperature).
	UnitK uom = "K"
	// UnitPercent is a percentage value.
	UnitPercent uom = "Percent"
)

// IsValid checks if the UnitOfMeasure value is valid per OCPP 1.6.
func (u UnitOfMeasure) IsValid() bool {
	switch u {
	case UnitWh,
		UnitKWh,
		UnitVarh,
		UnitKvarh,
		UnitW,
		UnitKW,
		UnitVA,
		UnitKVA,
		UnitVar,
		UnitKvar,
		UnitA,
		UnitV,
		UnitCelsius,
		UnitFahrenheit,
		UnitK,
		UnitPercent:
		return true
	default:
		return false
	}
}

// String returns the string representation of UnitOfMeasure.
func (u UnitOfMeasure) String() string {
	return string(u)
}
