package authorize

import (
	"errors"
	"strings"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

const errWrapMsg = "expected error to wrap ErrInvalidRequestIdTag, got: %v"

func TestRequest_validPayload(t *testing.T) {
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

func TestRequest_emptyIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: ""}
	if err := input.Validate(); err == nil {
		t.Error("expected validation error for empty IdTag, got nil")
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for empty IdTag, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}

func TestRequest_tooLongIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: strings.Repeat("A", 21)}
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for IdTag > 20 characters, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}

func TestRequest_nonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := authorizetypes.RequestPayload{IdTag: "مرحباOCPP"} // Arabic characters
	_, err := Request(input)

	if err == nil {
		t.Fatal("expected error for non-ASCII IdTag, got nil")
	}

	if !errors.Is(err, ErrInvalidRequestIdTag) {
		t.Errorf(errWrapMsg, err)
	}
}
