package updatefirmware_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/updatefirmware"
)

// ExampleConf demonstrates creating an UpdateFirmware.conf message.
// UpdateFirmware.conf is an empty confirmation message per OCPP 1.6
// specification.
func ExampleConf() {
	input := updatefirmware.ConfInput{}

	_, err := updatefirmware.Conf(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("UpdateFirmware.conf created successfully")
	// Output:
	// UpdateFirmware.conf created successfully
}
