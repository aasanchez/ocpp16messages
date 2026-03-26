package statusnotification_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/statusnotification"
)

// ExampleConf demonstrates creating a StatusNotification.conf message.
// This message has no fields per OCPP 1.6 specification.
func ExampleConf() {
	_, err := statusnotification.Conf(statusnotification.ConfInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("StatusNotification.conf created successfully")
	// Output:
	// StatusNotification.conf created successfully
}
