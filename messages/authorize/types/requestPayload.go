package authorizetypes

import (
	"errors"
	"fmt"
)

type RequestPayload struct {
	IdTag string
}

func (r RequestPayload) Validate() error {
	if r.IdTag == "" {
		return fmt.Errorf("%w", errors.New("missing required field: idTag"))
	}
	return nil
}

func (r RequestPayload) Value() string {
	return r.IdTag
}
