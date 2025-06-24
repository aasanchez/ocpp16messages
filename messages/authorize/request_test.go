package authorize

import (
	"strings"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestAuthorizeRequest_validPayload(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "ABC1234567890XYZ78"}
	if err := input.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	msg, err := Request(input)
	if err != nil {
		t.Fatalf("Request() returned unexpected error: %v", err)
	}

	if msg.IdTag.Value() != input.IdTag {
		t.Errorf("expected IdTag %q, got %q", input.IdTag, msg.IdTag.Value())
	}
}

func TestAuthorizeRequest_emptyIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: ""}
	if err := input.Validate(); err == nil {
		t.Error("expected validation error for empty IdTag, got nil")
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for empty IdTag, got nil")
	}

	expected := "request -> invalid idTag -> ciString.Validate: value must not be empty"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf(sharedtypes.ErrContainsFmt, expected, err.Error())
	}
}

func TestAuthorizeRequest_tooLongIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: strings.Repeat("A", 21)}
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for IdTag > 20 characters, got nil")
	}

	expected := "request -> invalid idTag -> ciString.Validate: value exceeds maximum allowed length (got length 21, max 20)"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf(sharedtypes.ErrContainsFmt, expected, err.Error())
	}
}

func TestAuthorizeRequest_nonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "مرحباOCPP"}
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for non-ASCII IdTag, got nil")
	}

	expected := "request -> invalid idTag -> ciString.Validate: value contains non-printable ASCII characters"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf(sharedtypes.ErrContainsFmt, expected, err.Error())
	}
}
