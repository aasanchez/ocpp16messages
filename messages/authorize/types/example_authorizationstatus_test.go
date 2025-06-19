package authorizetypes_test

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleAuthorizationStatus_valid() {
	status, err := authorizetypes.AuthorizationStatus("Accepted")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("AuthorizationStatus:", status.Value())

	// Output:
	// AuthorizationStatus: Accepted
}

func ExampleAuthorizationStatus_invalid() {
	_, err := authorizetypes.AuthorizationStatus("NotARealStatus")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output:
	// Error: invalid authorization status: "NotARealStatus"
}
