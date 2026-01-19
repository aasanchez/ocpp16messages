package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/statusNotification/types"
)

const labelStatus = "Status:"

// ExampleChargePointStatus demonstrates using ChargePointStatus values.
func ExampleChargePointStatus() {
	status := types.ChargePointStatusAvailable

	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Available
	// IsValid: true
}

// ExampleChargePointStatus_charging demonstrates a Charging status.
func ExampleChargePointStatus_charging() {
	status := types.ChargePointStatusCharging

	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Charging
	// IsValid: true
}

// ExampleChargePointStatus_invalid demonstrates an invalid status.
func ExampleChargePointStatus_invalid() {
	status := types.ChargePointStatus("InvalidStatus")

	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: InvalidStatus
	// IsValid: false
}
