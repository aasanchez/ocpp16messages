package authorize_test

//revive:disable:line-length-limit

import (
	"fmt"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
)

// ExampleNewRequest demonstrates creating a valid Authorize request
// with a properly formatted ID tag using the Input struct.
func ExampleNewRequest() {
	req, err := ma.NewRequest(ma.Input{IdTag: "RFID-TAG-12345"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: RFID-TAG-12345
}

// ExampleNewRequest_emptyIdTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleNewRequest_emptyIdTag() {
	_, err := ma.NewRequest(ma.Input{IdTag: ""})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: NewIdToken: IdToken: value cannot be empty
}

// ExampleNewRequest_idTagTooLong demonstrates the error returned when
// the ID tag exceeds the maximum length of 20 characters.
func ExampleNewRequest_idTagTooLong() {
	_, err := ma.NewRequest(ma.Input{IdTag: "RFID-ABC123456789012345"}) // 23 chars
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: CiString Error on Construct (len=23, max=20): exceeds maximum length
}

// ExampleNewRequest_invalidCharacters demonstrates the error returned when
// the ID tag contains non-printable ASCII characters.
func ExampleNewRequest_invalidCharacters() {
	_, err := ma.NewRequest(ma.Input{IdTag: "RFID\x00TAG"}) // Contains null byte
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: CiString: value contains non-printable ASCII characters
}

// ExampleNewRequest_shortIdTag demonstrates that short ID tags
// (within the 20 character limit) are valid.
func ExampleNewRequest_shortIdTag() {
	req, err := ma.NewRequest(ma.Input{IdTag: "TAG1"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: TAG1
}
