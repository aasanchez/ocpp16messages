package getCompositeSchedule

import (
	"errors"
	"fmt"
	"strconv"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a
// GetCompositeSchedule.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: Status of the request (Accepted or Rejected)
	Status string
	// Optional: The connector ID for which the schedule is reported
	ConnectorId *int
	// Optional: Time at which the schedule starts (RFC3339 format)
	ScheduleStart *string
	// Optional: The composite charging schedule
	ChargingSchedule *gt.ChargingScheduleInput
}

// ConfMessage represents an OCPP 1.6 GetCompositeSchedule.conf message.
type ConfMessage struct {
	Status           gt.GetCompositeScheduleStatus
	ConnectorId      *st.Integer
	ScheduleStart    *st.DateTime
	ChargingSchedule *gt.ChargingSchedule
}

// Conf creates a GetCompositeSchedule.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Status is not a valid GetCompositeScheduleStatus value
//   - ConnectorId (if provided) is negative or exceeds uint16 max (65535)
//   - ScheduleStart (if provided) is not a valid RFC3339 timestamp
//   - ChargingSchedule (if provided) is invalid
func Conf(input ConfInput) (ConfMessage, error) {
	var errs []error

	status, err := confValidateStatus(input.Status)
	if err != nil {
		errs = append(errs, err)
	}

	connectorId, err := confValidateConnectorId(input.ConnectorId)
	if err != nil {
		errs = append(errs, err)
	}

	scheduleStart, err := confValidateScheduleStart(input.ScheduleStart)
	if err != nil {
		errs = append(errs, err)
	}

	chargingSchedule, err := confValidateChargingSchedule(
		input.ChargingSchedule,
	)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > errCountZero {
		return ConfMessage{}, errors.Join(errs...)
	}

	return ConfMessage{
		Status:           status,
		ConnectorId:      connectorId,
		ScheduleStart:    scheduleStart,
		ChargingSchedule: chargingSchedule,
	}, nil
}

// confValidateStatus validates the status field.
func confValidateStatus(
	statusStr string,
) (gt.GetCompositeScheduleStatus, error) {
	status := gt.GetCompositeScheduleStatus(statusStr)
	if !status.IsValid() {
		return "", fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return status, nil
}

// confValidateConnectorId validates the optional connector ID field.
func confValidateConnectorId(connectorId *int) (*st.Integer, error) {
	if connectorId == nil {
		return nil, nil
	}

	cid, err := st.NewInteger(strconv.Itoa(*connectorId))
	if err != nil {
		return nil, fmt.Errorf("connectorId: %w", err)
	}

	return &cid, nil
}

// confValidateScheduleStart validates the optional schedule start field.
func confValidateScheduleStart(scheduleStart *string) (*st.DateTime, error) {
	if scheduleStart == nil {
		return nil, nil
	}

	ss, err := st.NewDateTime(*scheduleStart)
	if err != nil {
		return nil, fmt.Errorf("scheduleStart: %w", err)
	}

	return &ss, nil
}

// confValidateChargingSchedule validates the optional charging schedule field.
func confValidateChargingSchedule(
	schedule *gt.ChargingScheduleInput,
) (*gt.ChargingSchedule, error) {
	if schedule == nil {
		return nil, nil
	}

	cs, err := gt.NewChargingSchedule(*schedule)
	if err != nil {
		return nil, fmt.Errorf("chargingSchedule: %w", err)
	}

	return &cs, nil
}
