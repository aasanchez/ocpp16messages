package diagnosticsStatusNotification_test

import (
	"fmt"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification"
)

// ExampleConf demonstrates creating a valid DiagnosticsStatusNotification.conf
// message. This message has no fields per OCPP 1.6 specification.
func ExampleConf() {
	_, err := dsn.Conf(dsn.ConfInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Confirmation created successfully")
	// Output:
	// Confirmation created successfully
}
