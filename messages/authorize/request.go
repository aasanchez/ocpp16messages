package authorize

import (
	"errors"
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ErrInvalidRequestIDTag is returned when the request IDTag is invalid.
var ErrInvalidRequestIDTag = errors.New("request -> invalid idTag")

// Request represents an OCPP 1.6 Authorize request message.
type Request struct {
	IDTag mat.IDToken
}

// NewRequest creates a new Authorize request from the input payload.
// Returns an error if the IDTag is invalid or cannot be converted.
func NewRequest(input mat.Request) (Request, error) {
	str, err := st.NewCiString20(input.IDTag)
	if err != nil {
		return Request{}, fmt.Errorf("request -> invalid idTag -> %w", err)
	}

	idToken, err := mat.NewIDToken(str)
	if err != nil {
		return Request{}, fmt.Errorf("request -> invalid idToken -> %w", err)
	}

	return Request{IDTag: idToken}, nil
}
