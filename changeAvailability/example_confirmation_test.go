package changeAvailability_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeAvailability"
)

const statusLabel = "Status:"

// ExampleConf demonstrates creating a valid ChangeAvailability.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := changeAvailability.Conf(changeAvailability.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_rejected demonstrates creating a ChangeAvailability.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := changeAvailability.Conf(changeAvailability.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_scheduled demonstrates creating a ChangeAvailability.conf message
// with a Scheduled status (when a transaction is in progress).
func ExampleConf_scheduled() {
	conf, err := changeAvailability.Conf(changeAvailability.ConfInput{
		Status: "Scheduled",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: Scheduled
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := changeAvailability.Conf(changeAvailability.ConfInput{
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
	_, err := changeAvailability.Conf(changeAvailability.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
