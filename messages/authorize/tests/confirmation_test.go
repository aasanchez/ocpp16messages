package authorize_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	StatusAccepted = "Accepted"
	ErrParentIdTag = "parentIdTag"
	ErrExpiryDate  = "expiryDate"
	ErrWantParent  = "Conf() error = %v, want error with 'parentIdTag'"
	ErrWantExpiry  = "Conf() error = %v, want error with 'expiryDate'"
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      StatusAccepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != StatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			StatusAccepted,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_ValidBlocked(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Blocked",
		ExpiryDate:  nil,
		ParentIdTag: nil,
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

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Expired",
		ExpiryDate:  nil,
		ParentIdTag: nil,
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

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Invalid",
		ExpiryDate:  nil,
		ParentIdTag: nil,
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

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "ConcurrentTx",
		ExpiryDate:  nil,
		ParentIdTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Unknown",
		ExpiryDate:  nil,
		ParentIdTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIdTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &invalidDate,
		ParentIdTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), "expiryDate") {
		t.Errorf("Conf() error = %v, want error containing 'expiryDate'", err)
	}
}

func TestConf_WithParentIdTag(t *testing.T) {
	t.Parallel()

	parentTag := "PARENT-TAG-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIdTag: &parentTag,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIdTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIdTag too long")
	}

	if !strings.Contains(err.Error(), ErrParentIdTag) {
		t.Errorf(ErrWantParent, err)
	}
}

func TestConf_WithEmptyParentIdTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIdTag: &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIdTag")
	}

	if !strings.Contains(err.Error(), ErrParentIdTag) {
		t.Errorf(ErrWantParent, err)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIdTag: &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Accepted" {
		t.Errorf(st.ErrorMismatch, "Accepted", conf.IdTagInfo.Status.String())
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Invalid-Status",
		ExpiryDate:  &invalidDate,
		ParentIdTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	// Check that all errors are present
	errStr := err.Error()
	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, ErrExpiryDate) {
		t.Errorf(ErrWantExpiry, err)
	}

	if !strings.Contains(errStr, ErrParentIdTag) {
		t.Errorf(ErrWantParent, err)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "invalid"

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "BadStatus",
		ExpiryDate:  &invalidDate,
		ParentIdTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, ErrExpiryDate) {
		t.Errorf(ErrWantExpiry, err)
	}
}
