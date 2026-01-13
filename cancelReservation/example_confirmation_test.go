package cancelReservation_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/cancelReservation"
)

// ExampleConf demonstrates creating a valid CancelReservation.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := cancelReservation.Conf(cancelReservation.ConfInput{
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
	conf, err := cancelReservation.Conf(cancelReservation.ConfInput{
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
	_, err := cancelReservation.Conf(cancelReservation.ConfInput{
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
	_, err := cancelReservation.Conf(cancelReservation.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
