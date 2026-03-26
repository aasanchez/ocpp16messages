package unlockconnector_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
	"github.com/aasanchez/ocpp16messages/unlockconnector"
	ut "github.com/aasanchez/ocpp16messages/unlockconnector/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Unlocked(t *testing.T) {
	t.Parallel()

	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "Unlocked",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != ut.UnlockStatusUnlocked {
		t.Errorf(
			st.ErrorMismatch,
			ut.UnlockStatusUnlocked,
			conf.Status,
		)
	}
}

func TestConf_Valid_UnlockFailed(t *testing.T) {
	t.Parallel()

	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "UnlockFailed",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != ut.UnlockStatusUnlockFailed {
		t.Errorf(
			st.ErrorMismatch,
			ut.UnlockStatusUnlockFailed,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != ut.UnlockStatusNotSupported {
		t.Errorf(
			st.ErrorMismatch,
			ut.UnlockStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "unlocked",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Accepted(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "Accepted",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Accepted (invalid for UnlockConnector)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
