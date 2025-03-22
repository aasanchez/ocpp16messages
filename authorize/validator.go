// Package authorize provides validation logic for Authorize.req and Authorize.conf messages.
package authorize

import (
	"fmt"
	"time"

	"github.com/aasanchez/ocpp16_messages/core"
)

// ValidateReq checks if the Authorize.req message complies with the OCPP 1.6J specification.
//
// This includes checking the validity of the IdTag field.
func ValidateReq(req Req) error {
	if err := req.IdTag.Validate(); err != nil {
		return core.NewFieldError("idTag", err)
	}
	return nil
}

// ValidateConf checks if the Authorize.conf message complies with the OCPP 1.6J specification.
//
// This includes validating the status, and optionally expiryDate and parentIdTag if present.
func ValidateConf(conf Conf) error {
	if !conf.IdTagInfo.Status.IsValid() {
		return core.NewFieldError("status", fmt.Errorf("invalid status: %s", conf.IdTagInfo.Status))
	}

	if conf.IdTagInfo.ExpiryDate != nil {
		if conf.IdTagInfo.ExpiryDate.Before(time.Now().Add(-365 * 24 * time.Hour)) {
			return core.NewFieldError("expiryDate", fmt.Errorf("date appears to be in the past"))
		}
	}

	if conf.IdTagInfo.ParentIdTag != nil {
		if err := conf.IdTagInfo.ParentIdTag.Validate(); err != nil {
			return core.NewFieldError("parentIdTag", err)
		}
	}

	return nil
}
