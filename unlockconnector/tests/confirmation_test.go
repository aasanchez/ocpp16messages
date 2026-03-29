package unlockconnector_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/unlockconnector"
	types "github.com/aasanchez/ocpp16types"
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UnlockStatusUnlocked {
		t.Errorf(
			types.ErrorMismatch,
			types.UnlockStatusUnlocked,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UnlockStatusUnlockFailed {
		t.Errorf(
			types.ErrorMismatch,
			types.UnlockStatusUnlockFailed,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UnlockStatusNotSupported {
		t.Errorf(
			types.ErrorMismatch,
			types.UnlockStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "unlocked",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Accepted(t *testing.T) {
	t.Parallel()

	_, err := unlockconnector.Conf(unlockconnector.ConfInput{
		Status: "Accepted",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Accepted (invalid for UnlockConnector)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
