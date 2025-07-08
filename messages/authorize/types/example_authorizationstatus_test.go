package authorizetypes_test

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleAuthorizationStatus_valid() {
	status, err := mat.AuthorizationStatus("Accepted")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("AuthorizationStatus:", status.Value())

	// Output:
	// AuthorizationStatus: Accepted
}

func ExampleAuthorizationStatus_invalid() {
	_, err := mat.AuthorizationStatus("NotARealStatus")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	// Output:
	// Error: invalid authorization status: "NotARealStatus"
}
