package getDiagnostics_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getDiagnostics"
)

const (
	exampleLocationValue  = "https://example.com/diagnostics"
	exampleRetriesValue   = 3
	exampleRetryIntValue  = 60
	exampleStartTimeValue = "2025-01-01T00:00:00Z"
	exampleStopTimeValue  = "2025-01-02T00:00:00Z"
	exampleNegativeValue  = -1

	outLocationLabel = "Location:"
)

// ExampleReq demonstrates creating a GetDiagnostics.req message
// with required fields only.
func ExampleReq() {
	req, err := getDiagnostics.Req(getDiagnostics.ReqInput{
		Location:      exampleLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	// Output:
	// Location: https://example.com/diagnostics
}

// ExampleReq_withRetries demonstrates creating a GetDiagnostics.req message
// with the optional Retries field.
func ExampleReq_withRetries() {
	retries := exampleRetriesValue

	req, err := getDiagnostics.Req(getDiagnostics.ReqInput{
		Location:      exampleLocationValue,
		Retries:       &retries,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	fmt.Println("Retries:", req.Retries.Value())
	// Output:
	// Location: https://example.com/diagnostics
	// Retries: 3
}

// ExampleReq_withAllFields demonstrates creating a GetDiagnostics.req message
// with all optional fields.
func ExampleReq_withAllFields() {
	retries := exampleRetriesValue
	retryInterval := exampleRetryIntValue
	startTime := exampleStartTimeValue
	stopTime := exampleStopTimeValue

	req, err := getDiagnostics.Req(getDiagnostics.ReqInput{
		Location:      exampleLocationValue,
		Retries:       &retries,
		RetryInterval: &retryInterval,
		StartTime:     &startTime,
		StopTime:      &stopTime,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outLocationLabel, req.Location.Value())
	fmt.Println("Retries:", req.Retries.Value())
	fmt.Println("RetryInterval:", req.RetryInterval.Value())
	// Output:
	// Location: https://example.com/diagnostics
	// Retries: 3
	// RetryInterval: 60
}

// ExampleReq_invalidLocation demonstrates the error returned when
// an empty Location is provided.
func ExampleReq_invalidLocation() {
	_, err := getDiagnostics.Req(getDiagnostics.ReqInput{
		Location:      "",
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
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

	_, err := getDiagnostics.Req(getDiagnostics.ReqInput{
		Location:      exampleLocationValue,
		Retries:       &retries,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		fmt.Println("Error: invalid retries")
	}
	// Output:
	// Error: invalid retries
}
