package authorizetypes_test

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleConfirmationPayload_Validate_valid() {
	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      mat.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	err := payload.Validate()
	if err != nil {
		fmt.Println("unexpected validation error:", err)

		return
	}

	fmt.Println("Valid payload")
	// Output:
	// Valid payload
}

func ExampleConfirmationPayload_Validate_invalid() {
	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	err := payload.Validate()
	if err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("This should not print")
	// Output:
	// Validation failed: invalid idTagInfo
}

func ExampleConfirmationPayload_Value() {
	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      mat.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	val := payload.Value()
	fmt.Printf("IdTagInfo.Status: %s\n", val.IdTagInfo.Status)

	// Output:
	// IdTagInfo.Status: Accepted
}
