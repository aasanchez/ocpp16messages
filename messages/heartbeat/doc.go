// Heartbeat handling rules:
//
// A Charge Point SHALL send Heartbeat.req messages at a configurable interval
// to inform the Central System that it is still connected.
//
// Upon receipt, the Central System SHALL respond with Heartbeat.conf containing
// its current time. The Charge Point is RECOMMENDED to use this time to
// synchronize its internal clock.
//
// A Charge Point MAY skip sending Heartbeat.req if another PDU has been sent
// within the configured heartbeat interval. The Central System SHOULD assume
// the Charge Point is available whenever any PDU is received.
//
// When using JSON over WebSocket, heartbeats are not mandatory, but at least
// one heartbeat per 24 hours is advised for time synchronization.
package heartbeat
