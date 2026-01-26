package startTransaction

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a StartTransaction.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: The transaction identifier assigned by the Central System.
	TransactionId int
	// Required: AuthorizationStatus value.
	Status string
	// Optional: RFC3339 formatted expiry date for the authorization.
	ExpiryDate *string
	// Optional: Parent ID tag (max 20 chars).
	ParentIdTag *string
}

// ConfMessage represents an OCPP 1.6 StartTransaction.conf message.
type ConfMessage struct {
	TransactionId st.Integer
	IdTagInfo     st.IdTagInfo
}

// confValidation holds validated fields during Conf construction.
type confValidation struct {
	transactionId st.Integer
	idTagInfo     st.IdTagInfo
	expiryDate    st.DateTime
	parentIdToken st.IdToken
}

// Conf creates a StartTransaction.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - TransactionId is negative or exceeds uint16 max value (65535)
//   - Status is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIdTag (if provided) exceeds 20 characters or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	validated, errs := validateConfInput(input)

	if len(errs) > errCountZero {
		return ConfMessage{}, errors.Join(errs...)
	}

	return buildConfMessage(input, validated), nil
}

// validateConfInput validates all fields in ConfInput and returns validated
// values along with any errors.
func validateConfInput(input ConfInput) (confValidation, []error) {
	var errs []error

	var validated confValidation

	// Validate transactionId (required)
	validated.transactionId, errs = validateTransactionId(
		input.TransactionId,
		errs,
	)

	// Validate status (required)
	validated.idTagInfo, errs = validateStatus(input.Status, errs)

	// Validate expiryDate (optional)
	if input.ExpiryDate != nil {
		validated.expiryDate, errs = validateExpiryDate(*input.ExpiryDate, errs)
	}

	// Validate parentIdTag (optional)
	if input.ParentIdTag != nil {
		validated.parentIdToken, errs = validateParentIdTag(
			*input.ParentIdTag,
			errs,
		)
	}

	return validated, errs
}

// validateTransactionId validates the transactionId field.
func validateTransactionId(
	transactionId int,
	errs []error,
) (st.Integer, []error) {
	val, err := st.NewInteger(transactionId)
	if err != nil {
		return st.Integer{}, append(errs, fmt.Errorf("transactionId: %w", err))
	}

	return val, errs
}

// validateStatus validates the status field and returns the IdTagInfo.
func validateStatus(status string, errs []error) (st.IdTagInfo, []error) {
	info, err := st.NewIdTagInfo(st.AuthorizationStatus(status))
	if err != nil {
		return st.IdTagInfo{
			Status:      "",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		}, append(errs, fmt.Errorf("status: %w", err))
	}

	return info, errs
}

// validateExpiryDate validates the expiry date string and returns the DateTime.
func validateExpiryDate(date string, errs []error) (st.DateTime, []error) {
	expiryDate, err := st.NewDateTime(date)
	if err != nil {
		return st.DateTime{}, append(errs, fmt.Errorf("expiryDate: %w", err))
	}

	return expiryDate, errs
}

// validateParentIdTag validates the parent ID tag and returns the token.
func validateParentIdTag(tag string, errs []error) (st.IdToken, []error) {
	ciStr, err := st.NewCiString20Type(tag)
	if err != nil {
		return st.IdToken{}, append(errs, fmt.Errorf("parentIdTag: %w", err))
	}

	return st.NewIdToken(ciStr), errs
}

// buildConfMessage constructs the final ConfMessage with validated fields.
func buildConfMessage(input ConfInput, validated confValidation) ConfMessage {
	idTagInfo := validated.idTagInfo

	if input.ExpiryDate != nil {
		idTagInfo = idTagInfo.WithExpiryDate(validated.expiryDate)
	}

	if input.ParentIdTag != nil {
		idTagInfo = idTagInfo.WithParentIdTag(validated.parentIdToken)
	}

	return ConfMessage{
		TransactionId: validated.transactionId,
		IdTagInfo:     idTagInfo,
	}
}
