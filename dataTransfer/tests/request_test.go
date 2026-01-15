package dataTransfer_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/dataTransfer"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidVendorId  = "com.example.vendor"
	testValidMessageId = "CustomMessage"
	testValidData      = `{"key": "value"}`
)

func TestReq_Valid_VendorIdOnly(t *testing.T) {
	t.Parallel()

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.VendorId.Value() != testValidVendorId {
		t.Errorf(st.ErrorMismatch, testValidVendorId, req.VendorId.Value())
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

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &messageId,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.MessageId == nil {
		t.Error("Req() MessageId = nil, want non-nil")

		return
	}

	if req.MessageId.Value() != testValidMessageId {
		t.Errorf(st.ErrorMismatch, testValidMessageId, req.MessageId.Value())
	}
}

func TestReq_Valid_WithData(t *testing.T) {
	t.Parallel()

	data := testValidData

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: nil,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(st.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_Valid_Complete(t *testing.T) {
	t.Parallel()

	messageId := testValidMessageId
	data := testValidData

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &messageId,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.VendorId.Value() != testValidVendorId {
		t.Errorf(st.ErrorMismatch, testValidVendorId, req.VendorId.Value())
	}

	if req.MessageId == nil {
		t.Error("Req() MessageId = nil, want non-nil")

		return
	}

	if req.MessageId.Value() != testValidMessageId {
		t.Errorf(st.ErrorMismatch, testValidMessageId, req.MessageId.Value())
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(st.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_EmptyVendorId(t *testing.T) {
	t.Parallel()

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "",
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty vendorId")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestReq_VendorIdTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 255 characters
	longVendorId := strings.Repeat("a", 256)

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  longVendorId,
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for vendorId too long")
	}

	if !strings.Contains(err.Error(), "vendorId") {
		t.Errorf(st.ErrorWantContains, err, "vendorId")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(st.ErrorWantContains, err, "exceeds maximum length")
	}
}

func TestReq_VendorIdInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "vendor\x00id",
		MessageId: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in vendorId")
	}

	if !strings.Contains(err.Error(), "vendorId") {
		t.Errorf(st.ErrorWantContains, err, "vendorId")
	}
}

func TestReq_EmptyMessageId(t *testing.T) {
	t.Parallel()

	emptyMessageId := ""

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &emptyMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty messageId")
	}

	if !strings.Contains(err.Error(), "messageId") {
		t.Errorf(st.ErrorWantContains, err, "messageId")
	}
}

func TestReq_MessageIdTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 50 characters
	longMessageId := strings.Repeat("m", 51)

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &longMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for messageId too long")
	}

	if !strings.Contains(err.Error(), "messageId") {
		t.Errorf(st.ErrorWantContains, err, "messageId")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(st.ErrorWantContains, err, "exceeds maximum length")
	}
}

func TestReq_MessageIdInvalidChars(t *testing.T) {
	t.Parallel()

	invalidMessageId := "msg\x00id"

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  testValidVendorId,
		MessageId: &invalidMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in messageId")
	}

	if !strings.Contains(err.Error(), "messageId") {
		t.Errorf(st.ErrorWantContains, err, "messageId")
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidMessageId := ""

	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "",
		MessageId: &invalidMessageId,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "vendorId") {
		t.Errorf(st.ErrorWantContains, err, "vendorId")
	}

	if !strings.Contains(errStr, "messageId") {
		t.Errorf(st.ErrorWantContains, err, "messageId")
	}
}
