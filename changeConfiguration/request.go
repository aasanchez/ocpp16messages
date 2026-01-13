package changeConfiguration

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a ChangeConfiguration.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	Key   string // Required: Configuration key name (max 50 chars)
	Value string // Required: New configuration value (max 500 chars)
}

// ReqMessage represents an OCPP 1.6 ChangeConfiguration.req message.
type ReqMessage struct {
	Key   st.CiString50Type
	Value st.CiString500Type
}

// Req creates a ChangeConfiguration.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Key is empty or exceeds 50 characters
//   - Key contains non-printable ASCII characters
//   - Value exceeds 500 characters
//   - Value contains non-printable ASCII characters
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	key, err := st.NewCiString50Type(input.Key)
	if err != nil {
		errs = append(errs, fmt.Errorf("key: %w", err))
	}

	value, err := st.NewCiString500Type(input.Value)
	if err != nil {
		errs = append(errs, fmt.Errorf("value: %w", err))
	}

	if errs != nil {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		Key:   key,
		Value: value,
	}, nil
}
