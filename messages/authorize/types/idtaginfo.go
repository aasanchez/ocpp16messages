package authorizetypes

import (
	"errors"
	"fmt"
	"time"
)

var (
	errInvalidAuthorizationStatus = errors.New("invalid authorization status")
	errInvalidExpiryDate          = errors.New("expiryDate is present but not a valid timestamp")
	errInvalidParentIdTag         = errors.New("invalid parentIdTag")
)

type IdTagInfoType struct {
	ExpiryDate  *time.Time
	ParentIdTag *IdTokenType
	Status      AuthorizationStatus
}

func IdTagInfo(status AuthorizationStatus) (IdTagInfoType, error) {
	if !status.IsValid() {
		return IdTagInfoType{}, fmt.Errorf("%w: %s", errInvalidAuthorizationStatus, status)
	}

	return IdTagInfoType{
		Status:      status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}, nil
}

func (info IdTagInfoType) Validate() error {
	if !info.Status.IsValid() {
		return fmt.Errorf("%w: %s", errInvalidAuthorizationStatus, info.Status)
	}

	if info.ExpiryDate != nil && info.ExpiryDate.IsZero() {
		return errInvalidExpiryDate
	}

	if info.ParentIdTag != nil {
		if err := info.ParentIdTag.Validate(); err != nil {
			return fmt.Errorf("%w: %w", errInvalidParentIdTag, err)
		}
	}

	return nil
}

func (info IdTagInfoType) String() string {
	str := "{status:" + string(info.Status)

	if info.ExpiryDate != nil {
		str += ", expiryDate:" + info.ExpiryDate.Format(time.RFC3339)
	}

	if info.ParentIdTag != nil {
		str += ", parentIdTag:" + info.ParentIdTag.String()
	}

	str += "}"

	return str
}
