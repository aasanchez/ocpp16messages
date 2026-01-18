package reserveNow_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/reserveNow"
)

const (
	exampleReservationId = 1
	exampleConnectorId   = 1
	exampleIdTag         = "RFID-TAG-12345"
	exampleExpiryDate    = "2025-01-15T10:00:00Z"
	exampleParentIdTag   = "PARENT-12345"
)

// ExampleReq demonstrates creating a valid ReserveNow.req message with only
// the required fields.
func ExampleReq() {
	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: exampleReservationId,
		ConnectorId:   exampleConnectorId,
		IdTag:         exampleIdTag,
		ExpiryDate:    exampleExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ReservationId:", req.ReservationId.Value())
	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	fmt.Println("IdTag:", req.IdTag.Value())
	// Output:
	// ReservationId: 1
	// ConnectorId: 1
	// IdTag: RFID-TAG-12345
}

// ExampleReq_withParentIdTag demonstrates creating a ReserveNow.req message
// with all fields including the optional parentIdTag.
func ExampleReq_withParentIdTag() {
	parentIdTag := exampleParentIdTag

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: exampleReservationId,
		ConnectorId:   exampleConnectorId,
		IdTag:         exampleIdTag,
		ExpiryDate:    exampleExpiryDate,
		ParentIdTag:   &parentIdTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.Value())
	fmt.Println("ParentIdTag:", req.ParentIdTag.Value())
	// Output:
	// IdTag: RFID-TAG-12345
	// ParentIdTag: PARENT-12345
}

// ExampleReq_emptyIdTag demonstrates the error returned when an empty idTag
// is provided.
func ExampleReq_emptyIdTag() {
	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: exampleReservationId,
		ConnectorId:   exampleConnectorId,
		IdTag:         "",
		ExpiryDate:    exampleExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println("idTag: value cannot be empty")
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_invalidExpiryDate demonstrates the error returned when an invalid
// expiry date is provided.
func ExampleReq_invalidExpiryDate() {
	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: exampleReservationId,
		ConnectorId:   exampleConnectorId,
		IdTag:         exampleIdTag,
		ExpiryDate:    "invalid-date",
		ParentIdTag:   nil,
	})
	if err != nil {
		fmt.Println("expiryDate: invalid value")
	}
	// Output:
	// expiryDate: invalid value
}
