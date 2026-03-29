package datatransfer_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/datatransfer"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testValidVendorId   = "com.example.vendor"
	testValidMessageId  = "CustomMessage"
	testValidData       = `{"key": "value"}`
	errFieldVendorId    = "vendorId"
	errFieldMessageId   = "messageId"
	errExceedsMaxLength = "exceeds maximum length"
	vendorIdMaxPlusOne  = 256
	messageIdMaxPlusOne = 51
)

func TestReq_Valid_VendorIdOnly(t *testing.T) {
	t.Parallel()

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorId.Value() != testValidVendorId {
		t.Errorf(types.ErrorMismatch, testValidVendorId, req.VendorId.Value())
	}

	if req.MessageId != nil {
		t.Error("Req() MessageId != nil, want nil")
	}

	if req.Data != nil {
		t.Error("Req() Data != nil, want nil")
	}
}

func TestReq_Valid_WithMessageId(t *testing.T) {
	t.Parallel()

	messageId := testValidMessageId

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &messageId,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.MessageId == nil {
		t.Error("Req() MessageId = nil, want non-nil")

		return
	}

	if req.MessageId.Value() != testValidMessageId {
		t.Errorf(types.ErrorMismatch, testValidMessageId, req.MessageId.Value())
	}
}

func TestReq_Valid_WithData(t *testing.T) {
	t.Parallel()

	data := testValidData

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: nil,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(types.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_Valid_Complete(t *testing.T) {
	t.Parallel()

	messageId := testValidMessageId
	data := testValidData

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &messageId,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorId.Value() != testValidVendorId {
		t.Errorf(types.ErrorMismatch, testValidVendorId, req.VendorId.Value())
	}

	if req.MessageId == nil {
		t.Error("Req() MessageId = nil, want non-nil")

		return
	}

	if req.MessageId.Value() != testValidMessageId {
		t.Errorf(types.ErrorMismatch, testValidMessageId, req.MessageId.Value())
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(types.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_EmptyVendorId(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  "",
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty vendorId")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_VendorIdTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 255 characters
	longVendorId := strings.Repeat("a", vendorIdMaxPlusOne)

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  longVendorId,
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for vendorId too long")
	}

	if !strings.Contains(err.Error(), errFieldVendorId) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorId)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_VendorIdInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  "vendor\x00id",
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in vendorId")
	}

	if !strings.Contains(err.Error(), errFieldVendorId) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorId)
	}
}

func TestReq_EmptyMessageId(t *testing.T) {
	t.Parallel()

	emptyMessageId := ""

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &emptyMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty messageId")
	}

	if !strings.Contains(err.Error(), errFieldMessageId) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageId)
	}
}

func TestReq_MessageIdTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 50 characters
	longMessageId := strings.Repeat("m", messageIdMaxPlusOne)

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &longMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for messageId too long")
	}

	if !strings.Contains(err.Error(), errFieldMessageId) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageId)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_MessageIdInvalidChars(t *testing.T) {
	t.Parallel()

	invalidMessageId := "msg\x00id"

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &invalidMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in messageId")
	}

	if !strings.Contains(err.Error(), errFieldMessageId) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageId)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidMessageId := ""

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  "",
		MessageId: &invalidMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, errFieldVendorId) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorId)
	}

	if !strings.Contains(errStr, errFieldMessageId) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageId)
	}
}
