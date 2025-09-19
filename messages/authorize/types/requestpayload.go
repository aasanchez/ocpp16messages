package authorizetypes

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

var errMissingIdTag = errors.New("error: Missing IdTag")

type RequestPayload struct {
	idTag string
}

func (input RequestPayload) Validate() error {
	if input.idTag == "" {
		return fmt.Errorf(st.ErrorWrapper, errMissingIdTag, input)
	}

	return nil
}

func (input RequestPayload) IdTag() string {
	return input.idTag
}
