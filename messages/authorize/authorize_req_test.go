package authorize

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

func TestAuthorizeRequestValid(t *testing.T) {
	t.Parallel()

	input := "ABC1234567890XYZ78"
	req, err := Request(input)

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

func TestAuthorizeRequestEmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := Request("")
	if err == nil {
		t.Error("Expected error for empty idTag, got nil")
	}
}

func TestAuthorizeRequestTooLongIdTag(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 21) // >20 chars
	_, err := Request(input)

	if err == nil {
		t.Error("Expected error for too long idTag, got nil")
	}
}

func TestAuthorizeRequestNonASCIIIdTag(t *testing.T) {
	t.Parallel()

	input := "مرحباOCPP"
	_, err := Request(input)

	if err == nil {
		t.Error("Expected error for non-ASCII idTag, got nil")
	}
}

func TestAuthorizeRequestValidateFailsWithInvalidIdTag(t *testing.T) {
	t.Parallel()

	req := RequestMessage{
		IdTag: types.IdTokenType{}, // zero value is invalid
	}

	err := req.Validate()
	if err == nil {
		t.Error("Expected Validate() to fail for zero-value IdTag")
	}
}

func TestAuthorizeRequestString(t *testing.T) {
	t.Parallel()

	req, err := Request("TAG123456789012345")
	if err != nil {
		t.Fatalf("unexpected error creating request: %v", err)
	}

	output := req.String()
	if !strings.Contains(output, "TAG123456789012345") {
		t.Errorf("expected String() to include idTag, got: %s", output)
	}
}
