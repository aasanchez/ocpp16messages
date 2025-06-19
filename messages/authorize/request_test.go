package authorize

import (
	"errors"
	"strings"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

const errWrapMsg = "expected ErrInvalidRequestIdTag, got: %v"

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

	if req.IdTag.Value() != input.IdTag {
		t.Errorf("expected IdTag %q, got %q", input.IdTag, req.IdTag.Value())
	}
}

func TestRequest_EmptyIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: ""}
	if err := input.Validate(); err == nil {
		t.Error("expected validation error for empty idTag, got nil")
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected Request() to return error for empty idTag, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}

func TestRequest_TooLongIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: strings.Repeat("A", 21)}
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for idTag longer than 20 characters, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}

func TestRequest_NonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "مرحباOCPP"}
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for non-ASCII idTag, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}
