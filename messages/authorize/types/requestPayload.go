package authorizetypes

import (
	"errors"
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var (
	ErrInvalidIdTag = errors.New("invalid idTag")
)

type RequestPayload struct {
	IdTag string
}

func (r RequestPayload) Validate() error {
	if r.IdTag == "" {
		return fmt.Errorf("request payload: %w", ErrInvalidIdTag)
	}

	_, err := sharedtypes.CiString20(r.IdTag)
	if err != nil {
		return fmt.Errorf("request payload: %w", err)
	}

	return nil
}

func (r RequestPayload) Value() string {
	return r.IdTag
}
