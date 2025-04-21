package authorize_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type ConfirmationPayload struct {
	Status      authorizetypes.AuthorizationStatus `json:"status"`
	ExpiryDate  *string                            `json:"expiryDate,omitempty"`
	ParentIdTag *string                            `json:"parentIdTag,omitempty"`
}

func ExampleConfirmation() {
	statusStr := "Accepted"
	expiryStr := "2027-04-12T10:03:04-04:00"
	parentStr := "A632-E2BB0231072C"

	parent, err := authorizetypes.IdToken(parentStr)
	if err != nil {
		panic(err)
	}

	expiry, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		panic(err)
	}

	status := authorizetypes.AuthorizationStatus(statusStr)

	idTagInfo := authorizetypes.IdTagInfoType{
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
	// Authorize.conf: {status:Accepted, expiryDate:2027-04-12T10:03:04-04:00, parentIdTag:A632-E2BB0231072C}
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

	var parent *authorizetypes.IdTokenType

	parentIdTag, err := authorizetypes.IdToken(*payload.ParentIdTag)
	if err != nil {
		log.Fatalf("invalid parentIdTag: %v", err)
	}

	parent = &parentIdTag

	idTagInfo := authorizetypes.IdTagInfoType{
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
	// Authorize.conf: {status:Accepted, expiryDate:2027-04-12T10:03:04-04:00, parentIdTag:B85A-50CBE9678EC6}
}

func ExampleConfirmationMessage_Validate() {
	info := authorizetypes.IdTagInfoType{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	err := msg.Validate()
	if err != nil {
		fmt.Printf("Expected validation failure: %v\n", err)
	} else {
		fmt.Print("Expected validation to fail, but it passed")
	}

	// Output:
	// Expected validation to fail, but it passed
}

func ExampleConfirmationMessage_Validate_invalid() {
	info := authorizetypes.IdTagInfoType{
		Status:      "InvalidStatus",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	err := msg.Validate()
	if err != nil {
		fmt.Printf("Expected validation failure: %v\n", err)
	} else {
		log.Fatal("Expected validation to fail, but it passed")
	}

	// Output:
	// Expected validation failure: ConfirmationMessage validation failed: invalid authorization status: InvalidStatus
}
