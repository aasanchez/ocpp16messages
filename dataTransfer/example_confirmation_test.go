package dataTransfer_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/dataTransfer"
)

const labelStatus = "Status:"

// ExampleConf demonstrates creating a valid DataTransfer.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "Accepted",
		Data:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_rejected demonstrates creating a DataTransfer.conf message
// with a Rejected status.
func ExampleConf_rejected() {
	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "Rejected",
		Data:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_unknownVendor demonstrates creating a DataTransfer.conf message
// with an UnknownVendor status, used when the vendorId is not recognized.
func ExampleConf_unknownVendor() {
	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "UnknownVendor",
		Data:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: UnknownVendor
}

// ExampleConf_unknownMessageId demonstrates creating a DataTransfer.conf
// message with an UnknownMessageId status.
func ExampleConf_unknownMessageId() {
	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "UnknownMessageId",
		Data:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	// Output:
	// Status: UnknownMessageId
}

// ExampleConf_withData demonstrates creating a DataTransfer.conf message
// with an optional data payload.
func ExampleConf_withData() {
	data := `{"temperature": 25.5}`

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "Accepted",
		Data:   &data,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.Status.String())
	fmt.Println("Data:", *conf.Data)
	// Output:
	// Status: Accepted
	// Data: {"temperature": 25.5}
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "InvalidStatus",
		Data:   nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: invalid value
}
