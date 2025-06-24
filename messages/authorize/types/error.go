package authorizetypes

import (
	"errors"
)

var (
	ErrInvalidIdToken             = errors.New("invalid IdToken")
	ErrInvalidIdTagInfo           = errors.New("invalid idTagInfo")
	ErrInvalidAuthorizationStatus = errors.New("invalid authorization status")
	ErrMissingStatus              = errors.New("missing required field: idTagInfo.Status")
	ErrInvalidExpiryDate          = errors.New("invalid expiryDate")
	ErrInvalidParentIdTag         = errors.New("invalid parentIdTag")
)
