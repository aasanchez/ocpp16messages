package authorizetypes_test

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleIdToken_valid() {
	str, err := st.SetCiString20Type("ABC123456789")
	if err != nil {
		fmt.Println("Error creating CiString20:", err)

		return
	}

	idToken, err := mat.IdToken(str)
	if err != nil {
		fmt.Println("Error creating IdToken:", err)

		return
	}

	fmt.Println("IdToken value:", idToken.Value())

	// Output:
	// IdToken value: ABC123456789
}
