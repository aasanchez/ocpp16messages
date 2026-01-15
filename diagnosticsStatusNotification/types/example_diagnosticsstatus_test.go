package types_test

import (
	"fmt"

	dn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification/types"
)

const (
	labelStatus  = "Status:"
	labelIsValid = "IsValid:"
)

// ExampleDiagnosticsStatusIdle demonstrates using the Idle status,
// sent when no diagnostics upload is in progress.
func ExampleDiagnosticsStatusIdle() {
	status := dn.DiagnosticsStatusIdle
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Idle
	// IsValid: true
}

// ExampleDiagnosticsStatusUploading demonstrates using the Uploading status,
// sent when diagnostics upload is in progress.
func ExampleDiagnosticsStatusUploading() {
	status := dn.DiagnosticsStatusUploading
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Uploading
	// IsValid: true
}

// ExampleDiagnosticsStatusUploaded demonstrates using the Uploaded status,
// sent when diagnostics upload has completed successfully.
func ExampleDiagnosticsStatusUploaded() {
	status := dn.DiagnosticsStatusUploaded
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Uploaded
	// IsValid: true
}

// ExampleDiagnosticsStatusUploadFailed demonstrates using the UploadFailed
// status, sent when diagnostics upload has failed.
func ExampleDiagnosticsStatusUploadFailed() {
	status := dn.DiagnosticsStatusUploadFailed
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: UploadFailed
	// IsValid: true
}

// ExampleDiagnosticsStatus_IsValid_invalid demonstrates that invalid status
// values are rejected.
func ExampleDiagnosticsStatus_IsValid_invalid() {
	status := dn.DiagnosticsStatus("InvalidStatus")
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// IsValid: false
}
