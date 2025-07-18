package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationPayload struct {
	IdTagInfo IdTagInfoPayload
}

func (conf ConfirmationPayload) Validate() error {
	if conf.IdTagInfo.Status == "" {
		return fmt.Errorf(st.ErrorMissing, "IdTagInfo")
	}

	return nil
}
