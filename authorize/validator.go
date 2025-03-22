// Package authorize provides validation logic for Authorize.req and Authorize.conf
// messages in the OCPP 1.6J protocol.
//
// This file contains validation functions that enforce schema correctness
// as defined by the OCPP specification.
package authorize

import (
	"github.com/aasanchez/ocpp16_messages/core"
)

// ValidateAuthorizeReq validates the AuthorizeReq message.
// It checks that the idTag field conforms to the CiString20Type constraint.
func ValidateAuthorizeReq(req AuthorizeReq) error {
	tag := core.CiString20Type(req.IdTag)
	if err := tag.Validate(); err != nil {
		return core.NewFieldError("idTag", err.Error())
	}
	return nil
}

// ValidateAuthorizeConf validates the AuthorizeConf message.
// It checks that the status is valid and any optional fields are well-formed.
func ValidateAuthorizeConf(conf AuthorizeConf) error {
	if !conf.IdTagInfo.Status.IsValid() {
		return core.NewFieldError("status", "must be a valid AuthorizationStatus")
	}
	if conf.IdTagInfo.ParentIdTag != nil {
		if err := conf.IdTagInfo.ParentIdTag.Validate(); err != nil {
			return core.NewFieldError("parentIdTag", err.Error())
		}
	}
	return nil
}
