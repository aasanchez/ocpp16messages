package types

import (
	"testing"
)

func TestIdTagValidInput(t *testing.T) {
	t.Parallel()

	validStr := "ABCDEFG1234567890"
	idTag, err := IdTag(validStr)

	if err != nil {
		t.Errorf("Expected no error for valid IdTag, got: %v", err)
	}

	if got := idTag.String(); got != validStr {
		t.Errorf("Expected String() = %q, got %q", validStr, got)
	}

	if err := idTag.Validate(); err != nil {
		t.Errorf("Expected Validate() to succeed, got error: %v", err)
	}
}

func TestIdTagInvalidInputEmpty(t *testing.T) {
	t.Parallel()

	_, err := IdTag("")

	if err == nil {
		t.Error("Expected error for empty IdTag, got nil")
	}
}

func TestIdTagInvalidInputTooLong(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	_, err := IdTag(input)

	if err == nil {
		t.Errorf("Expected error for too long IdTag (%d chars), got nil", len(input))
	}
}

func TestIdTagInvalidInputNonASCII(t *testing.T) {
	t.Parallel()

	input := "ТестТаг123456" // Cyrillic
	_, err := IdTag(input)

	if err == nil {
		t.Error("Expected error for non-ASCII IdTag, got nil")
	}
}

func TestIdTagRevalidateAfterConstruction(t *testing.T) {
	t.Parallel()

	input := "12345678901234567890"
	idTag, err := IdTag(input)

	if err != nil {
		t.Fatalf("Unexpected error constructing IdTag: %v", err)
	}

	if err := idTag.Validate(); err != nil {
		t.Errorf("Expected Validate() to pass, got error: %v", err)
	}
}
