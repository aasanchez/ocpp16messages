package bootnotification_test

import (
	"testing"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func validPayload() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2025-08-08T08:08:08Z",
		Interval:    "42",
		Status:      "Accepted",
	}
}

func invalidPayloadMissingCurrentTime() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "",
		Interval:    "42",
		Status:      "Accepted",
	}
}

func invalidPayloadMissingInterval() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2026-08-08T08:08:08Z",
		Interval:    "",
		Status:      "Accepted",
	}
}

func invalidPayloadMissingStatus() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2024-08-08T08:08:08Z",
		Interval:    "42",
		Status:      "",
	}
}

func BenchmarkConfirmationPayload_Validate(b *testing.B) {
	payload := validPayload()
	for range b.N {
		if err := payload.Validate(); err != nil {
			b.Fatalf("unexpected error for valid: %v", err)
		}
	}
}

func BenchmarkConfirmationPayload_Validate_MissingCurrentTime(b *testing.B) {
	payload := invalidPayloadMissingCurrentTime()
	for range b.N {
		if err := payload.Validate(); err == nil {
			b.Fatalf("expected error for missing current time, got nil")
		}
	}
}

func BenchmarkConfirmationPayload_Validate_MissingInterval(b *testing.B) {
	payload := invalidPayloadMissingInterval()
	for range b.N {
		if err := payload.Validate(); err == nil {
			b.Fatalf("expected error for missing interval, got nil")
		}
	}
}

func BenchmarkConfirmationPayload_Validate_MissingStatus(b *testing.B) {
	payload := invalidPayloadMissingStatus()
	for range b.N {
		if err := payload.Validate(); err == nil {
			b.Fatalf("expected error for missing status, got nil")
		}
	}
}

func BenchmarkConfirmationPayload_Validate_MixedParallel(b *testing.B) {
	allPayloads := []bootnotificationtypes.ConfirmationPayload{
		validPayload(),
		invalidPayloadMissingCurrentTime(),
		invalidPayloadMissingInterval(),
		invalidPayloadMissingStatus(),
	}

	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		idx := 0
		for pb.Next() {
			pl := allPayloads[idx%len(allPayloads)]
			_ = pl.Validate()
			idx++
		}
	})
}
