package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/messages"
)

// FullMessage represents the full structure of a typical OCPP message
type FullMessage struct {
	MessageTypeId int                              `json:"MessageTypeId"`
	UniqueId      string                           `json:"UniqueId"`
	Action        string                           `json:"Action"`
	Payload       messages.AuthorizeRequestMessage `json:"Payload"`
}

// ValidateJSON validates if the input JSON matches the schema (e.g., idTag is required and less than 20 chars)
func ValidateJSON(data []byte) error {
	var msg FullMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Validate Payload (idTag) max length constraint and required field
	if len(msg.Payload.IdTag.String()) == 0 {
		return errors.New("idTag is required")
	}

	if len(msg.Payload.IdTag.String()) > 20 {
		return errors.New("idTag exceeds maximum length of 20")
	}

	// You can add more field-specific validation here if necessary.

	return nil
}

func main() {
	// Simulate receiving a full JSON message (header + payload)
	receivedJSON := `{
		2,
		"12345",
		"Authorize",
		"Payload": {
			"IdTag": "ABC1234567890XYZ78"
		}
	}`

	// Step 1: Validate JSON against the schema
	if err := ValidateJSON([]byte(receivedJSON)); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Step 2: Deserialize the full message (header + payload)
	var fullMsg FullMessage
	err := json.Unmarshal([]byte(receivedJSON), &fullMsg)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON into FullMessage: %v", err)
	}

	// Step 3: Verify the Action and MessageTypeId (assume these are valid for simplicity)
	fmt.Printf("Received Message Type: %d\n", fullMsg.MessageTypeId)
	fmt.Printf("Received Action: %s\n", fullMsg.Action)
	fmt.Printf("Received UniqueId: %s\n", fullMsg.UniqueId)

	// Step 4: Validate the payload (the AuthorizeRequestMessage)
	if err := fullMsg.Payload.Validate(); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Step 5: Handle the deserialized request (e.g., print it out)
	fmt.Printf("Received and Validated AuthorizeRequestMessage: %s\n", fullMsg.Payload.String())
}
