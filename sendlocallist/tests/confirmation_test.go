package sendlocallist_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/sendlocallist"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UpdateStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Failed(t *testing.T) {
	t.Parallel()

	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Failed",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UpdateStatusFailed {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateStatusFailed,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UpdateStatusNotSupported {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_Valid_VersionMismatch(t *testing.T) {
	t.Parallel()

	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "VersionMismatch",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.UpdateStatusVersionMismatch {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateStatusVersionMismatch,
			conf.Status,
		)
	}
}

func TestConf_InvalidStatus_Empty(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "Unknown",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Conf(sendlocallist.ConfInput{
		Status: "accepted",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
