package changeAvailability_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeAvailability"
)

const (
	connectorIdLabel = "ConnectorId:"
	typeLabel        = "Type:"
)

// ExampleReq demonstrates creating a valid ChangeAvailability.req message
// to set a connector to Inoperative.
func ExampleReq() {
	req, err := changeAvailability.Req(changeAvailability.ReqInput{
		ConnectorId: 1,
		Type:        "Inoperative",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(connectorIdLabel, req.ConnectorId.Value())
	fmt.Println(typeLabel, req.Type.String())
	// Output:
	// ConnectorId: 1
	// Type: Inoperative
}

// ExampleReq_operative demonstrates creating a ChangeAvailability.req message
// to set a connector to Operative.
func ExampleReq_operative() {
	req, err := changeAvailability.Req(changeAvailability.ReqInput{
		ConnectorId: 2,
		Type:        "Operative",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(connectorIdLabel, req.ConnectorId.Value())
	fmt.Println(typeLabel, req.Type.String())
	// Output:
	// ConnectorId: 2
	// Type: Operative
}

// ExampleReq_entireChargePoint demonstrates creating a ChangeAvailability.req
// message for the entire Charge Point (connectorId = 0).
func ExampleReq_entireChargePoint() {
	req, err := changeAvailability.Req(changeAvailability.ReqInput{
		ConnectorId: 0,
		Type:        "Inoperative",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(connectorIdLabel, req.ConnectorId.Value())
	fmt.Println(typeLabel, req.Type.String())
	// Output:
	// ConnectorId: 0
	// Type: Inoperative
}

// ExampleReq_invalidType demonstrates the error returned when
// an invalid availability type is provided.
func ExampleReq_invalidType() {
	_, err := changeAvailability.Req(changeAvailability.ReqInput{
		ConnectorId: 1,
		Type:        "Unknown",
	})
	if err != nil {
		fmt.Println("Error: invalid type")
	}
	// Output:
	// Error: invalid type
}

// ExampleReq_negativeConnectorId demonstrates the error returned when
// a negative connector ID is provided.
func ExampleReq_negativeConnectorId() {
	_, err := changeAvailability.Req(changeAvailability.ReqInput{
		ConnectorId: -1,
		Type:        "Operative",
	})
	if err != nil {
		fmt.Println("Error: invalid connector ID")
	}
	// Output:
	// Error: invalid connector ID
}
