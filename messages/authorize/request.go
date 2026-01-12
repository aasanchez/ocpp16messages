package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// Input represents the raw input data for creating an Authorize.req message.
// The constructor Req validates all fields automatically.
type Input struct {
	IdTag string
}

// Message represents an OCPP 1.6 Authorize.req message.
type Message struct {
	IdTag mat.IdToken
}

// Req creates an Authorize.req message from the given input.
// It validates all fields automatically and returns an error if:
//   - IdTag is empty
//   - IdTag exceeds 20 characters
//   - IdTag contains non-printable ASCII characters
func Req(input Input) (Message, error) {
	str, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		return Message{}, fmt.Errorf("idTag: %w", err)
	}

	idToken, err := mat.NewIdToken(str)
	if err != nil {
		return Message{}, fmt.Errorf("idTag: %w", err)
	}

	return Message{IdTag: idToken}, nil
}
