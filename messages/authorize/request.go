package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// Input represents the raw input data for creating an Authorize request.
// The constructor NewRequest validates all fields automatically.
type Input struct {
	IdTag string
}

// Request represents an OCPP 1.6 Authorize request message.
type Request struct {
	IdTag mat.IdToken
}

// NewRequest creates a new Authorize request from the given input.
// It validates all fields automatically and returns an error if:
//   - IdTag is empty
//   - IdTag exceeds 20 characters
//   - IdTag contains non-printable ASCII characters
func NewRequest(input Input) (Request, error) {
	str, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		return Request{}, fmt.Errorf("idTag: %w", err)
	}

	idToken, err := mat.NewIdToken(str)
	if err != nil {
		return Request{}, fmt.Errorf("idTag: %w", err)
	}

	return Request{IdTag: idToken}, nil
}
