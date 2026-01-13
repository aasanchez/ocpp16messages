package clearCache_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearCache"
)

// ExampleConf demonstrates creating a valid ClearCache.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := clearCache.Conf(clearCache.ConfInput{
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

// ExampleConf_rejected demonstrates creating a ClearCache.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := clearCache.Conf(clearCache.ConfInput{
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
	_, err := clearCache.Conf(clearCache.ConfInput{
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
	_, err := clearCache.Conf(clearCache.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
