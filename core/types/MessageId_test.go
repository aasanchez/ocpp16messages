package types

import "testing"

func TestNewMecdssageIDValid(t *testing.T) {
	validIDs := []string{
		"1234567890",
		"550e8400-e29b-41d4-a716-446655440000", // valid UUID
		"simple-id",
		"abcdefghijklmnopqrstuvwxyz0123456789", // 36 chars
	}
	for _, id := range validIDs {
		mid, err := NewMessageID(id)
		if err != nil {
			t.Errorf("expected valid MessageID for input %q, got error: %v", id, err)
		}
		if mid.String() != id {
			t.Errorf("expected %q, got %q", id, mid.String())
		}
	}
}

func TestNewMessageIDEmpty(t *testing.T) {
	_, err := NewMessageID("")
	if err == nil {
		t.Error("expected error for empty MessageID, got nil")
	}
}

func TestNewMessageIDTooLong(t *testing.T) {
	tooLong := "abcdefghijklmnopqrstuvwxyz01234567890" // 37 chars
	_, err := NewMessageID(tooLong)
	if err == nil {
		t.Error("expected error for MessageID longer than 36 characters, got nil")
	}
}
