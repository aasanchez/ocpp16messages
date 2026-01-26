package types

import (
	"errors"
	"fmt"
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
	// Optional: Starting point of an absolute schedule (RFC3339 format)
	StartSchedule *string
}

// ChargingSchedule represents a composite charging schedule as defined in
// OCPP 1.6.
type ChargingSchedule struct {
	duration               *Integer
	startSchedule          *DateTime
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
//   - StartSchedule (if provided) is not a valid RFC3339 timestamp
func NewChargingSchedule(
	input ChargingScheduleInput,
) (ChargingSchedule, error) {
	var errs []error

	duration, err := validateScheduleDuration(input.Duration)
	if err != nil {
		errs = append(errs, err)
	}

	startSchedule, err := validateStartSchedule(input.StartSchedule)
	if err != nil {
		errs = append(errs, err)
	}

	chargingRateUnit, err := validateChargingRateUnit(input.ChargingRateUnit)
	if err != nil {
		errs = append(errs, err)
	}

	periods, periodErrs := validateSchedulePeriods(input.ChargingSchedulePeriod)
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
		startSchedule:          startSchedule,
		chargingRateUnit:       chargingRateUnit,
		chargingSchedulePeriod: periods,
		minChargingRate:        minChargingRate,
	}, nil
}

// validateScheduleDuration validates the optional duration field.
func validateScheduleDuration(duration *int) (*Integer, error) {
	if duration == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	d, err := NewIntegerFromInt(*duration)
	if err != nil {
		return nil, fmt.Errorf("duration: %w", err)
	}

	return &d, nil
}

// validateStartSchedule validates the optional start schedule field.
func validateStartSchedule(startSchedule *string) (*DateTime, error) {
	if startSchedule == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	ss, err := NewDateTime(*startSchedule)
	if err != nil {
		return nil, fmt.Errorf("startSchedule: %w", err)
	}

	return &ss, nil
}

// validateChargingRateUnit validates the charging rate unit field.
func validateChargingRateUnit(unit string) (ChargingRateUnit, error) {
	chargingRateUnit := ChargingRateUnit(unit)
	if !chargingRateUnit.IsValid() {
		return "", fmt.Errorf("chargingRateUnit: %w", ErrInvalidValue)
	}

	return chargingRateUnit, nil
}

// validateSchedulePeriods validates the charging schedule periods.
func validateSchedulePeriods(
	periods []ChargingSchedulePeriodInput,
) ([]ChargingSchedulePeriod, []error) {
	if len(periods) == periodsLenZero {
		return nil, []error{
			fmt.Errorf("chargingSchedulePeriod: %w", ErrEmptyValue),
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
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	if *rate < minChargingRateZero {
		return nil, fmt.Errorf("minChargingRate: %w", ErrInvalidValue)
	}

	return rate, nil
}

// Duration returns the duration in seconds, or nil if not specified.
func (c ChargingSchedule) Duration() *Integer {
	return c.duration
}

// StartSchedule returns the start schedule, or nil if not specified.
func (c ChargingSchedule) StartSchedule() *DateTime {
	return c.startSchedule
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
