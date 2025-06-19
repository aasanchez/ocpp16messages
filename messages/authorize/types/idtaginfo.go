package authorizetypes

import (
	"errors"
	"fmt"

	sharedTypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var (
	errInvalidAuthorizationStatus = errors.New("invalid authorization status")
	errInvalidExpiryDate          = errors.New("invalid expiryDate")
	errInvalidParentIdTag         = errors.New("invalid parentIdTag")
)

type IdTagInfo struct {
	ExpiryDate  *sharedTypes.DateTimeType
	ParentIdTag *IdToken
	Status      AuthorizationStatusType
}

func NewIdTagInfo(rawStatus string) (IdTagInfo, error) {
	status, err := AuthorizationStatus(rawStatus)
	if err != nil {
		return IdTagInfo{}, fmt.Errorf("%w: %w", errInvalidAuthorizationStatus, err)
	}

	return IdTagInfo{
		Status:      status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}, nil
}

func (info IdTagInfo) Validate() error {
	if err := info.Status.Validate(); err != nil {
		return fmt.Errorf("%w: %w", errInvalidAuthorizationStatus, err)
	}

	if info.ExpiryDate != nil {
		if err := info.ExpiryDate.Validate(); err != nil {
			return fmt.Errorf("%w: %w", errInvalidExpiryDate, err)
		}
	}

	if info.ParentIdTag != nil {
		if err := info.ParentIdTag.Validate(); err != nil {
			return fmt.Errorf("%w: %w", errInvalidParentIdTag, err)
		}
	}

	return nil
}
