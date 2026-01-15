package diagnosticsStatusNotification

// ConfInput represents the raw input data for creating a
// DiagnosticsStatusNotification.conf message.
// The constructor Conf validates all fields automatically.
// This message has no fields per OCPP 1.6 specification.
type ConfInput struct{}

// ConfMessage represents an OCPP 1.6 DiagnosticsStatusNotification.conf
// message. This message has no fields per OCPP 1.6 specification.
type ConfMessage struct{}

// Conf creates a DiagnosticsStatusNotification.conf message from the given
// input. This message has no fields, so it always succeeds.
func Conf(_ ConfInput) (ConfMessage, error) {
	return ConfMessage{}, nil
}
