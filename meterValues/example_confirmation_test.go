package meterValues_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues"
)

// ExampleConf demonstrates creating a MeterValues.conf message.
// MeterValues.conf is an empty confirmation message per OCPP 1.6 specification.
func ExampleConf() {
	input := meterValues.ConfInput{}

	_, err := meterValues.Conf(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("MeterValues.conf created successfully")
	// Output:
	// MeterValues.conf created successfully
}
