package main

import (
	"log"
	"os"
	"time"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/validators"
)

func main() {
	// Set log format to include time
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	req := chargePoint.BootNotificationReq{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorY",
	}

	exitOnError(validators.ValidateBootNotificationReq(req), "BootNotificationReq validation failed")

	log.Println("✅ Valid BootNotification request")

	conf := chargePoint.BootNotificationConf{
		Status:      "Accepted", // try "Invalid" to test
		CurrentTime: time.Now(),
		Interval:    30,
	}

	exitOnError(validators.ValidateBootNotificationConf(conf), "BootNotificationConf validation failed")

	log.Println("✅ Valid BootNotification confirmation")
}

// exitOnError logs the error and exits if err is not nil
func exitOnError(err error, context string) {
	if err != nil {
		log.Printf("❌ %s: %v\n", context, err)
		os.Exit(1)
	}
}
