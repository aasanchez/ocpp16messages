package clearCache

// ReqInput represents the raw input data for creating a ClearCache.req message.
// ClearCache.req has no fields per OCPP 1.6 specification.
type ReqInput struct{}

// ReqMessage represents an OCPP 1.6 ClearCache.req message.
// This message has no fields as per OCPP 1.6 specification.
type ReqMessage struct{}

// Req creates a ClearCache.req message from the given input.
// ClearCache.req has no fields, so this always succeeds.
func Req(_ ReqInput) (ReqMessage, error) {
	return ReqMessage{}, nil
}
