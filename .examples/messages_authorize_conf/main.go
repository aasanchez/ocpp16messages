package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

func main() {
	fmt.Println("=== EXAMPLE 1: Build a message from separate variables ===")
	exampleBuildFromSeparateVars()

	fmt.Println("\n=== EXAMPLE 2: Validate incoming message (status only) ===")
	exampleValidateStatusOnly()

	fmt.Println("\n=== EXAMPLE 3: Validate incoming message (status + expiryDate) ===")
	exampleValidateWithExpiry()

	fmt.Println("\n=== EXAMPLE 4: Validate incoming message (status + parentIdTag) ===")
	exampleValidateWithParent()

	fmt.Println("\n=== EXAMPLE 5: Validate incoming message (fully populated) ===")
	exampleValidateFullyPopulated()

	fmt.Println("\n=== EXAMPLE 6: Invalid input should fail validation ===")
	exampleInvalidStatusFails()
}

func exampleBuildFromSeparateVars() {
	// Example input values (e.g. from your CSMS logic)
	status := types.Accepted
	expiry := time.Now().Add(48 * time.Hour).UTC()
	parentStr := "PARENT1234"

	// Create optional parentIdTag
	parent, err := types.IdToken(parentStr)
	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	// Compose IdTagInfoType
	info := types.IdTagInfoType{
		Status:      status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	// Build confirmation message
	msg, err := authorize.Confirmation(info)
	if err != nil {
		log.Fatalf("failed to build confirmation message: %v", err)
	}

	fmt.Printf("Built Authorize.conf: %s\n", msg.String())
}

func exampleValidateStatusOnly() {
	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed ✅ (status only)")
}

func exampleValidateWithExpiry() {
	expiry := time.Now().Add(2 * time.Hour).UTC()

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed ✅ (with expiryDate)")
}

func exampleValidateWithParent() {
	parent, err := types.IdToken("GROUPXYZ")
	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &parent,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed ✅ (with parentIdTag)")
}

func exampleValidateFullyPopulated() {
	expiry := time.Now().Add(1 * time.Hour).UTC()
	parent, err := types.IdToken("FULLTAG")
	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed ✅ (fully populated)")
}

func exampleInvalidStatusFails() {
	info := types.IdTagInfoType{
		Status:      "InvalidStatus",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	err := msg.Validate()
	if err != nil {
		fmt.Printf("Expected validation failure ❌: %v\n", err)
	} else {
		log.Fatal("Expected validation to fail, but it passed")
	}
}
