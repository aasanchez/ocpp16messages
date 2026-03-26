package metervalues

// ConfInput represents the raw input data for creating a MeterValues.conf
// message. The constructor Conf validates all fields automatically.
// MeterValues.conf is an empty message per OCPP 1.6 specification.
type ConfInput struct{}

// ConfMessage represents an OCPP 1.6 MeterValues.conf message.
// This is an empty message that acknowledges receipt of MeterValues.req.
type ConfMessage struct{}

// Conf creates a MeterValues.conf message from the given input.
// MeterValues.conf is an empty confirmation message per OCPP 1.6 specification.
func Conf(_ ConfInput) (ConfMessage, error) {
	return ConfMessage{}, nil
}
