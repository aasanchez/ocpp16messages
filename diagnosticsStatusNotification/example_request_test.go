package diagnosticsStatusNotification_test

import (
	"fmt"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification"
)

const labelStatus = "Status:"

// ExampleReq demonstrates creating a valid DiagnosticsStatusNotification.req
// message with an Idle status.
func ExampleReq() {
	req, err := dsn.Req(dsn.ReqInput{Status: "Idle"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Idle
}

// ExampleReq_uploading demonstrates creating a
// DiagnosticsStatusNotification.req message with an Uploading status.
func ExampleReq_uploading() {
	req, err := dsn.Req(dsn.ReqInput{Status: "Uploading"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Uploading
}

// ExampleReq_uploaded demonstrates creating a DiagnosticsStatusNotification.req
// message with an Uploaded status.
func ExampleReq_uploaded() {
	req, err := dsn.Req(dsn.ReqInput{Status: "Uploaded"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: Uploaded
}

// ExampleReq_uploadFailed demonstrates creating a
// DiagnosticsStatusNotification.req message with an UploadFailed status.
func ExampleReq_uploadFailed() {
	req, err := dsn.Req(dsn.ReqInput{Status: "UploadFailed"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// Status: UploadFailed
}

// ExampleReq_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleReq_invalidStatus() {
	_, err := dsn.Req(dsn.ReqInput{Status: "InvalidStatus"})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: invalid value
}
