// If the transaction is being stopped using a different identifier than the one
// used to start it, the Charge Point must send a new Authorize.req request.
//
// The Authorize.req message is designed exclusively for verifying whether a
// specific identifier is allowed to initiate charging.
//
//  Charge Points may also perform local authorization, bypassing the Central
// System, by referencing a locally stored list of valid identifiers
// (Local Authorization List). If the presented identifier (idTag) is not found
// in this list or in the Authorization Cache, the Charge Point must forward an
// Authorize.req message to the Central System for validation.
//
// Even if the identifier exists in the local list or cache, the Charge Point
// may still choose to send the request to the Central System.
//
// Upon receiving an Authorize.req, the Central System must respond with an
// Authorize.conf. This response indicates whether the idTag is valid. If the
// identifier is accepted, the response must include an authorization status
// and may optionally contain a parentIdTag.
//
// If the Charge Point maintains an Authorization Cache and the idTag is not
// part of the local list, then upon receiving an Authorize.conf, it must update
// the cache entry using the details provided in the response.

package bootnotification
