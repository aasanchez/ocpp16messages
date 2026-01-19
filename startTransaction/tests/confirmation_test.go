package startTransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/startTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAccepted     = "Accepted"
	errParentIdTag     = "parentIdTag"
	errExpiryDate      = "expiryDate"
	errTransactionId   = "transactionId"
	testTransactionId  = 12345
	testTransactionId2 = 12346
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expectedTransactionId := uint16(testTransactionId)
	if conf.TransactionId.Value() != expectedTransactionId {
		t.Errorf(
			st.ErrorMismatch,
			expectedTransactionId,
			conf.TransactionId.Value(),
		)
	}

	if conf.IdTagInfo.Status.String() != statusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			statusAccepted,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_ValidBlocked(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "Blocked",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Blocked" {
		t.Errorf(st.ErrorMismatch, "Blocked", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "Expired",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Expired" {
		t.Errorf(st.ErrorMismatch, "Expired", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "Invalid",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Invalid" {
		t.Errorf(st.ErrorMismatch, "Invalid", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "ConcurrentTx",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "ConcurrentTx" {
		t.Errorf(
			st.ErrorMismatch,
			"ConcurrentTx",
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "Unknown",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_TransactionIdNegative(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testValueNegativeOne,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for negative transactionId")
	}

	if !strings.Contains(err.Error(), errTransactionId) {
		t.Errorf(st.ErrorWantContains, err, errTransactionId)
	}
}

func TestConf_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    &expiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.ExpiryDate == nil {
		t.Error("Conf() ExpiryDate = nil, want non-nil")
	}
}

func TestConf_WithInvalidExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "not-a-date"

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    &invalidDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}
}

func TestConf_WithParentIdTag(t *testing.T) {
	t.Parallel()

	parentTag := "PARENT-TAG-123"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIdTag:   &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.ParentIdTag == nil {
		t.Error("Conf() ParentIdTag = nil, want non-nil")
	}

	if conf.IdTagInfo.ParentIdTag.String() != parentTag {
		t.Errorf(
			st.ErrorMismatch,
			parentTag,
			conf.IdTagInfo.ParentIdTag.String(),
		)
	}
}

func TestConf_WithParentIdTagTooLong(t *testing.T) {
	t.Parallel()

	longTag := "PARENT-TAG-123456789012345" // 26 chars, max is 20

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIdTag:   &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIdTag too long")
	}

	if !strings.Contains(err.Error(), errParentIdTag) {
		t.Errorf(st.ErrorWantContains, err, errParentIdTag)
	}
}

func TestConf_WithEmptyParentIdTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIdTag:   &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIdTag")
	}

	if !strings.Contains(err.Error(), errParentIdTag) {
		t.Errorf(st.ErrorWantContains, err, errParentIdTag)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId2,
		Status:        statusAccepted,
		ExpiryDate:    &expiryDate,
		ParentIdTag:   &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	expectedTransactionId := uint16(testTransactionId2)
	if conf.TransactionId.Value() != expectedTransactionId {
		t.Errorf(
			st.ErrorMismatch,
			expectedTransactionId,
			conf.TransactionId.Value(),
		)
	}

	if conf.IdTagInfo.Status.String() != statusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			statusAccepted,
			conf.IdTagInfo.Status.String(),
		)
	}

	if conf.IdTagInfo.ExpiryDate == nil {
		t.Error("Conf() ExpiryDate = nil, want non-nil")
	}

	if conf.IdTagInfo.ParentIdTag == nil {
		t.Error("Conf() ParentIdTag = nil, want non-nil")
	}

	if conf.IdTagInfo.ParentIdTag.String() != parentTag {
		t.Errorf(
			st.ErrorMismatch,
			parentTag,
			conf.IdTagInfo.ParentIdTag.String(),
		)
	}
}

func TestConf_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testValueNegativeOne,
		Status:        "Invalid-Status",
		ExpiryDate:    &invalidDate,
		ParentIdTag:   &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errTransactionId) {
		t.Errorf(st.ErrorWantContains, err, errTransactionId)
	}

	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}

	if !strings.Contains(errStr, errParentIdTag) {
		t.Errorf(st.ErrorWantContains, err, errParentIdTag)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "invalid"

	_, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: testTransactionId,
		Status:        "BadStatus",
		ExpiryDate:    &invalidDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}
}
