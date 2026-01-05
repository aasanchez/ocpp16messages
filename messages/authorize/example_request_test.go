package authorize_test

//revive:disable line-length-limit
import (
	"fmt"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

// ExampleNewRequest demonstrates creating a valid Authorize request
// with a properly formatted ID tag.
func ExampleNewRequest() {
	payload := mat.Request{
		IdTag: "RFID-TAG-12345",
	}

	req, err := ma.NewRequest(payload)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: RFID-TAG-12345
}

// ExampleNewRequest_emptyIdTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleNewRequest_emptyIdTag() {
	payload := mat.Request{
		IdTag: "",
	}

	_, err := ma.NewRequest(payload)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output:
	// Error: idToken: NewIdToken: IdToken: value cannot be empty
}

// ExampleNewRequest_idTagTooLong demonstrates the error returned when
// the ID tag exceeds the maximum length of 20 characters.
func ExampleNewRequest_idTagTooLong() {
	payload := mat.Request{
		IdTag: "RFID-ABC123456789012345", // 25 characters
	}

	_, err := ma.NewRequest(payload)
	if err != nil {
		fmt.Println("Error: ID tag exceeds maximum length")
	}
	// Output:
	// Error: ID tag exceeds maximum length
}

// ExampleNewRequest_invalidCharacters demonstrates the error returned when
// the ID tag contains non-printable ASCII characters.
func ExampleNewRequest_invalidCharacters() {
	payload := mat.Request{
		IdTag: "RFID\x00TAG", // Contains null byte
	}

	_, err := ma.NewRequest(payload)
	if err != nil {
		fmt.Println("Error: ID tag contains invalid characters")
	}
	// Output:
	// Error: ID tag contains invalid characters
}

// ExampleNewRequest_shortIdTag demonstrates that short ID tags
// (within the 20 character limit) are valid.
func ExampleNewRequest_shortIdTag() {
	payload := mat.Request{
		IdTag: "TAG1",
	}

	req, err := ma.NewRequest(payload)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("IdTag:", req.IdTag.String())
	// Output:
	// IdTag: TAG1
}
