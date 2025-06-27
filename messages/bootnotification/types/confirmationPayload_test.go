package bootnotificationtypes

import "testing"

func TestConfirmationPayload_Validate_AllFields(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		CurrentTime: "2026-01-01T15:04:05Z",
		Interval:    "300",
		Status:      "Accepted",
	}

	err := payload.Validate()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConfirmationPayload_Validate_MissingCurrentTime(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		CurrentTime: "",
		Interval:    "300",
		Status:      "Accepted",
	}

	err := payload.Validate()
	if err == nil || err.Error() != "confirmation payload: missing required field: CurrentTime" {
		t.Errorf("expected missing CurrentTime error, got %q", err)
	}
}

func TestConfirmationPayload_Validate_MissingInterval(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		CurrentTime: "2025-01-01T15:04:05Z",
		Interval:    "",
		Status:      "Accepted",
	}

	err := payload.Validate()
	if err == nil || err.Error() != "confirmation payload: missing required field: Interval" {
		t.Errorf("expected missing Interval error, got %q", err)
	}
}

func TestConfirmationPayload_Validate_MissingStatus(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		CurrentTime: "2025-04-01T15:04:05Z",
		Interval:    "300",
		Status:      "",
	}

	err := payload.Validate()
	if err == nil || err.Error() != "confirmation payload: missing required field: Status" {
		t.Errorf("expected missing Status error, got %q", err)
	}
}
