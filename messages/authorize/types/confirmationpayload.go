// Package authorizetypes provides types for the Authorize message.
package authorizetypes

// import (
// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// // ConfirmationPayload is the payload for the Authorize.conf message.
// // It contains the IdTagInfo which provides authorization details.
// // See OCPP 1.6J Part 2, Section 4.1.2.2, Table 13 for details.
// type ConfirmationPayload struct {
// 	IdTagInfo IdTagInfoPayload // IdTagInfo contains the authorization status and other relevant information. (Required)
// }

// // Validate checks that the ConfirmationPayload contains the required fields.
// // It returns an error if any mandatory field is missing or invalid.
// func (c ConfirmationPayload) Validate() error {
// 	if c.IdTagInfo.Status == "" {
// 		return fmt.Errorf(st.ErrorInvalid, "IdTagInfo", err)
// 	}

// 	return nil
// }

// // Value returns the ConfirmationPayload itself.
// func (c ConfirmationPayload) Value() ConfirmationPayload {
// 	return c
// }
