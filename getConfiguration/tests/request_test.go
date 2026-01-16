package getConfiguration_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/getConfiguration"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidKey          = "HeartbeatInterval"
	testValidKeyTwo       = "ConnectionTimeOut"
	testValidKeyThree     = "MeterValueSampleInterval"
	errFieldKey           = "key"
	errExceedsMaxLength   = "exceeds maximum length"
	keyMaxLength          = 50
	keyMaxLengthPlusOne   = 51
	expectedKeyCountOne   = 1
	expectedKeyCountThree = 3
	errIndexKeyZero       = "key[0]"
	errIndexKeyOne        = "key[1]"
)

func TestReq_Valid_NilKeys(t *testing.T) {
	t.Parallel()

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Key != nil {
		t.Errorf("Req() Key = %v, want nil", req.Key)
	}
}

func TestReq_Valid_EmptyKeys(t *testing.T) {
	t.Parallel()

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Key != nil {
		t.Errorf("Req() Key = %v, want nil", req.Key)
	}
}

func TestReq_Valid_SingleKey(t *testing.T) {
	t.Parallel()

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{testValidKey},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(req.Key) != expectedKeyCountOne {
		t.Errorf("Req() Key length = %d, want %d", len(req.Key), expectedKeyCountOne)
	}

	if req.Key[0].Value() != testValidKey {
		t.Errorf(st.ErrorMismatch, testValidKey, req.Key[0].Value())
	}
}

func TestReq_Valid_MultipleKeys(t *testing.T) {
	t.Parallel()

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{testValidKey, testValidKeyTwo, testValidKeyThree},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if len(req.Key) != expectedKeyCountThree {
		t.Errorf(
			"Req() Key len = %d, want %d",
			len(req.Key),
			expectedKeyCountThree,
		)
	}

	if req.Key[0].Value() != testValidKey {
		t.Errorf(st.ErrorMismatch, testValidKey, req.Key[0].Value())
	}

	if req.Key[1].Value() != testValidKeyTwo {
		t.Errorf(st.ErrorMismatch, testValidKeyTwo, req.Key[1].Value())
	}

	if req.Key[2].Value() != testValidKeyThree {
		t.Errorf(st.ErrorMismatch, testValidKeyThree, req.Key[2].Value())
	}
}

func TestReq_Valid_MaxLengthKey(t *testing.T) {
	t.Parallel()

	maxKey := strings.Repeat("k", keyMaxLength)

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{maxKey},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Key[0].Value() != maxKey {
		t.Errorf(st.ErrorMismatch, maxKey, req.Key[0].Value())
	}
}

func TestReq_EmptyKeyInList(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{testValidKey, ""},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty key in list")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}

	if !strings.Contains(err.Error(), errIndexKeyOne) {
		t.Errorf(st.ErrorWantContains, err, errIndexKeyOne)
	}
}

func TestReq_KeyTooLong(t *testing.T) {
	t.Parallel()

	longKey := strings.Repeat("k", keyMaxLengthPlusOne)

	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{longKey},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for key too long")
	}

	if !strings.Contains(err.Error(), errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_KeyInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"Invalid\x00Key"},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in key")
	}

	if !strings.Contains(err.Error(), errFieldKey) {
		t.Errorf(st.ErrorWantContains, err, errFieldKey)
	}
}

func TestReq_MultipleInvalidKeys(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"", "Invalid\x00Key"},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid keys")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errIndexKeyZero) {
		t.Errorf(st.ErrorWantContains, err, errIndexKeyZero)
	}

	if !strings.Contains(errStr, errIndexKeyOne) {
		t.Errorf(st.ErrorWantContains, err, errIndexKeyOne)
	}
}

func TestReq_FirstKeyEmptyRestValid(t *testing.T) {
	t.Parallel()

	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"", testValidKey},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty first key")
	}

	if !strings.Contains(err.Error(), errIndexKeyZero) {
		t.Errorf(st.ErrorWantContains, err, errIndexKeyZero)
	}
}
