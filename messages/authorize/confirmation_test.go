package authorize

import (
	"testing"
	"time"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestAuthorizeConfirmationValidInput(t *testing.T) {
	t.Parallel()

	parent, err := authorizetypes.IdToken("GROUP123")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	expiry := time.Now().Add(24 * time.Hour).UTC()
	info := authorizetypes.IdTagInfoType{
		Status:      authorizetypes.Accepted,
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

	invalidInfo := authorizetypes.IdTagInfoType{
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
		IdTagInfo: authorizetypes.IdTagInfoType{
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

	parent, err := authorizetypes.IdToken("GROUPABC")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	expiry := time.Now().UTC()
	info := authorizetypes.IdTagInfoType{
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
		Status:      authorizetypes.Accepted,
	}

	msg := ConfirmationMessage{
		IdTagInfo: info,
	}

	want := info.String()
	got := msg.String()

	if got != want {
		t.Errorf("unexpected String() output:\nwant: %s\ngot : %s", want, got)
	}
}
