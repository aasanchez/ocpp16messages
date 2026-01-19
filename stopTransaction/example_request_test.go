package stopTransaction_test

import (
	"fmt"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	"github.com/aasanchez/ocpp16messages/stopTransaction"
)

const (
	exampleTxId           = 12345
	exampleMeterStop      = 50000
	exampleTimestamp      = "2025-01-15T10:30:00Z"
	exampleReqErrorLabel  = "Error:"
	exampleTxIdLabel      = "TransactionId:"
	exampleMeterStopLabel = "MeterStop:"
)

// ExampleReq demonstrates creating a basic StopTransaction.req message.
func ExampleReq() {
	req, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId:   exampleTxId,
		IdTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionId.Value())
	fmt.Println(exampleMeterStopLabel, req.MeterStop.Value())
	// Output:
	// TransactionId: 12345
	// MeterStop: 50000
}

// ExampleReq_withIdTag demonstrates stopping a transaction with an ID tag.
func ExampleReq_withIdTag() {
	idTag := "RFID-ABC123"

	req, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId:   exampleTxId,
		IdTag:           &idTag,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionId.Value())
	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// TransactionId: 12345
	// IdTag: RFID-ABC123
}

// ExampleReq_withReason demonstrates stopping a transaction with a reason.
func ExampleReq_withReason() {
	reason := "Remote"

	req, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId:   exampleTxId,
		IdTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionId.Value())
	fmt.Println("Reason:", req.Reason.String())
	// Output:
	// TransactionId: 12345
	// Reason: Remote
}

// ExampleReq_withTransactionData shows including meter values in the request.
func ExampleReq_withTransactionData() {
	req, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId: exampleTxId,
		IdTag:         nil,
		MeterStop:     exampleMeterStop,
		Timestamp:     exampleTimestamp,
		Reason:        nil,
		TransactionData: []mt.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []mt.SampledValueInput{
					{
						Value:     "50000",
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
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionId.Value())
	fmt.Println("TransactionData count:", len(req.TransactionData))
	// Output:
	// TransactionId: 12345
	// TransactionData count: 1
}

// ExampleReq_invalidTimestamp demonstrates validation error for bad timestamp.
func ExampleReq_invalidTimestamp() {
	_, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId:   exampleTxId,
		IdTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       "invalid-timestamp",
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println("Validation failed as expected")
	}
	// Output:
	// Validation failed as expected
}
