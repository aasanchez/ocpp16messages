package cancelreservation_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/cancelreservation"
)

// ExampleConf demonstrates creating a valid CancelReservation.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := cancelreservation.Conf(cancelreservation.ConfInput{
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

// ExampleConf_rejected demonstrates creating a CancelReservation.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := cancelreservation.Conf(cancelreservation.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := cancelreservation.Conf(cancelreservation.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}

// ExampleConf_emptyStatus demonstrates the error returned when
// an empty status is provided.
func ExampleConf_emptyStatus() {
	_, err := cancelreservation.Conf(cancelreservation.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
