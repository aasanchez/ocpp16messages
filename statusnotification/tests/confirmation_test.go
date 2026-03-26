package statusnotification_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/statusnotification"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestConf_Valid(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Conf(statusnotification.ConfInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}
}

func TestConf_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	// StatusNotification.conf has no fields, so it should always succeed
	conf, err := statusnotification.Conf(statusnotification.ConfInput{})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	// Verify the message type is returned
	_ = conf
}
