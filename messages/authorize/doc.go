// Package authorize provides OCPP 1.6 Authorize message implementation.
//
// Authorization is required before a Charge Point may start or stop a charging
// transaction. Energy shall only be supplied after successful authorization.
//
// The Authorize.req message is used to validate an identifier (idTag) for
// charging. When stopping a transaction, Authorize.req shall only be sent if
// the identifier differs from the one used to start the transaction.
//
// A Charge Point may authorize an idTag locally using the Local Authorization
// List or Authorization Cache. If the idTag is not available locally, the
// Charge Point shall send an Authorize.req to the Central System. If the idTag
// is found locally, sending Authorize.req is optional.
//
// The Central System shall respond to Authorize.req with an Authorize.conf
// indicating whether the idTag is accepted or rejected. An accepted response
// must include an authorization status and may include a parentIdTag.
package authorize
