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

func TestNewChargingSchedulePeriod_NegativeStartPeriod(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: -1,
		Limit:       testPeriodLimit,
	})

	if err == nil {
		t.Error("NewChargingSchedulePeriod() error = nil, want error")
	}
}

func TestNewChargingSchedulePeriod_NegativeLimit(t *testing.T) {
	t.Parallel()

	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: testPeriodStartPeriod,
		Limit:       -1.0,
	})

	if err == nil {
		t.Error("NewChargingSchedulePeriod() error = nil, want error")
	}
}

func TestNewChargingSchedulePeriod_NumberPhasesTooLow(t *testing.T) {
	t.Parallel()

	numPhases := 0
	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	})

	if err == nil {
		t.Error("NewChargingSchedulePeriod() error = nil, want error")
	}
}

func TestNewChargingSchedulePeriod_NumberPhasesTooHigh(t *testing.T) {
	t.Parallel()

	numPhases := 4
	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	})

	if err == nil {
		t.Error("NewChargingSchedulePeriod() error = nil, want error")
	}
}
