package bootNotification_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/bootNotification"
)

const labelStatus = "Status:"

// ExampleConf demonstrates creating a valid BootNotification.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := bootNotification.Conf(bootNotification.ConfInput{
		Status:      "Accepted",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    300,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	fmt.Println("Interval:", conf.Interval.Value())
	// Output:
	// Status: Accepted
	// Interval: 300
}

// ExampleConf_pending demonstrates creating a BootNotification.conf message
// with a Pending status.
func ExampleConf_pending() {
	conf, err := bootNotification.Conf(bootNotification.ConfInput{
		Status:      "Pending",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    60,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: Pending
}

// ExampleConf_rejected demonstrates creating a BootNotification.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := bootNotification.Conf(bootNotification.ConfInput{
		Status:      "Rejected",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    600,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	fmt.Println("Interval:", conf.Interval.Value())
	// Output:
	// Status: Rejected
	// Interval: 600
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := bootNotification.Conf(bootNotification.ConfInput{
		Status:      "Unknown",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    300,
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}

// ExampleConf_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleConf_multipleErrors() {
	_, err := bootNotification.Conf(bootNotification.ConfInput{
		Status:      "Invalid-Status",
		CurrentTime: "not-a-date",
		Interval:    -1,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
