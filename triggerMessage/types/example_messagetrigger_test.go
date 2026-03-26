package types_test

import (
	"fmt"

	tt "github.com/aasanchez/ocpp16messages/triggermessage/types"
)

const (
	labelTrigger = "Trigger:"
	labelValid   = "Valid:"
)

// ExampleMessageTriggerBootNotification demonstrates using the BootNotification
// message trigger constant.
func ExampleMessageTriggerBootNotification() {
	trigger := tt.MessageTriggerBootNotification
	fmt.Println(labelTrigger, trigger.String())
	fmt.Println(labelValid, trigger.IsValid())
	// Output:
	// Trigger: BootNotification
	// Valid: true
}

// ExampleMessageTriggerHeartbeat demonstrates using the Heartbeat
// message trigger constant.
func ExampleMessageTriggerHeartbeat() {
	trigger := tt.MessageTriggerHeartbeat
	fmt.Println(labelTrigger, trigger.String())
	fmt.Println(labelValid, trigger.IsValid())
	// Output:
	// Trigger: Heartbeat
	// Valid: true
}

// ExampleMessageTriggerMeterValues demonstrates using the MeterValues
// message trigger constant.
func ExampleMessageTriggerMeterValues() {
	trigger := tt.MessageTriggerMeterValues
	fmt.Println(labelTrigger, trigger.String())
	fmt.Println(labelValid, trigger.IsValid())
	// Output:
	// Trigger: MeterValues
	// Valid: true
}

// ExampleMessageTriggerStatusNotification demonstrates using the
// StatusNotification message trigger constant.
func ExampleMessageTriggerStatusNotification() {
	trigger := tt.MessageTriggerStatusNotification
	fmt.Println(labelTrigger, trigger.String())
	fmt.Println(labelValid, trigger.IsValid())
	// Output:
	// Trigger: StatusNotification
	// Valid: true
}
