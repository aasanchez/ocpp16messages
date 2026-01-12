package authorize_test

//revive:disable:line-length-limit

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
)

// ExampleReq demonstrates creating a valid Authorize.req message
// with a properly formatted ID tag using the Input struct.
func ExampleReq() {
	req, err := authorize.Req(authorize.Input{IdTag: "RFID-TAG-12345"})
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
	_, err := authorize.Req(authorize.Input{IdTag: ""})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: NewIdToken: IdToken: value cannot be empty
}

// ExampleReq_idTagTooLong demonstrates the error returned when
// the ID tag exceeds the maximum length of 20 characters.
func ExampleReq_idTagTooLong() {
	_, err := authorize.Req(authorize.Input{IdTag: "RFID-ABC123456789012345"}) // 23 chars
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: CiString Error on Construct (len=23, max=20): exceeds maximum length
}

// ExampleReq_invalidCharacters demonstrates the error returned when
// the ID tag contains non-printable ASCII characters.
func ExampleReq_invalidCharacters() {
	_, err := authorize.Req(authorize.Input{IdTag: "RFID\x00TAG"}) // Contains null byte
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: CiString: value contains non-printable ASCII characters
}

// ExampleReq_shortIdTag demonstrates that short ID tags
// (within the 20 character limit) are valid.
func ExampleReq_shortIdTag() {
	req, err := authorize.Req(authorize.Input{IdTag: "TAG1"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: TAG1
}
