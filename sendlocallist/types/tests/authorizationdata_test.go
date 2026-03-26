package types_test

import (
	"strings"
	"testing"

	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	validIdTag        = "RFID12345"
	validStatus       = "Accepted"
	validExpiryDate   = "2025-12-31T23:59:59Z"
	validParentIdTag  = "PARENT123"
	errIdTag          = "IdTag"
	errIdTagInfo      = "IdTagInfo"
	errExpiryDate     = "ExpiryDate"
	errParentIdTag    = "ParentIdTag"
	errEmptyValue     = "cannot be empty"
	errInvalidValue   = "invalid value"
	idTagTooLong      = "ThisIdTagIsTooLongForCiString20"
	invalidStatusStr  = "InvalidStatus"
	invalidDateFormat = "not-a-date"
	fieldIdTagInfo    = "IdTagInfo"
)

func TestNewAuthorizationData_Valid_WithIdTagOnly(t *testing.T) {
	t.Parallel()

	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag:     validIdTag,
		IdTagInfo: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if authData.IdTag.String() != validIdTag {
		t.Errorf(st.ErrorMismatch, validIdTag, authData.IdTag.String())
	}

	if authData.IdTagInfo != nil {
		t.Errorf("IdTagInfo = %v, want nil", authData.IdTagInfo)
	}
}

func TestNewAuthorizationData_Valid_WithIdTagInfo(t *testing.T) {
	t.Parallel()

	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if authData.IdTag.String() != validIdTag {
		t.Errorf(st.ErrorMismatch, validIdTag, authData.IdTag.String())
	}

	if authData.IdTagInfo == nil {
		t.Errorf(st.ErrorWantNonNil, fieldIdTagInfo)
	}
}

func TestNewAuthorizationData_Valid_WithExpiryDate(t *testing.T) {
	t.Parallel()

	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  ptrString(validExpiryDate),
			ParentIdTag: nil,
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if authData.IdTagInfo == nil {
		t.Errorf(st.ErrorWantNonNil, fieldIdTagInfo)

		return
	}

	if authData.IdTagInfo.ExpiryDate == nil {
		t.Errorf(st.ErrorWantNonNil, "ExpiryDate")
	}
}

func TestNewAuthorizationData_Valid_WithParentIdTag(t *testing.T) {
	t.Parallel()

	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  nil,
			ParentIdTag: ptrString(validParentIdTag),
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if authData.IdTagInfo == nil {
		t.Errorf(st.ErrorWantNonNil, fieldIdTagInfo)

		return
	}

	if authData.IdTagInfo.ParentIdTag == nil {
		t.Errorf(st.ErrorWantNonNil, "ParentIdTag")
	}
}

func TestNewAuthorizationData_Valid_AllFields(t *testing.T) {
	t.Parallel()

	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  ptrString(validExpiryDate),
			ParentIdTag: ptrString(validParentIdTag),
		},
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if authData.IdTag.String() != validIdTag {
		t.Errorf(st.ErrorMismatch, validIdTag, authData.IdTag.String())
	}

	if authData.IdTagInfo == nil {
		t.Errorf(st.ErrorWantNonNil, fieldIdTagInfo)

		return
	}

	if authData.IdTagInfo.ExpiryDate == nil {
		t.Errorf(st.ErrorWantNonNil, "ExpiryDate")
	}

	if authData.IdTagInfo.ParentIdTag == nil {
		t.Errorf(st.ErrorWantNonNil, "ParentIdTag")
	}
}

func TestNewAuthorizationData_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag:     "",
		IdTagInfo: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty idTag")
	}

	if !strings.Contains(err.Error(), errIdTag) {
		t.Errorf(st.ErrorWantContains, err, errIdTag)
	}

	if !strings.Contains(err.Error(), errEmptyValue) {
		t.Errorf(st.ErrorWantContains, err, errEmptyValue)
	}
}

func TestNewAuthorizationData_IdTagTooLong(t *testing.T) {
	t.Parallel()

	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag:     idTagTooLong,
		IdTagInfo: nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "idTag too long")
	}

	if !strings.Contains(err.Error(), errIdTag) {
		t.Errorf(st.ErrorWantContains, err, errIdTag)
	}
}

func TestNewAuthorizationData_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      invalidStatusStr,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status")
	}

	if !strings.Contains(err.Error(), errIdTagInfo) {
		t.Errorf(st.ErrorWantContains, err, errIdTagInfo)
	}

	if !strings.Contains(err.Error(), errInvalidValue) {
		t.Errorf(st.ErrorWantContains, err, errInvalidValue)
	}
}

func TestNewAuthorizationData_InvalidExpiryDate(t *testing.T) {
	t.Parallel()

	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  ptrString(invalidDateFormat),
			ParentIdTag: nil,
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid expiry date")
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}
}

func TestNewAuthorizationData_InvalidParentIdTag(t *testing.T) {
	t.Parallel()

	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: validIdTag,
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      validStatus,
			ExpiryDate:  nil,
			ParentIdTag: ptrString(idTagTooLong),
		},
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid parent idTag")
	}

	if !strings.Contains(err.Error(), errParentIdTag) {
		t.Errorf(st.ErrorWantContains, err, errParentIdTag)
	}
}

func ptrString(stringVal string) *string {
	return &stringVal
}
