package triggerMessage_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/triggerMessage"
)

const (
	labelStatus = "Status:"
)

// ExampleConf demonstrates creating a valid TriggerMessage.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_rejected demonstrates creating a TriggerMessage.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_notImplemented demonstrates creating a TriggerMessage.conf
// message with a NotImplemented status.
func ExampleConf_notImplemented() {
	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "NotImplemented",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: NotImplemented
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := triggerMessage.Conf(triggerMessage.ConfInput{
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
	_, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "",
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
