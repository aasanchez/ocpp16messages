package cancelReservation_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/cancelReservation"
)

const reservationIdLabel = "ReservationId:"

// ExampleReq demonstrates creating a valid CancelReservation.req message
// with a reservation ID.
func ExampleReq() {
	req, err := cancelReservation.Req(cancelReservation.ReqInput{
		ReservationId: 123,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(reservationIdLabel, req.ReservationId.Value())
	// Output:
	// ReservationId: 123
}

// ExampleReq_negativeId demonstrates the error returned when
// a negative reservation ID is provided.
func ExampleReq_negativeId() {
	_, err := cancelReservation.Req(cancelReservation.ReqInput{
		ReservationId: -1,
	})
	if err != nil {
		fmt.Println("Error: invalid reservation ID")
	}
	// Output:
	// Error: invalid reservation ID
}

// ExampleReq_zeroId demonstrates that zero is a valid reservation ID.
func ExampleReq_zeroId() {
	req, err := cancelReservation.Req(cancelReservation.ReqInput{
		ReservationId: 0,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(reservationIdLabel, req.ReservationId.Value())
	// Output:
	// ReservationId: 0
}

// ExampleReq_maxValue demonstrates the maximum valid reservation ID (65535).
func ExampleReq_maxValue() {
	req, err := cancelReservation.Req(cancelReservation.ReqInput{
		ReservationId: 65535,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(reservationIdLabel, req.ReservationId.Value())
	// Output:
	// ReservationId: 65535
}

// ExampleReq_exceedsMax demonstrates the error returned when
// the reservation ID exceeds the maximum value (65535).
func ExampleReq_exceedsMax() {
	_, err := cancelReservation.Req(cancelReservation.ReqInput{
		ReservationId: 65536,
	})
	if err != nil {
		fmt.Println("Error: reservation ID exceeds maximum")
	}
	// Output:
	// Error: reservation ID exceeds maximum
}
