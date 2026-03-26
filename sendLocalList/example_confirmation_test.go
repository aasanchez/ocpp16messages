package sendlocallist_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/sendlocallist"
)

const fmtStatusStr = "Status: %s\n"

// ExampleConf demonstrates creating a valid SendLocalList.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatusStr, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_failed demonstrates creating a SendLocalList.conf message
// with a Failed status.
func ExampleConf_failed() {
	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Failed",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatusStr, conf.Status.String())
	// Output:
	// Status: Failed
}

// ExampleConf_versionMismatch demonstrates creating a SendLocalList.conf
// message with a VersionMismatch status.
func ExampleConf_versionMismatch() {
	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "VersionMismatch",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatusStr, conf.Status.String())
	// Output:
	// Status: VersionMismatch
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
