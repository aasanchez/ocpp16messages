package messages

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// AuthorizeRequestMessage represents the OCPP 1.6J Authorize.req message.
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
type AuthorizeRequestMessage struct {
	// IdTag is the authorization identifier being validated by the Central System.
	//
	// It must comply with the constraints of the OCPP 1.6J `idTag` field,
	// defined as a CiString[20]: max 20 characters, printable ASCII only.
	IdTag types.IdTagType
}

// AuthorizeRequest constructs a new AuthorizeRequestMessage with a validated IdTag.
//
// Returns an error if the provided string is invalid (e.g. too long, empty, or
// contains non-ASCII characters). This helper ensures that only valid messages
// are created prior to serialization or transmission.
//
// Example usage:
//
//	req, err := messages.AuthorizeRequest("ABC123")
//	if err != nil {
//	    log.Fatalf("invalid request: %v", err)
//	}
func AuthorizeRequest(idTag string) (AuthorizeRequestMessage, error) {
	tag, err := types.IdTag(idTag)
	if err != nil {
		return AuthorizeRequestMessage{}, fmt.Errorf("failed to create AuthorizeRequestMessage: %w", err)
	}

	return AuthorizeRequestMessage{IdTag: tag}, nil
}

// String returns a human-readable representation of the AuthorizeRequestMessage.
//
// Useful for logging, debugging, and developer tools.
func (r AuthorizeRequestMessage) String() string {
	return fmt.Sprintf("{idTag=%s}", r.IdTag.String())
}

// Validate performs a revalidation of the AuthorizeRequestMessage fields.
//
// This is typically used after deserialization or transformation to ensure the
// message still complies with OCPP constraints before processing.
func (r AuthorizeRequestMessage) Validate() error {
	if err := r.IdTag.Validate(); err != nil {
		return fmt.Errorf("AuthorizeRequestMessage validation failed: %w", err)
	}

	return nil
}
