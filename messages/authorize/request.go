package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// Request represents an OCPP 1.6 Authorize request message.
type Request struct {
	IdTag mat.IdToken
}

// NewRequest creates a new Authorize request from the input payload.
// Returns an error if the IdTag is invalid or cannot be converted.
func NewRequest(input mat.Request) (Request, error) {
	str, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		return Request{}, fmt.Errorf("idTag: %w", err)
	}

	idToken, err := mat.NewIdToken(str)
	if err != nil {
		return Request{}, fmt.Errorf("idToken: %w", err)
	}

	return Request{IdTag: idToken}, nil
}
