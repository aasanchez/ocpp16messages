package authorize

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestAuthorizeRequest_Valid(t *testing.T) {
	t.Parallel()

	input := "ABC1234567890XYZ78"
	req, err := AuthorizeRequest(input)

	if err != nil {
		t.Fatalf("Expected no error for valid idTag, got: %v", err)
	}

	if req.IdTag.String() != input {
		t.Errorf("Expected idTag to be %q, got %q", input, req.IdTag.String())
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Expected Validate() to succeed, got error: %v", err)
	}
}

func TestAuthorizeRequest_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := AuthorizeRequest("")
	if err == nil {
		t.Error("Expected error for empty idTag, got nil")
	}
}

func TestAuthorizeRequest_TooLongIdTag(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 21) // >20 chars
	_, err := AuthorizeRequest(input)

	if err == nil {
		t.Error("Expected error for too long idTag, got nil")
	}
}

func TestAuthorizeRequest_NonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := "مرحباOCPP"
	_, err := AuthorizeRequest(input)

	if err == nil {
		t.Error("Expected error for non-ASCII idTag, got nil")
	}
}

func TestAuthorizeRequest_ValidateFails_WithInvalidIdTag(t *testing.T) {
	t.Parallel()

	req := AuthorizeRequestMessage{
		IdTag: types.IdTag{}, // zero value is invalid
	}

	err := req.Validate()
	if err == nil {
		t.Error("Expected Validate() to fail for zero-value IdTag")
	}
}

func TestAuthorizeRequest_String(t *testing.T) {
	t.Parallel()

	req, err := AuthorizeRequest("TAG123456789012345")
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	output := req.String()
	if !strings.Contains(output, "TAG123456789012345") {
		t.Errorf("expected String() to include idTag, got: %s", output)
	}
}
