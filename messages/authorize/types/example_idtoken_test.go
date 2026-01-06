package types_test

//revive:disable line-length-limit

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ExampleNewIdToken demonstrates creating a valid IdToken from a
// CiString20Type.
func ExampleNewIdToken() {
	ciStr, _ := st.NewCiString20Type("RFID-TAG-12345")

	idToken, err := mat.NewIdToken(ciStr)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: RFID-TAG-12345
}

// ExampleNewIdToken_empty demonstrates the error returned when trying to
// create an IdToken from an empty CiString20Type value.
func ExampleNewIdToken_empty() {
	ciStr, _ := st.NewCiString20Type("")

	_, err := mat.NewIdToken(ciStr)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// NewIdToken: IdToken: value cannot be empty
}

// ExampleNewIdToken_shortTag demonstrates creating an IdToken with a short
// tag value, showing that any non-empty value is acceptable.
func ExampleNewIdToken_shortTag() {
	ciStr, _ := st.NewCiString20Type("TAG1")

	idToken, err := mat.NewIdToken(ciStr)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: TAG1
}
