package stoptransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/stoptransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAccepted       = "Accepted"
	statusBlocked        = "Blocked"
	statusExpired        = "Expired"
	statusInvalid        = "Invalid"
	statusConcurrentTx   = "ConcurrentTx"
	errConfStatus        = "status"
	errConfExpiryDate    = "expiryDate"
	errConfParentId      = "parentIdTag"
	errConfIdTagInfoNil  = "Conf() IdTagInfo = nil, want non-nil"
	errConfExpiryDateNil = "Conf() ExpiryDate = nil, want non-nil"
	errConfParentIdNil   = "Conf() ParentIdTag = nil, want non-nil"
)

func TestConf_NoIdTagInfo(t *testing.T) {
	t.Parallel()

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      nil,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo != nil {
		t.Error("Conf() IdTagInfo != nil, want nil for empty input")
	}
}

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	status := statusAccepted

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
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

	status := statusBlocked

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.Status.String() != statusBlocked {
		t.Errorf(
			st.ErrorMismatch,
			statusBlocked,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	status := statusExpired

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.Status.String() != statusExpired {
		t.Errorf(
			st.ErrorMismatch,
			statusExpired,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	status := statusInvalid

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.Status.String() != statusInvalid {
		t.Errorf(
			st.ErrorMismatch,
			statusInvalid,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	status := statusConcurrentTx

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.Status.String() != statusConcurrentTx {
		t.Errorf(
			st.ErrorMismatch,
			statusConcurrentTx,
			conf.IdTagInfo.Status.String(),
		)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	status := "Unknown"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
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

	status := ""

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
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

	status := statusAccepted
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.ExpiryDate == nil {
		t.Error(errConfExpiryDateNil)
	}
}

func TestConf_WithInvalidExpiryDate(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	invalidDate := "not-a-date"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &invalidDate,
		ParentIdTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), errConfExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errConfExpiryDate)
	}
}

func TestConf_WithParentIdTag(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	parentTag := "PARENT-TAG-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.ParentIdTag == nil {
		t.Error(errConfParentIdNil)
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

	status := statusAccepted
	longTag := "PARENT-TAG-123456789012345" // 26 chars, max is 20

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIdTag too long")
	}

	if !strings.Contains(err.Error(), errConfParentId) {
		t.Errorf(st.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_WithEmptyParentIdTag(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	emptyTag := ""

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIdTag")
	}

	if !strings.Contains(err.Error(), errConfParentId) {
		t.Errorf(st.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIdTag: &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo == nil {
		t.Error(errConfIdTagInfoNil)

		return
	}

	if conf.IdTagInfo.Status.String() != statusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			statusAccepted,
			conf.IdTagInfo.Status.String(),
		)
	}

	if conf.IdTagInfo.ExpiryDate == nil {
		t.Error(errConfExpiryDateNil)
	}

	if conf.IdTagInfo.ParentIdTag == nil {
		t.Error(errConfParentIdNil)
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

	invalidStatus := "Invalid-Status"
	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &invalidStatus,
		ExpiryDate:  &invalidDate,
		ParentIdTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errConfStatus) {
		t.Errorf(st.ErrorWantContains, err, errConfStatus)
	}

	if !strings.Contains(errStr, errConfExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errConfExpiryDate)
	}

	if !strings.Contains(errStr, errConfParentId) {
		t.Errorf(st.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidStatus := "BadStatus"
	invalidDate := "invalid"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &invalidStatus,
		ExpiryDate:  &invalidDate,
		ParentIdTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errConfStatus) {
		t.Errorf(st.ErrorWantContains, err, errConfStatus)
	}

	if !strings.Contains(errStr, errConfExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errConfExpiryDate)
	}
}
