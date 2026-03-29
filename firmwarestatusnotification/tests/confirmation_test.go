package firmwarestatusnotification_test

import (
	"testing"

	fsn "github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
	types "github.com/aasanchez/ocpp16types"
)

const repeatCount = 3

func TestConf_Valid(t *testing.T) {
	t.Parallel()

	_, err := fsn.Conf(fsn.ConfInput{})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}
}

func TestConf_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	// Call multiple times to ensure it always succeeds
	for range repeatCount {
		_, err := fsn.Conf(fsn.ConfInput{})
		if err != nil {
			t.Errorf(types.ErrorUnexpectedError, err)
		}
	}
}
