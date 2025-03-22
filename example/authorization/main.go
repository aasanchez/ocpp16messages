// Command main demonstrates how to decode and validate an Authorize.req and Authorize.conf
// message using the ocpp16_messages package.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Example raw CALL message for Authorize.req
	raw := []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)

	result, err := core.ValidateRawMessage(raw)
	if err != nil {
		log.Fatalf("❌ Failed to validate Authorize.req: %v", err)
	}

	req, ok := result.Payload.(authorize.AuthorizeReq)
	if !ok {
		log.Fatalf("❌ Payload is not an AuthorizeReq")
	}

	fmt.Println("✅ Valid Authorize.req message")
	fmt.Printf("  Unique ID : %s\n", result.UniqueID)
	fmt.Printf("  Action    : %s\n", result.Action)
	fmt.Printf("  IdTag     : %s\n", req.IdTag)

	// Example Authorize.conf response
	conf := authorize.AuthorizeConf{
		IdTagInfo: authorize.IdTagInfo{
			Status: authorize.AuthorizationAccepted,
		},
	}

	if err := authorize.ValidateAuthorizeConf(conf); err != nil {
		log.Fatalf("❌ Invalid Authorize.conf: %v", err)
	}

	output, _ := json.MarshalIndent(conf, "", "  ")
	fmt.Println("✅ Valid Authorize.conf message:")
	fmt.Println(string(output))
}
