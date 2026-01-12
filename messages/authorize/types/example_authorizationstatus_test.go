package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

// ExampleAuthorizationStatus demonstrates using the AuthorizationStatus
// enumeration constants defined in OCPP 1.6.
func ExampleAuthorizationStatus() {
	status := types.AuthorizationStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleAuthorizationStatus_allValues demonstrates all valid
// AuthorizationStatus values as defined in OCPP 1.6.
func ExampleAuthorizationStatus_allValues() {
	statuses := []types.AuthorizationStatus{
		types.AuthorizationStatusAccepted,
		types.AuthorizationStatusBlocked,
		types.AuthorizationStatusExpired,
		types.AuthorizationStatusInvalid,
		types.AuthorizationStatusConcurrentTx,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Blocked
	// Expired
	// Invalid
	// ConcurrentTx
}

// ExampleAuthorizationStatus_IsValid demonstrates validation of
// AuthorizationStatus values.
func ExampleAuthorizationStatus_IsValid() {
	valid := types.AuthorizationStatusAccepted
	invalid := types.AuthorizationStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
