package firmwareStatusNotification

// ConfInput represents the raw input data for creating a
// FirmwareStatusNotification.conf message.
// The constructor Conf validates all fields automatically.
// This message has no fields per OCPP 1.6 specification.
type ConfInput struct{}

// ConfMessage represents an OCPP 1.6 FirmwareStatusNotification.conf
// message. This message has no fields per OCPP 1.6 specification.
type ConfMessage struct{}

// Conf creates a FirmwareStatusNotification.conf message from the given
// input. This message has no fields, so it always succeeds.
func Conf(_ ConfInput) (ConfMessage, error) {
	return ConfMessage{}, nil
}
