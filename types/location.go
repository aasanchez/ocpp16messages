package types

// Location represents the location of measurement as defined in OCPP 1.6.
type Location string

// Alias for shorter constant names within this package.
type loc = Location

// Location enumeration values as defined in OCPP 1.6.
const (
	// LocationBody is a measurement inside the charge point body.
	LocationBody loc = "Body"
	// LocationCable is a measurement on the cable between charge point and EV.
	LocationCable loc = "Cable"
	// LocationEV is a measurement inside the EV.
	LocationEV loc = "EV"
	// LocationInlet is a measurement at the network/grid inlet connection.
	LocationInlet loc = "Inlet"
	// LocationOutlet is a measurement at the connector/outlet to EV.
	LocationOutlet loc = "Outlet"
)

// IsValid checks if the Location value is valid per OCPP 1.6.
func (l Location) IsValid() bool {
	switch l {
	case LocationBody, LocationCable, LocationEV, LocationInlet, LocationOutlet:
		return true
	default:
		return false
	}
}

// String returns the string representation of Location.
func (l Location) String() string {
	return string(l)
}
