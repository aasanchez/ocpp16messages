package bootNotification

import (
	"errors"
	"fmt"
	"strconv"

	mbt "github.com/aasanchez/ocpp16messages/messages/bootNotification/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfInput represents the raw input data for creating a BootNotification.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	Status      string // Required: RegistrationStatus value
	CurrentTime string // Required: Central System's current time (RFC3339)
	Interval    int    // Required: Heartbeat interval in seconds
}

// ConfMessage represents an OCPP 1.6 BootNotification.conf message.
type ConfMessage struct {
	Status      mbt.RegistrationStatus
	CurrentTime st.DateTime
	Interval    st.Integer
}

// confValidation holds validated fields during Conf construction.
type confValidation struct {
	status      mbt.RegistrationStatus
	currentTime st.DateTime
	interval    st.Integer
}

// Conf creates a BootNotification.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - Status is not a valid RegistrationStatus value
//   - CurrentTime is not a valid RFC3339 date
//   - Interval is negative or exceeds uint16 max value
func Conf(input ConfInput) (ConfMessage, error) {
	validated, errs := validateConfInput(input)

	if errs != nil {
		return ConfMessage{
			Status:      "",
			CurrentTime: st.DateTime{},
			Interval:    st.Integer{},
		}, errors.Join(errs...)
	}

	return ConfMessage{
		Status:      validated.status,
		CurrentTime: validated.currentTime,
		Interval:    validated.interval,
	}, nil
}

// validateConfInput validates all fields in ConfInput and returns validated
// values along with any errors.
func validateConfInput(input ConfInput) (confValidation, []error) {
	var errs []error

	var validated confValidation

	// Validate status (required)
	validated.status, errs = validateStatus(input.Status, errs)

	// Validate currentTime (required)
	validated.currentTime, errs = validateCurrentTime(input.CurrentTime, errs)

	// Validate interval (required)
	validated.interval, errs = validateInterval(input.Interval, errs)

	return validated, errs
}

// validateStatus validates the status field and returns the RegistrationStatus.
func validateStatus(
	status string,
	errs []error,
) (mbt.RegistrationStatus, []error) {
	regStatus := mbt.RegistrationStatus(status)

	if !regStatus.IsValid() {
		return "", append(
			errs,
			fmt.Errorf("status: %w", st.ErrInvalidValue),
		)
	}

	return regStatus, errs
}

// validateCurrentTime validates the currentTime field.
func validateCurrentTime(
	timeStr string,
	errs []error,
) (st.DateTime, []error) {
	dateTime, err := st.NewDateTime(timeStr)
	if err != nil {
		return st.DateTime{}, append(
			errs,
			fmt.Errorf("currentTime: %w", err),
		)
	}

	return dateTime, errs
}

// validateInterval validates the interval field.
func validateInterval(interval int, errs []error) (st.Integer, []error) {
	intVal, err := st.NewInteger(strconv.Itoa(interval))
	if err != nil {
		return st.Integer{}, append(
			errs,
			fmt.Errorf("interval: %w", err),
		)
	}

	return intVal, errs
}
