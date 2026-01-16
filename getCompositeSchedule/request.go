package getCompositeSchedule

import (
	"errors"
	"fmt"
	"strconv"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a
// GetCompositeSchedule.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The ID of the connector for which the schedule is requested.
	// Use 0 to request the schedule for the entire Charge Point.
	ConnectorId int
	// Required: Duration in seconds of the requested schedule.
	Duration int
	// Optional: Preferred unit of measure for the charging rate (W or A).
	ChargingRateUnit *string
}

// ReqMessage represents an OCPP 1.6 GetCompositeSchedule.req message.
type ReqMessage struct {
	ConnectorId      st.Integer
	Duration         st.Integer
	ChargingRateUnit *gt.ChargingRateUnit
}

// Req creates a GetCompositeSchedule.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - Duration is negative or exceeds uint16 max value (65535)
//   - ChargingRateUnit (if provided) is not a valid value ("W" or "A")
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	connectorId, err := st.NewInteger(strconv.Itoa(input.ConnectorId))
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	duration, err := st.NewInteger(strconv.Itoa(input.Duration))
	if err != nil {
		errs = append(errs, fmt.Errorf("duration: %w", err))
	}

	var chargingRateUnit *gt.ChargingRateUnit

	if input.ChargingRateUnit != nil {
		unit := gt.ChargingRateUnit(*input.ChargingRateUnit)
		if !unit.IsValid() {
			errs = append(
				errs,
				fmt.Errorf("chargingRateUnit: %w", st.ErrInvalidValue),
			)
		} else {
			chargingRateUnit = &unit
		}
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId:      connectorId,
		Duration:         duration,
		ChargingRateUnit: chargingRateUnit,
	}, nil
}
