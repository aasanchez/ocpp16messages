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

	idToken := mat.NewIdToken(ciStr)

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: RFID-TAG-12345
}

// ExampleNewIdToken_empty demonstrates that empty strings are rejected at the
// CiString20Type level, not at the IdToken level.
func ExampleNewIdToken_empty() {
	_, err := st.NewCiString20Type("")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// value cannot be empty
}

// ExampleNewIdToken_shortTag demonstrates creating an IdToken with a short
// tag value, showing that any non-empty value is acceptable.
func ExampleNewIdToken_shortTag() {
	ciStr, _ := st.NewCiString20Type("TAG1")

	idToken := mat.NewIdToken(ciStr)

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: TAG1
}
