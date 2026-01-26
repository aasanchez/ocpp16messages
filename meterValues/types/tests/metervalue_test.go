package types_test

import (
	"strings"
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	validTimestamp       = "2025-01-02T15:00:00Z"
	validMeterValue      = "100"
	invalidTimestamp     = "not-a-timestamp"
	expectedSampledOne   = 1
	expectedSampledThree = 3
	emptyValueStr        = ""
	fieldTimestamp       = "Timestamp"
	fieldSampledValue    = "SampledValue"
	fieldSampledValueIdx = "sampledValue[0]"
)

func TestNewMeterValue_ValidSingleSampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
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

	meterValue, err := mt.NewMeterValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(meterValue.SampledValue) != expectedSampledOne {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedSampledOne,
			len(meterValue.SampledValue),
		)
	}
}

func TestNewMeterValue_ValidMultipleSampledValues(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{
				Value:     "100",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     "200",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     "300",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	meterValue, err := mt.NewMeterValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(meterValue.SampledValue) != expectedSampledThree {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedSampledThree,
			len(meterValue.SampledValue),
		)
	}
}

func TestNewMeterValue_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: "",
		SampledValue: []mt.SampledValueInput{
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

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty timestamp")
	}

	if !strings.Contains(err.Error(), fieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, fieldTimestamp)
	}
}

func TestNewMeterValue_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: invalidTimestamp,
		SampledValue: []mt.SampledValueInput{
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

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid timestamp")
	}

	if !strings.Contains(err.Error(), fieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, fieldTimestamp)
	}
}

func TestNewMeterValue_EmptySampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp:    validTimestamp,
		SampledValue: []mt.SampledValueInput{},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty sampled value")
	}

	if !strings.Contains(err.Error(), fieldSampledValue) {
		t.Errorf(st.ErrorWantContains, err, fieldSampledValue)
	}
}

func TestNewMeterValue_NilSampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp:    validTimestamp,
		SampledValue: nil,
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "nil sampled value")
	}

	if !strings.Contains(err.Error(), fieldSampledValue) {
		t.Errorf(st.ErrorWantContains, err, fieldSampledValue)
	}
}

func TestNewMeterValue_InvalidSampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{
				Value:     emptyValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid sampled value")
	}

	if !strings.Contains(err.Error(), fieldSampledValueIdx) {
		t.Errorf(st.ErrorWantContains, err, fieldSampledValueIdx)
	}
}

func TestNewMeterValue_MultipleInvalidSampledValues(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{
				Value:     emptyValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     "valid",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
			{
				Value:     emptyValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid sampled values")
	}

	if !strings.Contains(err.Error(), fieldSampledValueIdx) {
		t.Errorf(st.ErrorWantContains, err, fieldSampledValueIdx)
	}

	if !strings.Contains(err.Error(), "sampledValue[2]") {
		t.Errorf(st.ErrorWantContains, err, "sampledValue[2]")
	}
}

func TestNewMeterValue_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: invalidTimestamp,
		SampledValue: []mt.SampledValueInput{
			{
				Value:     emptyValueStr,
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), fieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, fieldTimestamp)
	}

	if !strings.Contains(err.Error(), fieldSampledValueIdx) {
		t.Errorf(st.ErrorWantContains, err, fieldSampledValueIdx)
	}
}

func TestNewMeterValue_TimestampNormalization(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: "2025-01-02T10:00:00-05:00",
		SampledValue: []mt.SampledValueInput{
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

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "non-UTC timestamp")
	}
}
