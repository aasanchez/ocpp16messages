package meterValues_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/meterValues"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	validConnectorId     = 1
	validTimestampReq    = "2025-01-02T15:00:00Z"
	validValueReq        = "100"
	connectorIdZero      = 0
	validTransactionId   = 123
	expectedMeterCount1  = 1
	expectedMeterCount2  = 2
	fieldConnectorId     = "ConnectorId"
	fieldTransactionId   = "TransactionId"
	fieldMeterValue      = "MeterValue"
	negativeConnectorId  = -1
	invalidTransactionId = -1
)

// Helper function to create int pointer.
func intPtr(i int) *int {
	return &i
}

func TestReq_ValidMinimalInput(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	req, err := meterValues.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != validConnectorId {
		t.Errorf(
			st.ErrorMismatchValue,
			validConnectorId,
			req.ConnectorId.Value(),
		)
	}

	if len(req.MeterValue) != expectedMeterCount1 {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedMeterCount1,
			len(req.MeterValue),
		)
	}
}

func TestReq_ValidWithTransactionId(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: intPtr(validTransactionId),
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	req, err := meterValues.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.TransactionId == nil {
		t.Errorf(st.ErrorWantNonNil, "TransactionId")
	}

	if req.TransactionId.Value() != validTransactionId {
		t.Errorf(
			st.ErrorMismatchValue,
			validTransactionId,
			req.TransactionId.Value(),
		)
	}
}

func TestReq_ValidConnectorIdZero(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   connectorIdZero,
		TransactionId: nil,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	req, err := meterValues.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != connectorIdZero {
		t.Errorf(
			st.ErrorMismatchValue,
			connectorIdZero,
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_ValidMultipleMeterValues(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: "2025-01-02T15:00:00Z",
				SampledValue: []st.SampledValueInput{
					{
						Value:     "100",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
			{
				Timestamp: "2025-01-02T15:05:00Z",
				SampledValue: []st.SampledValueInput{
					{
						Value:     "150",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	req, err := meterValues.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(req.MeterValue) != expectedMeterCount2 {
		t.Errorf(
			st.ErrorMismatchValue,
			expectedMeterCount2,
			len(req.MeterValue),
		)
	}
}

func TestReq_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   negativeConnectorId,
		TransactionId: nil,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative connector id")
	}

	if !strings.Contains(err.Error(), fieldConnectorId) {
		t.Errorf(st.ErrorWantContains, err, fieldConnectorId)
	}
}

func TestReq_NegativeTransactionId(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: intPtr(invalidTransactionId),
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative transaction id")
	}

	if !strings.Contains(err.Error(), fieldTransactionId) {
		t.Errorf(st.ErrorWantContains, err, fieldTransactionId)
	}
}

func TestReq_EmptyMeterValue(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue:    []st.MeterValueInput{},
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(st.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_NilMeterValue(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue:    nil,
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "nil meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(st.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_InvalidMeterValue(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []st.MeterValueInput{
			{
				Timestamp: "invalid-timestamp",
				SampledValue: []st.SampledValueInput{
					{
						Value:     validValueReq,
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid meter value")
	}

	if !strings.Contains(err.Error(), "meterValue[0]") {
		t.Errorf(st.ErrorWantContains, err, "meterValue[0]")
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := meterValues.ReqInput{
		ConnectorId:   negativeConnectorId,
		TransactionId: intPtr(invalidTransactionId),
		MeterValue:    nil,
	}

	_, err := meterValues.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), fieldConnectorId) {
		t.Errorf(st.ErrorWantContains, err, fieldConnectorId)
	}

	if !strings.Contains(err.Error(), fieldTransactionId) {
		t.Errorf(st.ErrorWantContains, err, fieldTransactionId)
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(st.ErrorWantContains, err, fieldMeterValue)
	}
}
