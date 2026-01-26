package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	testPeriodStartPeriod = 0
	testPeriodLimit       = 32.0
	testPeriodNumPhases   = 3
	errCreatePeriod       = "Failed to create ChargingSchedulePeriod: %v"
	errWantPeriodError    = "NewChargingSchedulePeriod(): want error"
	invalidNumPhasesLow   = 0
	invalidNumPhasesHigh  = 4
	invalidStartPeriod    = -1
	invalidLimitValue     = -1.0
)

func TestChargingSchedulePeriod_StartPeriod(t *testing.T) {
	t.Parallel()

	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: nil,
	}

	period, err := types.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errCreatePeriod, err)
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

	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: nil,
	}

	period, err := types.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errCreatePeriod, err)
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
	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	}

	period, err := types.NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errCreatePeriod, err)
	}

	numberPhases := period.NumberPhases()

	if numberPhases == nil {
		t.Error("ChargingSchedulePeriod.NumberPhases() = nil, want non-nil")
	}
}

func TestChargingSchedulePeriod_NumberPhases_WhenNil(t *testing.T) {
	t.Parallel()

	period, err := types.NewChargingSchedulePeriod(
		types.ChargingSchedulePeriodInput{
			StartPeriod:  testPeriodStartPeriod,
			Limit:        testPeriodLimit,
			NumberPhases: nil,
		},
	)
	if err != nil {
		t.Fatalf(errCreatePeriod, err)
	}

	if period.NumberPhases() != nil {
		t.Error("ChargingSchedulePeriod.NumberPhases() = non-nil, want nil")
	}
}

func TestNewChargingSchedulePeriod_NegativeStartPeriod(t *testing.T) {
	t.Parallel()

	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  invalidStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: nil,
	}

	_, err := types.NewChargingSchedulePeriod(input)
	if err == nil {
		t.Error(errWantPeriodError)
	}
}

func TestNewChargingSchedulePeriod_NegativeLimit(t *testing.T) {
	t.Parallel()

	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        invalidLimitValue,
		NumberPhases: nil,
	}

	_, err := types.NewChargingSchedulePeriod(input)
	if err == nil {
		t.Error(errWantPeriodError)
	}
}

func TestNewChargingSchedulePeriod_NumberPhasesTooLow(t *testing.T) {
	t.Parallel()

	numPhases := invalidNumPhasesLow

	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	})
	if err == nil {
		t.Error(errWantPeriodError)
	}
}

func TestNewChargingSchedulePeriod_NumberPhasesTooHigh(t *testing.T) {
	t.Parallel()

	numPhases := invalidNumPhasesHigh

	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  testPeriodStartPeriod,
		Limit:        testPeriodLimit,
		NumberPhases: &numPhases,
	})
	if err == nil {
		t.Error(errWantPeriodError)
	}
}
