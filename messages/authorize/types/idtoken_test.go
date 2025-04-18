package authorizetypes

import (
	"testing"
)

func Test_IdTokenFromString_Valid(t *testing.T) {
	t.Parallel()

	valid := "ABC1234567890123456" // 20 characters
	token, err := IdToken(valid)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := token.Validate(); err != nil {
		t.Errorf("unexpected validation failure: %v", err)
	}
}

func Test_IdTokenFromString_TooLong(t *testing.T) {
	t.Parallel()

	toolong := "ABC1234567890123456789" // > 20 characters
	_, err := IdToken(toolong)

	if err == nil {
		t.Errorf("expected error for over-length IdToken, got nil")
	}
}

func Test_IdTokenFromString_Empty(t *testing.T) {
	t.Parallel()

	_, err := IdToken("")
	if err == nil {
		t.Errorf("expected error for empty input, got nil")
	}
}
