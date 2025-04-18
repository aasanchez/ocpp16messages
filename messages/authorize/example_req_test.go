package authorize_test

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	types "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleRequest() {
	input := types.RequestMessageInput{IdTag: "B85A-50CBE9678EC6"}

	req, err := authorize.Request(input)
	if err != nil {
		fmt.Printf("Error: failed to construct request: %v\n", err)
		return
	}

	if err := req.Validate(); err != nil {
		fmt.Printf("Error: validation failed: %v\n", err)
		return
	}

	fmt.Printf("Authorize.req: {idTag:%s}\n", input.IdTag)

	// Output:
	// Authorize.req: {idTag:B85A-50CBE9678EC6}
}

func ExampleRequest_parse() {
	const raw = `
	[
		2,
  	"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
  	"Authorize",
  	{
    	"idTag":"A632-E2BB0231072C"
  	}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf("Error: invalid JSON array: %v\n", err)
		return
	}

	var input types.RequestMessageInput
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf("Error: invalid payload: %v\n", err)
		return
	}

	req, err := authorize.Request(input)
	if err != nil {
		fmt.Printf("Error: authorize.Request failed: %v\n", err)
		return
	}

	if err := req.Validate(); err != nil {
		fmt.Printf("Error: validation failed: %v\n", err)
		return
	}

	fmt.Printf("Authorize.req: {idTag:%s}\n", input.IdTag)

	// Output:
	// Authorize.req: {idTag:A632-E2BB0231072C}
}
