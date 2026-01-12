// ReserveNow handling rules:
//
// The Central System MAY request a Charge Point to reserve a connector for a specific idTag
// by sending ReserveNow.req. The Charge Point SHALL respond with ReserveNow.conf.
//
// Behavior and responses:
//   - If reservationId matches an existing reservation, the Charge Point SHALL replace it.
//   - If reservationId is new, the Charge Point SHALL respond with:
//   - Accepted: Reservation succeeded.
//   - Occupied: Connector or Charge Point is in use or already reserved.
//   - Faulted: Connector or Charge Point is in Faulted state.
//   - Unavailable: Connector or Charge Point is Unavailable.
//   - Rejected: Charge Point is configured not to accept reservations.
//
// Reserved connector behavior:
//   - Charge Point SHALL refuse charging for all incoming idTags on the reserved connector,
//     except when the idTag or parent idTag matches the reservation.
//   - If ReserveConnectorZeroSupported = true and connectorId = 0, the Charge Point SHALL
//     ensure one connector is always available for the reserved idTag. Otherwise, return Rejected.
//
// Parent idTag handling:
//   - If parent idTag is provided, the Charge Point MAY look it up in the Local Authorization List
//     or Authorization Cache. If not found, the Charge Point SHALL send Authorize.req to the Central System.
//
// Reservation termination:
//
//	A reservation SHALL end when any of the following occur:
//	  1. A transaction starts for the reserved idTag or parent idTag on the reserved connector
//	     (or any connector if connectorId = 0). The reservationId SHALL be included in StartTransaction.req.
//	  2. The expiryDate is reached.
//	  3. The connector or Charge Point enters Faulted or Unavailable state.
//
// Expired reservations:
//   - The Charge Point SHALL terminate the reservation and make the connector available.
//   - StatusNotification SHALL be sent to the Central System indicating availability.
//
// Authorization Cache:
//   - If implemented, the Charge Point SHALL update the cache with IdTagInfo from ReserveNow.conf
//     if the idTag is not in the Local Authorization List.
//
// Recommendation:
//   - Validate the idTag with an Authorize.req after receiving ReserveNow.req and before starting a transaction.
package reserveNow
