package getConfiguration_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getConfiguration"
	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
)

const (
	labelConfigKeyCount  = "ConfigKeyCount:"
	labelUnknownKeyCount = "UnknownKeyCount:"
	labelKey             = "Key:"
)

// ExampleConf demonstrates creating a valid GetConfiguration.conf message
// with a single configuration key.
func ExampleConf() {
	value := "300"

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "HeartbeatInterval",
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConfigKeyCount, len(conf.ConfigurationKey))
	fmt.Println(labelKey, conf.ConfigurationKey[0].Key().Value())
	fmt.Println("Readonly:", conf.ConfigurationKey[0].Readonly())
	fmt.Println("Value:", conf.ConfigurationKey[0].Value().Value())
	// Output:
	// ConfigKeyCount: 1
	// Key: HeartbeatInterval
	// Readonly: false
	// Value: 300
}

// ExampleConf_withReadonlyKey demonstrates creating a GetConfiguration.conf
// message with a read-only configuration key.
func ExampleConf_withReadonlyKey() {
	value := "1.6"

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "SupportedFeatureProfiles",
				Readonly: true,
				Value:    &value,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelKey, conf.ConfigurationKey[0].Key().Value())
	fmt.Println("Readonly:", conf.ConfigurationKey[0].Readonly())
	// Output:
	// Key: SupportedFeatureProfiles
	// Readonly: true
}

// ExampleConf_withUnknownKeys demonstrates creating a GetConfiguration.conf
// message with unknown keys.
func ExampleConf_withUnknownKeys() {
	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{"NonExistentKey", "AnotherUnknown"},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelUnknownKeyCount, len(conf.UnknownKey))
	fmt.Println("UnknownKey[0]:", conf.UnknownKey[0].Value())
	// Output:
	// UnknownKeyCount: 2
	// UnknownKey[0]: NonExistentKey
}

// ExampleConf_withKeyNoValue demonstrates creating a GetConfiguration.conf
// message with a known key that has no value set.
func ExampleConf_withKeyNoValue() {
	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "AuthorizationKey",
				Readonly: false,
				Value:    nil,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelKey, conf.ConfigurationKey[0].Key().Value())
	fmt.Println("HasValue:", conf.ConfigurationKey[0].Value() != nil)
	// Output:
	// Key: AuthorizationKey
	// HasValue: false
}

// ExampleConf_complete demonstrates creating a complete GetConfiguration.conf
// message with both configuration keys and unknown keys.
func ExampleConf_complete() {
	value := "60"

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "HeartbeatInterval",
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: []string{"VendorSpecificKey"},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConfigKeyCount, len(conf.ConfigurationKey))
	fmt.Println(labelUnknownKeyCount, len(conf.UnknownKey))
	// Output:
	// ConfigKeyCount: 1
	// UnknownKeyCount: 1
}

// ExampleConf_empty demonstrates creating an empty GetConfiguration.conf
// message (valid when Charge Point has no configuration).
func ExampleConf_empty() {
	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConfigKeyCount, len(conf.ConfigurationKey))
	fmt.Println(labelUnknownKeyCount, len(conf.UnknownKey))
	// Output:
	// ConfigKeyCount: 0
	// UnknownKeyCount: 0
}

// ExampleConf_invalidConfigKey demonstrates the error returned when
// a configuration key is invalid.
func ExampleConf_invalidConfigKey() {
	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "",
				Readonly: false,
				Value:    nil,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// configurationKey[0]: key: value cannot be empty
}

// ExampleConf_invalidUnknownKey demonstrates the error returned when
// an unknown key is invalid.
func ExampleConf_invalidUnknownKey() {
	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{"ValidKey", ""},
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// unknownKey[1]: value cannot be empty
}
