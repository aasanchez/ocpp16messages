package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/remotestarttransaction/types"
)

// ExampleRemoteStartTransactionStatus demonstrates using the
// RemoteStartTransactionStatus enumeration constants defined in OCPP 1.6.
func ExampleRemoteStartTransactionStatus() {
	status := types.RemoteStartTransactionStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleRemoteStartTransactionStatus_allValues demonstrates all valid
// RemoteStartTransactionStatus values as defined in OCPP 1.6.
func ExampleRemoteStartTransactionStatus_allValues() {
	statuses := []types.RemoteStartTransactionStatus{
		types.RemoteStartTransactionStatusAccepted,
		types.RemoteStartTransactionStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Rejected
}

// ExampleRemoteStartTransactionStatus_IsValid demonstrates validation of
// RemoteStartTransactionStatus values.
func ExampleRemoteStartTransactionStatus_IsValid() {
	valid := types.RemoteStartTransactionStatusAccepted
	invalid := types.RemoteStartTransactionStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
