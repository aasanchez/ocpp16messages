//go:build race

package race

import (
	"fmt"
	"testing"

	dt "github.com/aasanchez/ocpp16messages/datatransfer"
	st "github.com/aasanchez/ocpp16messages/types"
)

func TestRace_ChargingSchedulePeriodIsolatedFromMutation(t *testing.T) {
	t.Parallel()

	duration := 60
	startSchedule := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	alternatePeriod, err := st.NewChargingSchedulePeriod(
		st.ChargingSchedulePeriodInput{
			StartPeriod: 0,
			Limit:       1,
		},
	)
	if err != nil {
		t.Fatalf("NewChargingSchedulePeriod: %v", err)
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, _ int) error {
		current := schedule.ChargingSchedulePeriod()
		if len(current) == 0 {
			return fmt.Errorf("ChargingSchedulePeriod: expected at least one period")
		}

		if worker%2 == 0 {
			current[0] = alternatePeriod
			return nil
		}

		_ = current[0].StartPeriod().Value()
		_ = current[0].Limit()

		return nil
	})
}

func TestRace_ChargingScheduleMinChargingRateIsolatedFromInputPointer(t *testing.T) {
	t.Parallel()

	duration := 60
	minChargingRate := 1.0
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          nil,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	if schedule.MinChargingRate() == nil {
		t.Fatalf("MinChargingRate: expected non-nil")
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		if worker == 0 {
			minChargingRate = float64(iteration)
			return nil
		}

		value := schedule.MinChargingRate()
		if value == nil {
			return fmt.Errorf("MinChargingRate: expected non-nil")
		}

		_ = *value

		return nil
	})
}

func TestRace_ChargingScheduleMinChargingRateGetterReturnsCopy(t *testing.T) {
	t.Parallel()

	duration := 60
	minChargingRate := 1.0
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: 0, Limit: 16},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          nil,
	})
	if err != nil {
		t.Fatalf("NewChargingSchedule: %v", err)
	}

	const want = 1.0

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		if worker == 0 {
			value := schedule.MinChargingRate()
			if value == nil {
				return fmt.Errorf("MinChargingRate: expected non-nil")
			}

			*value = float64(iteration)

			return nil
		}

		value := schedule.MinChargingRate()
		if value == nil {
			return fmt.Errorf("MinChargingRate: expected non-nil")
		}

		if *value != want {
			return fmt.Errorf("MinChargingRate = %v, want %v", *value, want)
		}

		return nil
	})
}

func TestRace_DataTransferReqDataIsolatedFromInputPointer(t *testing.T) {
	t.Parallel()

	data := "payload"

	message, err := dt.Req(dt.ReqInput{
		VendorId:  "Vendor-1",
		MessageId: nil,
		Data:      &data,
	})
	if err != nil {
		t.Fatalf("DataTransfer.Req: %v", err)
	}

	if message.Data == nil {
		t.Fatalf("DataTransfer.Req: expected Data to be non-nil")
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		if worker == 0 {
			data = fmt.Sprintf("payload-%d", iteration)
			return nil
		}

		_ = *message.Data
		return nil
	})
}

func TestRace_DataTransferConfDataIsolatedFromInputPointer(t *testing.T) {
	t.Parallel()

	data := "payload"

	message, err := dt.Conf(dt.ConfInput{Status: "Accepted", Data: &data})
	if err != nil {
		t.Fatalf("DataTransfer.Conf: %v", err)
	}

	if message.Data == nil {
		t.Fatalf("DataTransfer.Conf: expected Data to be non-nil")
	}

	runConcurrent(t, raceWorkers, raceIterations, func(worker int, iteration int) error {
		if worker == 0 {
			data = fmt.Sprintf("payload-%d", iteration)
			return nil
		}

		_ = *message.Data
		return nil
	})
}
