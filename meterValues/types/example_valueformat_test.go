package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues/types"
)

// ExampleValueFormat demonstrates using the ValueFormat enumeration constants
// defined in OCPP 1.6.
func ExampleValueFormat() {
	format := types.Raw
	fmt.Println("Format:", format.String())
	fmt.Println("IsValid:", format.IsValid())
	// Output:
	// Format: Raw
	// IsValid: true
}

// ExampleValueFormat_allValues demonstrates all valid ValueFormat values
// as defined in OCPP 1.6.
func ExampleValueFormat_allValues() {
	formats := []types.ValueFormat{
		types.Raw,
		types.SignedData,
	}

	for _, f := range formats {
		fmt.Println(f.String())
	}
	// Output:
	// Raw
	// SignedData
}
