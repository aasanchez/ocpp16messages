package bootnotification_test

import (
	"testing"

	bn "github.com/aasanchez/ocpp16messages/messages/bootnotification"
	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func validPayload() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2024-08-08T08:08:08Z",
		Interval:    "42",
		Status:      "Accepted",
	}
}

func BenchmarkConfirmation(b *testing.B) {
	input := validPayload()
	for range b.N {
		_, err := bn.Confirmation(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
