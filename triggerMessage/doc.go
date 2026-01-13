// Package triggerMessage implements the OCPP 1.6 TriggerMessage message.
//
// # Handling Rules
//
// The Central System MAY request a Charge Point to send a Charge
// Point-initiated message using TriggerMessage.req. The request specifies
// which message the Central System wishes to receive and optionally the
// connectorId to which the request applies.
//
// # ConnectorId Rules
//
//   - If the connectorId is irrelevant for the requested message, it SHALL
//     be ignored.
//   - If the connectorId is relevant but absent, the request is interpreted
//     as "for all allowed connectorId values".
//   - Examples:
//   - Triggering a StatusNotification for connectorId 0 → request for
//     overall Charge Point status.
//   - Triggering a StatusNotification without connectorId → request for
//     status of Charge Point and all connectors.
//
// # Charge Point Response
//
//   - The Charge Point SHALL first respond with TriggerMessage.conf before
//     sending the requested message.
//   - TriggerMessage.conf SHALL indicate ACCEPTED, REJECTED, or
//     NOT_IMPLEMENTED.
//   - ACCEPTED → Charge Point SHOULD send the requested message.
//   - REJECTED → Charge Point chooses not to send the message.
//   - NOT_IMPLEMENTED → message is unknown or not supported.
//
// # Message Sending Rules
//
//   - Messages accepted in TriggerMessage SHOULD be sent.
//   - If the same message is triggered by normal operation before being
//     sent, it MAY count as satisfying the request.
//   - TriggerMessage is intended for current information only; it is NOT
//     for retrieving historic data.
//   - Example: MeterValues.req triggered this way SHALL return the most
//     recent measurements as per MeterValuesSampledData configuration.
//
// # Exclusions
//
//   - StartTransaction and StopTransaction are not supported via
//     TriggerMessage because they represent state transitions rather than
//     current state.
package triggerMessage
