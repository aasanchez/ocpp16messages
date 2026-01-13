package changeConfiguration_test

import (
	"strings"
	"testing"

	cc "github.com/aasanchez/ocpp16messages/messages/changeConfiguration"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	errKey   = "key"
	errValue = "value"

	validKey   = "HeartbeatInterval"
	validValue = "900"

	emptyString     = ""
	repeatCharA     = "a"
	repeatCharX     = "x"
	maxKeyLen       = 50
	maxValueLen     = 500
	exceedsKeyLen   = 51
	exceedsValueLen = 501
)

func TestReq_Valid_SimpleKeyValue(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{Key: validKey, Value: validValue}

	req, err := cc.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Key.Value() != validKey {
		t.Errorf(st.ErrorMismatchValue, validKey, req.Key.Value())
	}

	if req.Value.Value() != validValue {
		t.Errorf(st.ErrorMismatchValue, validValue, req.Value.Value())
	}
}

func TestReq_EmptyValue(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{Key: validKey, Value: emptyString}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty value")
	}

	if !strings.Contains(err.Error(), errValue) {
		t.Errorf(st.ErrorWantContains, err, errValue)
	}
}

func TestReq_Valid_MaxKeyLength(t *testing.T) {
	t.Parallel()

	maxKey := strings.Repeat(repeatCharA, maxKeyLen)
	input := cc.ReqInput{Key: maxKey, Value: validValue}

	req, err := cc.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Key.Value() != maxKey {
		t.Errorf(st.ErrorMismatchValue, maxKey, req.Key.Value())
	}
}

func TestReq_Valid_MaxValueLength(t *testing.T) {
	t.Parallel()

	maxValue := strings.Repeat(repeatCharX, maxValueLen)
	input := cc.ReqInput{Key: validKey, Value: maxValue}

	req, err := cc.Req(input)
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Value.Value() != maxValue {
		t.Errorf(st.ErrorMismatchValue, maxValue, req.Value.Value())
	}
}

func TestReq_EmptyKey(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{Key: emptyString, Value: validValue}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty key")
	}

	if !strings.Contains(err.Error(), errKey) {
		t.Errorf(st.ErrorWantContains, err, errKey)
	}
}

func TestReq_KeyTooLong(t *testing.T) {
	t.Parallel()

	longKey := strings.Repeat(repeatCharA, exceedsKeyLen)
	input := cc.ReqInput{Key: longKey, Value: validValue}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "key too long")
	}

	if !strings.Contains(err.Error(), errKey) {
		t.Errorf(st.ErrorWantContains, err, errKey)
	}
}

func TestReq_ValueTooLong(t *testing.T) {
	t.Parallel()

	longValue := strings.Repeat(repeatCharX, exceedsValueLen)
	input := cc.ReqInput{Key: validKey, Value: longValue}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "value too long")
	}

	if !strings.Contains(err.Error(), errValue) {
		t.Errorf(st.ErrorWantContains, err, errValue)
	}
}

func TestReq_KeyWithNonPrintableChar(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{Key: "Test\x00Key", Value: validValue}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "key with non-printable char")
	}

	if !strings.Contains(err.Error(), errKey) {
		t.Errorf(st.ErrorWantContains, err, errKey)
	}
}

func TestReq_ValueWithNonPrintableChar(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{Key: validKey, Value: "Test\x7FValue"}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "value with non-printable char")
	}

	if !strings.Contains(err.Error(), errValue) {
		t.Errorf(st.ErrorWantContains, err, errValue)
	}
}

func TestReq_MultipleErrors_EmptyKeyAndValueTooLong(t *testing.T) {
	t.Parallel()

	longValue := strings.Repeat(repeatCharX, exceedsValueLen)
	input := cc.ReqInput{Key: emptyString, Value: longValue}

	_, err := cc.Req(input)
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errKey) {
		t.Errorf(st.ErrorWantContains, err, errKey)
	}

	if !strings.Contains(err.Error(), errValue) {
		t.Errorf(st.ErrorWantContains, err, errValue)
	}
}
