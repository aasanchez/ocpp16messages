// Package clearcache implements the Open Charge Point Protocol (OCPP) 1.6
// ClearCache message for EV charging.
//
// # Handling Rules
//
// The Central System MAY request a Charge Point to clear its Authorization
// Cache by sending a ClearCache.req.
//
// Upon receipt of ClearCache.req, the Charge Point SHALL respond with
// ClearCache.conf indicating whether the Authorization Cache was
// successfully cleared.
package clearcache
