package authorizetypes

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

var errIdTagInfoPayload = errors.New("invalid IdTagInfoPayload")

type IdTagInfoPayload struct {
	ExpiryDate  string
	ParentIdTag string
	Status      string
}

func (input IdTagInfoPayload) Validate() error {
	if input.Status == "" {
		return fmt.Errorf(st.ErrorWrapper, errIdTagInfoPayload, input)
	}

	return nil
}
