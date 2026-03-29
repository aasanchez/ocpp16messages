package metervalues_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/metervalues"
	types "github.com/aasanchez/ocpp16types"
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

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
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

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != validConnectorId {
		t.Errorf(
			types.ErrorMismatchValue,
			validConnectorId,
			req.ConnectorId.Value(),
		)
	}

	if len(req.MeterValue) != expectedMeterCount1 {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedMeterCount1,
			len(req.MeterValue),
		)
	}
}

func TestReq_ValidWithTransactionId(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: intPtr(validTransactionId),
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
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

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionId == nil {
		t.Errorf(types.ErrorWantNonNil, "TransactionId")
	}

	if req.TransactionId.Value() != validTransactionId {
		t.Errorf(
			types.ErrorMismatchValue,
			validTransactionId,
			req.TransactionId.Value(),
		)
	}
}

func TestReq_ValidConnectorIdZero(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   connectorIdZero,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
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

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != connectorIdZero {
		t.Errorf(
			types.ErrorMismatchValue,
			connectorIdZero,
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_ValidMultipleMeterValues(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: "2025-01-02T15:00:00Z",
				SampledValue: []types.SampledValueInput{
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
				SampledValue: []types.SampledValueInput{
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

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.MeterValue) != expectedMeterCount2 {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedMeterCount2,
			len(req.MeterValue),
		)
	}
}

func TestReq_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   negativeConnectorId,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
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

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connector id")
	}

	if !strings.Contains(err.Error(), fieldConnectorId) {
		t.Errorf(types.ErrorWantContains, err, fieldConnectorId)
	}
}

func TestReq_NegativeTransactionId(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: intPtr(invalidTransactionId),
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
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

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative transaction id")
	}

	if !strings.Contains(err.Error(), fieldTransactionId) {
		t.Errorf(types.ErrorWantContains, err, fieldTransactionId)
	}
}

func TestReq_EmptyMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue:    []types.MeterValueInput{},
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_NilMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue:    nil,
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "nil meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_InvalidMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   validConnectorId,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: "invalid-timestamp",
				SampledValue: []types.SampledValueInput{
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

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid meter value")
	}

	if !strings.Contains(err.Error(), "meterValue[0]") {
		t.Errorf(types.ErrorWantContains, err, "meterValue[0]")
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorId:   negativeConnectorId,
		TransactionId: intPtr(invalidTransactionId),
		MeterValue:    nil,
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), fieldConnectorId) {
		t.Errorf(types.ErrorWantContains, err, fieldConnectorId)
	}

	if !strings.Contains(err.Error(), fieldTransactionId) {
		t.Errorf(types.ErrorWantContains, err, fieldTransactionId)
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}
