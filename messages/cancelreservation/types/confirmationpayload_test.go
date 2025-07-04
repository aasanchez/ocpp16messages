package cancelreservationtypes

import (
	"strings"
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestConfirmationPayload_Validate_OK(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{Status: "Accepted"}

	err := payload.Validate()
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}
}

func TestConfirmationPayload_Validate_EmptyStatus(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{Status: ""}

	err := payload.Validate()
	if err == nil {
		t.Fatalf(sharedtypes.ErrExpectedError, err)
	}

	if !strings.Contains(err.Error(), "missing required field: status") {
		t.Errorf(sharedtypes.ErrExpectedNoError, err)
	}
}
