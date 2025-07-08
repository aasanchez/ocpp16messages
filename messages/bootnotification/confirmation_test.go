package bootnotification

import (
	"testing"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func validPayload() bootnotificationtypes.ConfirmationPayload {
	return bootnotificationtypes.ConfirmationPayload{
		CurrentTime: "2024-08-08T08:08:08Z",
		Interval:    "42",
		Status:      "Accepted",
	}
}

func TestConfirmation(t *testing.T) {
	t.Parallel()

	input := validPayload()

	conf, err := Confirmation(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if conf.Interval != 42 {
		t.Errorf("expected interval 42, got %d", conf.Interval)
	}

	if got := conf.Status.Value(); got != "Accepted" {
		t.Errorf("expected status 'Accepted', got %q", got)
	}

	if want := input.CurrentTime; conf.CurrentTime.String() != want {
		t.Errorf("expected time %q, got %q", want, conf.CurrentTime.String())
	}
}

func TestConfirmation_MissingCurrentTime(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.CurrentTime = ""

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected error for missing CurrentTime, got nil")
	}
}

func TestConfirmation_MissingInterval(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.Interval = ""

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected error for missing Interval, got nil")
	}
}

func TestConfirmation_MissingStatus(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.Status = ""

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected error for missing Status, got nil")
	}
}

func TestConfirmation_InvalidDateTime(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.CurrentTime = "notadatetime"

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected failure for invalid datetime, got nil")
	}
}

func TestConfirmation_InvalidInterval(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.Interval = "notanint"

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected failure for invalid interval, got nil")
	}
}

func TestConfirmation_InvalidStatus(t *testing.T) {
	t.Parallel()

	input := validPayload()
	input.Status = "notastatus"

	_, err := Confirmation(input)
	if err == nil {
		t.Error("expected failure for invalid status, got nil")
	}
}
