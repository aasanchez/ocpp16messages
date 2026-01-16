package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
)

// ExampleGetCompositeScheduleStatus demonstrates using the
// GetCompositeScheduleStatus enumeration constants defined in OCPP 1.6.
func ExampleGetCompositeScheduleStatus() {
	status := types.GetCompositeScheduleStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleGetCompositeScheduleStatus_allValues demonstrates all valid
// GetCompositeScheduleStatus values as defined in OCPP 1.6.
func ExampleGetCompositeScheduleStatus_allValues() {
	statuses := []types.GetCompositeScheduleStatus{
		types.GetCompositeScheduleStatusAccepted,
		types.GetCompositeScheduleStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Rejected
}

// ExampleGetCompositeScheduleStatus_IsValid demonstrates validation of
// GetCompositeScheduleStatus values.
func ExampleGetCompositeScheduleStatus_IsValid() {
	valid := types.GetCompositeScheduleStatusAccepted
	invalid := types.GetCompositeScheduleStatus("Invalid")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Invalid IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Invalid IsValid: false
}
