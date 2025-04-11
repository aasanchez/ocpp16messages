package bootnotification_test

import (
	"fmt"

	bn "github.com/aasanchez/ocpp16messages/messages/bootnotification"
	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func ExampleConfirmation() {
	input := bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2024-08-08T08:08:08Z",
		Interval:    "42",
		Status:      "Accepted",
	}

	conf, err := bn.Confirmation(input)
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println(conf.Interval)
	fmt.Println(conf.Status.Value())
	fmt.Println(conf.CurrentTime.String())
	// Output:
	// 42
	// Accepted
	// 2024-08-08T08:08:08Z
}

func ExampleConfirmation_invalid() {
	input := bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "", // Missing field
		Interval:    "42",
		Status:      "Accepted",
	}

	_, err := bn.Confirmation(input)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Output:
	// error: bootnotificationtypes.Confirmation: invalid payload: confirmation payload: missing required field: CurrentTime
}
