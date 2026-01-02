package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdTagInfo contains status information about an identifier.
// It is returned in Authorize, StartTransaction, and StopTransaction responses.
type IdTagInfo struct {
	Status      AuthorizationStatus
	ExpiryDate  *st.DateTime // Optional
	ParentIdTag *IdToken     // Optional
}

// NewIdTagInfo creates a new IdTagInfo with the given status.
// ExpiryDate and ParentIdTag are optional and can be set separately.
func NewIdTagInfo(status AuthorizationStatus) (IdTagInfo, error) {
	if !status.IsValid() {
		return IdTagInfo{}, fmt.Errorf(
			"NewIdTagInfo: %w",
			st.ErrInvalid("AuthorizationStatus"),
		)
	}

	return IdTagInfo{
		Status:      status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}, nil
}

// WithExpiryDate sets the expiry date and returns the IdTagInfo.
func (i IdTagInfo) WithExpiryDate(expiryDate st.DateTime) IdTagInfo {
	i.ExpiryDate = &expiryDate
	return i
}

// WithParentIdTag sets the parent ID tag and returns the IdTagInfo.
func (i IdTagInfo) WithParentIdTag(parentIdTag IdToken) IdTagInfo {
	i.ParentIdTag = &parentIdTag
	return i
}
