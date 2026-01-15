package types_test

import (
	"fmt"

	dt "github.com/aasanchez/ocpp16messages/dataTransfer/types"
)

const (
	labelStatus  = "Status:"
	labelIsValid = "IsValid:"
)

// ExampleDataTransferStatusAccepted demonstrates using the Accepted status.
func ExampleDataTransferStatusAccepted() {
	status := dt.DataTransferStatusAccepted
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleDataTransferStatusRejected demonstrates using the Rejected status.
func ExampleDataTransferStatusRejected() {
	status := dt.DataTransferStatusRejected
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: Rejected
	// IsValid: true
}

// ExampleDataTransferStatusUnknownMessageId demonstrates the UnknownMessageId
// status used when a messageId is not recognized.
func ExampleDataTransferStatusUnknownMessageId() {
	status := dt.DataTransferStatusUnknownMessageId
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: UnknownMessageId
	// IsValid: true
}

// ExampleDataTransferStatusUnknownVendor demonstrates the UnknownVendor status
// used when a vendorId is not recognized.
func ExampleDataTransferStatusUnknownVendor() {
	status := dt.DataTransferStatusUnknownVendor
	fmt.Println(labelStatus, status.String())
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// Status: UnknownVendor
	// IsValid: true
}

// ExampleDataTransferStatus_IsValid_invalid demonstrates that invalid status
// values are rejected.
func ExampleDataTransferStatus_IsValid_invalid() {
	status := dt.DataTransferStatus("InvalidStatus")
	fmt.Println(labelIsValid, status.IsValid())
	// Output:
	// IsValid: false
}
