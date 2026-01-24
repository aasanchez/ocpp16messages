package types

// Phase represents the phase of the electrical measurement as defined in
// OCPP 1.6.
type Phase string

// Alias for shorter constant names within this package.
type ph = Phase

// Phase enumeration values as defined in OCPP 1.6.
const (
	// PhaseL1 is the measurement on L1.
	PhaseL1 ph = "L1"
	// PhaseL2 is the measurement on L2.
	PhaseL2 ph = "L2"
	// PhaseL3 is the measurement on L3.
	PhaseL3 ph = "L3"
	// PhaseN is the measurement on Neutral.
	PhaseN ph = "N"
	// PhaseL1N is the measurement between L1 and Neutral.
	PhaseL1N ph = "L1-N"
	// PhaseL2N is the measurement between L2 and Neutral.
	PhaseL2N ph = "L2-N"
	// PhaseL3N is the measurement between L3 and Neutral.
	PhaseL3N ph = "L3-N"
	// PhaseL1L2 is the measurement between L1 and L2.
	PhaseL1L2 ph = "L1-L2"
	// PhaseL2L3 is the measurement between L2 and L3.
	PhaseL2L3 ph = "L2-L3"
	// PhaseL3L1 is the measurement between L3 and L1.
	PhaseL3L1 ph = "L3-L1"
)

// IsValid checks if the Phase value is valid per OCPP 1.6.
func (p Phase) IsValid() bool {
	switch p {
	case PhaseL1,
		PhaseL2,
		PhaseL3,
		PhaseN,
		PhaseL1N,
		PhaseL2N,
		PhaseL3N,
		PhaseL1L2,
		PhaseL2L3,
		PhaseL3L1:
		return true
	default:
		return false
	}
}

// String returns the string representation of Phase.
func (p Phase) String() string {
	return string(p)
}
