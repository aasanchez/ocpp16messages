package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	testSampledValueValue = "100"
)

func TestNewSampledValue_Valid(t *testing.T) {
	t.Parallel()

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value: testSampledValueValue,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	_, err := types.NewSampledValue(types.SampledValueInput{
		Value: "",
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for empty value")
	}
}

func TestNewSampledValue_InvalidContext(t *testing.T) {
	t.Parallel()

	invalidContext := "InvalidContext"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:   testSampledValueValue,
		Context: &invalidContext,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid context")
	}
}

func TestNewSampledValue_ValidContext(t *testing.T) {
	t.Parallel()

	validContext := "Sample.Periodic"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value:   testSampledValueValue,
		Context: &validContext,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Context == nil {
		t.Error("SampledValue.Context = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidFormat(t *testing.T) {
	t.Parallel()

	invalidFormat := "InvalidFormat"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:  testSampledValueValue,
		Format: &invalidFormat,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid format")
	}
}

func TestNewSampledValue_ValidFormat(t *testing.T) {
	t.Parallel()

	validFormat := "Raw"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value:  testSampledValueValue,
		Format: &validFormat,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Format == nil {
		t.Error("SampledValue.Format = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidMeasurand(t *testing.T) {
	t.Parallel()

	invalidMeasurand := "InvalidMeasurand"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Measurand: &invalidMeasurand,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid measurand")
	}
}

func TestNewSampledValue_ValidMeasurand(t *testing.T) {
	t.Parallel()

	validMeasurand := "Energy.Active.Import.Register"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value:     testSampledValueValue,
		Measurand: &validMeasurand,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Measurand == nil {
		t.Error("SampledValue.Measurand = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidPhase(t *testing.T) {
	t.Parallel()

	invalidPhase := "InvalidPhase"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value: testSampledValueValue,
		Phase: &invalidPhase,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid phase")
	}
}

func TestNewSampledValue_ValidPhase(t *testing.T) {
	t.Parallel()

	validPhase := "L1"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value: testSampledValueValue,
		Phase: &validPhase,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Phase == nil {
		t.Error("SampledValue.Phase = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidLocation(t *testing.T) {
	t.Parallel()

	invalidLocation := "InvalidLocation"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value:    testSampledValueValue,
		Location: &invalidLocation,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid location")
	}
}

func TestNewSampledValue_ValidLocation(t *testing.T) {
	t.Parallel()

	validLocation := "Outlet"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value:    testSampledValueValue,
		Location: &validLocation,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Location == nil {
		t.Error("SampledValue.Location = nil, want non-nil")
	}
}

func TestNewSampledValue_InvalidUnit(t *testing.T) {
	t.Parallel()

	invalidUnit := "InvalidUnit"
	_, err := types.NewSampledValue(types.SampledValueInput{
		Value: testSampledValueValue,
		Unit:  &invalidUnit,
	})

	if err == nil {
		t.Error("NewSampledValue() error = nil, want error for invalid unit")
	}
}

func TestNewSampledValue_ValidUnit(t *testing.T) {
	t.Parallel()

	validUnit := "Wh"
	sv, err := types.NewSampledValue(types.SampledValueInput{
		Value: testSampledValueValue,
		Unit:  &validUnit,
	})
	if err != nil {
		t.Errorf("NewSampledValue() error = %v, want nil", err)
	}

	if sv.Unit == nil {
		t.Error("SampledValue.Unit = nil, want non-nil")
	}
}
