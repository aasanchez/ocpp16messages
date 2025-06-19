package authorize_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

const (
	errFormat           = "Error: failed to construct request: %v\n"
	errRequestFormat    = "Error: authorize.Request failed: %v\n"
	errInvalidJSONInput = "Error: invalid JSON array: %v\n"
	errInvalidPayload   = "Error: invalid payload: %v\n"
	outputFormat        = "Authorize.req: {idTag:%s}\n"
)

func ExampleRequest() {
	input := authorizetypes.RequestPayload{IdTag: "B85A-50CBE9678EC6"}

	req, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errFormat, err)
	}

	fmt.Printf(outputFormat, req.IdTag.Value())

	// Output:
	// Authorize.req: {idTag:B85A-50CBE9678EC6}
}

func ExampleRequest_emptyIdTag() {
	input := authorizetypes.RequestPayload{IdTag: ""}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("This should not print")
	// Output:
	// Error: request: invalid idTag: ciString.Validate: value must not be empty
}

func ExampleRequest_too_long_idtag() {
	input := authorizetypes.RequestPayload{IdTag: strings.Repeat("A", 25)}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errFormat, err)
	}

	// Output:
	// Error: failed to construct request: request: invalid idTag: ciString.Validate: value exceeds maximum allowed length: got length 25, max 20
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
		fmt.Printf(errInvalidJSONInput, err)
	}

	var input authorizetypes.RequestPayload
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf(errInvalidPayload, err)
	}

	req, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errRequestFormat, err)
	}

	fmt.Printf(outputFormat, req.IdTag.Value())

	// Output:
	// Authorize.req: {idTag:A632-E2BB0231072C}
}

func ExampleRequest_parse_invalid() {
	const raw = `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{
			"idTag":"72366A26-ECC7-432B-BD30-A3DB781CBCB0"
		}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf(errInvalidJSONInput, err)
	}

	var input authorizetypes.RequestPayload
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf(errInvalidPayload, err)
	}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errRequestFormat, err)
	}

	// Output:
	// Error: authorize.Request failed: request: invalid idTag: ciString.Validate: value exceeds maximum allowed length: got length 36, max 20
}

func ExampleRequest_parse_empty() {
	const raw = `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf(errInvalidJSONInput, err)
	}

	var input authorizetypes.RequestPayload
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf(errInvalidPayload, err)
	}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errRequestFormat, err)
	}

	// Output:
	// Error: authorize.Request failed: request: invalid idTag: ciString.Validate: value must not be empty
}

func ExampleRequest_parse_idTag_empty() {
	const raw = `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{
			"idTag":""
		}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf(errInvalidJSONInput, err)
	}

	var input authorizetypes.RequestPayload
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf(errInvalidPayload, err)
	}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errRequestFormat, err)
	}

	// Output:
	// Error: authorize.Request failed: request: invalid idTag: ciString.Validate: value must not be empty
}

func ExampleRequest_parse_idTag_NotFound() {
	const raw = `
	[
		2,
		"0FDD6C86-1382-4E98-A36F-362B3D6960D7",
		"Authorize",
		{
			"Tag":"A632-E2BB0231072C"
		}
	]`

	var msg []json.RawMessage
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		fmt.Printf(errInvalidJSONInput, err)
	}

	var input authorizetypes.RequestPayload
	if err := json.Unmarshal(msg[3], &input); err != nil {
		fmt.Printf(errInvalidPayload, err)
	}

	_, err := authorize.Request(input)
	if err != nil {
		fmt.Printf(errRequestFormat, err)
	}

	// Output:
	// Error: authorize.Request failed: request: invalid idTag: ciString.Validate: value must not be empty
}
