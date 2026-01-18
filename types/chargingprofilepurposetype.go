package types

// ChargingProfilePurposeType represents the purpose of a charging profile
// as defined in OCPP 1.6.
type ChargingProfilePurposeType string

// Alias for shorter constant names within this package.
type cppt = ChargingProfilePurposeType

// ChargingProfilePurposeType enumeration values as defined in OCPP 1.6.
const (
	// ChargePointMaxProfile limits overall charge point power.
	ChargePointMaxProfile cppt = "ChargePointMaxProfile"
	// TxDefaultProfile is the default profile for transactions.
	TxDefaultProfile cppt = "TxDefaultProfile"
	// TxProfile is for a specific transaction.
	TxProfile cppt = "TxProfile"
)

// IsValid checks if the ChargingProfilePurposeType value is valid per OCPP 1.6.
func (c ChargingProfilePurposeType) IsValid() bool {
	switch c {
	case ChargePointMaxProfile,
		TxDefaultProfile,
		TxProfile:
		return true
	default:
		return false
	}
}

// String returns the string representation of ChargingProfilePurposeType.
func (c ChargingProfilePurposeType) String() string {
	return string(c)
}
