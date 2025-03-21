package validators

import (
	"fmt"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
)

func ValidateAuthorizeReq(req chargePoint.AuthorizeReq) error {
	if err := req.IdToken.Validate(); err != nil {
		return fmt.Errorf("IdToken: %w", err)
	}
	return nil
}

func ValidateAuthorizeConf(conf chargePoint.AuthorizeConf) error {
	if !conf.IdTagInfo.Status.IsValid() {
		return fmt.Errorf("invalid AuthorizationStatus: %s", conf.IdTagInfo.Status)
	}
	if conf.IdTagInfo.ParentIdTag != nil {
		if err := conf.IdTagInfo.ParentIdTag.Validate(); err != nil {
			return fmt.Errorf("ParentIdTag: %w", err)
		}
	}
	return nil
}
