//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	gc "github.com/aasanchez/ocpp16messages/getconfiguration"
	gt "github.com/aasanchez/ocpp16messages/getconfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetConfigurationConf(f *testing.F) {
	f.Add(uint8(0), "", false, false, "", uint8(0), "")
	f.Add(uint8(2), "TestKey", true, true, "TestValue", uint8(0), "")
	f.Add(uint8(2), "", true, true, "TestValue", uint8(0), "")
	f.Add(uint8(2), "TestKey", true, true, "", uint8(0), "")
	f.Add(uint8(0), "", false, false, "", uint8(2), "UnknownKey")
	f.Add(uint8(0), "", false, false, "", uint8(2), "")

	f.Fuzz(func(
		t *testing.T,
		configMode uint8,
		key string,
		readonly bool,
		hasValue bool,
		value string,
		unknownMode uint8,
		unknownKey string,
	) {
		if len(key) > maxFuzzStringLen ||
			len(value) > maxFuzzStringLen ||
			len(unknownKey) > maxFuzzStringLen {
			t.Skip()
		}

		var configurationKey []gt.KeyValueInput
		switch configMode % 3 {
		case 0:
			configurationKey = nil
		case 1:
			configurationKey = []gt.KeyValueInput{}
		default:
			var valuePtr *string
			if hasValue {
				valuePtr = &value
			}

			configurationKey = []gt.KeyValueInput{
				{
					Key:      key,
					Readonly: readonly,
					Value:    valuePtr,
				},
			}
		}

		var unknownKeys []string
		switch unknownMode % 3 {
		case 0:
			unknownKeys = nil
		case 1:
			unknownKeys = []string{}
		default:
			unknownKeys = []string{unknownKey}
		}

		conf, err := gc.Conf(gc.ConfInput{
			ConfigurationKey: configurationKey,
			UnknownKey:       unknownKeys,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		switch configMode % 3 {
		case 0, 1:
			if conf.ConfigurationKey != nil {
				t.Fatal("ConfigurationKey != nil, want nil")
			}
		default:
			if len(conf.ConfigurationKey) == 0 {
				t.Fatal("ConfigurationKey is empty, want at least one")
			}
			if conf.ConfigurationKey[0].Key().String() != key {
				t.Fatalf(
					"ConfigurationKey[0].Key = %q, want %q",
					conf.ConfigurationKey[0].Key().String(),
					key,
				)
			}
		}

		switch unknownMode % 3 {
		case 0, 1:
			if conf.UnknownKey != nil {
				t.Fatal("UnknownKey != nil, want nil")
			}
		default:
			if len(conf.UnknownKey) == 0 {
				t.Fatal("UnknownKey is empty, want at least one")
			}
			if conf.UnknownKey[0].String() != unknownKey {
				t.Fatalf(
					"UnknownKey[0] = %q, want %q",
					conf.UnknownKey[0].String(),
					unknownKey,
				)
			}
		}
	})
}
