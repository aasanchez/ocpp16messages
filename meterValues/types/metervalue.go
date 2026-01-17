package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// MeterValueInput represents the raw input data for creating a MeterValue.
type MeterValueInput struct {
	// Required: The timestamp of the measurement in RFC3339 format.
	Timestamp string
	// Required: One or more sampled values.
	SampledValue []SampledValueInput
}

// MeterValue represents a collection of sampled values at a specific timestamp
// as defined in OCPP 1.6.
type MeterValue struct {
	Timestamp    st.DateTime
	SampledValue []SampledValue
}

// meterValueValidation holds validated fields during construction.
type meterValueValidation struct {
	timestamp    st.DateTime
	sampledValue []SampledValue
}

// NewMeterValue creates a new MeterValue from the given input.
// It validates all fields and accumulates errors, returning them together.
func NewMeterValue(input MeterValueInput) (MeterValue, error) {
	validated, errs := validateMeterValueInput(input)
	if errs != nil {
		return MeterValue{
			Timestamp:    st.DateTime{},
			SampledValue: nil,
		}, fmt.Errorf("NewMeterValue: %w", joinErrors(errs))
	}

	return MeterValue{
		Timestamp:    validated.timestamp,
		SampledValue: validated.sampledValue,
	}, nil
}

func validateMeterValueInput(
	input MeterValueInput,
) (meterValueValidation, []error) {
	var errs []error

	var validated meterValueValidation

	validated.timestamp, errs = validateMeterValueTimestamp(
		input.Timestamp,
		errs,
	)
	validated.sampledValue, errs = validateMeterValueSampledValues(
		input.SampledValue,
		errs,
	)

	return validated, errs
}

func validateMeterValueTimestamp(
	timestamp string,
	errs []error,
) (st.DateTime, []error) {
	dateTime, err := st.NewDateTime(timestamp)
	if err != nil {
		return st.DateTime{}, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Timestamp", err),
		)
	}

	return dateTime, errs
}

func validateMeterValueSampledValues(
	sampledValues []SampledValueInput,
	errs []error,
) ([]SampledValue, []error) {
	if len(sampledValues) == errCountZero {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "SampledValue", st.ErrEmptyValue),
		)
	}

	var validValues []SampledValue

	for i, svInput := range sampledValues {
		sampledValue, err := NewSampledValue(svInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("sampledValue[%d]: %w", i, err))
		} else {
			validValues = append(validValues, sampledValue)
		}
	}

	return validValues, errs
}
