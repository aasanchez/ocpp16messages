package updatefirmware

// ConfInput represents the raw input data for creating an UpdateFirmware.conf
// message. This message has no fields per OCPP 1.6 specification.
type ConfInput struct{}

// ConfMessage represents an OCPP 1.6 UpdateFirmware.conf message.
// This is an empty acknowledgment message with no fields.
type ConfMessage struct{}

// Conf creates an UpdateFirmware.conf message from the given input.
// Per OCPP 1.6 specification, UpdateFirmware.conf has no fields.
// This function always succeeds and returns an empty ConfMessage.
func Conf(_ ConfInput) (ConfMessage, error) {
	return ConfMessage{}, nil
}
