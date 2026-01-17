package types

// Location represents the location of measurement as defined in OCPP 1.6.
type Location string

// Alias for shorter constant names within this package.
type loc = Location

// Location enumeration values as defined in OCPP 1.6.
//
//nolint:varnamelen // OCPP 1.6 standard naming convention
const (
	// Body is a measurement inside the charge point body.
	Body loc = "Body"
	// Cable is a measurement on the cable between charge point and EV.
	Cable loc = "Cable"
	// EV is a measurement inside the EV.
	EV loc = "EV"
	// Inlet is a measurement at the network/grid inlet connection.
	Inlet loc = "Inlet"
	// Outlet is a measurement at the connector/outlet to EV.
	Outlet loc = "Outlet"
)

// IsValid checks if the Location value is valid per OCPP 1.6.
func (l Location) IsValid() bool {
	switch l {
	case Body, Cable, EV, Inlet, Outlet:
		return true
	default:
		return false
	}
}

// String returns the string representation of Location.
func (l Location) String() string {
	return string(l)
}
