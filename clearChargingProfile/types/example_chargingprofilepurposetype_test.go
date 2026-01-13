package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
)

// ExampleChargingProfilePurposeType demonstrates using the
// ChargingProfilePurposeType enumeration constants defined in OCPP 1.6.
func ExampleChargingProfilePurposeType() {
	purpose := types.TxProfile
	fmt.Println("Purpose:", purpose.String())
	fmt.Println("IsValid:", purpose.IsValid())
	// Output:
	// Purpose: TxProfile
	// IsValid: true
}

// ExampleChargingProfilePurposeType_allValues demonstrates all valid
// ChargingProfilePurposeType values as defined in OCPP 1.6.
func ExampleChargingProfilePurposeType_allValues() {
	purposes := []types.ChargingProfilePurposeType{
		types.ChargePointMaxProfile,
		types.TxDefaultProfile,
		types.TxProfile,
	}

	for _, p := range purposes {
		fmt.Println(p.String())
	}
	// Output:
	// ChargePointMaxProfile
	// TxDefaultProfile
	// TxProfile
}

// ExampleChargingProfilePurposeType_IsValid demonstrates validation of
// ChargingProfilePurposeType values.
func ExampleChargingProfilePurposeType_IsValid() {
	valid := types.TxDefaultProfile
	invalid := types.ChargingProfilePurposeType("Invalid")

	fmt.Println("TxDefaultProfile IsValid:", valid.IsValid())
	fmt.Println("Invalid IsValid:", invalid.IsValid())
	// Output:
	// TxDefaultProfile IsValid: true
	// Invalid IsValid: false
}
