package getconfiguration

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a GetConfiguration.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Optional: List of configuration keys to retrieve (max 50 chars each).
	// If nil or empty, all configuration settings will be returned.
	Key []string
}

// ReqMessage represents an OCPP 1.6 GetConfiguration.req message.
type ReqMessage struct {
	Key []types.CiString50Type
}

// Req creates a GetConfiguration.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Any key in the list is empty, exceeds 50 characters, or contains invalid
//     chars
func Req(input ReqInput) (ReqMessage, error) {
	if len(input.Key) == errCountZero {
		return ReqMessage{Key: nil}, nil
	}

	var errs []error

	var keys []types.CiString50Type

	for i, keyStr := range input.Key {
		key, err := types.NewCiString50Type(keyStr)
		if err != nil {
			errs = append(errs, fmt.Errorf("key[%d]: %w", i, err))
		} else {
			keys = append(keys, key)
		}
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{Key: keys}, nil
}
