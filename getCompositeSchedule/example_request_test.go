package getCompositeSchedule_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule"
)

const (
	exampleConnectorIdOne    = 1
	exampleConnectorIdZero   = 0
	exampleDurationThreeHund = 300
	exampleDurationSixHund   = 600
	exampleNegativeValue     = -1

	outConnectorId = "ConnectorId:"
	outDuration    = "Duration:"
)

// ExampleReq demonstrates creating a GetCompositeSchedule.req message
// with required fields only.
func ExampleReq() {
	req, err := getCompositeSchedule.Req(getCompositeSchedule.ReqInput{
		ConnectorId:      exampleConnectorIdOne,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorId, req.ConnectorId.Value())
	fmt.Println(outDuration, req.Duration.Value())
	// Output:
	// ConnectorId: 1
	// Duration: 300
}

// ExampleReq_withChargingRateUnit demonstrates creating a
// GetCompositeSchedule.req message with the optional ChargingRateUnit field.
func ExampleReq_withChargingRateUnit() {
	unit := "W"

	req, err := getCompositeSchedule.Req(getCompositeSchedule.ReqInput{
		ConnectorId:      exampleConnectorIdOne,
		Duration:         exampleDurationSixHund,
		ChargingRateUnit: &unit,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorId, req.ConnectorId.Value())
	fmt.Println(outDuration, req.Duration.Value())
	fmt.Println("ChargingRateUnit:", req.ChargingRateUnit.String())
	// Output:
	// ConnectorId: 1
	// Duration: 600
	// ChargingRateUnit: W
}

// ExampleReq_entireChargePoint demonstrates requesting a composite schedule
// for the entire Charge Point by using ConnectorId 0.
func ExampleReq_entireChargePoint() {
	req, err := getCompositeSchedule.Req(getCompositeSchedule.ReqInput{
		ConnectorId:      exampleConnectorIdZero,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorId, req.ConnectorId.Value())
	fmt.Println(outDuration, req.Duration.Value())
	// Output:
	// ConnectorId: 0
	// Duration: 300
}

// ExampleReq_invalidConnectorId demonstrates the error returned when
// a negative ConnectorId is provided.
func ExampleReq_invalidConnectorId() {
	_, err := getCompositeSchedule.Req(getCompositeSchedule.ReqInput{
		ConnectorId:      exampleNegativeValue,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid connector ID")
	}
	// Output:
	// Error: invalid connector ID
}

// ExampleReq_invalidChargingRateUnit demonstrates the error returned when
// an invalid ChargingRateUnit is provided.
func ExampleReq_invalidChargingRateUnit() {
	unit := "X"

	_, err := getCompositeSchedule.Req(getCompositeSchedule.ReqInput{
		ConnectorId:      exampleConnectorIdOne,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: &unit,
	})
	if err != nil {
		fmt.Println("Error: invalid charging rate unit")
	}
	// Output:
	// Error: invalid charging rate unit
}
