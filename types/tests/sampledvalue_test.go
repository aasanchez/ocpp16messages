package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	testSampledValueValue = "100"
	errSampledValueNil    = "NewSampledValue() error = %v, want nil"
)

func TestNewSampledValue_Valid(t *testing.T) {
	t.Parallel()

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     "",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for empty value")
	}
}

func TestNewSampledValue_InvalidContext(t *testing.T) {
	t.Parallel()

	invalidContext := "InvalidContext"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   &invalidContext,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid context")
	}
}

func TestNewSampledValue_ValidContext(t *testing.T) {
	t.Parallel()

	validContext := "Sample.Periodic"

	valueWithContext, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   &validContext,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithContext.Context == nil {
		t.Error("SampledValue.Context = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidFormat(t *testing.T) {
	t.Parallel()

	invalidFormat := "InvalidFormat"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    &invalidFormat,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid format")
	}
}

func TestNewSampledValue_ValidFormat(t *testing.T) {
	t.Parallel()

	validFormat := "Raw"

	valueWithFormat, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    &validFormat,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithFormat.Format == nil {
		t.Error("SampledValue.Format = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidMeasurand(t *testing.T) {
	t.Parallel()

	invalidMeasurand := "InvalidMeasurand"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: &invalidMeasurand,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want invalid measurand")
	}
}

func TestNewSampledValue_ValidMeasurand(t *testing.T) {
	t.Parallel()

	validMeasurand := "Energy.Active.Import.Register"

	valueWithMeasurand, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: &validMeasurand,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithMeasurand.Measurand == nil {
		t.Error("SampledValue.Measurand = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidPhase(t *testing.T) {
	t.Parallel()

	invalidPhase := "InvalidPhase"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     &invalidPhase,
		Location:  nil,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid phase")
	}
}

func TestNewSampledValue_ValidPhase(t *testing.T) {
	t.Parallel()

	validPhase := "L1"

	valueWithPhase, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     &validPhase,
		Location:  nil,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithPhase.Phase == nil {
		t.Error("SampledValue.Phase = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidLocation(t *testing.T) {
	t.Parallel()

	invalidLocation := "InvalidLocation"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  &invalidLocation,
		Unit:      nil,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for bad location")
	}
}

func TestNewSampledValue_ValidLocation(t *testing.T) {
	t.Parallel()

	validLocation := "Outlet"

	valueWithLocation, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  &validLocation,
		Unit:      nil,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithLocation.Location == nil {
		t.Error("SampledValue.Location = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidUnit(t *testing.T) {
	t.Parallel()

	invalidUnit := "InvalidUnit"

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      &invalidUnit,
	})
	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid unit")
	}
}

func TestNewSampledValue_ValidUnit(t *testing.T) {
	t.Parallel()

	validUnit := "Wh"

	valueWithUnit, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      &validUnit,
	})
	if err != nil {
		t.Errorf(errSampledValueNil, err)
	}

	if valueWithUnit.Unit == nil {
		t.Error("SampledValue.Unit = nil, want non-nil")
	}
}
