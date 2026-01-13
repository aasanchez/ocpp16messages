package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/clearCache/types"
)

// ExampleClearCacheStatus demonstrates using the ClearCacheStatus
// enumeration constants defined in OCPP 1.6.
func ExampleClearCacheStatus() {
	status := types.ClearCacheStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleClearCacheStatus_allValues demonstrates all valid
// ClearCacheStatus values as defined in OCPP 1.6.
func ExampleClearCacheStatus_allValues() {
	statuses := []types.ClearCacheStatus{
		types.ClearCacheStatusAccepted,
		types.ClearCacheStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Rejected
}

// ExampleClearCacheStatus_IsValid demonstrates validation of
// ClearCacheStatus values.
func ExampleClearCacheStatus_IsValid() {
	valid := types.ClearCacheStatusAccepted
	invalid := types.ClearCacheStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
