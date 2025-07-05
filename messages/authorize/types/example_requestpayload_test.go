package authorizetypes_test

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleRequestPayload_valid() {
	payload := authorizetypes.RequestPayload{IdTag: "ABC123XYZ789"}

	if err := payload.Validate(); err != nil {
		fmt.Println("unexpected error:", err)

		return
	}

	fmt.Println("IdTag is valid:", payload.Value())
	// Output:
	// IdTag is valid: ABC123XYZ789
}

func ExampleRequestPayload_invalid_emptyIdTag() {
	payload := authorizetypes.RequestPayload{IdTag: ""}

	if err := payload.Validate(); err != nil {
		fmt.Println("Validation failed unexpectedly:", err)

		return
	}

	fmt.Println("Validation passed for empty IdTag")
	// Output:
	// Validation passed for empty IdTag
}
