package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/reservenow/types"
)

// ExampleReservationStatus demonstrates using the ReservationStatus enumeration
// constants defined in OCPP 1.6.
func ExampleReservationStatus() {
	status := types.ReservationStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleReservationStatus_allValues demonstrates all valid ReservationStatus
// values as defined in OCPP 1.6.
func ExampleReservationStatus_allValues() {
	statuses := []types.ReservationStatus{
		types.ReservationStatusAccepted,
		types.ReservationStatusFaulted,
		types.ReservationStatusOccupied,
		types.ReservationStatusRejected,
		types.ReservationStatusUnavailable,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Faulted
	// Occupied
	// Rejected
	// Unavailable
}

// ExampleReservationStatus_IsValid demonstrates validation of ReservationStatus
// values.
func ExampleReservationStatus_IsValid() {
	valid := types.ReservationStatusAccepted
	invalid := types.ReservationStatus("Unknown")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Unknown IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Unknown IsValid: false
}
