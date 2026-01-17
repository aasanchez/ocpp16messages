package heartbeat_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/heartbeat"
)

// ExampleReq demonstrates creating a valid Heartbeat.req message.
// Heartbeat.req has no fields - it simply signals that the Charge Point
// is still connected.
func ExampleReq() {
	_, err := heartbeat.Req(heartbeat.ReqInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Heartbeat request created successfully")
	// Output:
	// Heartbeat request created successfully
}
