package types

import (
	"fmt"
	"testing"
)

func TestIdTag_ValidInput(t *testing.T) {
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

func TestIdTag_InvalidInput_Empty(t *testing.T) {
	_, err := IdTag("")
	fmt.Println(err)
	if err == nil {
		t.Error("Expected error for empty IdTag, got nil")
	}
}

func TestIdTag_InvalidInput_TooLong(t *testing.T) {
	// More than 20 characters
	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	_, err := IdTag(input)
	if err == nil {
		t.Errorf("Expected error for too long IdTag (%d chars), got nil", len(input))
	}
}

func TestIdTag_InvalidInput_NonASCII(t *testing.T) {
	input := "ТестТаг123456" // Cyrillic characters
	_, err := IdTag(input)
	if err == nil {
		t.Error("Expected error for non-ASCII IdTag, got nil")
	}
}

func TestIdTag_RevalidateAfterConstruction(t *testing.T) {
	input := "12345678901234567890"
	idTag, err := IdTag(input)
	if err != nil {
		t.Fatalf("Unexpected error constructing IdTag: %v", err)
	}
	if err := idTag.Validate(); err != nil {
		t.Errorf("Expected Validate() to pass, got error: %v", err)
	}
}
