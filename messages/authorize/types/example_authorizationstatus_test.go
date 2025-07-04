package authorizetypes_test

import (
	"fmt"

	at "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleAuthorizationStatus_valid() {
	status, err := at.SetAuthorizationStatus("Accepted")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("AuthorizationStatus:", status.Value())

	// Output:
	// AuthorizationStatus: Accepted
}

func ExampleAuthorizationStatus_invalid() {
	_, err := at.SetAuthorizationStatus("NotARealStatus")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	// Output:
	// Error: invalid authorization status: "NotARealStatus"
}
