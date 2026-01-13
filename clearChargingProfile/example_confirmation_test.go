package clearChargingProfile_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearChargingProfile"
)

// ExampleConf demonstrates creating a valid ClearChargingProfile.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := clearChargingProfile.Conf(clearChargingProfile.ConfInput{
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

// ExampleConf_unknown demonstrates creating a ClearChargingProfile.conf message
// with an Unknown status.
func ExampleConf_unknown() {
	conf, err := clearChargingProfile.Conf(clearChargingProfile.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.Status.String())
	// Output:
	// Status: Unknown
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := clearChargingProfile.Conf(clearChargingProfile.ConfInput{
		Status: "Invalid",
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
	_, err := clearChargingProfile.Conf(clearChargingProfile.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
