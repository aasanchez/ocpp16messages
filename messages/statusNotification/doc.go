// Package statusNotification implements the OCPP 1.6 StatusNotification
// message.
//
// # Handling Rules
//
// The Charge Point SHALL send StatusNotification.req to the Central System
// to inform about status changes or errors within the Charge Point.
//
// # Status Changes
//
// The Charge Point MAY send a StatusNotification.req when transitioning
// from a previous status to a new status. Specific status transitions
// determine when notifications are appropriate.
//
// # Status Definitions
//
// The previous "Occupied" status from older OCPP versions is split into:
//   - Preparing
//   - Charging
//   - SuspendedEV
//   - SuspendedEVSE
//   - Finishing
//
// # Notes
//
// "EVSE" terminology is used in StatusNotification instead of "Socket" or
// "Charge Point" for forward compatibility with future OCPP versions.
package statusNotification
