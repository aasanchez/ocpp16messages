package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating an Authorize.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	IdTag string
}

// ReqMessage represents an OCPP 1.6 Authorize.req message.
type ReqMessage struct {
	IdTag mat.IdToken
}

// Req creates an Authorize.req message from the given input.
// It validates all fields automatically and returns an error if:
//   - IdTag is empty
//   - IdTag exceeds 20 characters
//   - IdTag contains non-printable ASCII characters
func Req(input ReqInput) (ReqMessage, error) {
	str, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		return ReqMessage{}, fmt.Errorf("idTag: %w", err)
	}

	idToken := mat.NewIdToken(str)

	return ReqMessage{IdTag: idToken}, nil
}
