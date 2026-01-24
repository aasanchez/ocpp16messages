package types

// ValueFormat represents the format of a sampled value as defined in OCPP 1.6.
type ValueFormat string

// Alias for shorter constant names within this package.
type vf = ValueFormat

// ValueFormat enumeration values as defined in OCPP 1.6.
const (
	// ValueFormatRaw indicates the data is a raw value (default).
	ValueFormatRaw vf = "Raw"
	// ValueFormatSignedData indicates the data contains a digitally signed
	// binary block. This format is EXPERIMENTAL and may be deprecated in
	// future versions.
	ValueFormatSignedData vf = "SignedData"
)

// IsValid checks if the ValueFormat value is valid per OCPP 1.6.
func (v ValueFormat) IsValid() bool {
	switch v {
	case ValueFormatRaw, ValueFormatSignedData:
		return true
	default:
		return false
	}
}

// String returns the string representation of ValueFormat.
func (v ValueFormat) String() string {
	return string(v)
}
