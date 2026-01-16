package types_test

import (
	"strings"
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStartPeriod  = "startPeriod"
	errLimit        = "limit"
	errNumberPhases = "numberPhases"

	valueZero       = 0
	valueOne        = 1
	valueTwo        = 2
	valueThree      = 3
	valueFour       = 4
	valueThirty     = 30
	valueNegative   = -1
	valueExceedsMax = 65536
	limitTen        = 10.0
	limitTwenty     = 20.0
	limitNegative   = -1.0
)

func intPtr(v int) *int {
	return &v
}

func TestChargingSchedulePeriod_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	period, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod: valueZero,
		Limit:       limitTen,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if period.StartPeriod().Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, period.StartPeriod().Value())
	}

	if period.Limit() != limitTen {
		t.Errorf(st.ErrorMismatchValue, limitTen, period.Limit())
	}

	if period.NumberPhases() != nil {
		t.Errorf("NumberPhases should be nil, got %v", period.NumberPhases())
	}
}

func TestChargingSchedulePeriod_Valid_WithStartPeriodNonZero(t *testing.T) {
	t.Parallel()

	period, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod: valueThirty,
		Limit:       limitTwenty,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if period.StartPeriod().Value() != valueThirty {
		t.Errorf(st.ErrorMismatchValue, valueThirty, period.StartPeriod().Value())
	}
}

func TestChargingSchedulePeriod_Valid_WithNumberPhasesOne(t *testing.T) {
	t.Parallel()

	period, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueOne),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if period.NumberPhases() == nil {
		t.Fatal("NumberPhases should not be nil")
	}

	if period.NumberPhases().Value() != valueOne {
		t.Errorf(st.ErrorMismatchValue, valueOne, period.NumberPhases().Value())
	}
}

func TestChargingSchedulePeriod_Valid_WithNumberPhasesTwo(t *testing.T) {
	t.Parallel()

	period, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueTwo),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if period.NumberPhases() == nil {
		t.Fatal("NumberPhases should not be nil")
	}

	if period.NumberPhases().Value() != valueTwo {
		t.Errorf(st.ErrorMismatchValue, valueTwo, period.NumberPhases().Value())
	}
}

func TestChargingSchedulePeriod_Valid_WithNumberPhasesThree(t *testing.T) {
	t.Parallel()

	period, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if period.NumberPhases() == nil {
		t.Fatal("NumberPhases should not be nil")
	}

	if period.NumberPhases().Value() != valueThree {
		t.Errorf(st.ErrorMismatchValue, valueThree, period.NumberPhases().Value())
	}
}

func TestChargingSchedulePeriod_Invalid_NegativeStartPeriod(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod: valueNegative,
		Limit:       limitTen,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative startPeriod")
	}

	if !strings.Contains(err.Error(), errStartPeriod) {
		t.Errorf(st.ErrorWantContains, err, errStartPeriod)
	}
}

func TestChargingSchedulePeriod_Invalid_StartPeriodExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod: valueExceedsMax,
		Limit:       limitTen,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "startPeriod exceeds max")
	}

	if !strings.Contains(err.Error(), errStartPeriod) {
		t.Errorf(st.ErrorWantContains, err, errStartPeriod)
	}
}

func TestChargingSchedulePeriod_Invalid_NegativeLimit(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod: valueZero,
		Limit:       limitNegative,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative limit")
	}

	if !strings.Contains(err.Error(), errLimit) {
		t.Errorf(st.ErrorWantContains, err, errLimit)
	}
}

func TestChargingSchedulePeriod_Invalid_NumberPhasesZero(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueZero),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "numberPhases is 0")
	}

	if !strings.Contains(err.Error(), errNumberPhases) {
		t.Errorf(st.ErrorWantContains, err, errNumberPhases)
	}
}

func TestChargingSchedulePeriod_Invalid_NumberPhasesFour(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueFour),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "numberPhases is 4")
	}

	if !strings.Contains(err.Error(), errNumberPhases) {
		t.Errorf(st.ErrorWantContains, err, errNumberPhases)
	}
}

func TestChargingSchedulePeriod_Invalid_NumberPhasesNegative(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedulePeriod(gt.ChargingSchedulePeriodInput{
		StartPeriod:  valueZero,
		Limit:        limitTen,
		NumberPhases: intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative numberPhases")
	}

	if !strings.Contains(err.Error(), errNumberPhases) {
		t.Errorf(st.ErrorWantContains, err, errNumberPhases)
	}
}
