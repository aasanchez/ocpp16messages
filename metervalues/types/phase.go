package types

// Phase represents the phase of the electrical measurement as defined in
// OCPP 1.6.
type Phase string

// Alias for shorter constant names within this package.
type ph = Phase

// Phase enumeration values as defined in OCPP 1.6.
//
//nolint:varnamelen // OCPP 1.6 standard naming convention
const (
	// L1 is the measurement on L1.
	L1 ph = "L1"
	// L2 is the measurement on L2.
	L2 ph = "L2"
	// L3 is the measurement on L3.
	L3 ph = "L3"
	// N is the measurement on Neutral.
	N ph = "N"
	// L1N is the measurement between L1 and Neutral.
	L1N ph = "L1-N"
	// L2N is the measurement between L2 and Neutral.
	L2N ph = "L2-N"
	// L3N is the measurement between L3 and Neutral.
	L3N ph = "L3-N"
	// L1L2 is the measurement between L1 and L2.
	L1L2 ph = "L1-L2"
	// L2L3 is the measurement between L2 and L3.
	L2L3 ph = "L2-L3"
	// L3L1 is the measurement between L3 and L1.
	L3L1 ph = "L3-L1"
)

// IsValid checks if the Phase value is valid per OCPP 1.6.
func (p Phase) IsValid() bool {
	switch p {
	case L1,
		L2,
		L3,
		N,
		L1N,
		L2N,
		L3N,
		L1L2,
		L2L3,
		L3L1:
		return true
	default:
		return false
	}
}

// String returns the string representation of Phase.
func (p Phase) String() string {
	return string(p)
}
