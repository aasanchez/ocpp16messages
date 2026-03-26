package types_test

import (
	"fmt"

	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
)

// ExampleUpdateTypeFull demonstrates using the Full update type constant.
func ExampleUpdateTypeFull() {
	updateType := slt.UpdateTypeFull
	fmt.Println("Type:", updateType.String())
	fmt.Println("Valid:", updateType.IsValid())
	// Output:
	// Type: Full
	// Valid: true
}

// ExampleUpdateTypeDifferential demonstrates using the Differential update
// type constant.
func ExampleUpdateTypeDifferential() {
	updateType := slt.UpdateTypeDifferential
	fmt.Println("Type:", updateType.String())
	fmt.Println("Valid:", updateType.IsValid())
	// Output:
	// Type: Differential
	// Valid: true
}
