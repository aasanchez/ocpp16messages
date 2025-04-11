//go:build example

package authorize

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aasanchez/ocpp16messages/types"
)

// --- Transport structs with JSON tags for OCPP 1.6J compliance ---

type RequestPayload struct {
	IdTag string `json:"idTag"`
}

type ConfirmationPayload struct {
	IdTagInfo struct {
		Status      types.AuthorizationStatus `json:"status"`
		ExpiryDate  *string                   `json:"expiryDate,omitempty"`
		ParentIdTag *string                   `json:"parentIdTag,omitempty"`
	} `json:"idTagInfo"`
}

func ExampleAuthorizeWorkflow() {
	fmt.Println("=== Full OCPP 1.6J-Compliant Authorize Workflow ===")

	// --- Step 1: Build domain-level Request ---
	idTagRaw := "ABC123456789"
	messageID := "msg-001"

	reqMsg, err := Request(idTagRaw)
	if err != nil {
		log.Fatalf("failed to construct request: %v", err)
	}

	// --- Step 2: Convert to JSON-compliant payload ---
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

	// --- Step 3: Simulate receiving JSON CALL ---
	var receivedCall []json.RawMessage
	if err := json.Unmarshal(callJSON, &receivedCall); err != nil {
		log.Fatalf("failed to parse CALL: %v", err)
	}

	var payload RequestPayload
	if err := json.Unmarshal(receivedCall[3], &payload); err != nil {
		log.Fatalf("failed to parse request payload: %v", err)
	}

	domainReq, err := Request(payload.IdTag)
	if err != nil {
		log.Fatalf("domain validation failed: %v", err)
	}
	if err := domainReq.Validate(); err != nil {
		log.Fatalf("request validation failed: %v", err)
	}
	fmt.Println("[CSMS] Authorize.req is valid ✅")

	// --- Step 4: Build domain-level AuthorizeConfirmation ---
	status := types.Accepted
	expiry := time.Now().Add(24 * time.Hour).UTC()
	parent, _ := types.IdToken("PARENT01")

	confInfo := types.IdTagInfoType{
		Status:      status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	confMsg, err := Confirmation(confInfo)
	if err != nil {
		log.Fatalf("failed to build response: %v", err)
	}
	if err := confMsg.Validate(); err != nil {
		log.Fatalf("response validation failed: %v", err)
	}
	fmt.Println("[CSMS] Authorize.conf is valid ✅")

	// --- Step 5: Convert to JSON-compliant CALLRESULT ---
	var respPayload ConfirmationPayload
	respPayload.IdTagInfo.Status = confMsg.IdTagInfo.Status

	if confMsg.IdTagInfo.ExpiryDate != nil {
		s := confMsg.IdTagInfo.ExpiryDate.Format(time.RFC3339)
		respPayload.IdTagInfo.ExpiryDate = &s
	}
	if confMsg.IdTagInfo.ParentIdTag != nil {
		s := confMsg.IdTagInfo.ParentIdTag.String()
		respPayload.IdTagInfo.ParentIdTag = &s
	}

	callResult := []any{3, messageID, respPayload}
	callResultJSON, err := json.MarshalIndent(callResult, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal CALLRESULT: %v", err)
	}

	fmt.Println("\n[CSMS -> CP] Sending CALLRESULT (Authorize.conf):")
	fmt.Println(string(callResultJSON))

	// --- Step 6: Simulate Charge Point parsing CALLRESULT ---
	var parsedResult []json.RawMessage
	if err := json.Unmarshal(callResultJSON, &parsedResult); err != nil {
		log.Fatalf("failed to parse CALLRESULT: %v", err)
	}

	var resultPayload ConfirmationPayload
	if err := json.Unmarshal(parsedResult[2], &resultPayload); err != nil {
		log.Fatalf("failed to parse response payload: %v", err)
	}

	var parsedParent *types.IdTokenType
	if resultPayload.IdTagInfo.ParentIdTag != nil {
		parentTag, err := types.IdToken(*resultPayload.IdTagInfo.ParentIdTag)
		if err != nil {
			log.Fatalf("invalid parentIdTag received: %v", err)
		}
		parsedParent = &parentTag
	}

	var parsedExpiry *time.Time
	if resultPayload.IdTagInfo.ExpiryDate != nil {
		ts, err := time.Parse(time.RFC3339, *resultPayload.IdTagInfo.ExpiryDate)
		if err != nil {
			log.Fatalf("invalid expiryDate received: %v", err)
		}
		parsedExpiry = &ts
	}

	parsedConf := ConfirmationMessage{
		IdTagInfo: types.IdTagInfoType{
			Status:      resultPayload.IdTagInfo.Status,
			ExpiryDate:  parsedExpiry,
			ParentIdTag: parsedParent,
		},
	}

	if err := parsedConf.Validate(); err != nil {
		log.Fatalf("[CP] Response validation failed: %v", err)
	}

	fmt.Println("[CP] Authorize.conf received and validated ✅")

	// Output:
}
