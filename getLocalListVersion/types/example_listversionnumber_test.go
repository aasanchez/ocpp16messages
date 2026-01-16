package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
)

// ExampleNewListVersionNumber demonstrates creating a ListVersionNumber
// with a positive version number.
func ExampleNewListVersionNumber() {
	listVersion, err := types.NewListVersionNumber("5")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Value:", listVersion.Value())
	fmt.Println("String:", listVersion.String())
	// Output:
	// Value: 5
	// String: 5
}

// ExampleNewListVersionNumber_unsupported demonstrates creating a
// ListVersionNumber indicating the Charge Point does not support
// Local Authorization Lists.
func ExampleNewListVersionNumber_unsupported() {
	listVersion, err := types.NewListVersionNumber("-1")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Value:", listVersion.Value())
	fmt.Println("IsUnsupported:", listVersion.IsUnsupported())
	// Output:
	// Value: -1
	// IsUnsupported: true
}

// ExampleNewListVersionNumber_emptyList demonstrates creating a
// ListVersionNumber indicating the local authorization list is empty.
func ExampleNewListVersionNumber_emptyList() {
	listVersion, err := types.NewListVersionNumber("0")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Value:", listVersion.Value())
	fmt.Println("IsEmpty:", listVersion.IsEmpty())
	// Output:
	// Value: 0
	// IsEmpty: true
}

// ExampleNewListVersionNumber_invalid demonstrates the error returned
// when an invalid value is provided.
func ExampleNewListVersionNumber_invalid() {
	_, err := types.NewListVersionNumber("invalid")
	if err != nil {
		fmt.Println("Error: invalid value")
	}
	// Output:
	// Error: invalid value
}
