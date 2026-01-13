package clearChargingProfile_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearChargingProfile"
)

const (
	exampleIdValue         = 123
	exampleConnectorId     = 1
	exampleStackLevelZero  = 0
	exampleIdOne           = 1
	exampleConnectorIdTwo  = 2
	exampleStackLevelThree = 3
	exampleNegativeId      = -1
)

// ExampleReq demonstrates creating a ClearChargingProfile.req message
// with no optional fields (clears all charging profiles).
func ExampleReq() {
	_, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ClearChargingProfile.req created successfully")
	// Output:
	// ClearChargingProfile.req created successfully
}

// ExampleReq_withId demonstrates creating a ClearChargingProfile.req message
// with a specific charging profile ID.
func ExampleReq_withId() {
	profileId := exampleIdValue

	req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     &profileId,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Id:", req.Id.Value())
	// Output:
	// Id: 123
}

// ExampleReq_withConnectorId demonstrates creating a ClearChargingProfile.req
// message with a connector ID.
func ExampleReq_withConnectorId() {
	connectorId := exampleConnectorId

	req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     nil,
		ConnectorId:            &connectorId,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	// Output:
	// ConnectorId: 1
}

// ExampleReq_withPurpose demonstrates creating a ClearChargingProfile.req
// message with a charging profile purpose.
func ExampleReq_withPurpose() {
	purpose := "TxProfile"

	req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: &purpose,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Purpose:", req.ChargingProfilePurpose.String())
	// Output:
	// Purpose: TxProfile
}

// ExampleReq_withStackLevel demonstrates creating a ClearChargingProfile.req
// message with a stack level.
func ExampleReq_withStackLevel() {
	stackLevel := exampleStackLevelZero

	req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             &stackLevel,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("StackLevel:", req.StackLevel.Value())
	// Output:
	// StackLevel: 0
}

// ExampleReq_withAllFields demonstrates creating a ClearChargingProfile.req
// message with all optional fields.
func ExampleReq_withAllFields() {
	profileId := exampleIdOne
	connectorId := exampleConnectorIdTwo
	purpose := "TxDefaultProfile"
	stackLevel := exampleStackLevelThree

	req, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     &profileId,
		ConnectorId:            &connectorId,
		ChargingProfilePurpose: &purpose,
		StackLevel:             &stackLevel,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Id:", req.Id.Value())
	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	fmt.Println("Purpose:", req.ChargingProfilePurpose.String())
	fmt.Println("StackLevel:", req.StackLevel.Value())
	// Output:
	// Id: 1
	// ConnectorId: 2
	// Purpose: TxDefaultProfile
	// StackLevel: 3
}

// ExampleReq_invalidPurpose demonstrates the error returned when
// an invalid charging profile purpose is provided.
func ExampleReq_invalidPurpose() {
	purpose := "Invalid"

	_, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: &purpose,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println("Error: invalid charging profile purpose")
	}
	// Output:
	// Error: invalid charging profile purpose
}

// ExampleReq_negativeId demonstrates the error returned when
// a negative ID is provided.
func ExampleReq_negativeId() {
	profileId := exampleNegativeId

	_, err := clearChargingProfile.Req(clearChargingProfile.ReqInput{
		Id:                     &profileId,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println("Error: invalid ID")
	}
	// Output:
	// Error: invalid ID
}
