package updateFirmware_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/updateFirmware"
)

const (
	exampleLocationValue     = "https://example.com/firmware/v1.2.3.bin"
	exampleRetrieveDateValue = "2025-01-15T10:00:00Z"
	exampleRetriesValue      = 3
	exampleRetryIntValue     = 60
	exampleNegativeValue     = -1

	outLocationLabel = "Location:"
)

// ExampleReq demonstrates creating an UpdateFirmware.req message
// with required fields only.
func ExampleReq() {
	req, err := updateFirmware.Req(updateFirmware.ReqInput{
		Location:      exampleLocationValue,
		RetrieveDate:  exampleRetrieveDateValue,
		Retries:       nil,
		RetryInterval: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	// Output:
	// Location: https://example.com/firmware/v1.2.3.bin
}

// ExampleReq_withRetries demonstrates creating an UpdateFirmware.req message
// with the optional Retries field.
func ExampleReq_withRetries() {
	retries := exampleRetriesValue

	req, err := updateFirmware.Req(updateFirmware.ReqInput{
		Location:      exampleLocationValue,
		RetrieveDate:  exampleRetrieveDateValue,
		Retries:       &retries,
		RetryInterval: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	fmt.Println("Retries:", req.Retries.Value())
	// Output:
	// Location: https://example.com/firmware/v1.2.3.bin
	// Retries: 3
}

// ExampleReq_withAllFields demonstrates creating an UpdateFirmware.req message
// with all optional fields.
func ExampleReq_withAllFields() {
	retries := exampleRetriesValue
	retryInterval := exampleRetryIntValue

	req, err := updateFirmware.Req(updateFirmware.ReqInput{
		Location:      exampleLocationValue,
		RetrieveDate:  exampleRetrieveDateValue,
		Retries:       &retries,
		RetryInterval: &retryInterval,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	fmt.Println("Retries:", req.Retries.Value())
	fmt.Println("RetryInterval:", req.RetryInterval.Value())
	// Output:
	// Location: https://example.com/firmware/v1.2.3.bin
	// Retries: 3
	// RetryInterval: 60
}

// ExampleReq_invalidLocation demonstrates the error returned when
// an empty Location is provided.
func ExampleReq_invalidLocation() {
	_, err := updateFirmware.Req(updateFirmware.ReqInput{
		Location:      "",
		RetrieveDate:  exampleRetrieveDateValue,
		Retries:       nil,
		RetryInterval: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid location")
	}
	// Output:
	// Error: invalid location
}

// ExampleReq_invalidRetries demonstrates the error returned when
// a negative Retries value is provided.
func ExampleReq_invalidRetries() {
	retries := exampleNegativeValue

	_, err := updateFirmware.Req(updateFirmware.ReqInput{
		Location:      exampleLocationValue,
		RetrieveDate:  exampleRetrieveDateValue,
		Retries:       &retries,
		RetryInterval: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid retries")
	}
	// Output:
	// Error: invalid retries
}
