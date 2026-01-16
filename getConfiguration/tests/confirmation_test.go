package getConfiguration_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/getConfiguration"
	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidConfigKey      = "HeartbeatInterval"
	testValidConfigValue    = "300"
	testUnknownKey          = "NonExistentKey"
	testUnknownKeyTwo       = "AnotherUnknown"
	errFieldConfigKey       = "configurationKey"
	errFieldUnknownKey      = "unknownKey"
	valueMaxLength          = 500
	valueMaxLengthPlusOne   = 501
	expectedCountZero       = 0
	expectedCountOne        = 1
	expectedCountTwo        = 2
	errConfigKeyLenFormat   = "Conf() ConfigurationKey length = %d, want %d"
	errUnknownKeyLenFormat  = "Conf() UnknownKey length = %d, want %d"
	errIndexedConfigKey     = "configurationKey[1]"
	errIndexedUnknownKey    = "unknownKey[1]"
	errInvalidCharsConfig   = "Conf() err = nil, want error for invalid chars"
	errInvalidCharsUnknown  = "Conf() err = nil, want error for invalid chars"
	errMultipleInvalidField = "Conf() err = nil, want multiple invalid fields"
)

func TestConf_Valid_Empty(t *testing.T) {
	t.Parallel()

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(conf.ConfigurationKey) != expectedCountZero {
		t.Errorf(
			errConfigKeyLenFormat,
			len(conf.ConfigurationKey),
			expectedCountZero,
		)
	}

	if len(conf.UnknownKey) != expectedCountZero {
		t.Errorf(
			errUnknownKeyLenFormat,
			len(conf.UnknownKey),
			expectedCountZero,
		)
	}
}

func TestConf_Valid_SingleConfigKey(t *testing.T) {
	t.Parallel()

	value := testValidConfigValue

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      testValidConfigKey,
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(conf.ConfigurationKey) != expectedCountOne {
		t.Errorf(
			errConfigKeyLenFormat,
			len(conf.ConfigurationKey),
			expectedCountOne,
		)

		return
	}

	if conf.ConfigurationKey[0].Key().Value() != testValidConfigKey {
		t.Errorf(
			st.ErrorMismatch,
			testValidConfigKey,
			conf.ConfigurationKey[0].Key().Value(),
		)
	}

	if conf.ConfigurationKey[0].Value().Value() != testValidConfigValue {
		t.Errorf(
			st.ErrorMismatch,
			testValidConfigValue,
			conf.ConfigurationKey[0].Value().Value(),
		)
	}
}

func TestConf_Valid_ReadonlyKey(t *testing.T) {
	t.Parallel()

	value := "Core"

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
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if !conf.ConfigurationKey[0].Readonly() {
		t.Error("Conf() ConfigurationKey[0].Readonly() = false, want true")
	}
}

func TestConf_Valid_KeyWithNoValue(t *testing.T) {
	t.Parallel()

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
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.ConfigurationKey[0].Value() != nil {
		t.Error("Conf() ConfigurationKey[0].Value() != nil, want nil")
	}
}

func TestConf_Valid_SingleUnknownKey(t *testing.T) {
	t.Parallel()

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{testUnknownKey},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(conf.UnknownKey) != expectedCountOne {
		t.Errorf(errUnknownKeyLenFormat, len(conf.UnknownKey), expectedCountOne)

		return
	}

	if conf.UnknownKey[0].Value() != testUnknownKey {
		t.Errorf(st.ErrorMismatch, testUnknownKey, conf.UnknownKey[0].Value())
	}
}

func TestConf_Valid_MultipleUnknownKeys(t *testing.T) {
	t.Parallel()

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{testUnknownKey, testUnknownKeyTwo},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(conf.UnknownKey) != expectedCountTwo {
		t.Errorf(errUnknownKeyLenFormat, len(conf.UnknownKey), expectedCountTwo)
	}
}

