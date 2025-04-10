package main

import (
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/types"
)

func main() {
	// Valid IdTag (20-characte
	input := "ABC1234567890XYZ7890"

	// Create a validated IdTag
	idTag, err := types.IdTag(input)
	if err != nil {
		log.Fatalf("❌ Failed to create IdTag: %v", err)
	}

	// Use the IdTag
	fmt.Println("✅ Validated IdTag:", idTag.String())
}
