package getConfiguration

import (
	"errors"
	"fmt"

	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a GetConfiguration.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Optional: List of known configuration keys with their values
	ConfigurationKey []gt.KeyValueInput
	// Optional: List of requested keys that are unknown to the Charge Point
	UnknownKey []string
}

// ConfMessage represents an OCPP 1.6 GetConfiguration.conf message.
type ConfMessage struct {
	ConfigurationKey []gt.KeyValue
	UnknownKey       []st.CiString50Type
}

// Conf creates a GetConfiguration.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Any ConfigurationKey entry has invalid key or value
//   - Any UnknownKey is empty, exceeds 50 characters, or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	configKeys, configErrs := confValidateConfigurationKeys(
		input.ConfigurationKey,
	)

	unknownKeys, unknownErrs := confValidateUnknownKeys(input.UnknownKey)

	errs := append(configErrs, unknownErrs...) //nolint:gocritic // intentional

	if len(errs) > errCountZero {
		return ConfMessage{}, errors.Join(errs...)
	}

	return ConfMessage{
		ConfigurationKey: configKeys,
		UnknownKey:       unknownKeys,
	}, nil
}

// confValidateConfigurationKeys validates the optional configuration keys list.
func confValidateConfigurationKeys(
	keys []gt.KeyValueInput,
) ([]gt.KeyValue, []error) {
	if len(keys) == errCountZero {
		return nil, nil
	}

	var errs []error

	var validKeys []gt.KeyValue

	for i, keyInput := range keys {
		keyValue, err := gt.NewKeyValue(keyInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("configurationKey[%d]: %w", i, err))
		} else {
			validKeys = append(validKeys, keyValue)
		}
	}

	return validKeys, errs
}

// confValidateUnknownKeys validates the optional unknown keys list.
func confValidateUnknownKeys(keys []string) ([]st.CiString50Type, []error) {
	if len(keys) == errCountZero {
		return nil, nil
	}

	var errs []error

	var validKeys []st.CiString50Type

	for i, keyStr := range keys {
		key, err := st.NewCiString50Type(keyStr)
		if err != nil {
			errs = append(errs, fmt.Errorf("unknownKey[%d]: %w", i, err))
		} else {
			validKeys = append(validKeys, key)
		}
	}

	return validKeys, errs
}
