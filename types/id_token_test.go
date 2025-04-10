package types

import (
	"testing"
)

func TestIdTokenFromString_Valid(t *testing.T) {
	t.Parallel()

	valid := "ABC1234567890123456" // 20 characters
	token, err := IdTokenFromString(valid)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if token.String() != valid {
		t.Errorf("expected %s, got %s", valid, token.String())
	}

	if err := token.Validate(); err != nil {
		t.Errorf("unexpected validation failure: %v", err)
	}
}

func TestIdTokenFromString_TooLong(t *testing.T) {
	t.Parallel()

	toolong := "ABC1234567890123456789" // > 20 characters
	_, err := IdTokenFromString(toolong)

	if err == nil {
		t.Errorf("expected error for over-length IdToken, got nil")
	}
}

func TestIdTokenFromString_NonASCII(t *testing.T) {
	t.Parallel()

	_, err := IdTokenFromString("你好世界")

	if err == nil {
		t.Errorf("expected error for non-ASCII input, got nil")
	}
}

func TestIdTokenFromString_Empty(t *testing.T) {
	t.Parallel()

	_, err := IdTokenFromString("")
	if err != nil {
		t.Errorf("expected no error for empty input, got %v", err)
	}
}
