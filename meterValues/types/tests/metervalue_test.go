package types_test

import (
	"strings"
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	validTimestamp   = "2025-01-02T15:00:00Z"
	validMeterValue  = "100"
	invalidTimestamp = "not-a-timestamp"
)

func TestNewMeterValue_ValidSingleSampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{Value: validMeterValue},
		},
	}

	meterValue, err := mt.NewMeterValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(meterValue.SampledValue) != 1 {
		t.Errorf(st.ErrorMismatchValue, 1, len(meterValue.SampledValue))
	}
}

func TestNewMeterValue_ValidMultipleSampledValues(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{Value: "100"},
			{Value: "200"},
			{Value: "300"},
		},
	}

	meterValue, err := mt.NewMeterValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expectedCount := 3
	if len(meterValue.SampledValue) != expectedCount {
		t.Errorf(st.ErrorMismatchValue, expectedCount, len(meterValue.SampledValue))
	}
}

func TestNewMeterValue_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: "",
		SampledValue: []mt.SampledValueInput{
			{Value: validMeterValue},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty timestamp")
	}

	if !strings.Contains(err.Error(), "Timestamp") {
		t.Errorf(st.ErrorWantContains, err, "Timestamp")
	}
}

func TestNewMeterValue_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: invalidTimestamp,
		SampledValue: []mt.SampledValueInput{
			{Value: validMeterValue},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid timestamp")
	}

	if !strings.Contains(err.Error(), "Timestamp") {
		t.Errorf(st.ErrorWantContains, err, "Timestamp")
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

	if !strings.Contains(err.Error(), "SampledValue") {
		t.Errorf(st.ErrorWantContains, err, "SampledValue")
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

	if !strings.Contains(err.Error(), "SampledValue") {
		t.Errorf(st.ErrorWantContains, err, "SampledValue")
	}
}

func TestNewMeterValue_InvalidSampledValue(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{Value: ""},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid sampled value")
	}

	if !strings.Contains(err.Error(), "sampledValue[0]") {
		t.Errorf(st.ErrorWantContains, err, "sampledValue[0]")
	}
}

func TestNewMeterValue_MultipleInvalidSampledValues(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: validTimestamp,
		SampledValue: []mt.SampledValueInput{
			{Value: ""},
			{Value: "valid"},
			{Value: ""},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid sampled values")
	}

	if !strings.Contains(err.Error(), "sampledValue[0]") {
		t.Errorf(st.ErrorWantContains, err, "sampledValue[0]")
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
			{Value: ""},
		},
	}

	_, err := mt.NewMeterValue(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), "Timestamp") {
		t.Errorf(st.ErrorWantContains, err, "Timestamp")
	}

	if !strings.Contains(err.Error(), "sampledValue[0]") {
		t.Errorf(st.ErrorWantContains, err, "sampledValue[0]")
	}
}

func TestNewMeterValue_TimestampNormalization(t *testing.T) {
	t.Parallel()

	input := mt.MeterValueInput{
		Timestamp: "2025-01-02T10:00:00-05:00",
		SampledValue: []mt.SampledValueInput{
			{Value: validMeterValue},
		},
	}

	meterValue, err := mt.NewMeterValue(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	timestampValue := meterValue.Timestamp.Value()
	expectedHour := 15
	if timestampValue.Hour() != expectedHour {
		t.Errorf(
			"timestamp should be normalized to UTC (hour 15), got hour %d",
			timestampValue.Hour(),
		)
	}
}
