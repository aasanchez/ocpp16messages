package authorize_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

type ConfirmationPayload struct {
	Status      types.AuthorizationStatus `json:"status"`
	ExpiryDate  *string                   `json:"expiryDate,omitempty"`
	ParentIdTag *string                   `json:"parentIdTag,omitempty"`
}

func testIdTag() string {
	idtag := "B85A-50CBE9678EC6"

	return idtag
}

func testConfirmationFields() (string, string, string) {
	status := "Accepted"
	expiry := "2027-04-12T10:03:04-04:00"
	parent := "A632-E2BB0231072C"

	return status, expiry, parent
}

func ExampleRequest() {
	idTag := testIdTag()

	authorizeReq, err := authorize.Request(idTag)
	if err != nil {
		log.Fatalf("failed to construct request: %v", err)
	}

	fmt.Printf("Authorize.req: %s\n", authorizeReq.String())
	// Output:
	// Authorize.req: {idTag=B85A-50CBE9678EC6}
}

func ExampleRequest_parse() {
	const raw = `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{ "idTag": "A632-E2BB0231072C" }
	]`

	var msg []json.RawMessage

	err := json.Unmarshal([]byte(raw), &msg)
	if err != nil {
		log.Fatalf("invalid JSON array: %v", err)
	}

	var payload struct {
		IdTag string `json:"idTag"`
	}

	err = json.Unmarshal(msg[3], &payload)
	if err != nil {
		log.Fatalf("invalid payload: %v", err)
	}

	req, err := authorize.Request(payload.IdTag)
	if err != nil {
		log.Fatalf("authorize.Request failed: %v", err)
	}

	fmt.Println("Authorize.req:", req)
	// Output:
	// Authorize.req: {idTag=A632-E2BB0231072C}
}

func ExampleConfirmation() {
	statusStr, expiryStr, parentStr := testConfirmationFields()

	parent, err := types.IdToken(parentStr)
	if err != nil {
		panic(err)
	}

	expiry, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		panic(err)
	}

	status := types.AuthorizationStatus(statusStr)

	idTagInfo := types.IdTagInfoType{
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
		Status:      status,
	}

	authorizeMsg, err := authorize.Confirmation(idTagInfo)
	if err != nil {
		log.Fatalf("failed to build confirmation message: %v", err)
	}

	fmt.Printf("Authorize.conf: %s\n", authorizeMsg.String())
	// Output:
	// Authorize.conf: {status=Accepted, expiryDate=2027-04-12T10:03:04-04:00, parentIdTag=A632-E2BB0231072C}
}

func ExampleConfirmation_parse() {
	receivedJSON := `[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{
			"status": "Accepted",
			"expiryDate": "2027-04-12T10:03:04-04:00",
			"parentIdTag": "B85A-50CBE9678EC6"
		}
	]`

	var msg []json.RawMessage

	err := json.Unmarshal([]byte(receivedJSON), &msg)
	if err != nil {
		log.Fatalf("invalid JSON array: %v", err)
	}

	var payload ConfirmationPayload
	if err := json.Unmarshal(msg[3], &payload); err != nil {
		log.Fatalf("invalid payload: %v", err)
	}

	var expiry *time.Time

	time, err := time.Parse(time.RFC3339, *payload.ExpiryDate)
	if err != nil {
		log.Fatalf("invalid expiryDate format: %v", err)
	}

	expiry = &time

	var parent *types.IdTokenType

	parentIdTag, err := types.IdToken(*payload.ParentIdTag)
	if err != nil {
		log.Fatalf("invalid parentIdTag: %v", err)
	}

	parent = &parentIdTag

	idTagInfo := types.IdTagInfoType{
		Status:      payload.Status,
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}

	conf, err := authorize.Confirmation(idTagInfo)
	if err != nil {
		log.Fatalf("authorize.Confirmation failed: %v", err)
	}

	fmt.Println("Authorize.conf:", conf)
	// Output:
	// Authorize.conf: {status=Accepted, expiryDate=2027-04-12T10:03:04-04:00, parentIdTag=B85A-50CBE9678EC6}
}

// func ExampleConfirmation_parse2() {
// 	status := types.Accepted
// 	location, _ := time.LoadLocation("America/Caracas")
// 	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
// 	parent, err := types.IdToken("PARENT1234")

// 	if err != nil {
// 		log.Fatalf("failed to create parentIdTag: %v", err)
// 	}

// 	idTagInfo := types.IdTagInfoType{
// 		ExpiryDate:  &expiry,
// 		ParentIdTag: &parent,
// 		Status:      status,
// 	}

// 	authorizeMsg, err := authorize.Confirmation(idTagInfo)
// 	if err != nil {
// 		log.Fatalf("failed to build confirmation message: %v", err)
// 	}

// 	fmt.Printf("Authorize.conf: %s\n", authorizeMsg.String())
// 	// Output:
// 	// Authorize.conf: {status=Accepted, expiryDate=2027-04-12T10:03:04-04:00, parentIdTag=PARENT1234}
// }
