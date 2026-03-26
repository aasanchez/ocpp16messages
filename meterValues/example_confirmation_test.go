package metervalues_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/metervalues"
)

// ExampleConf demonstrates creating a MeterValues.conf message.
// MeterValues.conf is an empty confirmation message per OCPP 1.6 specification.
func ExampleConf() {
	input := metervalues.ConfInput{}

	_, err := metervalues.Conf(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("MeterValues.conf created successfully")
	// Output:
	// MeterValues.conf created successfully
}
