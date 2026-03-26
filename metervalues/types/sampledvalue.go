package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// SampledValueInput represents the raw input data for creating a SampledValue.
type SampledValueInput struct {
	// Required: The measurement value as a string.
	Value string
	// Optional: The reading context.
	Context *string
	// Optional: The value format.
	Format *string
	// Optional: The type of measurement.
	Measurand *string
	// Optional: The phase of the measurement.
	Phase *string
	// Optional: The location of the measurement.
	Location *string
	// Optional: The unit of the measurement.
	Unit *string
}

// SampledValue represents a single sampled value as defined in OCPP 1.6.
type SampledValue struct {
	Value     st.CiString500Type
	Context   *ReadingContext
	Format    *ValueFormat
	Measurand *Measurand
	Phase     *Phase
	Location  *Location
	Unit      *UnitOfMeasure
}

// sampledValueValidation holds validated fields during construction.
type sampledValueValidation struct {
	value     st.CiString500Type
	context   *ReadingContext
	format    *ValueFormat
	measurand *Measurand
	phase     *Phase
	location  *Location
	unit      *UnitOfMeasure
}

// NewSampledValue creates a new SampledValue from the given input.
// It validates all fields and accumulates errors, returning them together.
func NewSampledValue(input SampledValueInput) (SampledValue, error) {
	validated, errs := validateSampledValueInput(input)
	if errs != nil {
		return SampledValue{
			Value:     st.CiString500Type{},
			Context:   nil,
			Format:    nil,
			Measurand: nil,
			Phase:     nil,
			Location:  nil,
			Unit:      nil,
		}, fmt.Errorf("NewSampledValue: %w", joinErrors(errs))
	}

	return SampledValue{
		Value:     validated.value,
		Context:   validated.context,
		Format:    validated.format,
		Measurand: validated.measurand,
		Phase:     validated.phase,
		Location:  validated.location,
		Unit:      validated.unit,
	}, nil
}

func validateSampledValueInput(
	input SampledValueInput,
) (sampledValueValidation, []error) {
	var errs []error

	var validated sampledValueValidation

	validated.value, errs = validateSampledValueValue(input.Value, errs)
	validated.context, errs = validateSampledValueContext(input.Context, errs)
	validated.format, errs = validateSampledValueFormat(input.Format, errs)
	validated.measurand, errs = validateSampledValueMeasurand(
		input.Measurand,
		errs,
	)
	validated.phase, errs = validateSampledValuePhase(input.Phase, errs)
	validated.location, errs = validateSampledValueLocation(
		input.Location,
		errs,
	)
	validated.unit, errs = validateSampledValueUnit(input.Unit, errs)

	return validated, errs
}

func validateSampledValueValue(
	value string,
	errs []error,
) (st.CiString500Type, []error) {
	ciValue, err := st.NewCiString500Type(value)
	if err != nil {
		return st.CiString500Type{}, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Value", err),
		)
	}

	return ciValue, errs
}

func validateSampledValueContext(
	context *string,
	errs []error,
) (*ReadingContext, []error) {
	if context == nil {
		return nil, errs
	}

	readingCtx := ReadingContext(*context)
	if !readingCtx.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Context", st.ErrInvalidValue),
		)
	}

	return &readingCtx, errs
}

func validateSampledValueFormat(
	format *string,
	errs []error,
) (*ValueFormat, []error) {
	if format == nil {
		return nil, errs
	}

	valueFormat := ValueFormat(*format)
	if !valueFormat.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Format", st.ErrInvalidValue),
		)
	}

	return &valueFormat, errs
}

func validateSampledValueMeasurand(
	measurand *string,
	errs []error,
) (*Measurand, []error) {
	if measurand == nil {
		return nil, errs
	}

	measurandVal := Measurand(*measurand)
	if !measurandVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Measurand", st.ErrInvalidValue),
		)
	}

	return &measurandVal, errs
}

func validateSampledValuePhase(
	phase *string,
	errs []error,
) (*Phase, []error) {
	if phase == nil {
		return nil, errs
	}

	phaseVal := Phase(*phase)
	if !phaseVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Phase", st.ErrInvalidValue),
		)
	}

	return &phaseVal, errs
}

func validateSampledValueLocation(
	location *string,
	errs []error,
) (*Location, []error) {
	if location == nil {
		return nil, errs
	}

	locationVal := Location(*location)
	if !locationVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Location", st.ErrInvalidValue),
		)
	}

	return &locationVal, errs
}

func validateSampledValueUnit(
	unit *string,
	errs []error,
) (*UnitOfMeasure, []error) {
	if unit == nil {
		return nil, errs
	}

	unitOfMeasure := UnitOfMeasure(*unit)
	if !unitOfMeasure.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "Unit", st.ErrInvalidValue),
		)
	}

	return &unitOfMeasure, errs
}
