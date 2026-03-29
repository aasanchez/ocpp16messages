package diagnosticsstatusnotification_test

import (
	"testing"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsstatusnotification"
	types "github.com/aasanchez/ocpp16types"
)

const repeatCount = 3

func TestConf_Valid(t *testing.T) {
	t.Parallel()

	_, err := dsn.Conf(dsn.ConfInput{})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}
}

func TestConf_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	// Call multiple times to ensure it always succeeds
	for range repeatCount {
		_, err := dsn.Conf(dsn.ConfInput{})
		if err != nil {
			t.Errorf(types.ErrorUnexpectedError, err)
		}
	}
}
