package setchargingprofile_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/setchargingprofile"
)

const fmtStatus = "Status: %s\n"

// ExampleConf demonstrates creating a valid SetChargingProfile.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatus, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_rejected demonstrates creating a SetChargingProfile.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatus, conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_notSupported demonstrates creating a SetChargingProfile.conf
// message with a NotSupported status.
func ExampleConf_notSupported() {
	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtStatus, conf.Status.String())
	// Output:
	// Status: NotSupported
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
