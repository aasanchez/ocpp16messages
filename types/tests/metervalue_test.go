package types_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const (
	validMeterValueTimestamp = "2025-01-02T15:00:00Z"
	validMeterValue          = "100"
	invalidMeterTimestamp    = "not-a-timestamp"
	emptyMeterValueStr       = ""
	fieldMeterTimestamp      = "Timestamp"
	fieldMeterSampledValue   = "SampledValue"
	fieldMeterSampledIdx     = "sampledValue[0]"
	expectedMeterSampledOne  = 1
)

func TestNewMeterValue_ValidSingleSampledValue(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: validMeterValueTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     validMeterValue,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	meterValue, err := types.NewMeterValue(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(meterValue.SampledValue) != expectedMeterSampledOne {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedMeterSampledOne,
			len(meterValue.SampledValue),
		)
	}
}

func TestNewMeterValue_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: "",
		SampledValue: []types.SampledValueInput{
			{
				Value:     validMeterValue,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want empty timestamp error")
	}

	if !strings.Contains(err.Error(), fieldMeterTimestamp) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterTimestamp)
	}
}

func TestNewMeterValue_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: invalidMeterTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     validMeterValue,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want invalid timestamp error")
	}

	if !strings.Contains(err.Error(), fieldMeterTimestamp) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterTimestamp)
	}
}

func TestNewMeterValue_EmptySampledValueArray(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp:    validMeterValueTimestamp,
		SampledValue: []types.SampledValueInput{},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want empty sampled error")
	}

	if !strings.Contains(err.Error(), fieldMeterSampledValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterSampledValue)
	}
}

func TestNewMeterValue_NilSampledValueArray(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp:    validMeterValueTimestamp,
		SampledValue: nil,
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want nil sampled error")
	}

	if !strings.Contains(err.Error(), fieldMeterSampledValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterSampledValue)
	}
}

func TestNewMeterValue_InvalidSampledValue(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: validMeterValueTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     emptyMeterValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want invalid sampled error")
	}

	if !strings.Contains(err.Error(), fieldMeterSampledIdx) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterSampledIdx)
	}
}

func TestNewMeterValue_MultipleInvalidSampledValues(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: validMeterValueTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     emptyMeterValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     validMeterValue,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     emptyMeterValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want invalid sampled error")
	}

	if !strings.Contains(err.Error(), fieldMeterSampledIdx) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterSampledIdx)
	}

	if !strings.Contains(err.Error(), "sampledValue[2]") {
		t.Errorf(types.ErrorWantContains, err, "sampledValue[2]")
	}
}

func TestNewMeterValue_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := types.MeterValueInput{
		Timestamp: invalidMeterTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     emptyMeterValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := types.NewMeterValue(input)
	if err == nil {
		t.Error("NewMeterValue() error = nil, want multiple errors")
	}

	if !strings.Contains(err.Error(), fieldMeterTimestamp) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterTimestamp)
	}

	if !strings.Contains(err.Error(), fieldMeterSampledIdx) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterSampledIdx)
	}
}
