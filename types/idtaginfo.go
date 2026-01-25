package types

import (
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*IdTagInfo)(nil)

// IdTagInfo contains status information about an identifier.
// It is returned in Authorize, StartTransaction, and StopTransaction responses.
type IdTagInfo struct {
	Status      AuthorizationStatus
	ExpiryDate  *DateTime // Optional
	ParentIdTag *IdToken  // Optional
}

// NewIdTagInfo creates a new IdTagInfo with the given status.
// ExpiryDate and ParentIdTag are optional and can be set separately.
func NewIdTagInfo(status AuthorizationStatus) (IdTagInfo, error) {
	if !status.IsValid() {
		return IdTagInfo{}, fmt.Errorf(
			"NewIdTagInfo: "+ErrorFieldFormat,
			"AuthorizationStatus",
			ErrInvalidValue,
		)
	}

	return IdTagInfo{
		Status:      status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}, nil
}

// WithExpiryDate sets the expiry date and returns the IdTagInfo.
func (i IdTagInfo) WithExpiryDate(expiryDate DateTime) IdTagInfo {
	i.ExpiryDate = &expiryDate

	return i
}

// WithParentIdTag sets the parent ID tag and returns the IdTagInfo.
func (i IdTagInfo) WithParentIdTag(parentIdTag IdToken) IdTagInfo {
	i.ParentIdTag = &parentIdTag

	return i
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the IdTagInfo for debugging purposes.
func (i IdTagInfo) String() string {
	result := "IdTagInfo{Status: " + i.Status.String()

	if i.ExpiryDate != nil {
		result += ", ExpiryDate: " + i.ExpiryDate.String()
	}

	if i.ParentIdTag != nil {
		result += ", ParentIdTag: " + i.ParentIdTag.String()
	}

	result += "}"

	return result
}
