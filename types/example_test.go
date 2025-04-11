//go:build example

package types_test

import (
	"fmt"
	"log"
)

func ExampleIdToken() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input

	idToken, err := IdToken(input)
	if err != nil {
		log.Fatalf("Failed to create IdToken: %v", err)
		return
	}

	fmt.Printf("Valid IdToken: %s\n", idToken.String())
}
