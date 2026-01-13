package changeConfiguration_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeConfiguration"
)

const statusLabel = "Status:"

// ExampleConf demonstrates creating a valid ChangeConfiguration.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
	conf, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
	conf, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
	conf, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
	_, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
	_, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
