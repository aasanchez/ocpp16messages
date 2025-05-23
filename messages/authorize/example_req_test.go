package authorize_test

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
)

func ExampleRequest() {
	idTag := "B85A-50CBE9678EC6"

	authorizeReq, err := authorize.Request(idTag)
	if err != nil {
		fmt.Printf("Error: failed to construct request: %v", err)
	}

	fmt.Printf("Authorize.req: %s\n", authorizeReq.String())
	// Output:
	// Authorize.req: {idTag:B85A-50CBE9678EC6}
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
		fmt.Printf("Error: invalid JSON array: %v", err)
	}

	var payload struct {
		IdTag string `json:"idTag"`
	}

	err = json.Unmarshal(msg[3], &payload)
	if err != nil {
		fmt.Printf("Error: invalid payload: %v", err)
	}

	req, err := authorize.Request(payload.IdTag)
	if err != nil {
		fmt.Printf("Error: authorize.Request failed: %v", err)
	}

	fmt.Println("Authorize.req:", req)
	// Output:
	// Authorize.req: {idTag:A632-E2BB0231072C}
}
