package clearChargingProfile_test

import (
	"strings"
	"testing"

	ccp "github.com/aasanchez/ocpp16messages/clearChargingProfile"
	ct "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := ccp.Conf(ccp.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != ct.ClearChargingProfileStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			ct.ClearChargingProfileStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Unknown(t *testing.T) {
	t.Parallel()

	conf, err := ccp.Conf(ccp.ConfInput{Status: "Unknown"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != ct.ClearChargingProfileStatusUnknown {
		t.Errorf(
			st.ErrorMismatch,
			ct.ClearChargingProfileStatusUnknown,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Invalid(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "Invalid"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Rejected(t *testing.T) {
	t.Parallel()

	_, err := ccp.Conf(ccp.ConfInput{Status: "Rejected"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Rejected (invalid for ClearChargingProfile)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
