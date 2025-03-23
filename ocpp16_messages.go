// Package ocpp16_messages provides decoding and validation capabilities
// for OCPP 1.6J and 1.6S messages using JSON and SOAP formats.
//
// It supports strict type validation for all official message types,
// and includes plugin-based extensibility to allow custom validators.
//
// Usage examples are available under the `example/authorization/json`
// and `example/authorization/soap` directories.
package ocpp16_messages

// "github.com/aasanchez/ocpp16_messages/authorize"
// "github.com/aasanchez/ocpp16_messages/core"

// RegisterBuiltInValidators registers all official OCPP 1.6J/S validators.
//
// This function should be called before using any message validation logic
// to ensure required message types are handled properly.
// func RegisterBuiltInValidators() {
// 	core.RegisterValidator("AuthorizeReq", authorize.ValidateAuthorizeReq)
// 	core.RegisterValidator("AuthorizeConf", authorize.ValidateAuthorizeConf)
// }

// ValidateRawJSON validates an incoming raw JSON OCPP message.
//
// This function parses the input as a JSON array of the form:
//
//	[2, "uid", "Action", {payload}]
//	[3, "uid", {payload}]
//	[4, "uid", "errorCode", "errorDescription", {details}]
//
// It returns the parsed message or a wrapped error in case of failure.

// func ValidateRawJSON(raw []byte) (*core.Message, error) {
// 	msg, err := core.ParseMessage(raw)
// 	if err != nil {
// 		return nil, err
// 	}

// 	switch msg.TypeID {
// 	case core.CALL, core.CALLRESULT:
// 		// Lookup validator by Action
// 		validator := core.GetRegisteredValidator(msg.Action)
// 		if validator == nil {
// 			return nil, core.NewFieldError(fmt.Errorf("no validator registered for action: %s", msg.Action), "action")
// 		}

// 		// Validate payload
// 		payload, err := validator.ValidateMessage(msg.Payload)
// 		if err != nil {
// 			return nil, err
// 		}
// 		msg.Payload = payload
// 	}

// 	return msg, nil
// }

// ValidateRawSOAP validates an incoming raw SOAP OCPP message.
//
// It performs minimal SOAP envelope parsing and validates the body
// using the corresponding registered message validator.
//
// It assumes the structure of the SOAP body corresponds to an OCPP message.
// func ValidateRawSOAP(raw []byte) (*core.Message, error) {
// 	msg, err := core.ParseSOAPMessage(raw)
// 	if err != nil {
// 		return nil, err
// 	}

// 	switch msg.TypeID {
// 	case core.CALL, core.CALLRESULT:
// 		validator := core.GetRegisteredValidator(msg.Action)
// 		if validator == nil {
// 			return nil, core.NewFieldError(fmt.Errorf("no validator registered for action: %s", msg.Action), "action")
// 		}

// 		payload, err := validator.ValidateMessage(msg.Payload)
// 		if err != nil {
// 			return nil, err
// 		}
// 		msg.Payload = payload
// 	}

// 	return msg, nil
// }
