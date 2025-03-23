package authorize

import (
	"encoding/json"
	"fmt"

	"github.com/aasanchez/ocpp16_messages/core"
)

// AuthorizeReq represents an Authorize request message.
type AuthorizeReq struct {
	IdTag core.IdToken `json:"idTag"`
}

// validateAuthorizeReq performs structural validation of the AuthorizeReq message.
func validateAuthorizeReq(raw json.RawMessage) (interface{}, error) {
	var req AuthorizeReq
	if err := json.Unmarshal(raw, &req); err != nil {
		return nil, core.NewFieldError("idTag", fmt.Errorf("invalid JSON structure: %w", err))
	}

	if req.IdTag.IdTag == "" {
		return nil, core.NewFieldError("idTag", fmt.Errorf("required"))
	}

	if err := req.IdTag.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}

// validateAuthorizeConf performs structural validation of the AuthorizeConf message.
func validateAuthorizeConf(raw json.RawMessage) (interface{}, error) {
	var conf Conf
	if err := json.Unmarshal(raw, &conf); err != nil {
		return nil, core.NewFieldError("idTagInfo", fmt.Errorf("invalid JSON structure: %w", err))
	}

	if conf.IdTagInfo.Status == "" {
		return nil, core.NewFieldError("idTagInfo.status", fmt.Errorf("required"))
	}

	if !conf.IdTagInfo.Status.IsValid() {
		return nil, core.NewFieldError("idTagInfo.status", fmt.Errorf("invalid status value"))
	}

	if conf.IdTagInfo.ParentIdTag != nil {
		if err := (*conf.IdTagInfo.ParentIdTag).Validate(); err != nil {
			return nil, core.NewFieldError("idTagInfo.parentIdTag", err)
		}
	}

	return conf, nil
}

// init registers the validators for AuthorizeReq and AuthorizeConf to the core registry.
func init() {
	core.RegisterValidator("Authorize", validateAuthorizeReq)
	core.RegisterValidator("AuthorizeResponse", validateAuthorizeConf)
}
