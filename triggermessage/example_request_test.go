package triggermessage_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/triggermessage"
)

const (
	labelRequestedMessage = "RequestedMessage:"
	labelConnectorId      = "ConnectorId:"
	connectorIdZero       = 0
	connectorIdOne        = 1
)

// ExampleReq demonstrates creating a valid TriggerMessage.req message
// to trigger a Heartbeat message.
func ExampleReq() {
	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	// Output:
	// RequestedMessage: Heartbeat
}

// ExampleReq_withConnectorId demonstrates creating a TriggerMessage.req
// message with an optional connectorId.
func ExampleReq_withConnectorId() {
	connectorId := connectorIdOne

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorId:      &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	fmt.Println(labelConnectorId, req.ConnectorId.Value())
	// Output:
	// RequestedMessage: StatusNotification
	// ConnectorId: 1
}

// ExampleReq_metervalues demonstrates triggering a MeterValues message.
func ExampleReq_metervalues() {
	connectorId := connectorIdZero

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorId:      &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	fmt.Println(labelConnectorId, req.ConnectorId.Value())
	// Output:
	// RequestedMessage: MeterValues
	// ConnectorId: 0
}

// ExampleReq_invalidMessage demonstrates the error returned when
// an invalid message trigger is provided.
func ExampleReq_invalidMessage() {
	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorId:      nil,
	})
	if err != nil {
		fmt.Println("Error: invalid requestedMessage")
	}
	// Output:
	// Error: invalid requestedMessage
}

// ExampleReq_emptyMessage demonstrates the error returned when
// an empty message trigger is provided.
func ExampleReq_emptyMessage() {
	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "",
		ConnectorId:      nil,
	})
	if err != nil {
		fmt.Println("Error: invalid requestedMessage")
	}
	// Output:
	// Error: invalid requestedMessage
}
