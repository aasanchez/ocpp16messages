package types_test

import (
	"fmt"

	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
)

// ExampleUpdateStatusAccepted demonstrates using the Accepted status constant.
func ExampleUpdateStatusAccepted() {
	status := slt.UpdateStatusAccepted
	fmt.Printf("Status: %s\n", status.String())
	fmt.Printf("Valid: %t\n", status.IsValid())
	// Output:
	// Status: Accepted
	// Valid: true
}

// ExampleUpdateStatusFailed demonstrates using the Failed status constant.
func ExampleUpdateStatusFailed() {
	status := slt.UpdateStatusFailed
	fmt.Printf("Status: %s\n", status.String())
	fmt.Printf("Valid: %t\n", status.IsValid())
	// Output:
	// Status: Failed
	// Valid: true
}

// ExampleUpdateStatusNotSupported demonstrates using the NotSupported status
// constant.
func ExampleUpdateStatusNotSupported() {
	status := slt.UpdateStatusNotSupported
	fmt.Printf("Status: %s\n", status.String())
	fmt.Printf("Valid: %t\n", status.IsValid())
	// Output:
	// Status: NotSupported
	// Valid: true
}

// ExampleUpdateStatusVersionMismatch demonstrates using the VersionMismatch
// status constant.
func ExampleUpdateStatusVersionMismatch() {
	status := slt.UpdateStatusVersionMismatch
	fmt.Printf("Status: %s\n", status.String())
	fmt.Printf("Valid: %t\n", status.IsValid())
	// Output:
	// Status: VersionMismatch
	// Valid: true
}
