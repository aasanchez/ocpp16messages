package types

// Full is the update type value for a full list replacement.
const Full = "Full"

// Differential is the update type value for a differential update.
const Differential = "Differential"

// UpdateType represents the type of update to apply to the local
// authorization list.
type UpdateType string

// UpdateType enumeration values as defined in OCPP 1.6.
const (
	UpdateTypeFull         UpdateType = Full
	UpdateTypeDifferential UpdateType = Differential
)

// IsValid checks if the UpdateType value is valid per OCPP 1.6.
func (u UpdateType) IsValid() bool {
	switch u {
	case UpdateTypeFull, UpdateTypeDifferential:
		return true
	default:
		return false
	}
}

// String returns the string representation of UpdateType.
func (u UpdateType) String() string {
	return string(u)
}
