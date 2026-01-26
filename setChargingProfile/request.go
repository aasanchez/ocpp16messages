package setChargingProfile

import (
	"errors"
	"fmt"

	spt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a SetChargingProfile.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The connector to which the charging profile applies.
	// connectorId 0 is associated with the entire Charge Point.
	ConnectorId int
	// Required: The charging profile to be set at the Charge Point.
	CsChargingProfiles spt.ChargingProfileInput
}

// ReqMessage represents an OCPP 1.6 SetChargingProfile.req message.
type ReqMessage struct {
	ConnectorId        types.Integer
	CsChargingProfiles spt.ChargingProfile
}

// Req creates a SetChargingProfile.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - CsChargingProfiles is invalid
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	connectorId, err := types.NewIntegerFromInt(input.ConnectorId)
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	csChargingProfiles, err := spt.NewChargingProfile(input.CsChargingProfiles)
	if err != nil {
		errs = append(errs, fmt.Errorf("csChargingProfiles: %w", err))
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId:        connectorId,
		CsChargingProfiles: csChargingProfiles,
	}, nil
}
