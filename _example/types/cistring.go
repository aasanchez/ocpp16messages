// Package main demonstrates how to use validated CiStringXXType types
// from the ocpp16messages/types package.
package main

import (
	"log"

	"github.com/aasanchez/ocpp16messages/types"
)

func main() {
	// Valid input: 20 ASCII characters
	input := "ABCDEFGHIJ1234567890"

	ciString, err := types.CiString20(input)
	if err != nil {
		log.Println("❌ CiString20 validation failed:", err)

		return
	}

	log.Println("✅ Validated CiString20 value:", ciString.String())

	// Optional: explicitly validate again later
	if err := ciString.Validate(); err != nil {
		log.Println("Unexpected validation error:", err)
	}

	// Invalid input: too long
	tooLong := input + "X"
	_, err = types.CiString20(tooLong)

	if err != nil {
		log.Println("❌ Expected error for too-long input:", err)
	}

	// Invalid input: contains non-ASCII character
	invalid := "Hello🚗World"
	_, err = types.CiString20(invalid)

	if err != nil {
		log.Println("❌ Expected error for non-ASCII input:", err)
	}
}
