package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/cancelReservation/types"
)

// ExampleCancelReservationStatus demonstrates using the CancelReservationStatus
// enumeration constants defined in OCPP 1.6.
func ExampleCancelReservationStatus() {
	status := types.CancelReservationStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleCancelReservationStatus_allValues demonstrates all valid
// CancelReservationStatus values as defined in OCPP 1.6.
func ExampleCancelReservationStatus_allValues() {
	statuses := []types.CancelReservationStatus{
		types.CancelReservationStatusAccepted,
		types.CancelReservationStatusRejected,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Rejected
}

// ExampleCancelReservationStatus_IsValid demonstrates validation of
// CancelReservationStatus values.
func ExampleCancelReservationStatus_IsValid() {
	valid := types.CancelReservationStatusAccepted
	invalid := types.CancelReservationStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
