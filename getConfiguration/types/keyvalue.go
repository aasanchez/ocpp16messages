package types

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// KeyValueInput represents the raw input data for creating a KeyValue.
// The constructor NewKeyValue validates all fields automatically.
type KeyValueInput struct {
	// Required: Configuration key name (max 50 chars)
	Key string
	// Required: Whether the configuration value is read-only
	Readonly bool
	// Optional: Configuration value (max 500 chars), nil if key is known but
	// not set
	Value *string
}

// KeyValue represents a configuration key-value pair as defined in OCPP 1.6.
// It is used in GetConfiguration.conf to return configuration settings.
type KeyValue struct {
	key      st.CiString50Type
	readonly bool
	value    *st.CiString500Type
}

// NewKeyValue creates a new KeyValue from input.
// Returns an error if:
//   - Key is empty or exceeds 50 characters or contains invalid chars
//   - Value (if provided) exceeds 500 characters or contains invalid chars
func NewKeyValue(input KeyValueInput) (KeyValue, error) {
	var errs []error

	key, err := st.NewCiString50Type(input.Key)
	if err != nil {
		errs = append(errs, fmt.Errorf("key: %w", err))
	}

	var value *st.CiString500Type

	if input.Value != nil {
		val, err := st.NewCiString500Type(*input.Value)
		if err != nil {
			errs = append(errs, fmt.Errorf("value: %w", err))
		} else {
			value = &val
		}
	}

	if len(errs) > errCountZero {
		return KeyValue{}, errors.Join(errs...)
	}

	return KeyValue{
		key:      key,
		readonly: input.Readonly,
		value:    value,
	}, nil
}

// Key returns the configuration key name.
func (k KeyValue) Key() st.CiString50Type {
	return k.key
}

// Readonly returns whether the configuration value is read-only.
func (k KeyValue) Readonly() bool {
	return k.readonly
}

// Value returns the configuration value, or nil if not set.
func (k KeyValue) Value() *st.CiString500Type {
	return k.value
}
