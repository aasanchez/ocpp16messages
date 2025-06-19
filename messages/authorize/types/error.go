package authorizetypes

import (
	"errors"
)

var (
	ErrInvalidAuthorizationStatus = errors.New("invalid authorization status")
	ErrMissingStatus              = errors.New("missing required field: idTagInfo.Status")
	ErrInvalidExpiryDate          = errors.New("invalid expiryDate")
	ErrInvalidParentIdTag         = errors.New("invalid parentIdTag")
)
