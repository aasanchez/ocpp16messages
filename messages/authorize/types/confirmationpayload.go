package authorizetypes

// import (
// 	"errors"
// 	"fmt"

// 	st "github.com/aasanchez/ocpp16messages/shared/types"
// )

// var errMissingIdTagInfo = errors.New("error: Missing IdTagInfo")

// type ConfirmationPayload struct {
// 	IdTagInfo IdTagInfoPayload
// }

// func (conf ConfirmationPayload) Validate() error {
// 	if conf.IdTagInfo.Status == "" {
// 		return fmt.Errorf(st.ErrorWrapper, errMissingIdTagInfo, conf)
// 	}

// 	return nil
// }
