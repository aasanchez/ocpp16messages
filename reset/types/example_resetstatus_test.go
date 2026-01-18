package types_test

import (
	"fmt"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
)

// ExampleResetStatusAccepted demonstrates using the Accepted status constant.
func ExampleResetStatusAccepted() {
	status := rt.ResetStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("Valid:", status.IsValid())
	// Output:
	// Status: Accepted
	// Valid: true
}

// ExampleResetStatusRejected demonstrates using the Rejected status constant.
func ExampleResetStatusRejected() {
	status := rt.ResetStatusRejected
	fmt.Println("Status:", status.String())
	fmt.Println("Valid:", status.IsValid())
	// Output:
	// Status: Rejected
	// Valid: true
}
