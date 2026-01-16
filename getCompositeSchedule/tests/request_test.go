package getCompositeSchedule_test

import (
	"strings"
	"testing"

	gcs "github.com/aasanchez/ocpp16messages/getCompositeSchedule"
	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errConnectorId      = "connectorId"
	errDuration         = "duration"
	errChargingRateUnit = "chargingRateUnit"

	valueZero       = 0
	valueOne        = 1
	valueThreeHund  = 300
	valueSixHund    = 600
	valueNegative   = -1
	valueExceedsMax = 65536

	chargingRateUnitNotNil = "ChargingRateUnit should not be nil"
)

func strPtr(v string) *string {
	return &v
}

func TestReq_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueOne {
		t.Errorf(st.ErrorMismatchValue, valueOne, req.ConnectorId.Value())
	}

	if req.Duration.Value() != valueThreeHund {
		t.Errorf(st.ErrorMismatchValue, valueThreeHund, req.Duration.Value())
	}

	if req.ChargingRateUnit != nil {
		t.Errorf("ChargingRateUnit should be nil, got %v", req.ChargingRateUnit)
	}
}

func TestReq_Valid_ConnectorIdZero(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueZero,
		Duration:         valueThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, req.ConnectorId.Value())
	}
}

func TestReq_Valid_WithChargingRateUnitWatts(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueThreeHund,
		ChargingRateUnit: strPtr("W"),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingRateUnit == nil {
		t.Fatal(chargingRateUnitNotNil)
	}

	if *req.ChargingRateUnit != gt.ChargingRateUnitWatts {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitWatts,
			*req.ChargingRateUnit,
		)
	}
}

func TestReq_Valid_WithChargingRateUnitAmperes(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueSixHund,
		ChargingRateUnit: strPtr("A"),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ChargingRateUnit == nil {
		t.Fatal(chargingRateUnitNotNil)
	}

	if *req.ChargingRateUnit != gt.ChargingRateUnitAmperes {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitAmperes,
			*req.ChargingRateUnit,
		)
	}
}

func TestReq_Invalid_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId: valueNegative,
		Duration:    valueThreeHund,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative ConnectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId: valueExceedsMax,
		Duration:    valueThreeHund,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "ConnectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_NegativeDuration(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId: valueOne,
		Duration:    valueNegative,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative Duration")
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}
}

func TestReq_Invalid_DurationExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId: valueOne,
		Duration:    valueExceedsMax,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Duration exceeds max")
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}
}

func TestReq_Invalid_ChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueThreeHund,
		ChargingRateUnit: strPtr("X"),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid ChargingRateUnit")
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}

func TestReq_Invalid_EmptyChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueThreeHund,
		ChargingRateUnit: strPtr(""),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty ChargingRateUnit")
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}

func TestReq_Invalid_LowercaseChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueOne,
		Duration:         valueThreeHund,
		ChargingRateUnit: strPtr("w"),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase ChargingRateUnit")
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}

func TestReq_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      valueNegative,
		Duration:         valueNegative,
		ChargingRateUnit: strPtr("X"),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}
