package authorize

import (
	"errors"
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfInput represents the raw input data for creating an Authorize.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	Status      string  // Required: AuthorizationStatus value
	ExpiryDate  *string // Optional: RFC3339 formatted expiry date
	ParentIdTag *string // Optional: Parent ID tag (max 20 chars)
}

// ConfMessage represents an OCPP 1.6 Authorize.conf message.
type ConfMessage struct {
	IdTagInfo mat.IdTagInfo
}

// confValidation holds validated fields during Conf construction.
type confValidation struct {
	idTagInfo     mat.IdTagInfo
	expiryDate    st.DateTime
	parentIdToken mat.IdToken
}

// Conf creates an Authorize.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - Status is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIdTag (if provided) exceeds 20 characters or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	validated, errs := validateConfInput(input)

	if errs != nil {
		return ConfMessage{
			IdTagInfo: mat.IdTagInfo{
				Status:      "",
				ExpiryDate:  nil,
				ParentIdTag: nil,
			},
		}, errors.Join(errs...)
	}

	return buildConfMessage(input, validated), nil
}

// validateConfInput validates all fields in ConfInput and returns validated
// values along with any errors.
func validateConfInput(input ConfInput) (confValidation, []error) {
	var errs []error

	var validated confValidation

	// Validate status (required)
	validated.idTagInfo, errs = validateStatus(input.Status, errs)

	// Validate expiryDate (optional)
	if input.ExpiryDate != nil {
		validated.expiryDate, errs = validateExpiryDate(*input.ExpiryDate, errs)
	}

	// Validate parentIdTag (optional)
	if input.ParentIdTag != nil {
		tag := *input.ParentIdTag
		validated.parentIdToken, errs = validateParentIdTag(tag, errs)
	}

	return validated, errs
}

// validateStatus validates the status field and returns the IdTagInfo.
func validateStatus(status string, errs []error) (mat.IdTagInfo, []error) {
	info, err := mat.NewIdTagInfo(mat.AuthorizationStatus(status))
	if err != nil {
		return mat.IdTagInfo{
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

// buildConfMessage constructs the final ConfMessage with validated fields.
func buildConfMessage(input ConfInput, validated confValidation) ConfMessage {
	idTagInfo := validated.idTagInfo

	if input.ExpiryDate != nil {
		idTagInfo = idTagInfo.WithExpiryDate(validated.expiryDate)
	}

	if input.ParentIdTag != nil {
		idTagInfo = idTagInfo.WithParentIdTag(validated.parentIdToken)
	}

	return ConfMessage{IdTagInfo: idTagInfo}
}

// validateParentIdTag validates the parent ID tag and returns the token.
func validateParentIdTag(tag string, errs []error) (mat.IdToken, []error) {
	ciStr, err := st.NewCiString20Type(tag)
	if err != nil {
		return mat.IdToken{}, append(errs, fmt.Errorf("parentIdTag: %w", err))
	}

	token, err := mat.NewIdToken(ciStr)
	if err != nil {
		return mat.IdToken{}, append(errs, fmt.Errorf("parentIdTag: %w", err))
	}

	return token, errs
}
