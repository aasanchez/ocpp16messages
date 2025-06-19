package authorize

import (
	"strings"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestRequest_Valid(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "ABC1234567890XYZ78"}
	if err := input.Validate(); err != nil {
		t.Fatalf("input validation failed unexpectedly: %v", err)
	}

	req, err := Request(input)
	if err != nil {
		t.Fatalf("Request() returned unexpected error: %v", err)
	}

	if err := req.Validate(); err != nil {
		t.Errorf("expected Validate() to succeed, got: %v", err)
	}
}

func TestRequest_EmptyIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestMessageInput{IdTag: ""}
	if err := input.Validate(); err == nil {
		t.Error("expected validation error for empty idTag, got nil")
	}

	_, err := Request(input)
	if err == nil {
		t.Error("expected Request() to return error for empty idTag, got nil")
	}
}

func TestRequest_TooLongIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: strings.Repeat("A", 21)}
	_, err := Request(input)

	if err == nil {
		t.Error("expected error for idTag longer than 20 characters, got nil")
	}
}

func TestRequest_NonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "مرحباOCPP"}
	_, err := Request(input)

	if err == nil {
		t.Error("expected error for non-ASCII idTag, got nil")
	}
}

func TestRequest_ValidateFailsWithInvalidIdTag(t *testing.T) {
	t.Parallel()

	req := RequestMessage{
		IdTag: authorizetypes.IdToken{}, // invalid, zero-value
	}

	if err := req.Validate(); err == nil {
		t.Error("expected Validate() to fail for invalid IdTag")
	}
}
