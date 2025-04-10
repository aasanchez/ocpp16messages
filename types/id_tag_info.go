package types

import (
	"errors"
	"fmt"
	"time"
)

// Static error definitions for validation.
var (
	ErrInvalidAuthorizationStatus = errors.New("invalid authorization status")
	ErrInvalidExpiryDate          = errors.New("expiryDate is present but not a valid timestamp")
	ErrInvalidParentIdTag         = errors.New("invalid parentIdTag")
)

// IdTagInfoType represents the authorization status and related metadata returned by the Central System
// in response to an Authorize.req message.
//
// This type aligns with the `idTagInfo` structure defined in the OCPP 1.6J specification, Section 5.2.
type IdTagInfoType struct {
	Status      AuthorizationStatus
	ExpiryDate  *time.Time
	ParentIdTag *IdTokenType
}

// NewIdTagInfo constructs a new IdTagInfoType with a validated AuthorizationStatus.
func IdTagInfo(status AuthorizationStatus) (IdTagInfoType, error) {
	if !status.IsValid() {
		return IdTagInfoType{}, fmt.Errorf("%w: %s", ErrInvalidAuthorizationStatus, status)
	}

	return IdTagInfoType{
		Status:      status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}, nil
}

// Validate checks the internal consistency of the IdTagInfoType struct.
func (info IdTagInfoType) Validate() error {
	if !info.Status.IsValid() {
		return fmt.Errorf("%w: %s", ErrInvalidAuthorizationStatus, info.Status)
	}

	if info.ExpiryDate != nil && info.ExpiryDate.IsZero() {
		return ErrInvalidExpiryDate
	}

	if info.ParentIdTag != nil {
		if err := info.ParentIdTag.Validate(); err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidParentIdTag, err)
		}
	}

	return nil
}

// String returns a human-readable representation of the IdTagInfoType.
func (info IdTagInfoType) String() string {
	str := "{status=" + string(info.Status)

	if info.ExpiryDate != nil {
		str += ", expiryDate=" + info.ExpiryDate.Format(time.RFC3339)
	}

	if info.ParentIdTag != nil {
		str += ", parentIdTag=" + info.ParentIdTag.String()
	}

	str += "}"

	return str
}
