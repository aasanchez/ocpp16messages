package types_test

//revive:disable line-length-limit

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// ExampleNewIdToken demonstrates creating a valid IdToken from a
// CiString20Type.
func ExampleNewIdToken() {
	ciStr, _ := types.NewCiString20Type("RFID-TAG-12345")

	idToken := types.NewIdToken(ciStr)

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: RFID-TAG-12345
}

// ExampleNewIdToken_empty demonstrates that empty strings are rejected at the
// CiString20Type level, not at the IdToken level.
func ExampleNewIdToken_empty() {
	_, err := types.NewCiString20Type("")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// value cannot be empty
}

// ExampleNewIdToken_shortTag demonstrates creating an IdToken with a short
// tag value, showing that any non-empty value is acceptable.
func ExampleNewIdToken_shortTag() {
	ciStr, _ := types.NewCiString20Type("TAG1")

	idToken := types.NewIdToken(ciStr)

	fmt.Println("IdToken:", idToken.String())
	// Output:
	// IdToken: TAG1
}
