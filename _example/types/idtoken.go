// Package main demonstrates how to use the IdToken type from the ocpp16messages/types package.
package main

import (
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/types"
)

func main() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input

	idToken, err := types.IdTokenFromString(input)
	if err != nil {
		log.Fatalf("❌ Failed to create IdToken: %v", err)
	}

	fmt.Printf("✅ Valid IdToken: %s\n", idToken.String())

	// Optional: validate it again later
	if err := idToken.Validate(); err != nil {
		log.Fatalf("❌ Unexpected validation error: %v", err)
	}

	// Now you can pass this IdToken into a request message (e.g., Authorize.req)
}
