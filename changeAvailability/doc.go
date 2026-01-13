// Package changeAvailability implements the OCPP 1.6 ChangeAvailability
// message.
//
// # Handling Rules
//
// The Central System MAY request a Charge Point to change its availability
// by sending a ChangeAvailability.req.
//
// A Charge Point is considered Available when it is charging or ready to
// charge, and Unavailable when it does not allow any charging.
//
// Upon receiving ChangeAvailability.req, the Charge Point SHALL respond
// with ChangeAvailability.conf indicating whether the requested change
// can be applied.
//
// If a transaction is in progress, the Charge Point SHALL respond with
// status Scheduled, indicating the change will occur after the
// transaction has finished.
//
// If the requested availability matches the current state, the Charge
// Point SHALL respond with status Accepted.
//
// Once the availability change has taken effect, the Charge Point SHALL
// notify the Central System by sending StatusNotification.req.
package changeAvailability
