package remoteStopTransaction

import (
	"fmt"
	"strconv"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a
// RemoteStopTransaction.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier of the transaction to stop.
	TransactionId int
}

// ReqMessage represents an OCPP 1.6 RemoteStopTransaction.req message.
type ReqMessage struct {
	TransactionId st.Integer
}

// Req creates a RemoteStopTransaction.req message from the given input.
// It validates all fields and returns an error if:
//   - TransactionId is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	transactionId, err := st.NewInteger(strconv.Itoa(input.TransactionId))
	if err != nil {
		return ReqMessage{}, fmt.Errorf("transactionId: %w", err)
	}

	return ReqMessage{
		TransactionId: transactionId,
	}, nil
}
