package heartbeat

// ReqInput represents the raw input data for creating a Heartbeat.req message.
// Heartbeat.req has no fields per OCPP 1.6 specification.
type ReqInput struct{}

// ReqMessage represents an OCPP 1.6 Heartbeat.req message.
// This message has no fields - it is used solely to signal that the
// Charge Point is still connected to the Central System.
type ReqMessage struct{}

// Req creates a Heartbeat.req message from the given input.
// Since Heartbeat.req has no fields, this function always succeeds.
func Req(_ ReqInput) (ReqMessage, error) {
	return ReqMessage{}, nil
}
