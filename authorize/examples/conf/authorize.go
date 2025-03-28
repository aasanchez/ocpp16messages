package main

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/authorize"
)

func main() {
	fmt.Println("--- Scenario A: Build Authorize.conf from variables ---")
	status := "Accepted"
	expiryDate := "2025-12-31T23:59:59Z"
	parentIdTag := "PARENT-456"

	conf, err := authorize.NewAuthorizeConf(status, &expiryDate, &parentIdTag)
	if err != nil {
		fmt.Println("Error building Authorize.conf:", err)
		return
	}
	if err := conf.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	fmt.Println("Authorize.conf built successfully:")
	outputA, _ := json.MarshalIndent(conf, "", "  ")
	fmt.Println(string(outputA))

	fmt.Println("\n--- Scenario B: Parse existing Authorize.conf JSON ---")
	raw := []byte(`[3, "res-1", {"idTagInfo": {"status": "Accepted", "expiryDate": "2025-12-31T23:59:59Z", "parentIdTag": "PARENT-456"}}]`)

	parsed, err := ocpp16_messages.ParseAndValidate(raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	confParsed, ok := parsed.Payload.(authorize.AuthorizeConf)
	if !ok {
		fmt.Println("Unexpected payload type")
		return
	}

	fmt.Println("Parsed Authorize.conf successfully:")
	fmt.Println("Unique ID:", parsed.UniqueID)
	fmt.Println("Status:", confParsed.IdTagInfo.Status)
	if confParsed.IdTagInfo.ExpiryDate != nil {
		fmt.Println("Expiry Date:", *confParsed.IdTagInfo.ExpiryDate)
	}
	if confParsed.IdTagInfo.ParentIdTag != nil {
		fmt.Println("Parent IdTag:", *confParsed.IdTagInfo.ParentIdTag)
	}
}
