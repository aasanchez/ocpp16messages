package authorize_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
)

// ExampleReq demonstrates creating a valid Authorize.req message
// with a properly formatted ID tag using the ReqInput struct.
func ExampleReq() {
	req, err := authorize.Req(authorize.ReqInput{IdTag: "RFID-TAG-12345"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: RFID-TAG-12345
}

// ExampleReq_emptyIdTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleReq_emptyIdTag() {
	_, err := authorize.Req(authorize.ReqInput{IdTag: ""})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: NewIdToken: IdToken: value cannot be empty
}

// ExampleReq_idTagTooLong demonstrates the error returned when
// the ID tag exceeds the maximum length of 20 characters.
func ExampleReq_idTagTooLong() {
	_, err := authorize.Req(authorize.ReqInput{
		IdTag: "RFID-ABC123456789012345",
	}) // 23 chars
	if err != nil {
		fmt.Println("idTag: Exceeds maximum length")
	}
	// Output:
	// idTag: Exceeds maximum length
}

// ExampleReq_invalidCharacters demonstrates the error returned when
// the ID tag contains non-printable ASCII characters.
func ExampleReq_invalidCharacters() {
	// Contains null byte
	_, err := authorize.Req(authorize.ReqInput{IdTag: "RFID\x00TAG"})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: CiString: value contains non-printable ASCII characters
}

// ExampleReq_shortIdTag demonstrates that short ID tags
// (within the 20 character limit) are valid.
func ExampleReq_shortIdTag() {
	req, err := authorize.Req(authorize.ReqInput{IdTag: "TAG1"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: TAG1
}
