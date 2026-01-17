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
)

// Helper function to create string pointer.
func strPtr(s string) *string {
	return &s
}

func TestNewSampledValue_ValidValueOnly(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value: validSampledValue,
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

func TestNewSampledValue_ValidWithAllOptionalFields(t *testing.T) {
	t.Parallel()

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
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	contextMismatch := sampledValue.Context == nil ||
		*sampledValue.Context != mt.SamplePeriodic
	if contextMismatch {
		t.Errorf(st.ErrorMismatchValue, mt.SamplePeriodic, sampledValue.Context)
	}

	if sampledValue.Format == nil || *sampledValue.Format != mt.Raw {
		t.Errorf(st.ErrorMismatchValue, mt.Raw, sampledValue.Format)
	}

	if sampledValue.Measurand == nil {
		t.Errorf(st.ErrorWantNonNil, fieldMeasurand)
	}

	if sampledValue.Phase == nil || *sampledValue.Phase != mt.L1 {
		t.Errorf(st.ErrorMismatchValue, mt.L1, sampledValue.Phase)
	}

	if sampledValue.Location == nil || *sampledValue.Location != mt.Outlet {
		t.Errorf(st.ErrorMismatchValue, mt.Outlet, sampledValue.Location)
	}

	if sampledValue.Unit == nil || *sampledValue.Unit != mt.Wh {
		t.Errorf(st.ErrorMismatchValue, mt.Wh, sampledValue.Unit)
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	input := mt.SampledValueInput{
		Value: "",
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

	if !strings.Contains(err.Error(), "Format") {
		t.Errorf(st.ErrorWantContains, err, "Format")
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

	if !strings.Contains(err.Error(), "Phase") {
		t.Errorf(st.ErrorWantContains, err, "Phase")
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

	if !strings.Contains(err.Error(), "Unit") {
		t.Errorf(st.ErrorWantContains, err, "Unit")
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
