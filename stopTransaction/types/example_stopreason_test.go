package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/stoptransaction/types"
)

const (
	exampleReasonLabel  = "Reason:"
	exampleIsValidLabel = "IsValid:"
)

// ExampleReason demonstrates creating and validating a Reason value.
func ExampleReason() {
	reason := types.ReasonLocal

	fmt.Println(exampleReasonLabel, reason.String())
	fmt.Println(exampleIsValidLabel, reason.IsValid())
	// Output:
	// Reason: Local
	// IsValid: true
}

// ExampleReason_remote demonstrates the Remote reason for remote stop commands.
func ExampleReason_remote() {
	reason := types.ReasonRemote

	fmt.Println(exampleReasonLabel, reason.String())
	fmt.Println(exampleIsValidLabel, reason.IsValid())
	// Output:
	// Reason: Remote
	// IsValid: true
}

// ExampleReason_invalid demonstrates that invalid reason values are rejected.
func ExampleReason_invalid() {
	reason := types.Reason("InvalidReason")

	fmt.Println(exampleReasonLabel, reason.String())
	fmt.Println(exampleIsValidLabel, reason.IsValid())
	// Output:
	// Reason: InvalidReason
	// IsValid: false
}
