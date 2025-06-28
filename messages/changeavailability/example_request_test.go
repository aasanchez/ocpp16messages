package changeavailability_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/changeavailability"
	changeavailabilitytypes "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func ExampleRequest() {
	input := changeavailabilitytypes.RequestPayload{ConnectorId: "1", Type: changeavailabilitytypes.AvailabilityTypeOperative}
	msg, err := changeavailability.Request(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("ConnectorId:", msg.ConnectorId)
	fmt.Println("Type:", msg.Type.Value())
	// Output:
	// ConnectorId: 1
	// Type: Operative
}
