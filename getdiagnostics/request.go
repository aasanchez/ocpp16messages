package getdiagnostics

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a GetDiagnostics.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: URI where the Charge Point SHALL upload the diagnostic data.
	Location string
	// Optional: Number of retries for uploading the diagnostics file.
	Retries *int
	// Optional: Interval (in seconds) between retry attempts.
	RetryInterval *int
	// Optional: Start time of the requested diagnostic report (RFC3339 format).
	StartTime *string
	// Optional: End time of the requested diagnostic report (RFC3339 format).
	StopTime *string
}

// ReqMessage represents an OCPP 1.6 GetDiagnostics.req message.
type ReqMessage struct {
	Location      types.CiString255Type
	Retries       *types.Integer
	RetryInterval *types.Integer
	StartTime     *types.DateTime
	StopTime      *types.DateTime
}

// Req creates a GetDiagnostics.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Location is empty or exceeds 255 characters
//   - Retries (if provided) is negative or exceeds uint16 max value (65535)
//   - RetryInterval (if provided) is negative or exceeds uint16 max (65535)
//   - StartTime (if provided) is not a valid RFC3339 timestamp
//   - StopTime (if provided) is not a valid RFC3339 timestamp
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	location, err := types.NewCiString255Type(input.Location)
	if err != nil {
		errs = append(errs, fmt.Errorf("location: %w", err))
	}

	retries, err := reqValidateRetries(input.Retries)
	if err != nil {
		errs = append(errs, err)
	}

	retryInterval, err := reqValidateRetryInterval(input.RetryInterval)
	if err != nil {
		errs = append(errs, err)
	}

	startTime, err := reqValidateStartTime(input.StartTime)
	if err != nil {
		errs = append(errs, err)
	}

	stopTime, err := reqValidateStopTime(input.StopTime)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		Location:      location,
		Retries:       retries,
		RetryInterval: retryInterval,
		StartTime:     startTime,
		StopTime:      stopTime,
	}, nil
}

// reqValidateRetries validates the optional retries field.
func reqValidateRetries(retries *int) (*types.Integer, error) {
	if retries == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	r, err := types.NewInteger(*retries)
	if err != nil {
		return nil, fmt.Errorf("retries: %w", err)
	}

	return &r, nil
}

// reqValidateRetryInterval validates the optional retry interval field.
func reqValidateRetryInterval(retryInterval *int) (*types.Integer, error) {
	if retryInterval == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	ri, err := types.NewInteger(*retryInterval)
	if err != nil {
		return nil, fmt.Errorf("retryInterval: %w", err)
	}

	return &ri, nil
}

// reqValidateStartTime validates the optional start time field.
func reqValidateStartTime(startTime *string) (*types.DateTime, error) {
	if startTime == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	parsedStartTime, err := types.NewDateTime(*startTime)
	if err != nil {
		return nil, fmt.Errorf("startTime: %w", err)
	}

	return &parsedStartTime, nil
}

// reqValidateStopTime validates the optional stop time field.
func reqValidateStopTime(stopTime *string) (*types.DateTime, error) {
	if stopTime == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	parsedStopTime, err := types.NewDateTime(*stopTime)
	if err != nil {
		return nil, fmt.Errorf("stopTime: %w", err)
	}

	return &parsedStopTime, nil
}
