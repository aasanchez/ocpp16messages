package meterValues_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues"
	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
)

const (
	exampleConnectorId   = 1
	exampleTransactionId = 42
	exampleTimestamp     = "2025-01-02T15:00:00Z"
	exampleValue         = "12500"
	outputError          = "Error:"
	outputConnectorId    = "ConnectorId:"
)

// ExampleReq demonstrates creating a valid MeterValues.req message with
// a single meter value containing one sampled value.
func ExampleReq() {
	input := meterValues.ReqInput{
		ConnectorId:   exampleConnectorId,
		TransactionId: nil,
		MeterValue: []mt.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []mt.SampledValueInput{
					{
						Value:     exampleValue,
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
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorId, req.ConnectorId.Value())
	fmt.Println("MeterValue count:", len(req.MeterValue))
	// Output:
	// ConnectorId: 1
	// MeterValue count: 1
}

// ExampleReq_withTransactionId demonstrates creating a MeterValues.req
// message with an associated transaction ID.
func ExampleReq_withTransactionId() {
	transactionId := exampleTransactionId

	input := meterValues.ReqInput{
		ConnectorId:   exampleConnectorId,
		TransactionId: &transactionId,
		MeterValue: []mt.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []mt.SampledValueInput{
					{
						Value:     exampleValue,
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
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorId, req.ConnectorId.Value())
	fmt.Println("TransactionId:", req.TransactionId.Value())
	// Output:
	// ConnectorId: 1
	// TransactionId: 42
}

// ExampleReq_withOptionalFields demonstrates creating a MeterValues.req
// message with all optional SampledValue fields populated.
func ExampleReq_withOptionalFields() {
	context := "Sample.Periodic"
	format := "Raw"
	measurand := "Energy.Active.Import.Register"
	phase := "L1"
	location := "Outlet"
	unit := "Wh"

	input := meterValues.ReqInput{
		ConnectorId:   exampleConnectorId,
		TransactionId: nil,
		MeterValue: []mt.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []mt.SampledValueInput{
					{
						Value:     exampleValue,
						Context:   &context,
						Format:    &format,
						Measurand: &measurand,
						Phase:     &phase,
						Location:  &location,
						Unit:      &unit,
					},
				},
			},
		},
	}

	req, err := meterValues.Req(input)
	if err != nil {
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorId, req.ConnectorId.Value())
	fmt.Println("Value:", req.MeterValue[0].SampledValue[0].Value.Value())
	// Output:
	// ConnectorId: 1
	// Value: 12500
}

// ExampleReq_emptyMeterValue demonstrates the error when MeterValue is empty.
func ExampleReq_emptyMeterValue() {
	input := meterValues.ReqInput{
		ConnectorId:   exampleConnectorId,
		TransactionId: nil,
		MeterValue:    []mt.MeterValueInput{},
	}

	_, err := meterValues.Req(input)
	if err != nil {
		fmt.Println("Error: MeterValue cannot be empty")
	}
	// Output:
	// Error: MeterValue cannot be empty
}
