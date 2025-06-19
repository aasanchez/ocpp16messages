package authorizetypes_test

import (
	"fmt"

	types "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleConfirmationPayload_Validate_valid() {
	payload := types.ConfirmationPayload{
		IdTagInfo: types.IdTagInfoPayload{
			Status:      types.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	if err := payload.Validate(); err != nil {
		fmt.Println("unexpected validation error:", err)
		return
	}

	fmt.Println("Valid payload")
	// Output:
	// Valid payload
}

func ExampleConfirmationPayload_Validate_invalid() {
	payload := types.ConfirmationPayload{
		IdTagInfo: types.IdTagInfoPayload{
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
	// Validation failed: confirmation payload: invalid idTagInfo
}

func ExampleConfirmationPayload_Value() {
	payload := types.ConfirmationPayload{
		IdTagInfo: types.IdTagInfoPayload{
			Status:      types.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	val := payload.Value()
	fmt.Printf("IdTagInfo.Status: %s\n", val.IdTagInfo.Status)

	// Output:
	// IdTagInfo.Status: Accepted
}
