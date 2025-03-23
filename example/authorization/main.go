package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

func main() {
	raw := []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)

	// Parse and validate the raw OCPP message
	msg, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Fatalf("❌ Failed to parse OCPP message: %v", err)
	}

	// Validate the message
	if err := core.ValidateMessage(msg.Action, msg.Payload); err != nil {
		log.Fatalf("❌ Validation error: %v", err)
	}

	// Check the decoded payload type
	req, ok := msg.DecodedPayload.(authorize.AuthorizeReq)
	if !ok {
		log.Fatalf("❌ Payload is not an AuthorizeReq")
	}

	fmt.Println("✅ Valid Authorize.req message")
	fmt.Printf("  ID     : %s\n", msg.ID)
	fmt.Printf("  idTag  : %s\n", req.IdTag)

	// Build a response
	resp := authorize.Conf{
		IdTagInfo: authorize.IdTagInfo{
			Status: core.AuthorizationAccepted,
		},
	}

	// Marshal the response
	encoded, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("❌ Failed to marshal response: %v", err)
	}
	fmt.Printf("✅ Encoded Authorize.conf: %s\n", string(encoded))
}
