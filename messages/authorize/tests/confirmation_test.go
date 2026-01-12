package authorize_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Accepted" {
		t.Errorf(st.ErrorMismatch, "Accepted", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidBlocked(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{Status: "Blocked"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Blocked" {
		t.Errorf(st.ErrorMismatch, "Blocked", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{Status: "Expired"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Expired" {
		t.Errorf(st.ErrorMismatch, "Expired", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{Status: "Invalid"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "Invalid" {
		t.Errorf(st.ErrorMismatch, "Invalid", conf.IdTagInfo.Status.String())
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{Status: "ConcurrentTx"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.Status.String() != "ConcurrentTx" {
		t.Errorf(st.ErrorMismatch, "ConcurrentTx", conf.IdTagInfo.Status.String())
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := authorize.Conf(authorize.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := authorize.Conf(authorize.ConfInput{Status: ""})
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
		Status:     "Accepted",
		ExpiryDate: &expiryDate,
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
		Status:     "Accepted",
		ExpiryDate: &invalidDate,
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
		ParentIdTag: &parentTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.IdTagInfo.ParentIdTag == nil {
		t.Error("Conf() ParentIdTag = nil, want non-nil")
	}

	if conf.IdTagInfo.ParentIdTag.String() != parentTag {
		t.Errorf(st.ErrorMismatch, parentTag, conf.IdTagInfo.ParentIdTag.String())
	}
}

func TestConf_WithParentIdTagTooLong(t *testing.T) {
	t.Parallel()

	longTag := "PARENT-TAG-123456789012345" // 26 chars, max is 20
	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ParentIdTag: &longTag,
	})

	if err == nil {
		t.Error("Conf() error = nil, want error for parentIdTag too long")
	}

	if !strings.Contains(err.Error(), "parentIdTag") {
		t.Errorf("Conf() error = %v, want error containing 'parentIdTag'", err)
	}
}

func TestConf_WithEmptyParentIdTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""
	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ParentIdTag: &emptyTag,
	})

	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIdTag")
	}

	if !strings.Contains(err.Error(), "parentIdTag") {
		t.Errorf("Conf() error = %v, want error containing 'parentIdTag'", err)
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
		t.Errorf(st.ErrorMismatch, parentTag, conf.IdTagInfo.ParentIdTag.String())
	}
}
