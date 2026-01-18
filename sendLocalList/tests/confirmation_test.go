package sendLocalList_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != slt.UpdateStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Failed(t *testing.T) {
	t.Parallel()

	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Failed",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != slt.UpdateStatusFailed {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateStatusFailed,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != slt.UpdateStatusNotSupported {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_Valid_VersionMismatch(t *testing.T) {
	t.Parallel()

	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "VersionMismatch",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != slt.UpdateStatusVersionMismatch {
		t.Errorf(
			st.ErrorMismatch,
			slt.UpdateStatusVersionMismatch,
			conf.Status,
		)
	}
}

func TestConf_InvalidStatus_Empty(t *testing.T) {
	t.Parallel()

	_, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "Unknown",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := sendLocalList.Conf(sendLocalList.ConfInput{
		Status: "accepted",
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
