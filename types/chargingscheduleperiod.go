package types

import (
	"fmt"
	"strconv"
)

const (
	// minNumberPhases is the minimum valid number of phases for charging.
	minNumberPhases = 1
	// maxNumberPhases is the maximum valid number of phases for charging.
	maxNumberPhases = 3
	// limitZero is the minimum valid limit value.
	limitZero = 0
)

// ChargingSchedulePeriodInput represents the raw input data for creating a
// ChargingSchedulePeriod. The constructor NewChargingSchedulePeriod validates
// all fields automatically.
type ChargingSchedulePeriodInput struct {
	// Required: Start of the period relative to schedule start (seconds)
	StartPeriod int
	// Required: Power limit during the schedule period (W or A depending on
	// chargingRateUnit)
	Limit float64
	// Optional: Number of phases (1-3) to use for charging
	NumberPhases *int
}

// ChargingSchedulePeriod represents a period within a charging schedule
// as defined in OCPP 1.6.
type ChargingSchedulePeriod struct {
	startPeriod  Integer
	limit        float64
	numberPhases *Integer
}

// NewChargingSchedulePeriod creates a new ChargingSchedulePeriod from input.
// Returns an error if:
//   - StartPeriod is negative or exceeds uint16 max value (65535)
//   - Limit is negative
//   - NumberPhases (if provided) is not between 1 and 3
func NewChargingSchedulePeriod(
	input ChargingSchedulePeriodInput,
) (ChargingSchedulePeriod, error) {
	startPeriod, err := NewInteger(strconv.Itoa(input.StartPeriod))
	if err != nil {
		return ChargingSchedulePeriod{}, fmt.Errorf("startPeriod: %w", err)
	}

	if input.Limit < limitZero {
		return ChargingSchedulePeriod{}, fmt.Errorf(
			"limit: %w",
			ErrInvalidValue,
		)
	}

	numberPhases, err := validateNumberPhases(input.NumberPhases)
	if err != nil {
		return ChargingSchedulePeriod{}, err
	}

	return ChargingSchedulePeriod{
		startPeriod:  startPeriod,
		limit:        input.Limit,
		numberPhases: numberPhases,
	}, nil
}

// validateNumberPhases validates the optional number of phases.
func validateNumberPhases(phases *int) (*Integer, error) {
	if phases == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	if *phases < minNumberPhases || *phases > maxNumberPhases {
		return nil, fmt.Errorf("numberPhases: %w", ErrInvalidValue)
	}

	// NewInteger cannot fail here: values 1-3 are always valid for uint16
	np, _ := NewInteger(strconv.Itoa(*phases))

	return &np, nil
}

// StartPeriod returns the start period in seconds relative to schedule start.
func (c ChargingSchedulePeriod) StartPeriod() Integer {
	return c.startPeriod
}

// Limit returns the power limit during this period.
func (c ChargingSchedulePeriod) Limit() float64 {
	return c.limit
}

// NumberPhases returns the number of phases to use for charging, or nil if not
// specified.
func (c ChargingSchedulePeriod) NumberPhases() *Integer {
	return c.numberPhases
}
