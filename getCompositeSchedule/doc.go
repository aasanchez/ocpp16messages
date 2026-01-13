// Package getCompositeSchedule implements the OCPP 1.6 GetCompositeSchedule
// message.
//
// # Handling Rules
//
// The Central System MAY request a Composite Charging Schedule by sending
// GetCompositeSchedule.req.
//
// The Charge Point SHALL calculate and return the Composite Charging
// Schedule in GetCompositeSchedule.conf, based on all active charging
// schedules and applicable local limits.
//
// The reported schedule SHALL cover the interval from the time the request
// is received (X) up to X + Duration.
//
// If ConnectorId is set to 0, the Charge Point SHALL report the total
// expected power or current consumption of the entire Charge Point for
// the requested period.
//
// The returned schedule is indicative at the time of reporting and MAY
// change due to external factors such as local load balancing.
//
// If the Charge Point cannot report the requested schedule (e.g. unknown
// ConnectorId), it SHALL respond with status Rejected.
package getCompositeSchedule
