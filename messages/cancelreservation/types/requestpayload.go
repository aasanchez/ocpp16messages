package cancelreservationtypes

import "fmt"

type RequestPayload struct {
	ReservationId string
}

func (r RequestPayload) Validate() error {
	if r.ReservationId == "" {
		return fmt.Errorf("request payload: missing required field: ReservationId")
	}

	return nil
}
