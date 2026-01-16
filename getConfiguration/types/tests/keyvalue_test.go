package types_test

import (
	"errors"
	"strings"
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidKey          = "HeartbeatInterval"
	testValidValue        = "300"
	errFieldKey           = "key"
	errFieldValue         = "value"
	errExceedsMaxLength   = "exceeds maximum length"
	keyMaxLength          = 50
	keyMaxLengthPlusOne   = 51
	valueMaxLength        = 500
	valueMaxLengthPlusOne = 501
	errInvalidCharsKey    = "NewKeyValue() err = nil, want error invalid key"
	errInvalidCharsValue  = "NewKeyValue() err = nil, want error invalid value"
	errMultipleInvalid    = "NewKeyValue() err = nil, want multiple errors"
)

func TestNewKeyValue_Valid_AllFields(t *testing.T) {
	t.Parallel()

	value := testValidValue

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      testValidKey,
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if keyValue.Key().Value() != testValidKey {
		t.Errorf(st.ErrorMismatch, testValidKey, keyValue.Key().Value())
	}

	if keyValue.Readonly() {
		t.Error("NewKeyValue() Readonly() = true, want false")
	}

	if keyValue.Value() == nil {
		t.Error("NewKeyValue() Value() = nil, want non-nil")

		return
	}

	if keyValue.Value().Value() != testValidValue {
		t.Errorf(st.ErrorMismatch, testValidValue, keyValue.Value().Value())
	}
}

func TestNewKeyValue_Valid_ReadonlyTrue(t *testing.T) {
	t.Parallel()

	value := "Core"

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "SupportedFeatureProfiles",
		Readonly: true,
		Value:    &value,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if !keyValue.Readonly() {
		t.Error("NewKeyValue() Readonly() = false, want true")
	}
}

func TestNewKeyValue_Valid_NoValue(t *testing.T) {
	t.Parallel()

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "AuthorizationKey",
		Readonly: false,
		Value:    nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if keyValue.Value() != nil {
		t.Error("NewKeyValue() Value() != nil, want nil")
	}
}

func TestNewKeyValue_Valid_MaxLengthKey(t *testing.T) {
	t.Parallel()

	maxKey := strings.Repeat("k", keyMaxLength)

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      maxKey,
		Readonly: false,
		Value:    nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if keyValue.Key().Value() != maxKey {
		t.Errorf(st.ErrorMismatch, maxKey, keyValue.Key().Value())
	}
}

func TestNewKeyValue_Valid_MaxLengthValue(t *testing.T) {
	t.Parallel()

	maxValue := strings.Repeat("v", valueMaxLength)

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      testValidKey,
		Readonly: false,
		Value:    &maxValue,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if keyValue.Value().Value() != maxValue {
		t.Errorf(st.ErrorMismatch, maxValue, keyValue.Value().Value())
	}
}

func TestNewKeyValue_EmptyKey(t *testing.T) {
	t.Parallel()

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "",
		Readonly: false,
		Value:    nil,
	})
	if err == nil {
		t.Error("NewKeyValue() error = nil, want error for empty key")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}

	if !strings.Contains(err.Error(), errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}
}

func TestNewKeyValue_KeyTooLong(t *testing.T) {
	t.Parallel()

	longKey := strings.Repeat("k", keyMaxLengthPlusOne)

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      longKey,
		Readonly: false,
		Value:    nil,
	})
	if err == nil {
		t.Error("NewKeyValue() error = nil, want error for key too long")
	}

	if !strings.Contains(err.Error(), errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestNewKeyValue_KeyInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "Invalid\x00Key",
		Readonly: false,
		Value:    nil,
	})
	if err == nil {
		t.Error(errInvalidCharsKey)
	}

	if !strings.Contains(err.Error(), errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}
}

func TestNewKeyValue_EmptyValue(t *testing.T) {
	t.Parallel()

	emptyValue := ""

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      testValidKey,
		Readonly: false,
		Value:    &emptyValue,
	})
	if err == nil {
		t.Error("NewKeyValue() error = nil, want error for empty value")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}

	if !strings.Contains(err.Error(), errFieldValue) {
		t.Errorf(st.ErrorWantContains, err, errFieldValue)
	}
}

func TestNewKeyValue_ValueTooLong(t *testing.T) {
	t.Parallel()

	longValue := strings.Repeat("v", valueMaxLengthPlusOne)

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      testValidKey,
		Readonly: false,
		Value:    &longValue,
	})
	if err == nil {
		t.Error("NewKeyValue() error = nil, want error for value too long")
	}

	if !strings.Contains(err.Error(), errFieldValue) {
		t.Errorf(st.ErrorWantContains, err, errFieldValue)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestNewKeyValue_ValueInvalidChars(t *testing.T) {
	t.Parallel()

	invalidValue := "Invalid\x00Value"

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      testValidKey,
		Readonly: false,
		Value:    &invalidValue,
	})
	if err == nil {
		t.Error(errInvalidCharsValue)
	}

	if !strings.Contains(err.Error(), errFieldValue) {
		t.Errorf(st.ErrorWantContains, err, errFieldValue)
	}
}

func TestNewKeyValue_MultipleErrors(t *testing.T) {
	t.Parallel()

	emptyValue := ""

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "",
		Readonly: false,
		Value:    &emptyValue,
	})
	if err == nil {
		t.Error(errMultipleInvalid)
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}

	if !strings.Contains(errStr, errFieldValue) {
		t.Errorf(st.ErrorWantContains, err, errFieldValue)
	}
}
