package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aasanchez/ocpp16messages/messages"
	"github.com/aasanchez/ocpp16messages/types"
)

// FullMessage represents the full structure of a typical OCPP message
type FullMessage struct {
	MessageTypeId int                              `json:"MessageTypeId"`
	UniqueId      string                           `json:"UniqueId"`
	Action        string                           `json:"Action"`
	Payload       messages.AuthorizeRequestMessage `json:"Payload"`
}

func main() {
	// Simulate receiving a full JSON message (header + payload)
	receivedJSON := `[
		2,
		"19223201",
		"Authorize",
		{
			"idTag": "ThisIsMySuperIDTag"
		}
	]`

	fmt.Println(strings.Repeat("#", 120))

	// Step 1: Deserialize the full message (header + payload)
	var rawMsg []interface{}
	err := json.Unmarshal([]byte(receivedJSON), &rawMsg)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Extract MessageTypeId, UniqueId, and Action
	messageTypeId, ok := rawMsg[0].(float64)
	if !ok {
		log.Fatalf("Invalid MessageTypeId type")
	}

	uniqueId, ok := rawMsg[1].(string)
	if !ok {
		log.Fatalf("Invalid UniqueId type")
	}

	action, ok := rawMsg[2].(string)
	if !ok {
		log.Fatalf("Invalid Action type")
	}

	// Step 2: Deserialize the payload part (the 'idTag' field)
	payloadJSON, err := json.Marshal(rawMsg[3])
	if err != nil {
		log.Fatalf("Error marshaling payload: %v", err)
	}

	// Deserialize the Payload into a map to extract 'idTag'
	var payloadMap map[string]interface{}
	err = json.Unmarshal(payloadJSON, &payloadMap)
	if err != nil {
		log.Fatalf("Error unmarshaling payload: %v", err)
	}

	// Manually extract and validate the 'idTag' field from the payload
	idTagStr, ok := payloadMap["idTag"].(string)
	if !ok {
		log.Fatalf("Error: idTag field is missing or invalid")
	}

	// Create an IdTag object using the extracted string
	idTag, err := types.IdToken(idTagStr)
	if err != nil {
		log.Fatalf("Error creating IdTag: %v", err)
	}

	// Create the AuthorizeRequestMessage with the validated IdTag
	payload := messages.AuthorizeRequestMessage{IdTag: idTag}

	// Step 3: Verify the Action and MessageTypeId (assume these are valid for simplicity)
	fmt.Printf("Received Message Type: %d\n", int(messageTypeId))
	fmt.Printf("Received Action: %s\n", action)
	fmt.Printf("Received UniqueId: %s\n", uniqueId)

	// Step 4: Validate the payload (the AuthorizeRequestMessage)
	if err := payload.Validate(); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Step 5: Handle the deserialized request (e.g., print it out)
	// This will print the formatted message with {idTag=...}
	fmt.Printf("Received and Validated AuthorizeRequestMessage: %s\n", payload.String())

	// Additional line to print just the idTag value
	// This will now print only the idTag value, as requested
	fmt.Printf("idTag: %s\n", payload.IdTag.String())

	fmt.Println(strings.Repeat("#", 120))
}
