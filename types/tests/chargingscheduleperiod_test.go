package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	testPeriodStartPeriod = 0
	testPeriodLimit       = 32.0
	testPeriodNumPhases   = 3
)

func TestChargingSchedulePeriod_StartPeriod(t *testing.T) {
	t.Parallel()

	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: testPeriodStartPeriod,
		Limit:       testPeriodLimit,
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedulePeriod: %v", err)
	}

	startPeriod := period.StartPeriod()

	if startPeriod.Value() != testPeriodStartPeriod {
		t.Errorf(
			"ChargingSchedulePeriod.StartPeriod().Value() = %v, want %v",
			startPeriod.Value(),
			testPeriodStartPeriod,
		)
	}
}

func TestChargingSchedulePeriod_Limit(t *testing.T) {
	t.Parallel()

	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: testPeriodStartPeriod,
		Limit:       testPeriodLimit,
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedulePeriod: %v", err)
	}

	limit := period.Limit()

	if limit != testPeriodLimit {
		t.Errorf(
			"ChargingSchedulePeriod.Limit() = %v, want %v",
			limit,
			testPeriodLimit,
		)
	}
}

func TestChargingSchedulePeriod_NumberPhases(t *testing.T) {
	t.Parallel()

	numPhases := testPeriodNumPhases
	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedulePeriod: %v", err)
	}

	numberPhases := period.NumberPhases()

	if numberPhases == nil {
		t.Error("ChargingSchedulePeriod.NumberPhases() = nil, want non-nil")
	}
}

func TestChargingSchedulePeriod_NumberPhases_WhenNil(t *testing.T) {
	t.Parallel()

	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: nil,
	})
	if err != nil {
		t.Fatalf("Failed to create ChargingSchedulePeriod: %v", err)
	}

	if period.NumberPhases() != nil {
		t.Error("ChargingSchedulePeriod.NumberPhases() = non-nil, want nil")
	}
}
