package sendLocalList_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/sendLocalList"
)

// ExampleConf demonstrates creating a valid SendLocalList.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Status: %s\n", conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_failed demonstrates creating a SendLocalList.conf message
// with a Failed status.
func ExampleConf_failed() {
	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Failed",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Status: %s\n", conf.Status.String())
	// Output:
	// Status: Failed
}

// ExampleConf_versionMismatch demonstrates creating a SendLocalList.conf
// message with a VersionMismatch status.
func ExampleConf_versionMismatch() {
	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "VersionMismatch",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("Status: %s\n", conf.Status.String())
	// Output:
	// Status: VersionMismatch
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
