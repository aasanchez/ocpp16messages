package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeConfiguration/types"
)

// ExampleConfigurationStatus_IsValid demonstrates validating a
// ConfigurationStatus.
func ExampleConfigurationStatus_IsValid() {
	accepted := types.ConfigurationStatus("Accepted")
	fmt.Println("Accepted valid:", accepted.IsValid())

	rejected := types.ConfigurationStatus("Rejected")
	fmt.Println("Rejected valid:", rejected.IsValid())

	rebootRequired := types.ConfigurationStatus("RebootRequired")
	fmt.Println("RebootRequired valid:", rebootRequired.IsValid())

	notSupported := types.ConfigurationStatus("NotSupported")
	fmt.Println("NotSupported valid:", notSupported.IsValid())

	invalid := types.ConfigurationStatus("Unknown")
	fmt.Println("Unknown valid:", invalid.IsValid())
	// Output:
	// Accepted valid: true
	// Rejected valid: true
	// RebootRequired valid: true
	// NotSupported valid: true
	// Unknown valid: false
}

// ExampleConfigurationStatus_String demonstrates getting the string
// representation of a ConfigurationStatus.
func ExampleConfigurationStatus_String() {
	rebootRequired := types.ConfigurationStatusRebootRequired
	fmt.Println(rebootRequired.String())
	// Output:
	// RebootRequired
}
