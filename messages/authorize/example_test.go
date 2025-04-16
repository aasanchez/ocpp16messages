package authorize_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

type ReqPayload struct {
	IdTag string `json:"idTag"`
}

// FullMessage represents the full structure of a typical OCPP message
type ConfirmationMessage struct {
	MessageTypeId int                 `json:"messageTypeId"`
	UniqueId      string              `json:"uniqueId"`
	Action        string              `json:"action"`
	Payload       ConfirmationPayload `json:"payload"`
}

type ConfirmationPayload struct {
	IdTagInfo struct {
		Status      types.AuthorizationStatus `json:"status"`
		ExpiryDate  *string                   `json:"expiryDate,omitempty"`
		ParentIdTag *string                   `json:"parentIdTag,omitempty"`
	} `json:"idTagInfo"`
}

var idTag = "A632-E2BB0231072C"

func ExampleRequest() {
	// Scope of this package
	authorizeReq, err := authorize.Request(idTag)
	if err != nil {
		log.Fatalf("failed to construct request: %v", err)
	}

	// End of Scope of this package
	authorizeRequestPayload := ReqPayload{
		IdTag: authorizeReq.IdTag.String(),
	}

	call := []any{
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		authorizeRequestPayload}

	callJSON, err := json.MarshalIndent(call, "", "  ")

	if err != nil {
		log.Fatalf("failed to marshal CALL: %v", err)
	}

	fmt.Println(string(callJSON))
	// Output:
	// [
	//   2,
	//   "0FDD6C86-1382-4E98-A36F-362B3D6960D7",
	//   "Authorize",
	//   {
	//     "idTag": "A632-E2BB0231072C"
	//   }
	// ]
}

func ExampleRequest_parse() {
	receivedJSON := `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{
			"idTag": "A632-E2BB0231072C"
		}
	]`

	var authorize []interface{}

	err := json.Unmarshal([]byte(receivedJSON), &authorize)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// fmt.Print(authorize)
	// messageTypeId, error := authorize[0].(float64)
	// if !error {
	// 	log.Fatalf("Invalid MessageTypeId type")
	// }

	// uniqueId, error := authorize[1].(string)
	// if !error {
	// 	log.Fatalf("Invalid UniqueId type")
	// }

	// action, okAction := authorize[2].(string)
	// if !okAction {
	// 	log.Fatalf("Invalid Action type")
	// }

	// payloadJSON, err := json.Marshal(authorize[3])
	// if err != nil {
	// 	log.Fatalf("Error marshaling payload: %v", err)
	// }

	// var payloadMap map[string]interface{}
	// err = json.Unmarshal(payloadJSON, &payloadMap)

	// if err != nil {
	// 	log.Fatalf("Error unmarshaling payload: %v", err)
	// }

	// idTagStr, ok := payloadMap["idTag"].(string)
	// if !ok {
	// 	log.Fatalf("Error: idTag field is missing or invalid")
	// }

	// idTag, err := types.IdToken(idTagStr)
	// if err != nil {
	// 	log.Fatalf("Error creating IdTag: %v", err)
	// }

	// payload := authorize.ReqPayload{IdTag: idTag}

	// fmt.Printf("Received Message Type: %d\n", int(messageTypeId))
	// fmt.Printf("Received Action: %s\n", action)
	// fmt.Printf("Received UniqueId: %s\n", uniqueId)

	// if err := payload.Validate(); err != nil {
	// 	log.Fatalf("Validation failed: %v", err)
	// }

	// fmt.Printf("Received and Validated RequestMessage: %s\n", payload.String())

	// fmt.Printf("idTag: %s\n", payload.IdTag.String())
	// Output:
	// Received Message Type: 2
	// Received Action: Authorize
	// Received UniqueId: 19223201
	// Received and Validated RequestMessage: {idTag=ThisIsMySuperIDTag}
	// idTag: ThisIsMySuperIDTag
}

// func ExampleConfirmation() {
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

// func ExampleConfirmation_onlyStatus() {
// 	idTagInfo := types.IdTagInfoType{
// 		ExpiryDate:  nil,
// 		ParentIdTag: nil,
// 		Status:      types.Accepted,
// 	}

// 	authorizeMsg := authorize.ConfirmationMessage{IdTagInfo: idTagInfo}

// 	if err := authorizeMsg.Validate(); err != nil {
// 		log.Fatalf("validation failed: %v", err)
// 	}

// 	fmt.Println("Validation passed (status only)")
// 	// Output:
// 	// Validation passed (status only)
// }

// func ExampleConfirmation_statusWithExpireDate() {
// 	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC)

// 	idTagInfo := types.IdTagInfoType{
// 		ExpiryDate:  &expiry,
// 		ParentIdTag: nil,
// 		Status:      types.Accepted,
// 	}

// 	msg := authorize.ConfirmationMessage{IdTagInfo: idTagInfo}

// 	if err := msg.Validate(); err != nil {
// 		log.Fatalf("validation failed: %v", err)
// 	}

// 	fmt.Print(msg)
// 	// Output:
// 	// {status=Accepted, expiryDate=2027-04-12T14:03:04Z}
// }

// func ExampleConfirmation_statusWithParentIdTag() {
// 	parent, err := types.IdToken("GROUPXYZ")
// 	if err != nil {
// 		log.Fatalf("failed to create parentIdTag: %v", err)
// 	}

// 	info := types.IdTagInfoType{
// 		ExpiryDate:  nil,
// 		ParentIdTag: &parent,
// 		Status:      types.Accepted,
// 	}

// 	msg := authorize.ConfirmationMessage{IdTagInfo: info}

// 	if err := msg.Validate(); err != nil {
// 		log.Fatalf("validation failed: %v", err)
// 	}

// 	fmt.Println("Validation passed (with parentIdTag)")
// 	// Output:
// 	// Validation passed (with parentIdTag)
// }

// func ExampleConfirmation_fully() {
// 	location, _ := time.LoadLocation("Europe/Madrid")
// 	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
// 	parent, err := types.IdToken("FULLTAG")

// 	if err != nil {
// 		log.Fatalf("failed to create parentIdTag: %v", err)
// 	}

// 	info := types.IdTagInfoType{
// 		Status:      types.Accepted,
// 		ExpiryDate:  &expiry,
// 		ParentIdTag: &parent,
// 	}

// 	msg := authorize.ConfirmationMessage{IdTagInfo: info}

// 	if err := msg.Validate(); err != nil {
// 		log.Fatalf("validation failed: %v", err)
// 	}

// 	fmt.Println("Validation passed (fully populated)")
// 	// Output:
// 	// Validation passed (fully populated)
// }

// func ExampleConfirmation_failValidation() {
// 	info := types.IdTagInfoType{
// 		Status:      "InvalidStatus",
// 		ExpiryDate:  nil,
// 		ParentIdTag: nil,
// 	}

// 	msg := authorize.ConfirmationMessage{IdTagInfo: info}

// 	err := msg.Validate()
// 	if err != nil {
// 		fmt.Printf("Expected validation failure: %v\n", err)
// 	} else {
// 		log.Fatal("Expected validation to fail, but it passed")
// 	}
// 	// Output:
// 	// Expected validation failure: ConfirmationMessage validation failed: invalid authorization status: InvalidStatus
// }
