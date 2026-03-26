package unlockconnector_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/unlockconnector"
)

const (
	outStatusLabel = "Status:"
)

// ExampleConf demonstrates creating a valid UnlockConnector.conf message
// with an Unlocked status.
func ExampleConf() {
	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "Unlocked",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outStatusLabel, conf.Status.String())
	// Output:
	// Status: Unlocked
}

// ExampleConf_unlockFailed demonstrates creating an UnlockConnector.conf
// message with an UnlockFailed status.
func ExampleConf_unlockFailed() {
	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "UnlockFailed",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outStatusLabel, conf.Status.String())
	// Output:
	// Status: UnlockFailed
}

// ExampleConf_notSupported demonstrates creating an UnlockConnector.conf
// message with a NotSupported status.
func ExampleConf_notSupported() {
	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outStatusLabel, conf.Status.String())
	// Output:
	// Status: NotSupported
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
