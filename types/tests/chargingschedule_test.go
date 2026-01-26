package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	testScheduleDuration    = 3600
	testScheduleStartTime   = "2025-01-01T00:00:00Z"
	testScheduleRateUnit    = "A"
	testScheduleStartPeriod = 0
	testScheduleLimit       = 32.0
	testScheduleMinRate     = 6.0
)

func createValidChargingSchedule(t *testing.T) types.ChargingSchedule {
	t.Helper()

	duration := testScheduleDuration
	startSchedule := testScheduleStartTime
	minChargingRate := testScheduleMinRate

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         &duration,
		StartSchedule:    &startSchedule,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
		MinChargingRate: &minChargingRate,
	})
	if err != nil {
		t.Fatalf("Failed to create valid ChargingSchedule: %v", err)
	}

	return schedule
}

func TestChargingSchedule_Duration(t *testing.T) {
	t.Parallel()

	schedule := createValidChargingSchedule(t)
	duration := schedule.Duration()

	if duration == nil {
		t.Error("ChargingSchedule.Duration() = nil, want non-nil")
	}
}

func TestChargingSchedule_Duration_WhenNil(t *testing.T) {
	t.Parallel()

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedule: %v", err)
	}

	if schedule.Duration() != nil {
		t.Error("ChargingSchedule.Duration() = non-nil, want nil")
	}
}

func TestChargingSchedule_StartSchedule(t *testing.T) {
	t.Parallel()

	schedule := createValidChargingSchedule(t)
	startSchedule := schedule.StartSchedule()

	if startSchedule == nil {
		t.Error("ChargingSchedule.StartSchedule() = nil, want non-nil")
	}
}

func TestChargingSchedule_StartSchedule_WhenNil(t *testing.T) {
	t.Parallel()

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		StartSchedule:    nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedule: %v", err)
	}

	if schedule.StartSchedule() != nil {
		t.Error("ChargingSchedule.StartSchedule() = non-nil, want nil")
	}
}

func TestChargingSchedule_MinChargingRate(t *testing.T) {
	t.Parallel()

	schedule := createValidChargingSchedule(t)
	minRate := schedule.MinChargingRate()

	if minRate == nil {
		t.Error("ChargingSchedule.MinChargingRate() = nil, want non-nil")
	}
}

func TestChargingSchedule_MinChargingRate_WhenNil(t *testing.T) {
	t.Parallel()

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
		MinChargingRate: nil,
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedule: %v", err)
	}

	if schedule.MinChargingRate() != nil {
		t.Error("ChargingSchedule.MinChargingRate() = non-nil, want nil")
	}
}

func TestNewChargingSchedule_InvalidDuration(t *testing.T) {
	t.Parallel()

	duration := -1
	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for negative duration")
	}
}

func TestNewChargingSchedule_InvalidStartSchedule(t *testing.T) {
	t.Parallel()

	invalidTime := "not-a-valid-time"
	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		StartSchedule:    &invalidTime,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for invalid time")
	}
}

func TestNewChargingSchedule_InvalidChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: "InvalidUnit",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for invalid unit")
	}
}

func TestNewChargingSchedule_EmptyPeriods(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit:       testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{},
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for empty periods")
	}
}

func TestNewChargingSchedule_InvalidPeriod(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: -1, Limit: testScheduleLimit},
		},
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for invalid period")
	}
}

func TestNewChargingSchedule_NegativeMinChargingRate(t *testing.T) {
	t.Parallel()

	negativeRate := -1.0
	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: testScheduleStartPeriod, Limit: testScheduleLimit},
		},
		MinChargingRate: &negativeRate,
	})

	if err == nil {
		t.Error("NewChargingSchedule() error = nil, want error for negative rate")
	}
}
