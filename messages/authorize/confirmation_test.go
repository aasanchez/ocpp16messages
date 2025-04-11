package authorize

import (
	"fmt"
	"testing"
	"time"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestAuthorizeConfirmationValidInput(t *testing.T) {
	t.Parallel()

	parent, err := types.IdToken("GROUP123")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	expiry := time.Now().Add(24 * time.Hour).UTC()
	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	msg, err := Confirmation(info)
	if err != nil {
		t.Fatalf("unexpected error constructing AuthorizeConfirmationMessage: %v", err)
	}

	if err := msg.Validate(); err != nil {
		t.Errorf("expected message to be valid, got error: %v", err)
	}
}

func TestAuthorizeConfirmationInvalidIdTagInfo(t *testing.T) {
	t.Parallel()

	invalidInfo := types.IdTagInfoType{
		Status:      "InvalidStatus",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	_, err := Confirmation(invalidInfo)
	if err == nil {
		t.Error("expected error for invalid IdTagInfo, got nil")
	}
}

func TestAuthorizeConfirmationValidateFailsWithInvalidData(t *testing.T) {
	t.Parallel()

	invalid := ConfirmationMessage{
		IdTagInfo: types.IdTagInfoType{
			Status:      "NotAValidStatus",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	err := invalid.Validate()
	if err == nil {
		t.Error("expected Validate() to fail, got nil")
	}
}

func TestAuthorizeConfirmationString(t *testing.T) {
	t.Parallel()

	parent, err := types.IdToken("GROUPABC")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	expiry := time.Now().UTC()
	info := types.IdTagInfoType{
		Status:      types.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	msg := ConfirmationMessage{
		IdTagInfo: info,
	}

	want := fmt.Sprintf("Authorize.conf{%s}", info.String())
	got := msg.String()

	if got != want {
		t.Errorf("unexpected String() output:\nwant: %s\ngot : %s", want, got)
	}
}
