package authorize

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// RequestMessage represents the OCPP 1.6J Authorize.req message.
//
// This message is sent from a Charge Point to the Central System (CSMS) to
// request authorization for an idTag prior to starting a charging session.
//
// According to the OCPP 1.6J specification, the `Authorize.req` message consists of:
//   - idTag: A case-insensitive identifier (RFID token, app identifier, etc.)
//     used to authenticate a user.
//
// Specification Reference:
// - OCPP 1.6J, Section 5.2: Authorize.req
type RequestMessage struct {
	// IdTag is the authorization identifier being validated by the Central System.
	//
	// It must comply with the constraints of the OCPP 1.6J `idTag` field,
	// defined as a CiString[20]: max 20 characters, printable ASCII only.
	IdTag types.IdTokenType
}

// Request constructs a new RequestMessage with a validated IdTag.
//
// Returns an error if the provided string is invalid (e.g. too long, empty, or
// contains non-ASCII characters). This helper ensures that only valid messages
// are created prior to serialization or transmission.
//
// Example usage:
//
//	req, err := messages.Request("ABC123")
//	if err != nil {
//	    log.Fatalf("invalid request: %v", err)
//	}
func Request(idTag string) (RequestMessage, error) {
	tok, err := types.IdToken(idTag)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("failed to create Authorize.req message, problem with idTag: %s: %w", idTag, err)
	}

	return RequestMessage{IdTag: tok}, nil
}

// String returns a human-readable representation of the RequestMessage.
//
// Useful for logging, debugging, and developer tools.
func (r RequestMessage) String() string {
	return fmt.Sprintf("{idTag=%s}", r.IdTag.String())
}

// Validate performs a revalidation of the RequestMessage fields.
//
// This is typically used after deserialization or transformation to ensure the
// message still complies with OCPP constraints before processing.
func (r RequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("RequestMessage validation failed: %w", err)
	}

	return nil
}
