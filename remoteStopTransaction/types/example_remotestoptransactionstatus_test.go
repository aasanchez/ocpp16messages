package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/remotestoptransaction/types"
)

// ExampleRemoteStopTransactionStatus demonstrates using the
// RemoteStopTransactionStatus enumeration constants defined in OCPP 1.6.
func ExampleRemoteStopTransactionStatus() {
	status := types.RemoteStopTransactionStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleRemoteStopTransactionStatus_allValues demonstrates all valid
// RemoteStopTransactionStatus values as defined in OCPP 1.6.
func ExampleRemoteStopTransactionStatus_allValues() {
	statuses := []types.RemoteStopTransactionStatus{
		types.RemoteStopTransactionStatusAccepted,
		types.RemoteStopTransactionStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Rejected
}

// ExampleRemoteStopTransactionStatus_IsValid demonstrates validation of
// RemoteStopTransactionStatus values.
func ExampleRemoteStopTransactionStatus_IsValid() {
	valid := types.RemoteStopTransactionStatusAccepted
	invalid := types.RemoteStopTransactionStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
