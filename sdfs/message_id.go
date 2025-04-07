package types

import (
	"errors"
	"fmt"
	"strings"
)

const maxMessageIDLength = 36

// MessageID represents the UniqueId of an OCPP CALL message.
type MessageID struct {
	value string
}

// NewMessageID creates a new MessageID after validation.
func NewMessageID(messageId string) (MessageID, error) {
	messageId = strings.TrimSpace(messageId)
	if len(messageId) == 0 {
		return MessageID{}, errors.New("message ID cannot be empty")
	}
	if len(messageId) > maxMessageIDLength {
		return MessageID{}, fmt.Errorf("message ID cannot exceed %d characters", maxMessageIDLength)
	}
	return MessageID{value: messageId}, nil
}

// String returns the underlying string value.
func (m MessageID) String() string {
	return m.value
}
