package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/bootNotification/types"
)

// ExampleRegistrationStatus demonstrates using the RegistrationStatus
// enumeration constants defined in OCPP 1.6.
func ExampleRegistrationStatus() {
	status := types.RegistrationStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleRegistrationStatus_allValues demonstrates all valid
// RegistrationStatus values as defined in OCPP 1.6.
func ExampleRegistrationStatus_allValues() {
	statuses := []types.RegistrationStatus{
		types.RegistrationStatusAccepted,
		types.RegistrationStatusPending,
		types.RegistrationStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Pending
	// Rejected
}

// ExampleRegistrationStatus_IsValid demonstrates validation of
// RegistrationStatus values.
func ExampleRegistrationStatus_IsValid() {
	valid := types.RegistrationStatusAccepted
	invalid := types.RegistrationStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
