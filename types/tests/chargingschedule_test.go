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
	testInvalidDuration     = -1
	testNegativeMinRate     = -1.0
	errFailedCreateSchedule = "Failed to create ChargingSchedule: %v"
	errScheduleWantNegDur   = "NewChargingSchedule() error = nil, want neg dur"
	errScheduleWantInvTime  = "NewChargingSchedule(): want error for bad time"
	errScheduleWantInvUnit  = "NewChargingSchedule(): want error for bad unit"
	errScheduleWantEmptyPer = "NewChargingSchedule(): want error empty periods"
	errScheduleWantInvPer   = "NewChargingSchedule(): want error bad period"
	errScheduleWantNegRate  = "NewChargingSchedule(): want error negative rate"
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
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
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
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	})
	if err != nil {
		t.Fatalf(errFailedCreateSchedule, err)
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
		Duration:         nil,
		StartSchedule:    nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
	})
	if err != nil {
		t.Fatalf(errFailedCreateSchedule, err)
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
		Duration:         nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	})
	if err != nil {
		t.Fatalf(errFailedCreateSchedule, err)
	}

	if schedule.MinChargingRate() != nil {
		t.Error("ChargingSchedule.MinChargingRate() = non-nil, want nil")
	}
}

func TestChargingSchedule_ChargingSchedulePeriod_WhenNil(t *testing.T) {
	t.Parallel()

	var schedule types.ChargingSchedule
	periods := schedule.ChargingSchedulePeriod()

	if periods != nil {
		t.Error(
			"ChargingSchedule.ChargingSchedulePeriod() != nil" +
				", want nil for zero-value",
		)
	}
}

func TestNewChargingSchedule_InvalidDuration(t *testing.T) {
	t.Parallel()

	duration := testInvalidDuration

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	})
	if err == nil {
		t.Error(errScheduleWantNegDur)
	}
}

func TestNewChargingSchedule_InvalidStartSchedule(t *testing.T) {
	t.Parallel()

	invalidTime := "not-a-valid-time"

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         nil,
		StartSchedule:    &invalidTime,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
	})
	if err == nil {
		t.Error(errScheduleWantInvTime)
	}
}

func TestNewChargingSchedule_InvalidChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: "InvalidUnit",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	})
	if err == nil {
		t.Error(errScheduleWantInvUnit)
	}
}

func TestNewChargingSchedule_EmptyPeriods(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:               nil,
		ChargingRateUnit:       testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{},
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err == nil {
		t.Error(errScheduleWantEmptyPer)
	}
}

func TestNewChargingSchedule_InvalidPeriod(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  -1,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	})
	if err == nil {
		t.Error(errScheduleWantInvPer)
	}
}

func TestNewChargingSchedule_NegativeMinChargingRate(t *testing.T) {
	t.Parallel()

	negativeRate := testNegativeMinRate

	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testScheduleRateUnit,
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{
				StartPeriod:  testScheduleStartPeriod,
				Limit:        testScheduleLimit,
				NumberPhases: nil,
			},
		},
		MinChargingRate: &negativeRate,
		StartSchedule:   nil,
	})
	if err == nil {
		t.Error(errScheduleWantNegRate)
	}
}
