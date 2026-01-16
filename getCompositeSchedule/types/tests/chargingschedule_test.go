package types_test

import (
	"strings"
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errDuration               = "duration"
	errChargingRateUnit       = "chargingRateUnit"
	errChargingSchedulePeriod = "chargingSchedulePeriod"
	errMinChargingRate        = "minChargingRate"

	durationSixHundred  = 600
	durationNegative    = -1
	durationExceedsMax  = 65536
	minChargingRateFive = 5.0
	minChargingRateNeg  = -1.0
)

func TestChargingSchedule_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if schedule.Duration() != nil {
		t.Errorf("Duration should be nil, got %v", schedule.Duration())
	}

	if schedule.ChargingRateUnit() != gt.ChargingRateUnitWatts {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitWatts,
			schedule.ChargingRateUnit(),
		)
	}

	if len(schedule.ChargingSchedulePeriod()) != valueOne {
		t.Errorf(
			st.ErrorMismatchValue,
			valueOne,
			len(schedule.ChargingSchedulePeriod()),
		)
	}

	if schedule.MinChargingRate() != nil {
		t.Errorf("MinChargingRate should be nil, got %v", schedule.MinChargingRate())
	}
}

func TestChargingSchedule_Valid_WithDuration(t *testing.T) {
	t.Parallel()

	duration := durationSixHundred

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if schedule.Duration() == nil {
		t.Fatal("Duration should not be nil")
	}

	if schedule.Duration().Value() != durationSixHundred {
		t.Errorf(
			st.ErrorMismatchValue,
			durationSixHundred,
			schedule.Duration().Value(),
		)
	}
}

func TestChargingSchedule_Valid_WithChargingRateUnitAmperes(t *testing.T) {
	t.Parallel()

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "A",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if schedule.ChargingRateUnit() != gt.ChargingRateUnitAmperes {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitAmperes,
			schedule.ChargingRateUnit(),
		)
	}
}

func TestChargingSchedule_Valid_WithMultiplePeriods(t *testing.T) {
	t.Parallel()

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
			{StartPeriod: valueThirty, Limit: limitTwenty},
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(schedule.ChargingSchedulePeriod()) != valueTwo {
		t.Errorf(
			st.ErrorMismatchValue,
			valueTwo,
			len(schedule.ChargingSchedulePeriod()),
		)
	}
}

func TestChargingSchedule_Valid_WithMinChargingRate(t *testing.T) {
	t.Parallel()

	minRate := minChargingRateFive

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
		MinChargingRate: &minRate,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if schedule.MinChargingRate() == nil {
		t.Fatal("MinChargingRate should not be nil")
	}

	if *schedule.MinChargingRate() != minChargingRateFive {
		t.Errorf(
			st.ErrorMismatchValue,
			minChargingRateFive,
			*schedule.MinChargingRate(),
		)
	}
}

func TestChargingSchedule_Valid_WithAllFields(t *testing.T) {
	t.Parallel()

	duration := durationSixHundred
	minRate := minChargingRateFive
	phases := valueThree

	schedule, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "A",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen, NumberPhases: &phases},
			{StartPeriod: valueThirty, Limit: limitTwenty},
		},
		MinChargingRate: &minRate,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if schedule.Duration() == nil {
		t.Fatal("Duration should not be nil")
	}

	if schedule.Duration().Value() != durationSixHundred {
		t.Errorf(
			st.ErrorMismatchValue,
			durationSixHundred,
			schedule.Duration().Value(),
		)
	}

	if schedule.ChargingRateUnit() != gt.ChargingRateUnitAmperes {
		t.Errorf(
			st.ErrorMismatch,
			gt.ChargingRateUnitAmperes,
			schedule.ChargingRateUnit(),
		)
	}

	if len(schedule.ChargingSchedulePeriod()) != valueTwo {
		t.Errorf(
			st.ErrorMismatchValue,
			valueTwo,
			len(schedule.ChargingSchedulePeriod()),
		)
	}

	if *schedule.MinChargingRate() != minChargingRateFive {
		t.Errorf(
			st.ErrorMismatchValue,
			minChargingRateFive,
			*schedule.MinChargingRate(),
		)
	}
}

func TestChargingSchedule_Invalid_NegativeDuration(t *testing.T) {
	t.Parallel()

	duration := durationNegative

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative duration")
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}
}

func TestChargingSchedule_Invalid_DurationExceedsMax(t *testing.T) {
	t.Parallel()

	duration := durationExceedsMax

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "duration exceeds max")
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}
}

func TestChargingSchedule_Invalid_ChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "X",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid chargingRateUnit")
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}

func TestChargingSchedule_Invalid_EmptyChargingRateUnit(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty chargingRateUnit")
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}
}

func TestChargingSchedule_Invalid_EmptyPeriods(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit:       "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty chargingSchedulePeriod")
	}

	if !strings.Contains(err.Error(), errChargingSchedulePeriod) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedulePeriod)
	}
}

func TestChargingSchedule_Invalid_NilPeriods(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit:       "W",
		ChargingSchedulePeriod: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "nil chargingSchedulePeriod")
	}

	if !strings.Contains(err.Error(), errChargingSchedulePeriod) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedulePeriod)
	}
}

func TestChargingSchedule_Invalid_InvalidPeriod(t *testing.T) {
	t.Parallel()

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueNegative, Limit: limitTen},
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid period")
	}

	if !strings.Contains(err.Error(), errChargingSchedulePeriod) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedulePeriod)
	}
}

func TestChargingSchedule_Invalid_NegativeMinChargingRate(t *testing.T) {
	t.Parallel()

	minRate := minChargingRateNeg

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
			{StartPeriod: valueZero, Limit: limitTen},
		},
		MinChargingRate: &minRate,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative minChargingRate")
	}

	if !strings.Contains(err.Error(), errMinChargingRate) {
		t.Errorf(st.ErrorWantContains, err, errMinChargingRate)
	}
}

func TestChargingSchedule_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	duration := durationNegative
	minRate := minChargingRateNeg

	_, err := gt.NewChargingSchedule(gt.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       "X",
		ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{},
		MinChargingRate:        &minRate,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errDuration) {
		t.Errorf(st.ErrorWantContains, err, errDuration)
	}

	if !strings.Contains(err.Error(), errChargingRateUnit) {
		t.Errorf(st.ErrorWantContains, err, errChargingRateUnit)
	}

	if !strings.Contains(err.Error(), errChargingSchedulePeriod) {
		t.Errorf(st.ErrorWantContains, err, errChargingSchedulePeriod)
	}

	if !strings.Contains(err.Error(), errMinChargingRate) {
		t.Errorf(st.ErrorWantContains, err, errMinChargingRate)
	}
}
