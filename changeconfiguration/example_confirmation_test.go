package changeconfiguration_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeconfiguration"
)

const statusLabel = "Status:"

// ExampleConf demonstrates creating a valid ChangeConfiguration.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
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

// ExampleConf_rejected demonstrates creating a ChangeConfiguration.conf
// message with a Rejected status.
func ExampleConf_rejected() {
	conf, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
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

// ExampleConf_rebootRequired demonstrates creating a ChangeConfiguration.conf
// message with a RebootRequired status.
func ExampleConf_rebootRequired() {
	conf, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
		Status: "RebootRequired",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: RebootRequired
}

// ExampleConf_notSupported demonstrates creating a ChangeConfiguration.conf
// message with a NotSupported status.
func ExampleConf_notSupported() {
	conf, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: NotSupported
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
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
	_, err := changeconfiguration.Conf(changeconfiguration.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
