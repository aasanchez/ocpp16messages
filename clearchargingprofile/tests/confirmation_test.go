package clearchargingprofile_test

import (
	"strings"
	"testing"

	ccp "github.com/aasanchez/ocpp16messages/clearchargingprofile"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := ccp.Conf(ccp.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ClearChargingProfileStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.ClearChargingProfileStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Unknown(t *testing.T) {
	t.Parallel()

	conf, err := ccp.Conf(ccp.ConfInput{Status: "Unknown"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ClearChargingProfileStatusUnknown {
		t.Errorf(
			types.ErrorMismatch,
			types.ClearChargingProfileStatusUnknown,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Invalid(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "Invalid"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Rejected(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "Rejected"})
	if err == nil {
		t.Errorf(
			types.ErrorWantNil, "Rejected (invalid for ClearChargingProfile)",
		)
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
