package types

// ValueFormat represents the format of a sampled value as defined in OCPP 1.6.
type ValueFormat string

// Alias for shorter constant names within this package.
type vf = ValueFormat

// ValueFormat enumeration values as defined in OCPP 1.6.
const (
	// Raw indicates the data is a raw value (default).
	Raw vf = "Raw"
	// SignedData indicates the data contains a digitally signed binary block.
	// This format is EXPERIMENTAL and may be deprecated in future versions.
	SignedData vf = "SignedData"
)

// IsValid checks if the ValueFormat value is valid per OCPP 1.6.
func (v ValueFormat) IsValid() bool {
	switch v {
	case Raw, SignedData:
		return true
	default:
		return false
	}
}

// String returns the string representation of ValueFormat.
func (v ValueFormat) String() string {
	return string(v)
}
