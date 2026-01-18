package types_test

import (
	"strings"
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	validSampledValue = "100"
	fieldValue        = "Value"
	fieldContext      = "Context"
	fieldLocation     = "Location"
	fieldMeasurand    = "Measurand"
	strUnit           = "Unit"
	strPhase          = "Phase"
	strFormat         = "Format"
)

// Helper function to create string pointer.
func strPtr(s string) *string {
	return &s
}

func TestNewSampledValue_ValidValueOnly(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	sampledValue, err := mt.NewSampledValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if sampledValue.Value.Value() != validSampledValue {
		t.Errorf(
			st.ErrorMismatch,
			validSampledValue,
			sampledValue.Value.Value(),
		)
	}
}

func createValidSampledValueWithAllFields(t *testing.T) mt.SampledValue {
	t.Helper()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   strPtr("Sample.Periodic"),
		Format:    strPtr("Raw"),
		Measurand: strPtr("Energy.Active.Import.Register"),
		Phase:     strPtr("L1"),
		Location:  strPtr("Outlet"),
		Unit:      strPtr("Wh"),
	}

	sampledValue, err := mt.NewSampledValue(input)
	if err != nil {
		t.Fatalf(st.ErrorUnexpectedError, err)
	}

	return sampledValue
}

func TestNewSampledValue_ValidWithAllFields_Context(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Context == nil {
		t.Errorf(st.ErrorWantNonNil, fieldContext)

		return
	}

	if *sampledValue.Context != mt.SamplePeriodic {
		t.Errorf(
			st.ErrorMismatchValue,
			mt.SamplePeriodic,
			*sampledValue.Context,
		)
	}
}

func TestNewSampledValue_ValidWithAllFields_Format(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Format == nil {
		t.Errorf(st.ErrorWantNonNil, strFormat)

		return
	}

	if *sampledValue.Format != mt.Raw {
		t.Errorf(st.ErrorMismatchValue, mt.Raw, *sampledValue.Format)
	}
}

func TestNewSampledValue_ValidWithAllFields_Measurand(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Measurand == nil {
		t.Errorf(st.ErrorWantNonNil, fieldMeasurand)
	}
}

func TestNewSampledValue_ValidWithAllFields_Phase(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Phase == nil {
		t.Errorf(st.ErrorWantNonNil, strPhase)

		return
	}

	if *sampledValue.Phase != mt.L1 {
		t.Errorf(st.ErrorMismatchValue, mt.L1, *sampledValue.Phase)
	}
}

func TestNewSampledValue_ValidWithAllFields_Location(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Location == nil {
		t.Errorf(st.ErrorWantNonNil, fieldLocation)

		return
	}

	if *sampledValue.Location != mt.Outlet {
		t.Errorf(st.ErrorMismatchValue, mt.Outlet, *sampledValue.Location)
	}
}

func TestNewSampledValue_ValidWithAllFields_Unit(t *testing.T) {
	t.Parallel()

	sampledValue := createValidSampledValueWithAllFields(t)

	if sampledValue.Unit == nil {
		t.Errorf(st.ErrorWantNonNil, strUnit)

		return
	}

	if *sampledValue.Unit != mt.Wh {
		t.Errorf(st.ErrorMismatchValue, mt.Wh, *sampledValue.Unit)
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     "",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty value")
	}

	if !strings.Contains(err.Error(), fieldValue) {
		t.Errorf(st.ErrorWantContains, err, fieldValue)
	}
}

func TestNewSampledValue_InvalidContext(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   strPtr("InvalidContext"),
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid context")
	}

	if !strings.Contains(err.Error(), fieldContext) {
		t.Errorf(st.ErrorWantContains, err, fieldContext)
	}
}

func TestNewSampledValue_InvalidFormat(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    strPtr("InvalidFormat"),
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid format")
	}

	if !strings.Contains(err.Error(), strFormat) {
		t.Errorf(st.ErrorWantContains, err, strFormat)
	}
}

func TestNewSampledValue_InvalidMeasurand(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    nil,
		Measurand: strPtr("InvalidMeasurand"),
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid measurand")
	}

	if !strings.Contains(err.Error(), fieldMeasurand) {
		t.Errorf(st.ErrorWantContains, err, fieldMeasurand)
	}
}

func TestNewSampledValue_InvalidPhase(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     strPtr("InvalidPhase"),
		Location:  nil,
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid phase")
	}

	if !strings.Contains(err.Error(), strPhase) {
		t.Errorf(st.ErrorWantContains, err, strPhase)
	}
}

func TestNewSampledValue_InvalidLocation(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  strPtr("InvalidLocation"),
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid location")
	}

	if !strings.Contains(err.Error(), fieldLocation) {
		t.Errorf(st.ErrorWantContains, err, fieldLocation)
	}
}

func TestNewSampledValue_InvalidUnit(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     validSampledValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      strPtr("InvalidUnit"),
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid unit")
	}

	if !strings.Contains(err.Error(), strUnit) {
		t.Errorf(st.ErrorWantContains, err, strUnit)
	}
}

func TestNewSampledValue_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value:     "",
		Context:   strPtr("InvalidContext"),
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  strPtr("InvalidLocation"),
		Unit:      nil,
	}

	_, err := mt.NewSampledValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), fieldValue) {
		t.Errorf(st.ErrorWantContains, err, fieldValue)
	}

	if !strings.Contains(err.Error(), fieldContext) {
		t.Errorf(st.ErrorWantContains, err, fieldContext)
	}

	if !strings.Contains(err.Error(), fieldLocation) {
		t.Errorf(st.ErrorWantContains, err, fieldLocation)
	}
}
