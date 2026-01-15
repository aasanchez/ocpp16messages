package firmwareStatusNotification_test

import (
	"fmt"

	fsn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification"
)

// ExampleConf demonstrates creating a valid FirmwareStatusNotification.conf
// message. This message has no fields per OCPP 1.6 specification.
func ExampleConf() {
	_, err := fsn.Conf(fsn.ConfInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Confirmation created successfully")
	// Output:
	// Confirmation created successfully
}
