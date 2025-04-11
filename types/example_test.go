//go:build example

package types_test

import (
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/types"
)

func ExampleIdTokenType() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input

	idToken, err := types.IdToken(input)
	if err != nil {
		log.Fatalf("Failed to create IdToken: %v", err)
		return
	}

	fmt.Printf("Valid IdToken: %s\n", idToken.String())
}

func ExampleCiString() {
	input := "Hello, OCPP!"
	maxLen := 20

	// Attempt to create a valid ciString
	cs, err := types.CiString(input, maxLen)
	if err != nil {
		log.Fatalf("Failed to create ciString: %v", err)
	}

	fmt.Println("ciString created:", cs.String())

	// Output:
	// ciString created: Hello, OCPP!
}
