package cancelreservation_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/cancelreservation"
	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func ExampleRequest_valid() {
	input := cancelreservationtypes.RequestPayload{
		ReservationId: "100",
	}

	msg, err := cancelreservation.Request(input)
	if err != nil {
		fmt.Println("unexpected error:", err)

		return
	}

	fmt.Printf("ReservationId: %d\n", msg.ReservationId)

	// Output:
	// ReservationId: 100
}

func ExampleRequest_invalid() {
	input := cancelreservationtypes.RequestPayload{
		ReservationId: "not-a-number",
	}

	_, err := cancelreservation.Request(input)
	if err != nil {
		fmt.Println("error:", err != nil)
	}

	// Output:
	// error: true
}
