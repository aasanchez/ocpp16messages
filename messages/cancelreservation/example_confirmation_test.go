package cancelreservation_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/cancelreservation"
	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func ExampleConfirmation_valid() {
	input := cancelreservationtypes.ConfirmationPayload{Status: cancelreservationtypes.CancelReservationAccepted}

	msg, err := cancelreservation.Confirmation(input)
	if err != nil {
		fmt.Println("unexpected error:", err)

		return
	}

	fmt.Printf("Status: %s\n", msg.Status.Value())

	// Output:
	// Status: Accepted
}

func ExampleConfirmation_invalid() {
	input := cancelreservationtypes.ConfirmationPayload{Status: "NotAStatus"}

	_, err := cancelreservation.Confirmation(input)
	if err != nil {
		fmt.Println("error:", err != nil)
	}

	// Output:
	// error: true
}
