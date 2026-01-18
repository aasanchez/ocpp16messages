package types_test

import (
	"fmt"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
)

// ExampleResetTypeHard demonstrates using the Hard reset type constant.
func ExampleResetTypeHard() {
	resetType := rt.ResetTypeHard
	fmt.Println("Type:", resetType.String())
	fmt.Println("Valid:", resetType.IsValid())
	// Output:
	// Type: Hard
	// Valid: true
}

// ExampleResetTypeSoft demonstrates using the Soft reset type constant.
func ExampleResetTypeSoft() {
	resetType := rt.ResetTypeSoft
	fmt.Println("Type:", resetType.String())
	fmt.Println("Valid:", resetType.IsValid())
	// Output:
	// Type: Soft
	// Valid: true
}
