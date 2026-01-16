package getCompositeSchedule_test

import (
	"strings"
	"testing"

	gcs "github.com/aasanchez/ocpp16messages/getCompositeSchedule"
	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus           = "status"
	errConnectorIdConf  = "connectorId"
	errScheduleStart    = "scheduleStart"
	errChargingSchedule = "chargingSchedule"

	validTimestamp    = "2025-01-15T10:00:00Z"
	invalidTimestamp  = "invalid-timestamp"
	durationConfValue = 3600
	limitConfValue    = 32.0

	connectorIdNotNil      = "ConnectorId should not be nil"
	scheduleStartNotNil    = "ScheduleStart should not be nil"
	chargingScheduleNotNil = "ChargingSchedule should not be nil"
)

func intPtr(v int) *int {
	return &v
}

func TestConf_Valid_AcceptedOnly(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != gt.GetCompositeScheduleStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			gt.GetCompositeScheduleStatusAccepted,
			conf.Status,
		)
	}

	if conf.ConnectorId != nil {
		t.Errorf("ConnectorId should be nil, got %v", conf.ConnectorId)
	}

	if conf.ScheduleStart != nil {
		t.Errorf("ScheduleStart should be nil, got %v", conf.ScheduleStart)
	}

	if conf.ChargingSchedule != nil {
		t.Errorf(
			"ChargingSchedule should be nil, got %v",
			conf.ChargingSchedule,
		)
	}
}

func TestConf_Valid_RejectedOnly(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Rejected",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != gt.GetCompositeScheduleStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			gt.GetCompositeScheduleStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_WithConnectorId(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      intPtr(valueOne),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ConnectorId == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ConnectorId.Value() != valueOne {
		t.Errorf(st.ErrorMismatchValue, valueOne, conf.ConnectorId.Value())
	}
}

func TestConf_Valid_WithConnectorIdZero(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      intPtr(valueZero),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ConnectorId == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ConnectorId.Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, conf.ConnectorId.Value())
	}
}

func TestConf_Valid_WithScheduleStart(t *testing.T) {
	t.Parallel()

	scheduleStart := validTimestamp

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      nil,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ScheduleStart == nil {
		t.Fatal(scheduleStartNotNil)
	}
}

func TestConf_Valid_WithChargingSchedule(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorId:   nil,
		ScheduleStart: nil,
		ChargingSchedule: &gt.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: "W",
			ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
				{
					StartPeriod:  valueZero,
					Limit:        limitConfValue,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ChargingSchedule == nil {
		t.Fatal(chargingScheduleNotNil)
	}

	if conf.ChargingSchedule.ChargingRateUnit() != gt.ChargingRateUnitWatts {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitWatts,
			conf.ChargingSchedule.ChargingRateUnit(),
		)
	}
}

func TestConf_Valid_WithAllFields(t *testing.T) {
	t.Parallel()

	scheduleStart := validTimestamp
	duration := durationConfValue

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorId:   intPtr(valueOne),
		ScheduleStart: &scheduleStart,
		ChargingSchedule: &gt.ChargingScheduleInput{
			Duration:         &duration,
			ChargingRateUnit: "A",
			ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
				{
					StartPeriod:  valueZero,
					Limit:        limitConfValue,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != gt.GetCompositeScheduleStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			gt.GetCompositeScheduleStatusAccepted,
			conf.Status,
		)
	}

	if conf.ConnectorId == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ScheduleStart == nil {
		t.Fatal(scheduleStartNotNil)
	}

	if conf.ChargingSchedule == nil {
		t.Fatal(chargingScheduleNotNil)
	}
}

func TestConf_Invalid_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Invalid",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "accepted",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      intPtr(valueNegative),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative ConnectorId")
	}

	if !strings.Contains(err.Error(), errConnectorIdConf) {
		t.Errorf(st.ErrorWantContains, err, errConnectorIdConf)
	}
}

func TestConf_Invalid_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      intPtr(valueExceedsMax),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "ConnectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorIdConf) {
		t.Errorf(st.ErrorWantContains, err, errConnectorIdConf)
	}
}

func TestConf_Invalid_InvalidScheduleStart(t *testing.T) {
	t.Parallel()

	scheduleStart := invalidTimestamp

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      nil,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid scheduleStart")
	}

	if !strings.Contains(err.Error(), errScheduleStart) {
		t.Errorf(st.ErrorWantContains, err, errScheduleStart)
	}
}

func TestConf_Invalid_InvalidChargingSchedule(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorId:   nil,
		ScheduleStart: nil,
		ChargingSchedule: &gt.ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid chargingSchedule")
	}

	if !strings.Contains(err.Error(), errChargingSchedule) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedule)
	}
}

func TestConf_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	scheduleStart := invalidTimestamp

	_, err := gcs.Conf(gcs.ConfInput{
		Status:        "Invalid",
		ConnectorId:   intPtr(valueNegative),
		ScheduleStart: &scheduleStart,
		ChargingSchedule: &gt.ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}

	if !strings.Contains(err.Error(), errConnectorIdConf) {
		t.Errorf(st.ErrorWantContains, err, errConnectorIdConf)
	}

	if !strings.Contains(err.Error(), errScheduleStart) {
		t.Errorf(st.ErrorWantContains, err, errScheduleStart)
	}

	if !strings.Contains(err.Error(), errChargingSchedule) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedule)
	}
}