func TestConf_Valid_Complete(t *testing.T) {
	t.Parallel()

	value := testValidConfigValue

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      testValidConfigKey,
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: []string{testUnknownKey},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(conf.ConfigurationKey) != expectedCountOne {
		t.Errorf(
			errConfigKeyLenFormat,
			len(conf.ConfigurationKey),
			expectedCountOne,
		)
	}

	if len(conf.UnknownKey) != expectedCountOne {
		t.Errorf(errUnknownKeyLenFormat, len(conf.UnknownKey), expectedCountOne)
	}
}

func TestConf_Valid_MaxLengthValue(t *testing.T) {
	t.Parallel()

	maxValue := strings.Repeat("v", valueMaxLength)

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      testValidConfigKey,
				Readonly: false,
				Value:    &maxValue,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	configValue := conf.ConfigurationKey[0].Value().Value()
	if configValue != maxValue {
		t.Errorf(st.ErrorMismatch, maxValue, configValue)
	}
}

func TestConf_EmptyConfigKey(t *testing.T) {
	t.Parallel()

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
	if err == nil {
		t.Error("Conf() error = nil, want error for empty config key")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}

	if !strings.Contains(err.Error(), errFieldConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldConfigKey)
	}
}

func TestConf_ConfigKeyTooLong(t *testing.T) {
	t.Parallel()

	longKey := strings.Repeat("k", keyMaxLengthPlusOne)

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      longKey,
				Readonly: false,
				Value:    nil,
			},
		},
		UnknownKey: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for config key too long")
	}

	if !strings.Contains(err.Error(), errFieldConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldConfigKey)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestConf_ConfigValueTooLong(t *testing.T) {
	t.Parallel()

	longValue := strings.Repeat("v", valueMaxLengthPlusOne)

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      testValidConfigKey,
				Readonly: false,
				Value:    &longValue,
			},
		},
		UnknownKey: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for config value too long")
	}

	if !strings.Contains(err.Error(), errFieldConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldConfigKey)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestConf_EmptyUnknownKey(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{"ValidKey", ""},
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty unknown key")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}

	if !strings.Contains(err.Error(), errFieldUnknownKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldUnknownKey)
	}
}

func TestConf_UnknownKeyTooLong(t *testing.T) {
	t.Parallel()

	longKey := strings.Repeat("k", keyMaxLengthPlusOne)

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{longKey},
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for unknown key too long")
	}

	if !strings.Contains(err.Error(), errFieldUnknownKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldUnknownKey)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestConf_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "",
				Readonly: false,
				Value:    nil,
			},
		},
		UnknownKey: []string{""},
	})
	if err == nil {
		t.Error(errMultipleInvalidField)
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldConfigKey)
	}

	if !strings.Contains(errStr, errFieldUnknownKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldUnknownKey)
	}
}

func TestConf_ConfigKeyInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{
				Key:      "Invalid\x00Key",
				Readonly: false,
				Value:    nil,
			},
		},
		UnknownKey: nil,
	})
	if err == nil {
		t.Error(errInvalidCharsConfig)
	}

	if !strings.Contains(err.Error(), errFieldConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldConfigKey)
	}
}

func TestConf_UnknownKeyInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: nil,
		UnknownKey:       []string{"Invalid\x00Key"},
	})
	if err == nil {
		t.Error(errInvalidCharsUnknown)
	}

	if !strings.Contains(err.Error(), errFieldUnknownKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldUnknownKey)
	}
}

func TestConf_IndexedErrorMessages(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gt.KeyValueInput{
			{Key: testValidConfigKey, Readonly: false, Value: nil},
			{Key: "", Readonly: false, Value: nil},
		},
		UnknownKey: []string{testUnknownKey, ""},
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for indexed invalid entries")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errIndexedConfigKey) {
		t.Errorf(st.ErrorWantContains, err, errIndexedConfigKey)
	}

	if !strings.Contains(errStr, errIndexedUnknownKey) {
		t.Errorf(st.ErrorWantContains, err, errIndexedUnknownKey)
	}
}
