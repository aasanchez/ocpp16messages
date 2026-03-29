package setchargingprofile_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/setchargingprofile"
	types "github.com/aasanchez/ocpp16types"
)

const errStatus = "status"

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ChargingProfileStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.ChargingProfileStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ChargingProfileStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.ChargingProfileStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "NotSupported",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ChargingProfileStatusNotSupported {
		t.Errorf(
			types.ErrorMismatch,
			types.ChargingProfileStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_Invalid_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_UnknownStatus(t *testing.T) {
	t.Parallel()

	_, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Unknown",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "accepted",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
