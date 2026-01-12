// Package bootNotification implements the OCPP 1.6 BootNotification message.
//
// # Handling Rules
//
// After sending a BootNotification.req, a Charge Point MUST NOT send any
// other requests to the Central System, including previously cached messages.
//
// If the Central System responds with status Accepted, the Charge Point
// SHALL update its heartbeat interval according to the response and SHOULD
// synchronize its internal clock with the Central System time.
//
// If the response status is not Accepted, the interval field defines the
// minimum wait time before retrying BootNotification.req. If the interval
// is zero, the Charge Point MUST choose a retry delay that avoids flooding
// the Central System. A new BootNotification.req MUST NOT be sent earlier
// unless explicitly triggered via TriggerMessage.req.
//
// If the status is Rejected, the Charge Point SHALL NOT send any OCPP
// messages until the retry interval expires. During this time, communication
// may be closed by either side, and the Charge Point SHALL NOT respond to
// Central System initiated messages. The Central System SHOULD NOT initiate
// any messages.
//
// If the status is Pending, the communication channel SHOULD remain open.
// The Central System MAY send requests to retrieve information or change
// configuration, and the Charge Point SHOULD respond. While Pending, the
// Charge Point SHALL NOT initiate requests unless triggered by
// TriggerMessage.req.
//
// While in Pending state, RemoteStartTransaction.req and
// RemoteStopTransaction.req are not allowed.
package bootNotification
