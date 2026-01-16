package types

import (
	"errors"
	"fmt"
	"strconv"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// minChargingRateZero is the minimum valid charging rate value.
	minChargingRateZero = 0
	// errCountZero is the empty error count.
	errCountZero = 0
	// periodsLenZero is the empty periods length.
	periodsLenZero = 0
)

// ChargingScheduleInput represents the raw input data for creating a
// ChargingSchedule. The constructor NewChargingSchedule validates all fields
// automatically.
type ChargingScheduleInput struct {
	// Optional: Duration of the charging schedule in seconds
	Duration *int
	// Required: Unit of measure for charging rate
	ChargingRateUnit string
	// Required: List of charging schedule periods (at least one)
	ChargingSchedulePeriod []ChargingSchedulePeriodInput
	// Optional: Minimum charging rate supported by the EV
	MinChargingRate *float64
}

// ChargingSchedule represents a composite charging schedule as defined in
// OCPP 1.6.
type ChargingSchedule struct {
	duration               *st.Integer
	chargingRateUnit       ChargingRateUnit
	chargingSchedulePeriod []ChargingSchedulePeriod
	minChargingRate        *float64
}

// NewChargingSchedule creates a new ChargingSchedule from input.
// Returns an error if:
//   - Duration (if provided) is negative or exceeds uint16 max value (65535)
//   - ChargingRateUnit is not a valid value ("W" or "A")
//   - ChargingSchedulePeriod is empty or contains invalid periods
//   - MinChargingRate (if provided) is negative
func NewChargingSchedule(
	input ChargingScheduleInput,
) (ChargingSchedule, error) {
	var errs []error

	duration, err := validateDuration(input.Duration)
	if err != nil {
		errs = append(errs, err)
	}

	chargingRateUnit, err := validateChargingRateUnit(input.ChargingRateUnit)
	if err != nil {
		errs = append(errs, err)
	}

	periods, periodErrs := validatePeriods(input.ChargingSchedulePeriod)
	errs = append(errs, periodErrs...)

	minChargingRate, err := validateMinChargingRate(input.MinChargingRate)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > errCountZero {
		return ChargingSchedule{}, errors.Join(errs...)
	}

	return ChargingSchedule{
		duration:               duration,
		chargingRateUnit:       chargingRateUnit,
		chargingSchedulePeriod: periods,
		minChargingRate:        minChargingRate,
	}, nil
}

// validateDuration validates the optional duration field.
func validateDuration(duration *int) (*st.Integer, error) {
	if duration == nil {
		return nil, nil
	}

	d, err := st.NewInteger(strconv.Itoa(*duration))
	if err != nil {
		return nil, fmt.Errorf("duration: %w", err)
	}

	return &d, nil
}

// validateChargingRateUnit validates the charging rate unit field.
func validateChargingRateUnit(unit string) (ChargingRateUnit, error) {
	chargingRateUnit := ChargingRateUnit(unit)
	if !chargingRateUnit.IsValid() {
		return "", fmt.Errorf("chargingRateUnit: %w", st.ErrInvalidValue)
	}

	return chargingRateUnit, nil
}

// validatePeriods validates the charging schedule periods.
func validatePeriods(
	periods []ChargingSchedulePeriodInput,
) ([]ChargingSchedulePeriod, []error) {
	if len(periods) == periodsLenZero {
		return nil, []error{
			fmt.Errorf("chargingSchedulePeriod: %w", st.ErrEmptyValue),
		}
	}

	var errs []error

	var validPeriods []ChargingSchedulePeriod

	for i, periodInput := range periods {
		period, err := NewChargingSchedulePeriod(periodInput)
		if err != nil {
			errs = append(
				errs,
				fmt.Errorf("chargingSchedulePeriod[%d]: %w", i, err),
			)
		} else {
			validPeriods = append(validPeriods, period)
		}
	}

	return validPeriods, errs
}

// validateMinChargingRate validates the optional minimum charging rate.
func validateMinChargingRate(rate *float64) (*float64, error) {
	if rate == nil {
		return nil, nil
	}

	if *rate < minChargingRateZero {
		return nil, fmt.Errorf("minChargingRate: %w", st.ErrInvalidValue)
	}

	return rate, nil
}

// Duration returns the duration in seconds, or nil if not specified.
func (c ChargingSchedule) Duration() *st.Integer {
	return c.duration
}

// ChargingRateUnit returns the unit of measure for the charging rate.
func (c ChargingSchedule) ChargingRateUnit() ChargingRateUnit {
	return c.chargingRateUnit
}

// ChargingSchedulePeriod returns the list of charging schedule periods.
func (c ChargingSchedule) ChargingSchedulePeriod() []ChargingSchedulePeriod {
	return c.chargingSchedulePeriod
}

// MinChargingRate returns the minimum charging rate, or nil if not specified.
func (c ChargingSchedule) MinChargingRate() *float64 {
	return c.minChargingRate
}
