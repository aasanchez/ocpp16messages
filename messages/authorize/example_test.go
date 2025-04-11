//go:build example

package authorize_test

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	"github.com/aasanchez/ocpp16messages/types"
)

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

func ExampleCiString500_valid() {
	input := strings.Repeat("A", 499)

	cistr, err := types.CiString500(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 499
}

func ExampleCiString500_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := types.CiString500(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 501, max 500
}
