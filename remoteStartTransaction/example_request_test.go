package remoteStartTransaction_test

import (
	"fmt"

	rst "github.com/aasanchez/ocpp16messages/remoteStartTransaction"
)

const (
	connectorIdOne        = 1
	connectorIdNegative   = -1
	testExampleValidIdTag = "RFID-TAG-12345"
)

// ExampleReq demonstrates creating a valid RemoteStartTransaction.req message
// with only the required idTag field.
func ExampleReq() {
	req, err := rst.Req(rst.ReqInput{
		IdTag:       testExampleValidIdTag,
		ConnectorId: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.Value())
	// Output:
	// IdTag: RFID-TAG-12345
}

// ExampleReq_withConnectorId demonstrates creating a RemoteStartTransaction.req
// message with both idTag and connectorId.
func ExampleReq_withConnectorId() {
	connectorId := connectorIdOne

	req, err := rst.Req(rst.ReqInput{
		IdTag:       testExampleValidIdTag,
		ConnectorId: &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.Value())
	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	// Output:
	// IdTag: RFID-TAG-12345
	// ConnectorId: 1
}

// ExampleReq_emptyIdTag demonstrates the error returned when
// an empty idTag is provided.
func ExampleReq_emptyIdTag() {
	_, err := rst.Req(rst.ReqInput{IdTag: "", ConnectorId: nil})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_idTagTooLong demonstrates the error returned when
// the idTag exceeds the maximum length of 20 characters.
func ExampleReq_idTagTooLong() {
	// 23 chars, max is 20
	_, err := rst.Req(rst.ReqInput{
		IdTag:       "RFID-ABC123456789012345",
		ConnectorId: nil,
	})
	if err != nil {
		fmt.Println("idTag: exceeds maximum length")
	}
	// Output:
	// idTag: exceeds maximum length
}

// ExampleReq_invalidConnectorId demonstrates the error returned when
// the connectorId is negative.
func ExampleReq_invalidConnectorId() {
	connectorId := connectorIdNegative

	_, err := rst.Req(rst.ReqInput{
		IdTag:       testExampleValidIdTag,
		ConnectorId: &connectorId,
	})
	if err != nil {
		fmt.Println("connectorId: invalid value")
	}
	// Output:
	// connectorId: invalid value
}
