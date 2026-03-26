package remotestoptransaction_test

import (
	"fmt"

	rst "github.com/aasanchez/ocpp16messages/remotestoptransaction"
)

const (
	exampleTransactionId    = 12345
	exampleTransactionIdNeg = -1
)

// ExampleReq demonstrates creating a valid RemoteStopTransaction.req message.
func ExampleReq() {
	req, err := rst.Req(rst.ReqInput{TransactionId: exampleTransactionId})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("TransactionId:", req.TransactionId.Value())
	// Output:
	// TransactionId: 12345
}

// ExampleReq_negativeTransactionId demonstrates the error returned when
// a negative transactionId is provided.
func ExampleReq_negativeTransactionId() {
	_, err := rst.Req(rst.ReqInput{TransactionId: exampleTransactionIdNeg})
	if err != nil {
		fmt.Println("transactionId: invalid value")
	}
	// Output:
	// transactionId: invalid value
}
