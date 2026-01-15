// Package dataTransfer implements the OCPP 1.6 DataTransfer message.
//
// # Handling Rules
//
// When a Charge Point needs to exchange information with the Central System
// for functionality not covered by OCPP, it SHALL use DataTransfer.req.
//
// The vendorId SHOULD uniquely identify the vendor-specific implementation
// and SHOULD be known to the Central System. It is RECOMMENDED to use a value
// from the reversed DNS namespace corresponding to the vendor’s registered
// primary domain.
//
// The optional messageId MAY be used to identify a specific message or
// implementation.
//
// The length of the data field in both request and response is undefined and
// MUST be agreed upon by the involved parties.
//
// If the recipient does not support the given vendorId, it SHALL return
// status UnknownVendor and SHALL NOT include the data field.
// If a messageId is provided and does not match, the recipient SHALL return
// status UnknownMessageId.
//
// In all other cases, the meaning of status Accepted or Rejected and the usage
// of the data field are defined by vendor-specific agreement.
package dataTransfer
