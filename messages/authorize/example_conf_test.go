package authorize_test

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type confirmationPayload struct {
	Status      string  `json:"status"`
	ExpiryDate  *string `json:"expiryDate,omitempty"`
	ParentIdTag *string `json:"parentIdTag,omitempty"`
}

func ExampleConfirmation() {
	expiryStr := "2027-04-12T10:03:04-04:00"
	parentStr := "A632-E2BB0231072C"
	statusStr := "Accepted"

	expiry, err := sharedtypes.DateTime(expiryStr)
	if err != nil {
		panic(err)
	}

	parent, err := authorizetypes.IdToken(parentStr)
	if err != nil {
		panic(err)
	}

	status, err := authorizetypes.AuthorizationStatus(statusStr)
	if err != nil {
		panic(err)
	}

	idTagInfo := authorizetypes.IdTagInfoType{
		Status:      status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	_, err = authorize.Confirmation(idTagInfo)
	if err != nil {
		fmt.Printf("Error: failed to build confirmation message: %v\n", err)
		return
	}

	fmt.Printf(
		"Authorize.conf: {status:%s, expiryDate:%s, parentIdTag:%s}\n",
		statusStr,
		expiryStr,
		parentStr,
	)

	// Output:
	// Authorize.conf: {status:Accepted, expiryDate:2027-04-12T10:03:04-04:00, parentIdTag:A632-E2BB0231072C}
}

func ExampleConfirmation_parse() {
	raw := `
	[
  	3,
  	"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
  	{
    	"status": "Accepted",
    	"expiryDate": "2027-04-12T10:03:04-04:00",
    	"parentIdTag": "B85A-50CBE9678EC6"
  	}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf("Error: invalid JSON array: %v\n", err)
		return
	}

	var payload confirmationPayload
	if err := json.Unmarshal(msg[2], &payload); err != nil {
		fmt.Printf("Error: invalid payload: %v\n", err)
		return
	}

	status, err := authorizetypes.AuthorizationStatus(payload.Status)
	if err != nil {
		fmt.Printf("Error: invalid status: %v\n", err)
		return
	}

	var expiry *sharedtypes.DateTimeType
	if payload.ExpiryDate != nil {
		t, err := sharedtypes.DateTime(*payload.ExpiryDate)
		if err != nil {
			fmt.Printf("Error: invalid expiryDate format: %v\n", err)
			return
		}

		expiry = &t
	}

	var parent *authorizetypes.IdTokenType
	if payload.ParentIdTag != nil {
		parentID, err := authorizetypes.IdToken(*payload.ParentIdTag)
		if err != nil {
			fmt.Printf("Error: invalid parentIdTag: %v\n", err)
			return
		}

		parent = &parentID
	}

	info := authorizetypes.IdTagInfoType{
		Status:      status,
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}

	_, err = authorize.Confirmation(info)
	if err != nil {
		fmt.Printf("Error: authorize.Confirmation failed: %v\n", err)
		return
	}

	fmt.Printf(
		"Authorize.conf: {status:%s, expiryDate:%s, parentIdTag:%s}\n",
		payload.Status,
		*payload.ExpiryDate,
		*payload.ParentIdTag,
	)

	// Output:
	// Authorize.conf: {status:Accepted, expiryDate:2027-04-12T10:03:04-04:00, parentIdTag:B85A-50CBE9678EC6}
}

func ExampleConfirmationMessage_Validate_valid() {
	status, _ := authorizetypes.AuthorizationStatus("Accepted")

	info := authorizetypes.IdTagInfoType{
		Status: status,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	err := msg.Validate()
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
	} else {
		fmt.Println("Validation passed")
	}

	// Output:
	// Validation passed
}

func ExampleConfirmationMessage_Validate_invalid() {
	status := authorizetypes.AuthorizationStatusType{}

	info := authorizetypes.IdTagInfoType{
		Status: status,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	err := msg.Validate()
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Expected failure, but validation passed")
	}

	// Output:
	// Validation failed: confirmationMessage.Validate: invalid authorization status: invalid authorization status: ""
}
