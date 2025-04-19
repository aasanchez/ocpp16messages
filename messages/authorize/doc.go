// Package authorize defines the Authorize.req and Authorize.conf messages used in OCPP 1.6.
//
// Authorization is required before a Charge Point can start or stop a charging
// session. Typically, the Charge Point sends an Authorize.req message to the
// Central System to verify whether a given idTag (user identifier) is allowed
// to proceed.
//
// If the idTag is found in the Local Authorization List or Authorization Cache,
// the Charge Point may authorize the user without consulting the Central
// System. Otherwise, it must send an Authorize.req request to obtain permission.
//
// The Central System responds with an Authorize.conf message, which includes
// the authorization status. If accepted, it may also provide a parentIdTag and
// optional expiry information. This confirmation may also be used by the
// Charge Point to update its local cache.
//
// Note: When stopping a transaction, a second authorization is only required
// if the idTag used to stop is different from the one that started the session.
//
// This package provides the message structures, constructors, and validation
// logic needed to correctly implement this authorization flow.
//
// This package should be imported using:
//
// import "github.com/aasanchez/ocpp16messages/types/authorize"
package authorize
