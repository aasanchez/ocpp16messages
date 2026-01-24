package types_test

import (
	"fmt"

	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
)

const (
	labelStatusTMS = "Status:"
	labelValidTMS  = "Valid:"
)

// ExampleTriggerMessageStatusAccepted demonstrates using the Accepted
// trigger message status constant.
func ExampleTriggerMessageStatusAccepted() {
	status := tt.TriggerMessageStatusAccepted
	fmt.Println(labelStatusTMS, status.String())
	fmt.Println(labelValidTMS, status.IsValid())
	// Output:
	// Status: Accepted
	// Valid: true
}

// ExampleTriggerMessageStatusRejected demonstrates using the Rejected
// trigger message status constant.
func ExampleTriggerMessageStatusRejected() {
	status := tt.TriggerMessageStatusRejected
	fmt.Println(labelStatusTMS, status.String())
	fmt.Println(labelValidTMS, status.IsValid())
	// Output:
	// Status: Rejected
	// Valid: true
}

// ExampleTriggerMessageStatusNotImplemented demonstrates using the
// NotImplemented trigger message status constant.
func ExampleTriggerMessageStatusNotImplemented() {
	status := tt.TriggerMessageStatusNotImplemented
	fmt.Println(labelStatusTMS, status.String())
	fmt.Println(labelValidTMS, status.IsValid())
	// Output:
	// Status: NotImplemented
	// Valid: true
}
