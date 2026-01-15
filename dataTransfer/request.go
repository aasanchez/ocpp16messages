package dataTransfer

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a DataTransfer.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	VendorId  string  // Required: Vendor identifier (max 255 chars)
	MessageId *string // Optional: Message identifier (max 50 chars)
	Data      *string // Optional: Data payload (unbounded per OCPP spec)
}

// ReqMessage represents an OCPP 1.6 DataTransfer.req message.
type ReqMessage struct {
	VendorId  st.CiString255Type
	MessageId *st.CiString50Type
	Data      *string
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	vendorId  st.CiString255Type
	messageId st.CiString50Type
}

// Req creates a DataTransfer.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - VendorId is empty
//   - VendorId exceeds 255 characters or contains invalid chars
//   - MessageId (if provided) exceeds 50 characters or contains invalid chars
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{}, errors.Join(errs...)
	}

	return buildReqMessage(input, validated), nil
}

// validateReqInput validates all fields in ReqInput and returns validated
// values along with any errors.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	// Validate vendorId (required)
	vendorId, err := st.NewCiString255Type(input.VendorId)
	if err != nil {
		errs = append(errs, fmt.Errorf("vendorId: %w", err))
	} else {
		validated.vendorId = vendorId
	}

	// Validate messageId (optional)
	if input.MessageId != nil {
		messageId, err := st.NewCiString50Type(*input.MessageId)
		if err != nil {
			errs = append(errs, fmt.Errorf("messageId: %w", err))
		} else {
			validated.messageId = messageId
		}
	}

	return validated, errs
}

// buildReqMessage constructs the final ReqMessage with validated fields.
func buildReqMessage(input ReqInput, validated reqValidation) ReqMessage {
	msg := ReqMessage{
		VendorId:  validated.vendorId,
		MessageId: nil,
		Data:      nil,
	}

	if input.MessageId != nil {
		msg.MessageId = &validated.messageId
	}

	if input.Data != nil {
		msg.Data = input.Data
	}

	return msg
}
