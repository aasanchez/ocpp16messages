package main

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/authorize"
)

func main() {
	// Simulate incoming Authorize.req JSON message
	raw := []byte(`[2, "req-1", "Authorize", {"idTag": "ABC123"}]`)

	parsed, err := ocpp16_messages.ParseAndValidate(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ensure it is an Authorize.req
	req, ok := parsed.Payload.(authorize.AuthorizeReq)
	if !ok {
		fmt.Println("Unexpected payload type")
		return
	}

	// Output parsed and validated data
	fmt.Println("Parsed Authorize.req successfully:")
	fmt.Println("Unique ID:", parsed.UniqueID)
	fmt.Println("Action:", parsed.Action)
	fmt.Println("IdTag:", req.IdTag.Value())

	// Marshal to verify structure (optional)
	pretty, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println("\nJSON Structure:")
	fmt.Println(string(pretty))
}
