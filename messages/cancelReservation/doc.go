// CancelReservation handling rules:
//
// To cancel a reservation, the Central System SHALL send a
// CancelReservation.req to the Charge Point.
//
// If the Charge Point has an active reservation matching the provided
// reservationId, it SHALL return status Accepted.
// If no matching reservation exists, it SHALL return status Rejected.
package cancelReservation
