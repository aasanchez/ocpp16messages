package authorizetypes_test

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleIdToken_valid() {
	str, err := sharedtypes.SetCiString20("ABC123456789")
	if err != nil {
		fmt.Println("Error creating CiString20:", err)

		return
	}

	idToken, err := authorizetypes.SetIdToken(str)
	if err != nil {
		fmt.Println("Error creating IdToken:", err)

		return
	}

	fmt.Println("IdToken value:", idToken.Value())

	// Output:
	// IdToken value: ABC123456789
}
