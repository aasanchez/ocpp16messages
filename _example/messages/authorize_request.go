package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
)

// FullMessage represents the full structure of a typical OCPP message
type FullMessage struct {
	MessageTypeId int                        `json:"MessageTypeId"`
	UniqueId      string                     `json:"UniqueId"`
	Action        string                     `json:"Action"`
	Payload       authorize.AuthorizeRequest `json:"Payload"`
}

func main() {
	// Simulate receiving a full JSON message (header + payload)
	receivedJSON := `[
		2,
		"12345",
		"Authorize",
		{
			"IdTag": "ABC1234567890XYZ78"
		}
	]`

	// Step 1: Deserialize the full message (header + payload)
	var fullMsg FullMessage
	err := json.Unmarshal([]byte(receivedJSON), &fullMsg)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON into FullMessage: %v", err)
	}

	// Step 2: Verify the Action and MessageTypeId (assume these are valid for simplicity)
	fmt.Printf("Received Message Type: %d\n", fullMsg.MessageTypeId)
	fmt.Printf("Received Action: %s\n", fullMsg.Action)
	fmt.Printf("Received UniqueId: %s\n", fullMsg.UniqueId)

	// Step 3: Validate the payload (the AuthorizeRequestMessage)
	if err := fullMsg.Payload.Validate(); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Step 4: Handle the deserialized request (e.g., print it out)
	fmt.Printf("Received and Validated AuthorizeRequestMessage: %s\n", fullMsg.Payload.String())
}
