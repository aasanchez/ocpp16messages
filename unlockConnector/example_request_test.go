package unlockConnector_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/unlockConnector"
)

const (
	exampleConnectorId     = 1
	exampleConnectorIdZero = 0
)

// ExampleReq demonstrates creating a valid UnlockConnector.req message.
func ExampleReq() {
	req, err := unlockConnector.Req(unlockConnector.ReqInput{
		ConnectorId: exampleConnectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	// Output:
	// ConnectorId: 1
}

// ExampleReq_zeroConnectorId demonstrates the error returned when
// connectorId is zero (connector 0 refers to the Charge Point itself).
func ExampleReq_zeroConnectorId() {
	_, err := unlockConnector.Req(unlockConnector.ReqInput{
		ConnectorId: exampleConnectorIdZero,
	})
	if err != nil {
		fmt.Println("connectorId: invalid value")
	}
	// Output:
	// connectorId: invalid value
}
