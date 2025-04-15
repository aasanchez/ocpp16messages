package authorize_test

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

type RequestPayload struct {
	IdTag string `json:"idTag"`
}

// FullMessage represents the full structure of a typical OCPP message
type FullMessage struct {
	MessageTypeId int                      `json:"messageTypeId"`
	UniqueId      string                   `json:"uniqueId"`
	Action        string                   `json:"action"`
	Payload       authorize.RequestMessage `json:"payload"`
}

type ConfirmationPayload struct {
	IdTagInfo struct {
		Status      types.AuthorizationStatus `json:"status"`
		ExpiryDate  *string                   `json:"expiryDate,omitempty"`
		ParentIdTag *string                   `json:"parentIdTag,omitempty"`
	} `json:"idTagInfo"`
}

func ExampleRequest() {
	idTagRaw := "ABC123456789"
	messageID := "msg-001"

	reqMsg, err := authorize.Request(idTagRaw)
	if err != nil {
		log.Fatalf("failed to construct request: %v", err)
	}

	reqPayload := RequestPayload{
		IdTag: reqMsg.IdTag.String(),
	}

	call := []any{2, messageID, "Authorize", reqPayload}
	callJSON, err := json.MarshalIndent(call, "", "  ")

	if err != nil {
		log.Fatalf("failed to marshal CALL: %v", err)
	}

	fmt.Println("\n[CP -> CSMS] Sending CALL (Authorize.req):")
	fmt.Println(string(callJSON))
	// Output:
	// [CP -> CSMS] Sending CALL (Authorize.req):
	// [
	//   2,
	//   "msg-001",
	//   "Authorize",
	//   {
	//     "idTag": "ABC123456789"
	//   }
	// ]
}

func ExampleConfirmation() {
	status := types.Accepted
	location, _ := time.LoadLocation("Europe/Madrid")
	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
	parent, err := types.IdToken("PARENT1234")

	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	idTagInfo := types.IdTagInfoType{
		Status:      status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	authorizeMsg, err := authorize.Confirmation(idTagInfo)
	if err != nil {
		log.Fatalf("failed to build confirmation message: %v", err)
	}

	fmt.Printf("Built Authorize.conf: %s\n", authorizeMsg.String())
	// Output:
	// Built Authorize.conf
}

func ExampleConfirmation_onlyStatus() {
	idTagInfo := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	authorizeMsg := authorize.ConfirmationMessage{IdTagInfo: idTagInfo}

	if err := authorizeMsg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed (status only)")
	// Output:
	// Validation passed (status only)
}

func ExampleConfirmation_statusWithExpireDate() {
	location, _ := time.LoadLocation("Europe/Madrid")
	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: nil,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed (with expiryDate)")
	// Output:
	// Validation passed (with expiryDate)
}

func ExampleConfirmation_statusWithParentIdTag() {
	parent, err := types.IdToken("GROUPXYZ")
	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &parent,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed (with parentIdTag)")
	// Output:
	// Validation passed (with parentIdTag)
}

func ExampleConfirmation_fully() {
	location, _ := time.LoadLocation("Europe/Madrid")
	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
	parent, err := types.IdToken("FULLTAG")

	if err != nil {
		log.Fatalf("failed to create parentIdTag: %v", err)
	}

	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	msg := authorize.ConfirmationMessage{IdTagInfo: info}

	if err := msg.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	fmt.Println("Validation passed (fully populated)")
	// Output:
	// Validation passed (fully populated)
}

func ExampleConfirmation_failValidation() {
	info := types.IdTagInfoType{
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
	// Expected validation failure
}

func ExampleRequest_build() {
	receivedJSON := `[ 2, "19223201", "Authorize", { "idTag": "ThisIsMySuperIDTag" } ]`

	var rawMsg []interface{}
	err := json.Unmarshal([]byte(receivedJSON), &rawMsg)

	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	messageTypeId, okTypeId := rawMsg[0].(float64)
	if !okTypeId {
		log.Fatalf("Invalid MessageTypeId type")
	}

	uniqueId, okUniqueId := rawMsg[1].(string)
	if !okUniqueId {
		log.Fatalf("Invalid UniqueId type")
	}

	action, okAction := rawMsg[2].(string)
	if !okAction {
		log.Fatalf("Invalid Action type")
	}

	payloadJSON, err := json.Marshal(rawMsg[3])
	if err != nil {
		log.Fatalf("Error marshaling payload: %v", err)
	}

	var payloadMap map[string]interface{}
	err = json.Unmarshal(payloadJSON, &payloadMap)

	if err != nil {
		log.Fatalf("Error unmarshaling payload: %v", err)
	}

	idTagStr, ok := payloadMap["idTag"].(string)
	if !ok {
		log.Fatalf("Error: idTag field is missing or invalid")
	}

	idTag, err := types.IdToken(idTagStr)
	if err != nil {
		log.Fatalf("Error creating IdTag: %v", err)
	}

	payload := authorize.RequestMessage{IdTag: idTag}

	fmt.Printf("Received Message Type: %d\n", int(messageTypeId))
	fmt.Printf("Received Action: %s\n", action)
	fmt.Printf("Received UniqueId: %s\n", uniqueId)

	if err := payload.Validate(); err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	fmt.Printf("Received and Validated RequestMessage: %s\n", payload.String())

	fmt.Printf("idTag: %s\n", payload.IdTag.String())
	// Output:
	// idTag
}
