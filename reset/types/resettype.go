package types

// Hard is the reset type value for a hard reset.
const Hard = "Hard"

// Soft is the reset type value for a soft reset.
const Soft = "Soft"

// ResetType represents the type of reset to perform on a Charge Point.
type ResetType string

// ResetType enumeration values as defined in OCPP 1.6.
const (
	ResetTypeHard ResetType = Hard
	ResetTypeSoft ResetType = Soft
)

// IsValid checks if the ResetType value is valid per OCPP 1.6.
func (r ResetType) IsValid() bool {
	switch r {
	case ResetTypeHard, ResetTypeSoft:
		return true
	default:
		return false
	}
}

// String returns the string representation of ResetType.
func (r ResetType) String() string {
	return string(r)
}
