// Command authorize_roundtrip demonstrates how to parse and validate both
// Authorize.req and Authorize.conf messages using the ocpp16_messages package.
package main

import (
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/authorize"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Example 1: Validate Authorize.req (CALL)
	call := []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)
	reqResult, err := ocpp16_messages.ValidateMessage(call)
	exitOnError(err, "Authorize.req validation failed")

	if req, ok := reqResult.Payload.(authorize.AuthorizeReq); ok {
		fmt.Println("✅ Authorize.req is valid:")
		fmt.Printf("  Unique ID : %s\n", reqResult.UniqueID)
		fmt.Printf("  IdTag     : %s\n", req.IdTag)
	} else {
		log.Fatalf("❌ Payload is not AuthorizeReq")
	}

	// Example 2: Validate Authorize.conf (CALLRESULT)
	conf := []byte(`[3, "01221201194032", {"idTagInfo":{"status":"Accepted"}}]`)
	confResult, err := ocpp16_messages.ValidateMessage(conf)
	exitOnError(err, "Authorize.conf validation failed")

	if res, ok := confResult.Payload.(authorize.AuthorizeConf); ok {
		fmt.Println("✅ Authorize.conf is valid:")
		fmt.Printf("  Unique ID : %s\n", confResult.UniqueID)
		fmt.Printf("  Status    : %s\n", res.IdTagInfo.Status)
	} else {
		log.Fatalf("❌ Payload is not AuthorizeConf")
	}
}

// exitOnError logs the error and exits if one occurs
func exitOnError(err error, context string) {
	if err != nil {
		log.Fatalf("❌ %s: %v", context, err)
	}
}
