// Command authorize_example demonstrates how to construct and validate
// Authorize.req and Authorize.conf messages using the ocpp16_messages package.
package main

import (
	"log"
	"os"
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/models"
	"github.com/aasanchez/ocpp16_messages/validators"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Example Authorize Request from the Charge Point
	req := chargePoint.AuthorizeReq{
		IdToken: models.IdToken{
			IdTag: models.CiString20Type("ABC123456"),
		},
	}
	exitOnError(validators.ValidateAuthorizeReq(req), "AuthorizeReq validation failed")
	log.Println("✅ Valid Authorize request")

	// Example Authorize Confirmation from the CSMS
	parent := models.CiString20Type("Parent001")
	expiry := time.Now().Add(2 * time.Hour)

	conf := chargePoint.AuthorizeConf{
		IdTagInfo: chargePoint.IdTagInfo{
			Status:      enums.AuthorizationAccepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}
	exitOnError(validators.ValidateAuthorizeConf(conf), "AuthorizeConf validation failed")
	log.Println("✅ Valid Authorize confirmation")
}

// exitOnError logs the error and exits the program if err is not nil.
func exitOnError(err error, context string) {
	if err != nil {
		log.Printf("❌ %s: %v\n", context, err)
		os.Exit(1)
	}
}
