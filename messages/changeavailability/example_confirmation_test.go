package changeavailability_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/changeavailability"
	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func ExampleConfirmation() {
	input := cat.ConfirmationPayload{Status: cat.ChangeAvailabilityStatusAccepted}

	msg, err := changeavailability.Confirmation(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Status:", msg.Status.Value())

	// Output:
	// Status: Accepted
}
