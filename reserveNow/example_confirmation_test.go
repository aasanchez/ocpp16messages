package reservenow_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/reservenow"
)

// ExampleConf demonstrates creating a valid ReserveNow.conf message with
// Accepted status.
func ExampleConf() {
	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_occupied demonstrates creating a ReserveNow.conf message with
// Occupied status.
func ExampleConf_occupied() {
	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: "Occupied",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.Status.String())
	// Output:
	// Status: Occupied
}

// ExampleConf_invalidStatus demonstrates the error returned when an invalid
// status is provided.
func ExampleConf_invalidStatus() {
	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("status: invalid value")
	}
	// Output:
	// status: invalid value
}
