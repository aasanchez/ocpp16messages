// Command ocpp16_example demonstrates how to construct and validate
// BootNotification messages (req/conf) using the ocpp16_messages package.
//
// This example uses the chargePoint messages and validators packages to simulate
// sending and validating BootNotification.req and BootNotification.conf as defined
// in the OCPP 1.6J specification.
package main

import (
	"log"
	"os"
	"time"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/validators"
)

func main() {
	// Configure log output to show timestamps and source file/line
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Example BootNotification.req from a Charge Point
	req := chargePoint.BootNotificationReq{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorY",
	}

	// Validate the request and exit if it's invalid
	exitOnError(validators.ValidateBootNotificationReq(req), "BootNotificationReq validation failed")

	log.Println("✅ Valid BootNotification request")

	// Example BootNotification.conf from the Central System
	conf := chargePoint.BootNotificationConf{
		Status:      "Accepted", // Change to an invalid string to test failure
		CurrentTime: time.Now(),
		Interval:    30,
	}

	// Validate the confirmation and exit if it's invalid
	exitOnError(validators.ValidateBootNotificationConf(conf), "BootNotificationConf validation failed")

	log.Println("✅ Valid BootNotification confirmation")
}

// exitOnError logs the given error with context and exits the program with status code 1.
// This is a utility function to simplify error handling in example code.
func exitOnError(err error, context string) {
	if err != nil {
		log.Printf("❌ %s: %v\n", context, err)
		os.Exit(1)
	}
}
