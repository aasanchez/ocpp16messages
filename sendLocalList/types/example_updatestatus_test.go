package types_test

import (
	"fmt"

	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
)

const (
	fmtStatus = "Status: %s\n"
	fmtValid  = "Valid: %t\n"
)

// ExampleUpdateStatusAccepted demonstrates using the Accepted status constant.
func ExampleUpdateStatusAccepted() {
	status := slt.UpdateStatusAccepted
	fmt.Printf(fmtStatus, status.String())
	fmt.Printf(fmtValid, status.IsValid())
	// Output:
	// Status: Accepted
	// Valid: true
}

// ExampleUpdateStatusFailed demonstrates using the Failed status constant.
func ExampleUpdateStatusFailed() {
	status := slt.UpdateStatusFailed
	fmt.Printf(fmtStatus, status.String())
	fmt.Printf(fmtValid, status.IsValid())
	// Output:
	// Status: Failed
	// Valid: true
}

// ExampleUpdateStatusNotSupported demonstrates using the NotSupported status
// constant.
func ExampleUpdateStatusNotSupported() {
	status := slt.UpdateStatusNotSupported
	fmt.Printf(fmtStatus, status.String())
	fmt.Printf(fmtValid, status.IsValid())
	// Output:
	// Status: NotSupported
	// Valid: true
}

// ExampleUpdateStatusVersionMismatch demonstrates using the VersionMismatch
// status constant.
func ExampleUpdateStatusVersionMismatch() {
	status := slt.UpdateStatusVersionMismatch
	fmt.Printf(fmtStatus, status.String())
	fmt.Printf(fmtValid, status.IsValid())
	// Output:
	// Status: VersionMismatch
	// Valid: true
}
