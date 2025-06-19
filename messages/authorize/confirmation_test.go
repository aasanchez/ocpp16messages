package authorize

// import (
// 	"testing"

// 	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// const errCreateStatusFmt = "failed to create status: %v"

// func TestConfirmation_validStatusOnly(t *testing.T) {
// 	t.Parallel()

// 	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
// 	if err != nil {
// 		t.Fatalf(errCreateStatusFmt, err)
// 	}

// 	info := authorizetypes.IdTagInfoType{
// 		Status:      status,
// 		ExpiryDate:  nil,
// 		ParentIdTag: nil,
// 	}

// 	msg, err := Confirmation(info)
// 	if err != nil {
// 		t.Fatalf(errCreateStatusFmt, err)
// 	}

// 	if err := msg.Validate(); err != nil {
// 		t.Errorf("expected message to be valid, got: %v", err)
// 	}
// }

// func TestConfirmation_withExpiryDate(t *testing.T) {
// 	t.Parallel()

// 	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
// 	if err != nil {
// 		t.Fatalf(errCreateStatusFmt, err)
// 	}

// 	expiry, err := sharedtypes.DateTime("2027-04-12T10:03:04Z")
// 	if err != nil {
// 		t.Fatalf("failed to parse DateTime: %v", err)
// 	}

// 	info := authorizetypes.IdTagInfoType{
// 		Status:      status,
// 		ExpiryDate:  &expiry,
// 		ParentIdTag: nil,
// 	}

// 	msg, err := Confirmation(info)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if err := msg.Validate(); err != nil {
// 		t.Errorf("expected message to be valid with expiryDate, got: %v", err)
// 	}
// }

// func TestConfirmation_withParentIdTag(t *testing.T) {
// 	t.Parallel()

// 	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
// 	if err != nil {
// 		t.Fatalf(errCreateStatusFmt, err)
// 	}

// 	parent, err := authorizetypes.IdToken("TOKEN123456789")
// 	if err != nil {
// 		t.Fatalf("failed to create IdToken: %v", err)
// 	}

// 	info := authorizetypes.IdTagInfoType{
// 		Status:      status,
// 		ExpiryDate:  nil,
// 		ParentIdTag: &parent,
// 	}

// 	msg, err := Confirmation(info)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if err := msg.Validate(); err != nil {
// 		t.Errorf("expected valid message, got: %v", err)
// 	}
// }

// func TestConfirmation_withAllFields(t *testing.T) {
// 	t.Parallel()

// 	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
// 	if err != nil {
// 		t.Fatalf(errCreateStatusFmt, err)
// 	}

// 	expiry, err := sharedtypes.DateTime("2028-01-01T00:00:00Z")
// 	if err != nil {
// 		t.Fatalf("failed to parse expiry date: %v", err)
// 	}

// 	parent, err := authorizetypes.IdToken("PARENT123")
// 	if err != nil {
// 		t.Fatalf("failed to create IdToken: %v", err)
// 	}

// 	info := authorizetypes.IdTagInfoType{
// 		Status:      status,
// 		ExpiryDate:  &expiry,
// 		ParentIdTag: &parent,
// 	}

// 	msg, err := Confirmation(info)
// 	if err != nil {
// 		t.Fatalf("unexpected error creating confirmation: %v", err)
// 	}

// 	if err := msg.Validate(); err != nil {
// 		t.Errorf("expected message to be valid with all fields, got: %v", err)
// 	}
// }

// func TestConfirmation_invalidStatus_shouldFail(t *testing.T) {
// 	t.Parallel()

// 	_, err := authorizetypes.AuthorizationStatus("Wrong")
// 	if err == nil {
// 		t.Error("expected error for invalid status, got nil")
// 	}
// }

// func TestConfirmation_validationFailsForZeroValueStatus(t *testing.T) {
// 	t.Parallel()

// 	info := authorizetypes.IdTagInfoType{
// 		Status:      authorizetypes.AuthorizationStatusType{},
// 		ExpiryDate:  nil,
// 		ParentIdTag: nil,
// 	}

// 	msg := ConfirmationMessage{IdTagInfo: info}

// 	if err := msg.Validate(); err == nil {
// 		t.Error("expected validation to fail for zero-value status, got nil")
// 	}
// }

// func TestConfirmation_invalidInfo_shouldReturnError(t *testing.T) {
// 	t.Parallel()

// 	invalidInfo := authorizetypes.IdTagInfoType{
// 		Status:      authorizetypes.AuthorizationStatusType{},
// 		ExpiryDate:  nil,
// 		ParentIdTag: nil,
// 	}

// 	_, err := Confirmation(invalidInfo)
// 	if err == nil {
// 		t.Fatal("expected error for invalid IdTagInfo, got nil")
// 	}

// 	expected := "confirmation: invalid IdTagInfo"
// 	if err != nil && err.Error()[:len(expected)] != expected {
// 		t.Errorf("unexpected error message: got %q, want prefix %q", err.Error(), expected)
// 	}
// }
