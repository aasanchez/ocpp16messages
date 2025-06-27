package bootnotificationtypes

import (
	"errors"
	"fmt"
)

type ConfirmationPayload struct {
	CurrentTime string
	Interval    string
	Status      string
}

var (
	ErrMissingCurrentTime = errors.New("missing required field: CurrentTime")
	ErrMissingInterval    = errors.New("missing required field: Interval")
	ErrMissingStatus      = errors.New("missing required field: Status")
)

const errFmtConfirmationFieldWrapped = "confirmation payload: %w"

func (c ConfirmationPayload) Validate() error {
	if c.CurrentTime == "" {
		return fmt.Errorf(errFmtConfirmationFieldWrapped, ErrMissingCurrentTime)
	}

	if c.Interval == "" {
		return fmt.Errorf(errFmtConfirmationFieldWrapped, ErrMissingInterval)
	}

	if c.Status == "" {
		return fmt.Errorf(errFmtConfirmationFieldWrapped, ErrMissingStatus)
	}

	return nil
}
