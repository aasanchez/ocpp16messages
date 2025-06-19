package authorizetypes

import (
	"testing"
)

func TestConfirmationMessageInput_validateSucceedsWithStatus(t *testing.T) {
	t.Parallel()

	input := ConfirmationMessageInput{
		IdTagInfo: IdTagInfoInput{
			Status: "Accepted",
		},
	}

	if err := input.Validate(); err != nil {
		t.Errorf("expected validation to succeed, got: %v", err)
	}
}

func TestConfirmationMessageInput_validateFailsWithoutStatus(t *testing.T) {
	t.Parallel()

	input := ConfirmationMessageInput{
		IdTagInfo: IdTagInfoInput{
			Status: "",
		},
	}

	if err := input.Validate(); err == nil {
		t.Error("expected validation to fail due to missing Status, got nil")
	}
}

func TestConfirmationMessageInput_optionalFieldsIgnoredByValidation(t *testing.T) {
	t.Parallel()

	expiry := "2027-04-12T10:03:04-04:00"
	parent := "PARENT123"

	input := ConfirmationMessageInput{
		IdTagInfo: IdTagInfoInput{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	if err := input.Validate(); err != nil {
		t.Errorf("expected validation to succeed with optional fields, got: %v", err)
	}
}
