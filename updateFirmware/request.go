package updateFirmware

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating an UpdateFirmware.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: URI where the Charge Point can retrieve the firmware.
	Location string
	// Required: Date and time after which the Charge Point is allowed to
	// retrieve the firmware (RFC3339 format).
	RetrieveDate string
	// Optional: Number of retries for downloading the firmware.
	Retries *int
	// Optional: Interval (in seconds) between retry attempts.
	RetryInterval *int
}

// ReqMessage represents an OCPP 1.6 UpdateFirmware.req message.
type ReqMessage struct {
	Location      st.CiString255Type
	RetrieveDate  st.DateTime
	Retries       *st.Integer
	RetryInterval *st.Integer
}

// Req creates an UpdateFirmware.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Location is empty or exceeds 255 characters
//   - RetrieveDate is not a valid RFC3339 timestamp
//   - Retries (if provided) is negative or exceeds uint16 max value (65535)
//   - RetryInterval (if provided) is negative or exceeds uint16 max (65535)
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	location, err := st.NewCiString255Type(input.Location)
	if err != nil {
		errs = append(errs, fmt.Errorf("location: %w", err))
	}

	retrieveDate, err := st.NewDateTime(input.RetrieveDate)
	if err != nil {
		errs = append(errs, fmt.Errorf("retrieveDate: %w", err))
	}

	retries, err := reqValidateRetries(input.Retries)
	if err != nil {
		errs = append(errs, err)
	}

	retryInterval, err := reqValidateRetryInterval(input.RetryInterval)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		Location:      location,
		RetrieveDate:  retrieveDate,
		Retries:       retries,
		RetryInterval: retryInterval,
	}, nil
}

// reqValidateRetries validates the optional retries field.
func reqValidateRetries(retries *int) (*st.Integer, error) {
	if retries == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	r, err := st.NewIntegerFromInt(*retries)
	if err != nil {
		return nil, fmt.Errorf("retries: %w", err)
	}

	return &r, nil
}

// reqValidateRetryInterval validates the optional retry interval field.
func reqValidateRetryInterval(retryInterval *int) (*st.Integer, error) {
	if retryInterval == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	ri, err := st.NewIntegerFromInt(*retryInterval)
	if err != nil {
		return nil, fmt.Errorf("retryInterval: %w", err)
	}

	return &ri, nil
}
