package starttransaction_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/starttransaction"
)

const testReservationId = 42

// ExampleReq demonstrates creating a valid StartTransaction.req message
// with all required fields.
func ExampleReq() {
	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationId: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	fmt.Println("IdTag:", req.IdTag.String())
	fmt.Println("MeterStart:", req.MeterStart.Value())
	// Output:
	// ConnectorId: 1
	// IdTag: RFID-TAG-12345
	// MeterStart: 1000
}

// ExampleReq_withReservation demonstrates creating a StartTransaction.req
// message that includes a reservation ID.
func ExampleReq_withReservation() {
	reservationId := testReservationId

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   2,
		IdTag:         "RFID-TAG-67890",
		MeterStart:    500,
		Timestamp:     "2025-01-15T11:00:00Z",
		ReservationId: &reservationId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorId:", req.ConnectorId.Value())
	fmt.Println("IdTag:", req.IdTag.String())
	fmt.Println("HasReservation:", req.ReservationId != nil)
	fmt.Println("ReservationId:", req.ReservationId.Value())
	// Output:
	// ConnectorId: 2
	// IdTag: RFID-TAG-67890
	// HasReservation: true
	// ReservationId: 42
}

// ExampleReq_emptyIdTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleReq_emptyIdTag() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   1,
		IdTag:         "",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationId: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_invalidTimestamp demonstrates the error returned when
// an invalid timestamp is provided.
func ExampleReq_invalidTimestamp() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "not-a-timestamp",
		ReservationId: nil,
	})
	if err != nil {
		fmt.Println("Invalid timestamp error")
	}
	// Output:
	// Invalid timestamp error
}

// ExampleReq_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleReq_multipleErrors() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   -1,
		IdTag:         "",
		MeterStart:    -1,
		Timestamp:     "invalid",
		ReservationId: nil,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
