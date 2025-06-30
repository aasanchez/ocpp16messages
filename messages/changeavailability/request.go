package changeavailability

import (
	"fmt"
	"strconv"

	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	ConnectorId uint32
	Type        cat.AvailabilityTypeType
}

func Request(input cat.RequestPayload) (RequestMessage, error) {
	if err := input.Validate(); err != nil {
		return RequestMessage{}, fmt.Errorf("changeavailability.Request: invalid payload: %w", err)
	}

	connIdTemp, err := strconv.ParseUint(input.ConnectorId, 10, 32)
	if err != nil {
		return RequestMessage{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to parse ConnectorId", err)
	}

	availabilityType, err := cat.AvailabilityType(input.Type)
	if err != nil {
		return RequestMessage{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to parse Type", err)
	}

	return RequestMessage{
		ConnectorId: uint32(connIdTemp),
		Type:        availabilityType,
	}, nil
}
