package remoteStartTransaction_test

import (
	"fmt"

	rst "github.com/aasanchez/ocpp16messages/remoteStartTransaction"
)

// ExampleConf demonstrates creating a valid RemoteStartTransaction.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := rst.Conf(rst.ConfInput{
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

// ExampleConf_rejected demonstrates creating a RemoteStartTransaction.conf
// message with a Rejected status.
func ExampleConf_rejected() {
	conf, err := rst.Conf(rst.ConfInput{
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
	_, err := rst.Conf(rst.ConfInput{
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
	_, err := rst.Conf(rst.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
