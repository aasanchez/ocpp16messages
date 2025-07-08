package authorizetypes_test

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleRequestPayload_valid() {
	payload := mat.RequestPayload{IdTag: "ABC123XYZ789"}

	err := payload.Validate()
	if err != nil {
		fmt.Println("unexpected error:", err)

		return
	}

	fmt.Println("IdTag is valid:", payload.Value())
	// Output:
	// IdTag is valid: ABC123XYZ789
}

func ExampleRequestPayload_invalid_emptyIdTag() {
	payload := mat.RequestPayload{IdTag: ""}

	err := payload.Validate()
	if err != nil {
		fmt.Println("Validation failed unexpectedly:", err)

		return
	}

	fmt.Println("Validation passed for empty IdTag")
	// Output:
	// Validation passed for empty IdTag
}
