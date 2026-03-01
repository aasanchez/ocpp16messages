//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	gc "github.com/aasanchez/ocpp16messages/getCompositeSchedule"
	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetCompositeScheduleConf(f *testing.F) {
	f.Add(
		gt.GetCompositeScheduleStatusAccepted.String(),
		false,
		0,
		false,
		"",
		false,
		"",
	)
	f.Add(
		gt.GetCompositeScheduleStatusRejected.String(),
		true,
		1,
		true,
		"2025-01-15T10:30:00Z",
		true,
		st.ChargingRateUnitWatts.String(),
	)
	f.Add(
		"invalid-status",
		false,
		0,
		false,
		"",
		false,
		"",
	)
	f.Add(
		gt.GetCompositeScheduleStatusAccepted.String(),
		true,
		-1,
		true,
		"bad-timestamp",
		true,
		"invalid-unit",
	)

	f.Fuzz(func(
		t *testing.T,
		status string,
		hasConnectorId bool,
		connectorId int,
		hasScheduleStart bool,
		scheduleStart string,
		hasChargingSchedule bool,
		chargingRateUnit string,
	) {
		if len(status) > maxFuzzStringLen ||
			len(scheduleStart) > maxFuzzStringLen ||
			len(chargingRateUnit) > maxFuzzStringLen {
			t.Skip()
		}

		var connectorIdPtr *int
		if hasConnectorId {
			connectorIdPtr = &connectorId
		}

		var scheduleStartPtr *string
		if hasScheduleStart {
			scheduleStartPtr = &scheduleStart
		}

		var chargingSchedulePtr *st.ChargingScheduleInput
		if hasChargingSchedule {
			chargingSchedule := st.ChargingScheduleInput{
				ChargingRateUnit: chargingRateUnit,
				ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
					{
						StartPeriod:  0,
						Limit:        0,
						NumberPhases: nil,
					},
				},
			}
			chargingSchedulePtr = &chargingSchedule
		}

		conf, err := gc.Conf(gc.ConfInput{
			Status:           status,
			ConnectorId:      connectorIdPtr,
			ScheduleStart:    scheduleStartPtr,
			ChargingSchedule: chargingSchedulePtr,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}

		if hasConnectorId {
			if conf.ConnectorId == nil {
				t.Fatal("ConnectorId = nil, want non-nil")
			}
		} else if conf.ConnectorId != nil {
			t.Fatal("ConnectorId != nil, want nil")
		}

		if hasScheduleStart {
			if conf.ScheduleStart == nil {
				t.Fatal("ScheduleStart = nil, want non-nil")
			}
			if conf.ScheduleStart.Value().Location() != time.UTC {
				t.Fatalf(
					"ScheduleStart location = %v, want UTC",
					conf.ScheduleStart.Value().Location(),
				)
			}
		} else if conf.ScheduleStart != nil {
			t.Fatal("ScheduleStart != nil, want nil")
		}

		if hasChargingSchedule {
			if conf.ChargingSchedule == nil {
				t.Fatal("ChargingSchedule = nil, want non-nil")
			}
			if len(conf.ChargingSchedule.ChargingSchedulePeriod()) == 0 {
				t.Fatal("ChargingSchedulePeriod is empty, want at least one")
			}
		} else if conf.ChargingSchedule != nil {
			t.Fatal("ChargingSchedule != nil, want nil")
		}
	})
}
