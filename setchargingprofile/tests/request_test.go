package setchargingprofile_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/setchargingprofile"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errConnectorId        = "connectorId"
	errCsChargingProfiles = "csChargingProfiles"

	valueZero        = 0
	valueOne         = 1
	valueTwo         = 2
	valueNegative    = -1
	valueExceedsMax  = 65536
	valueLimitThirty = 30.0
)

func validChargingProfileInput() types.ChargingProfileInput {
	return types.ChargingProfileInput{
		ChargingProfileId:      valueOne,
		TransactionId:          nil,
		StackLevel:             valueZero,
		ChargingProfilePurpose: "TxDefaultProfile",
		ChargingProfileKind:    "Absolute",
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: types.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: "W",
			ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
				{
					StartPeriod:  valueZero,
					Limit:        valueLimitThirty,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	}
}

func TestReq_Valid_MinimalInput(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ConnectorId.Value())
	}
}

func TestReq_Valid_WithConnectorId(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueTwo,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueTwo {
		t.Errorf(types.ErrorMismatchValue, valueTwo, req.ConnectorId.Value())
	}
}

func TestReq_Valid_ChargingProfileId(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.CsChargingProfiles.ChargingProfileId().Value() != valueOne {
		t.Errorf(
			types.ErrorMismatchValue,
			valueOne,
			req.CsChargingProfiles.ChargingProfileId().Value(),
		)
	}
}

func TestReq_Valid_StackLevel(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.CsChargingProfiles.StackLevel().Value() != valueZero {
		t.Errorf(
			types.ErrorMismatchValue,
			valueZero,
			req.CsChargingProfiles.StackLevel().Value(),
		)
	}
}

func TestReq_Valid_ChargingProfilePurpose(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	purpose := req.CsChargingProfiles.ChargingProfilePurpose()
	if purpose != types.TxDefaultProfile {
		t.Errorf(
			types.ErrorMismatch,
			types.TxDefaultProfile,
			purpose,
		)
	}
}

func TestReq_Valid_ChargingProfileKind(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.CsChargingProfiles.ChargingProfileKind() !=
		types.ChargingProfileKindAbsolute {
		t.Errorf(
			types.ErrorMismatch,
			types.ChargingProfileKindAbsolute,
			req.CsChargingProfiles.ChargingProfileKind(),
		)
	}
}

func TestReq_Invalid_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueNegative,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative ConnectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueExceedsMax,
		CsChargingProfiles: validChargingProfileInput(),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "ConnectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_InvalidChargingProfilePurpose(t *testing.T) {
	t.Parallel()

	input := validChargingProfileInput()
	input.ChargingProfilePurpose = "Invalid"

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: input,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errCsChargingProfiles) {
		t.Errorf(types.ErrorWantContains, err, errCsChargingProfiles)
	}
}

func TestReq_Invalid_InvalidChargingProfileKind(t *testing.T) {
	t.Parallel()

	input := validChargingProfileInput()
	input.ChargingProfileKind = "Invalid"

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: input,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid ChargingProfileKind")
	}

	if !strings.Contains(err.Error(), errCsChargingProfiles) {
		t.Errorf(types.ErrorWantContains, err, errCsChargingProfiles)
	}
}

func TestReq_Invalid_InvalidChargingRateUnit(t *testing.T) {
	t.Parallel()

	input := validChargingProfileInput()
	input.ChargingSchedule.ChargingRateUnit = "X"

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: input,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid ChargingRateUnit")
	}

	if !strings.Contains(err.Error(), errCsChargingProfiles) {
		t.Errorf(types.ErrorWantContains, err, errCsChargingProfiles)
	}
}

func TestReq_Invalid_EmptyChargingSchedulePeriod(t *testing.T) {
	t.Parallel()

	input := validChargingProfileInput()
	input.ChargingSchedule.ChargingSchedulePeriod = nil

	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId:        valueZero,
		CsChargingProfiles: input,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty ChargingSchedulePeriod")
	}

	if !strings.Contains(err.Error(), errCsChargingProfiles) {
		t.Errorf(types.ErrorWantContains, err, errCsChargingProfiles)
	}
}
