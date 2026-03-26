package types

// UnitOfMeasure represents the unit of a measurement as defined in OCPP 1.6.
type UnitOfMeasure string

// Alias for shorter constant names within this package.
type uom = UnitOfMeasure

// UnitOfMeasure enumeration values as defined in OCPP 1.6.
//
//nolint:varnamelen // OCPP 1.6 standard naming convention
const (
	// Wh is watt-hours (energy).
	Wh uom = "Wh"
	// KWh is kilowatt-hours (energy).
	KWh uom = "kWh"
	// Varh is var-hours (reactive energy).
	Varh uom = "varh"
	// Kvarh is kilovar-hours (reactive energy).
	Kvarh uom = "kvarh"
	// W is watts (power).
	W uom = "W"
	// KW is kilowatts (power).
	KW uom = "kW"
	// VA is volt-amperes (apparent power).
	VA uom = "VA"
	// KVA is kilovolt-amperes (apparent power).
	KVA uom = "kVA"
	// Var is vars (reactive power).
	Var uom = "var"
	// Kvar is kilovars (reactive power).
	Kvar uom = "kvar"
	// A is amperes (current).
	A uom = "A"
	// V is volts (voltage).
	V uom = "V"
	// Celsius is degrees Celsius (temperature).
	Celsius uom = "Celsius"
	// Fahrenheit is degrees Fahrenheit (temperature).
	Fahrenheit uom = "Fahrenheit"
	// K is Kelvin (temperature).
	K uom = "K"
	// Percent is a percentage value.
	Percent uom = "Percent"
)

// IsValid checks if the UnitOfMeasure value is valid per OCPP 1.6.
func (u UnitOfMeasure) IsValid() bool {
	switch u {
	case Wh,
		KWh,
		Varh,
		Kvarh,
		W,
		KW,
		VA,
		KVA,
		Var,
		Kvar,
		A,
		V,
		Celsius,
		Fahrenheit,
		K,
		Percent:
		return true
	default:
		return false
	}
}

// String returns the string representation of UnitOfMeasure.
func (u UnitOfMeasure) String() string {
	return string(u)
}
